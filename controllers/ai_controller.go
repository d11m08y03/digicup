package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const uploadDir = "./ai/upload"

func UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	if err := ctx.SaveUploadedFile(file, uploadDir+"/uploaded_image"); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

  resp, err := http.Get("http://localhost:8000/ai/classify")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

  result := string(body)

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
