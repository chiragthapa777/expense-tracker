package constant

type ResponseCode string

const (
	ResponseCodePasswordNotSet ResponseCode = "1001"
	ResponseCodeUserBlocked    ResponseCode = "1002"
)
