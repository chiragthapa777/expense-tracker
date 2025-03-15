package middleware

import (
	"fmt"
	"strings"

	fileModule "github.com/chiragthapa777/expense-tracker-api/internal/modules/file"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/chiragthapa777/expense-tracker-api/internal/utils"
	"github.com/gofiber/fiber/v2" // Replace with your actual types package
)

// FileCheckConfig defines the configuration for the middleware
type FileCheckConfig struct {
	MaxSize      int64    // Max file size in bytes (e.g., 5MB = 5 * 1024 * 1024)
	AllowedTypes []string // Allowed MIME types (e.g., "image/jpeg", "image/png")
	FieldName    string   // Form field name containing the files
	MaxFileCount int
}

var imageTypes = []string{
	"image/jpeg", // .jpg, .jpeg
	"image/png",  // .png
	"image/gif",  // .gif
}
var videoTypes = []string{
	// Videos
	"video/mp4",       // .mp4
	"video/avi",       // .avi
	"video/quicktime", // .mov
}

var documentTypes = []string{
	// PDF
	"application/pdf", // .pdf
	// CSV
	"text/csv", // .csv
	// Excel
	"application/vnd.ms-excel", // .xls (older Excel)
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", // .xlsx (modern Excel)
}

// Default config
var DefaultFileCheckConfig = FileCheckConfig{
	MaxSize:      5 * 1024 * 1024, // 5MB default
	AllowedTypes: utils.MergeSlices(imageTypes, videoTypes, documentTypes),
	FieldName:    "files",
	MaxFileCount: 10,
}

var ImageFileCheckConfig = FileCheckConfig{
	MaxSize:      5 * 1024 * 1024, // 5MB default
	AllowedTypes: utils.MergeSlices(imageTypes),
	FieldName:    "files",
	MaxFileCount: 10,
}

// FileCheck returns a middleware that validates file type and size
func FileCheck(config ...FileCheckConfig) fiber.Handler {
	// Use default config if none provided, otherwise use the first config
	cfg := DefaultFileCheckConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(c *fiber.Ctx) error {
		// Parse the multipart form
		form, err := c.MultipartForm()
		if err != nil {
			return response.SendError(c, types.ErrorResponseOption{
				Error:  fmt.Errorf("failed to parse multipart form: %v", err),
				Status: fiber.StatusBadRequest,
			})
		}

		// Get files from the specified field
		files, ok := form.File[cfg.FieldName]
		if !ok || len(files) == 0 {
			return response.SendError(c, types.ErrorResponseOption{
				Error:  fmt.Errorf("no files found in field '%s'", cfg.FieldName),
				Status: fiber.StatusBadRequest,
			})
		}

		// check for files count
		if len(files) > cfg.MaxFileCount {
			return response.SendError(c, types.ErrorResponseOption{
				Error:  fmt.Errorf("max file count is '%d'", cfg.MaxFileCount),
				Status: fiber.StatusBadRequest,
			})
		}

		// Validate each file
		for _, file := range files {
			// Check file size
			if file.Size > cfg.MaxSize {
				return response.SendError(c, types.ErrorResponseOption{
					Error: fmt.Errorf("file '%s' exceeds max size of %d bytes (got %d)",
						file.Filename, cfg.MaxSize, file.Size),
					Status: fiber.StatusBadRequest,
				})
			}
			// Detect content type
			contentType, err := fileModule.GetFileMimeType(file)
			if err != nil {
				return response.SendError(c, types.ErrorResponseOption{
					Error:  err,
					Status: fiber.StatusBadRequest,
				})
			}
			allowed := false
			for _, allowedType := range cfg.AllowedTypes {
				if strings.HasPrefix(contentType, allowedType) {
					allowed = true
					break
				}
			}
			if !allowed {
				return response.SendError(c, types.ErrorResponseOption{
					Error: fmt.Errorf("file '%s' has invalid type '%s' (allowed: %v)",
						file.Filename, contentType, cfg.AllowedTypes),
					Status: fiber.StatusBadRequest,
				})
			}

		}

		// All files passed validation, proceed to next handler
		return c.Next()
	}
}
