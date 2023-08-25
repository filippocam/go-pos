package handler

import (
	"github.com/gin-gonic/gin"
)

// errorResponse returns a JSON response with the error message and status code
func errorResponse(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

// successResponse returns a JSON response with the success message and data
func successResponse(ctx *gin.Context, statusCode int, data any) {
	if data == nil {
		ctx.JSON(statusCode, gin.H{
			"success": true,
			"message": "Success",
		})
	} else {
		ctx.JSON(statusCode, gin.H{
			"success": true,
			"message": "Success",
			"data":    data,
		})
	}
}

// meta represents metadata for a paginated response
type meta struct {
	Total uint64 `json:"total"`
	Limit uint64 `json:"limit"`
	Skip  uint64 `json:"skip"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip uint64) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}