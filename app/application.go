package app

import (
	"github.com/frech87/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application ...")
	_ = router.Run("http://api.bookstore.com:8080")
}
