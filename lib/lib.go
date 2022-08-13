package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateJSON(result interface{}) gin.H {
	return gin.H{
		"status": http.StatusOK,
		"data":   result,
	}
}
