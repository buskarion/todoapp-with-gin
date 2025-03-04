package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Server is up and running."})
}
