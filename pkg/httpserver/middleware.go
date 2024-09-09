package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	// You can customize the CORS configuration here
	// For example, to allow only specific origins:
	// config.AllowOrigins = []string{"http://example.com"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	return cors.New(config)
}

func DbTransactionMiddleware(handler func(Context), db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := db.Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// set db_tx variable
		ctx.Set("db_tx", tx)

		// before request
		ctx.Next()

		// after request
		convertToGinHandler(handler)(ctx)
		if len(ctx.Errors) > 0 {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}
