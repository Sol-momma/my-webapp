package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// FileInfo はファイル情報を表す構造体です
type FileInfo struct {
	ID       string    `json:"id"`
	FileName string    `json:"file_name"`
	FileURL  string    `json:"file_url"`
	Uploaded time.Time `json:"uploaded_at"`
}

// ListFilesHandler はアップロード済みファイルの一覧を返します
func ListFilesHandler(c *gin.Context) {
	// DBからファイル情報を取得する処理を実装してください
	// ここでは疑似的なデータを返します
	dummyFiles := []FileInfo{
		{
			ID:       "1",
			FileName: "example.jpg",
			FileURL:  "https://your-supabase-storage-url/example.jpg",
			Uploaded: time.Now().Add(-time.Hour),
		},
		{
			ID:       "2",
			FileName: "document.pdf",
			FileURL:  "https://your-supabase-storage-url/document.pdf",
			Uploaded: time.Now().Add(-2 * time.Hour),
		},
	}
	c.JSON(http.StatusOK, gin.H{"files": dummyFiles})
}

// DownloadHandler は指定したIDのファイルのダウンロード情報を返します
func DownloadHandler(c *gin.Context) {
	id := c.Param("id")
	// DBからIDに紐づくファイル情報を取得する処理を実装してください
	fileURL := "https://your-supabase-storage-url/dummy-file.jpg"
	c.JSON(http.StatusOK, gin.H{"id": id, "file_url": fileURL})
}

// DeleteFileHandler は指定したIDのファイルを削除します
func DeleteFileHandler(c *gin.Context) {
	id := c.Param("id")
	// DBおよびSupabase Storage上での削除処理を実装してください
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("ファイル %s を削除しました", id)})
}
