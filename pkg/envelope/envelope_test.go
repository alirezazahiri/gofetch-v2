package envelope_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alirezazahiri/gofetch-v2/pkg/envelope"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	w := httptest.NewRecorder()
	return router, w
}

// Example: Basic success response
func ExampleOK() {
	router, w := setupTestRouter()

	router.GET("/users/:id", func(c *gin.Context) {
		user := map[string]interface{}{
			"id":    c.Param("id"),
			"name":  "John Doe",
			"email": "john@example.com",
		}
		envelope.OK(c, user)
	})

	req := httptest.NewRequest("GET", "/users/123", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": true,
	//   "data": {
	//     "id": "123",
	//     "name": "John Doe",
	//     "email": "john@example.com"
	//   }
	// }
}

// Example: Created response
func ExampleCreated() {
	router, w := setupTestRouter()

	router.POST("/users", func(c *gin.Context) {
		newUser := map[string]interface{}{
			"id":         "456",
			"name":       "Jane Smith",
			"email":      "jane@example.com",
			"created_at": time.Now(),
		}
		envelope.Created(c, newUser)
	})

	req := httptest.NewRequest("POST", "/users", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": true,
	//   "data": {
	//     "id": "456",
	//     "name": "Jane Smith",
	//     "email": "jane@example.com",
	//     "created_at": "2024-01-01T00:00:00Z"
	//   }
	// }
}

// Example: Response with pagination metadata
func ExampleSuccessWithMeta() {
	router, w := setupTestRouter()

	router.GET("/users", func(c *gin.Context) {
		users := []map[string]interface{}{
			{"id": "1", "name": "User 1"},
			{"id": "2", "name": "User 2"},
		}

		meta := &envelope.Meta{
			Page:       1,
			PageSize:   20,
			TotalPages: 5,
			TotalItems: 100,
		}

		envelope.SuccessWithMeta(c, http.StatusOK, users, meta)
	})

	req := httptest.NewRequest("GET", "/users?page=1", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": true,
	//   "data": [
	//     {"id": "1", "name": "User 1"},
	//     {"id": "2", "name": "User 2"}
	//   ],
	//   "meta": {
	//     "page": 1,
	//     "page_size": 20,
	//     "total_pages": 5,
	//     "total_items": 100
	//   }
	// }
}

// Example: Bad request error
func ExampleBadRequest() {
	router, w := setupTestRouter()

	router.POST("/users", func(c *gin.Context) {
		envelope.BadRequest(c, "Invalid request format", map[string]string{
			"email": "invalid email format",
		})
	})

	req := httptest.NewRequest("POST", "/users", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": false,
	//   "error": {
	//     "code": "BAD_REQUEST",
	//     "message": "Invalid request format",
	//     "details": {
	//       "email": "invalid email format"
	//     }
	//   }
	// }
}

// Example: Validation error
func ExampleValidationError() {
	router, w := setupTestRouter()

	router.POST("/jobs", func(c *gin.Context) {
		envelope.ValidationError(c, "Validation failed", map[string]string{
			"urls":        "must not be empty",
			"concurrency": "must be greater than 0",
		})
	})

	req := httptest.NewRequest("POST", "/jobs", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": false,
	//   "error": {
	//     "code": "VALIDATION_ERROR",
	//     "message": "Validation failed",
	//     "details": {
	//       "urls": "must not be empty",
	//       "concurrency": "must be greater than 0"
	//     }
	//   }
	// }
}

// Example: Not found error
func ExampleNotFound() {
	router, w := setupTestRouter()

	router.GET("/users/:id", func(c *gin.Context) {
		envelope.NotFound(c, "User not found")
	})

	req := httptest.NewRequest("GET", "/users/999", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": false,
	//   "error": {
	//     "code": "NOT_FOUND",
	//     "message": "User not found"
	//   }
	// }
}

// Example: Internal server error
func ExampleInternalServerError() {
	router, w := setupTestRouter()

	router.GET("/users", func(c *gin.Context) {
		envelope.InternalServerError(c, "Database connection failed", nil)
	})

	req := httptest.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	// Response will be:
	// {
	//   "success": false,
	//   "error": {
	//     "code": "INTERNAL_SERVER_ERROR",
	//     "message": "Database connection failed"
	//   }
	// }
}

// Test: Verify response structure
func TestResponseStructure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	w := httptest.NewRecorder()

	router.GET("/test", func(c *gin.Context) {
		envelope.OK(c, map[string]string{"test": "data"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"success":true`)
	assert.Contains(t, w.Body.String(), `"data"`)
}

// Test: Verify error response structure
func TestErrorResponseStructure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	w := httptest.NewRecorder()

	router.GET("/test", func(c *gin.Context) {
		envelope.BadRequest(c, "test error", nil)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"success":false`)
	assert.Contains(t, w.Body.String(), `"error"`)
	assert.Contains(t, w.Body.String(), `"code":"BAD_REQUEST"`)
}
