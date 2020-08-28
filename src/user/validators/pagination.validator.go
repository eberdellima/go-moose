package validators

import (
	"go-moose/src/user/inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateRequestPagination checks whether the provided pagination
// parametes are valid, otherwise returns bad request
func ValidateRequestPagination() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var paginator inputs.Paginator

		if err := ctx.ShouldBindQuery(&paginator); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("paginator", paginator)
		ctx.Next()
	}
}
