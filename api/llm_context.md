You are a coding assistant, i will be asking you question and give answer for this project and answer should be with best industry practices
Project: Expense Tracker API  
Tech: Golang, Fiber  
Structure:  
/github.com/chiragthapa777/expense-tracker-api
├── /cmd/server/main.go         
├── /internal
│   ├── /config/config.go  
│   ├── /constant/response_constant.go
│   ├── /database/database.go  
│   ├── /dto/auth_dto.go        
│   ├── /logger/logger.go     
│   ├── /middleware/auth.go   
│   ├── /models/user.go   
│   ├── /modules/user    
│   │   ├── user_admin.go    
│   │   ├── user_user.go     
│   ├── /repository        
│   │   ├── base_repository.go  
│   │   ├── user_repository.go 
│   ├── /request/request.go   
│   ├── /response/response.go
│   ├── /routes/user.go   
│   ├── /types/response.go 
│   ├── /utils/jwt.go    
Key Components:
Main (cmd/server/main.go):
Initializes logger, config, DB, Fiber app. Sets up CORS middleware and routes. Starts server on configurable port.
Code: func main() { log := logger.GetLogger(); cfg := config.GetConfig(); database.InitDB(); app := fiber.New(); app.Use(middleware.CORSMiddleware()); routes.SetUpRoutes(app); log.Info("Server starting on port " + cfg.Port); app.Listen(":" + cfg.Port) }
Config (config/config.go):
Loads .env vars with defaults (e.g., Port: "3000", DBHost: "localhost"). Validates fields (e.g., Port, DBUser, JWTSecret) using validator. Singleton pattern with sync.Once.
Code: type Config struct { Port, DBHost, DBPort, DBUser, DBPass, DBName, JWTSecret string validate:"required" }; func GetConfig() Config {...}
Request (request/request.go):
Parses and validates request body/query using Fiber and validator.
Code: func LoadAndValidateBody(body any, c *fiber.Ctx) error {...}
Response (response/response.go):
Formats success/error responses with status codes and metadata. Handles validation errors, not-found cases.
Code: func SendError(c *fiber.Ctx, opt types.ErrorResponseOption) error {...}; func Send(c *fiber.Ctx, opt types.ResponseOption) error {...}
Routes (routes/user.go):
Defines user and admin routes with auth middleware. Example: /user/update-profile, /admin/user/:id/block.
Code: func SetupUserRoutes(app fiber.Router) { userRoute := app.Group("/user", middleware.AuthGuard); adminUserRoute := app.Group("/admin/user", middleware.AuthGuardAdminOnly); userRoute.Put("/update-profile", user.UpdateProfile); adminUserRoute.Patch("/:id/block", user.AdminBlockUser) }
Models (models/user.go):
Defines User with fields like FirstName, Email, Role (USER/ADMIN).
Code: type User struct { FirstName string; Email *string; Role UserRoleEnum; BlockedAt *time.Time }
User Module (modules/user/user_admin.go):
Handles admin actions (e.g., block, unblock, update user). Uses repository for DB ops, validates UUIDs.
Code: func AdminBlockUser(c *fiber.Ctx) error { userId := c.Params("id"); userRepo := repository.NewUserRepository(); user, err := userRepo.FindByID(userId, repository.Option{}); if user.BlockedAt != nil { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user already blocked")}) }; userRepo.BlockUser(userId, repository.Option{}); return response.Send(c, types.ResponseOption{Data: "user blocked"}) }
Repository (repository/):
BaseRepository: Generic CRUD with pagination. UserRepository: Extends base with user-specific methods (e.g., BlockUser).
Code: type BaseRepository[T any] struct { db *gorm.DB }; func (r *BaseRepository[T]) FindWithPagination(option Option, searchFields []string, validSortColumns map[string]string) (*PaginationResult[T], error) {...}; type UserRepository struct { *BaseRepository[models.User] }; func (r *UserRepository) BlockUser(userId string, option Option) error {...}
Base Repository (repository/base_repository.go):
Generic CRUD operations with pagination support for any model type. Uses GORM for DB interactions.
Code: type BaseRepository[T any] struct { db *gorm.DB }; func NewBaseRepositoryT any *BaseRepository[T] { if database.DB == nil { panic("database.DB is not initialized") }; return &BaseRepository[T]{db: database.DB} }; func (r *BaseRepository[T]) getDB(option Option) *gorm.DB { if option.Tx != nil { return option.Tx }; return r.db }; func (r *BaseRepository[T]) Create(entity *T, option Option) error { db := r.getDB(option); if idSetter, ok := any(entity).(interface { GetID() string; SetID(string) }); ok && idSetter.GetID() == "" { idSetter.SetID(uuid.New().String()) }; return db.Create(entity).Error }; func (r *BaseRepository[T]) FindByID(id string, option Option) (*T, error) { db := r.getDB(option); var entity T; err := db.Where("id = ?", id).First(&entity).Error; if err != nil { if err == gorm.ErrRecordNotFound { return nil, nil }; return nil, err }; return &entity, nil }; func (r *BaseRepository[T]) Update(entity *T, option Option) error { db := r.getDB(option); return db.Save(entity).Error }; func (r *BaseRepository[T]) Delete(id string, option Option) error { db := r.getDB(option); var entity T; return db.Where("id = ?", id).Delete(&entity).Error }; func (r *BaseRepository[T]) Find(option Option) ([]T, error) { db := r.getDB(option); var entities []T; if err := db.Find(&entities).Error; err != nil { return nil, err }; return entities, nil }; func (r *BaseRepository[T]) FindWithPagination(option Option, searchFields []string, validSortColumns map[string]string) (*PaginationResult[T], error) { db := r.getDB(option); queryBuilder := db.Model(new(T)); if option.PaginationDto.Search != "" { searchString := "%" + option.PaginationDto.Search + "%"; for i, field := range searchFields { if i == 0 { queryBuilder.Where(field+" ILIKE ?", searchString) } else { queryBuilder.Or(field+" ILIKE ?", searchString) } } }; var total int64; if err := queryBuilder.Count(&total).Error; err != nil { return nil, err }; limit, offset := GetLimitAndOffSet(*option.PaginationDto); order := GetOrderString(*option.PaginationDto, validSortColumns); var entities []T; if err := queryBuilder.Limit(limit).Offset(offset).Order(order).Find(&entities).Error; err != nil { return nil, err }; return &PaginationResult[T]{MetaData: types.ResponsePaginationMeta{Total: total, Limit: limit, CurrentPage: PageFromQueryString(option.PaginationDto.Page), TotalPages: int(math.Ceil(float64(total) / float64(limit)))}, Data: entities}, nil }
type UserRoleEnum string; const ( UserRoleUser UserRoleEnum = "USER"; UserRoleAdmin UserRoleEnum = "ADMIN" ); type User struct { BaseModel; FirstName string gorm:"column:first_name;type:varchar;not null" json:"firstName"; LastName *string gorm:"column:last_name;type:varchar" json:"lastName"; Email *string gorm:"type:varchar" json:"email"; EmailVerifiedAt *time.Time gorm:"column:email_verified_at" json:"emailVerifiedAt"; BlockedAt *time.Time gorm:"column:blocked_at" json:"blockedAt"; Password string gorm:"type:varchar;not null" json:"-"; IsPasswordSetByUser bool gorm:"type:boolean;not null;column:is_password_set_by_user" json:"isFirstLoggedIn"; Role UserRoleEnum; LoginLogs *[]LoginLog gorm:"foreignKey:user_id;references:id" json:"loginLogs,omitempty" }
User Repository (repository/user_repository.go):
Extends BaseRepository with user-specific methods (e.g., BlockUser).
Code: type UserRepository struct { *BaseRepository[models.User] }; func NewUserRepository() *UserRepository { return &UserRepository{BaseRepository: NewBaseRepositorymodels.User} }; func (r *UserRepository) FindWithPagination(option Option) (*PaginationResult[models.User], error) { searchFields := []string{"email", "first_name", "last_name"}; return r.BaseRepository.FindWithPagination(option, searchFields, map[string]string{"id": "id", "email": "email", "firstName": "first_name", "lastName": "last_name", "createdAt": "created_at"}) }; func (r *UserRepository) BlockUser(userId string, option Option) error { db := r.getDB(option); return db.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]any{"blocked_at": time.Now()}).Error }
Create User:
Code: func AdminCreateUser(c *fiber.Ctx) error { body := new(dto.UserCreateDto); if err := request.LoadAndValidateBody(body, c); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; userRepo := repository.NewUserRepository(); user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: body.Password, Role: models.UserRoleUser}; if err := userRepo.Create(&user, repository.Option{}); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; return response.Send(c, types.ResponseOption{Data: user, Status: fiber.StatusCreated}) }
Read Users (List):
Code: func AdminGetUsers(c *fiber.Ctx) error { query := dto.PaginationQueryDto{}; if err := request.LoadAndValidateQuery(&query, c); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; userRepo := repository.NewUserRepository(); result, err := userRepo.FindWithPagination(repository.Option{PaginationDto: &query}); if err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; return response.Send(c, types.ResponseOption{Data: result.Data, MetaData: &types.ResponseMetaData{PaginationMetaData: &result.MetaData}}) }
Read User (Single):
Code: func AdminGetUser(c *fiber.Ctx) error { userId := c.Params("id"); if err := uuid.Validate(userId); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest}) }; userRepo := repository.NewUserRepository(); user, err := userRepo.FindByID(userId, repository.Option{}); if err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; if user == nil { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user not found"), Status: fiber.StatusNotFound}) }; return response.Send(c, types.ResponseOption{Data: user}) }
Update User:
Code: func AdminUpdateUser(c *fiber.Ctx) error { body := new(dto.UserProfileUpdateDto); if err := request.LoadAndValidateBody(body, c); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; userId := c.Params("id"); if err := uuid.Validate(userId); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest}) }; userRepo := repository.NewUserRepository(); user, err := userRepo.FindByID(userId, repository.Option{}); if err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; if user == nil { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user not found"), Status: fiber.StatusNotFound}) }; user.FirstName = body.FirstName; user.LastName = body.LastName; if err := userRepo.Update(user, repository.Option{}); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; return response.Send(c, types.ResponseOption{Data: user}) }
Delete User:
Code: func AdminDeleteUser(c *fiber.Ctx) error { userId := c.Params("id"); if err := uuid.Validate(userId); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest}) }; userRepo := repository.NewUserRepository(); user, err := userRepo.FindByID(userId, repository.Option{}); if err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; if user == nil { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user not found"), Status: fiber.StatusNotFound}) }; if user.Role == models.UserRoleAdmin { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("cannot delete admin"), Status: fiber.StatusBadRequest}) }; if err := userRepo.Delete(userId, repository.Option{}); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; return response.Send(c, types.ResponseOption{Data: "user deleted", Status: fiber.StatusNoContent}) }
Block User (Existing):
Code: func AdminBlockUser(c *fiber.Ctx) error { userId := c.Params("id"); if err := uuid.Validate(userId); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest}) }; userRepo := repository.NewUserRepository(); user, err := userRepo.FindByID(userId, repository.Option{}); if err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; if user == nil { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user not found"), Status: fiber.StatusNotFound}) }; if user.BlockedAt != nil { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user already blocked"), Status: fiber.StatusBadRequest}) }; if user.Role == models.UserRoleAdmin { return response.SendError(c, types.ErrorResponseOption{Error: errors.New("cannot block admin"), Status: fiber.StatusBadRequest}) }; if err := userRepo.BlockUser(userId, repository.Option{}); err != nil { return response.SendError(c, types.ErrorResponseOption{Error: err}) }; return response.Send(c, types.ResponseOption{Data: "user blocked"}) }
Purpose: RESTful API for managing expenses with user authentication, admin controls, and DB integration. Modular, scalable design with separation of concerns.
