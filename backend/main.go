package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"my-webapp/backend/handlers"
)

func main() {
	// gin.Default() は Logger と Recovery ミドルウェアが組み込まれている
	router := gin.Default()

	// APIエンドポイントの登録
	router.POST("/upload", handlers.UploadHandler)
	router.GET("/files", handlers.ListFilesHandler)
	router.GET("/download/:id", handlers.DownloadHandler)
	router.DELETE("/files/:id", handlers.DeleteFileHandler)

	// PORT環境変数がなければ8080を使用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
