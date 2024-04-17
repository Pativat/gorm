package helper

type SuccessResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
}

type UnAuthorizeResponse struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"UNAUTHORIZED"`
}

type NotFoundResponse struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"NOT FOUND"`
}

type InternalServerErrorResponse struct {
	Code    int    `json:"code" example:"5000"`
	Message string `json:"message" example:"INTERNAL SERVER ERROR"`
}
