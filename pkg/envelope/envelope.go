package envelope

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents the standard API response envelope
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// Error represents the error details in the response
type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Meta represents metadata for responses (pagination, etc.)
type Meta struct {
	Page       int   `json:"page,omitempty"`
	PageSize   int   `json:"page_size,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
	TotalItems int64 `json:"total_items,omitempty"`
}

// Success sends a successful response with data
func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
	})
}

// SuccessWithMeta sends a successful response with data and metadata
func SuccessWithMeta(c *gin.Context, statusCode int, data interface{}, meta *Meta) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, code string, message string, details interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// BadRequest sends a 400 Bad Request error
func BadRequest(c *gin.Context, message string, details interface{}) {
	ErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", message, details)
}

// Unauthorized sends a 401 Unauthorized error
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", message, nil)
}

// Forbidden sends a 403 Forbidden error
func Forbidden(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, "FORBIDDEN", message, nil)
}

// NotFound sends a 404 Not Found error
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", message, nil)
}

// Conflict sends a 409 Conflict error
func Conflict(c *gin.Context, message string, details interface{}) {
	ErrorResponse(c, http.StatusConflict, "CONFLICT", message, details)
}

// InternalServerError sends a 500 Internal Server Error
func InternalServerError(c *gin.Context, message string, details interface{}) {
	ErrorResponse(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message, details)
}

// ValidationError sends a 422 Unprocessable Entity error for validation failures
func ValidationError(c *gin.Context, message string, details interface{}) {
	ErrorResponse(c, http.StatusUnprocessableEntity, "VALIDATION_ERROR", message, details)
}

// Created sends a 201 Created success response
func Created(c *gin.Context, data interface{}) {
	Success(c, http.StatusCreated, data)
}

// OK sends a 200 OK success response
func OK(c *gin.Context, data interface{}) {
	Success(c, http.StatusOK, data)
}

// NoContent sends a 204 No Content success response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// Accepted sends a 202 Accepted success response
func Accepted(c *gin.Context, data interface{}) {
	Success(c, http.StatusAccepted, data)
}
