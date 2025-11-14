# Envelope Package

The `envelope` package provides a standardized way to structure HTTP API responses in your Go application. It ensures consistency across all endpoints for both success and error responses.

## Response Structure

All responses follow this structure:

```json
{
  "success": true|false,
  "data": {...},        // Present on success
  "error": {...},       // Present on error
  "meta": {...}         // Optional metadata (pagination, etc.)
}
```

### Success Response

```json
{
  "success": true,
  "data": {
    "job_id": "abc123",
    "status": "running"
  }
}
```

### Error Response

```json
{
  "success": false,
  "error": {
    "code": "BAD_REQUEST",
    "message": "Invalid input provided",
    "details": {
      "field": "urls",
      "reason": "must be a non-empty array"
    }
  }
}
```

### Response with Metadata (Pagination)

```json
{
  "success": true,
  "data": [...],
  "meta": {
    "page": 1,
    "page_size": 20,
    "total_pages": 5,
    "total_items": 100
  }
}
```

## Usage

### Success Responses

```go
import "github.com/alirezazahiri/gofetch-v2/pkg/envelope"

// 200 OK
envelope.OK(c, data)

// 201 Created
envelope.Created(c, data)

// 202 Accepted
envelope.Accepted(c, data)

// 204 No Content
envelope.NoContent(c)

// With custom status code
envelope.Success(c, http.StatusOK, data)

// With metadata (pagination)
meta := &envelope.Meta{
    Page:       1,
    PageSize:   20,
    TotalPages: 5,
    TotalItems: 100,
}
envelope.SuccessWithMeta(c, http.StatusOK, data, meta)
```

### Error Responses

```go
// 400 Bad Request
envelope.BadRequest(c, "Invalid request format", validationErrors)

// 401 Unauthorized
envelope.Unauthorized(c, "Authentication required")

// 403 Forbidden
envelope.Forbidden(c, "Access denied")

// 404 Not Found
envelope.NotFound(c, "Resource not found")

// 409 Conflict
envelope.Conflict(c, "Resource already exists", nil)

// 422 Validation Error
envelope.ValidationError(c, "Validation failed", validationDetails)

// 500 Internal Server Error
envelope.InternalServerError(c, "An unexpected error occurred", nil)

// Custom error with custom code
envelope.ErrorResponse(c, http.StatusBadGateway, "GATEWAY_ERROR", "Upstream service unavailable", nil)
```

## Example: Using in Handlers

```go
package jobshandler

import (
	"github.com/alirezazahiri/gofetch-v2/pkg/envelope"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Check(c *gin.Context) {
	var request jobsdto.CheckRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		envelope.BadRequest(c, "Invalid request format", err.Error())
		return
	}

	// Validate request
	if len(request.Urls) == 0 {
		envelope.ValidationError(c, "Validation failed", map[string]string{
			"urls": "must not be empty",
		})
		return
	}

	// Process the request...
	response := jobsdto.CheckResponse{
		JobId:     "job-123",
		Status:    "running",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	envelope.Created(c, response)
}
```

## Benefits

1. **Consistency**: All API responses follow the same structure
2. **Type Safety**: Strongly typed response structures
3. **Easy to Parse**: Clients can easily determine success/failure and extract data
4. **Extensible**: Easy to add metadata like pagination
5. **Error Details**: Structured error information with codes and details
6. **DRY**: Reusable helper functions reduce boilerplate code

## Error Codes

Standard error codes used:
- `BAD_REQUEST`: Invalid request format or parameters
- `UNAUTHORIZED`: Authentication required
- `FORBIDDEN`: Insufficient permissions
- `NOT_FOUND`: Resource not found
- `CONFLICT`: Resource conflict (e.g., duplicate)
- `VALIDATION_ERROR`: Request validation failed
- `INTERNAL_SERVER_ERROR`: Unexpected server error

Custom error codes can be used with the `ErrorResponse` function.

