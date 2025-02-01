package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// FileResponse はアップロード後のレスポンス用構造体です
type FileResponse struct {
	FileID string `json:"file_id"`
}

// UploadHandler はフォームの"file"フィールドを処理してファイルをアップロードします
func UploadHandler(c *gin.Context) {
	// フォームから"file"フィールドを取得
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ファイルが見つかりません"})
		return
	}
	defer file.Close()

	// 一時ファイルとしてローカルに保存
	tempFileName := fmt.Sprintf("temp-%s", header.Filename)
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "一時ファイルの作成に失敗しました"})
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ファイルの保存に失敗しました"})
		return
	}

	// ※ここでSupabase Storage API などを呼び出し、実際のアップロードを行い、URLを取得する処理を実装してください

	// DBへファイルメタデータを登録する（ここでは疑似的にUUIDを生成）
	fileID := uuid.New().String()
	// 例：{ id, file_name, file_url, created_at } をDBに登録する処理を実装

	// 一時ファイルの削除
	os.Remove(tempFileName)

	c.JSON(http.StatusOK, FileResponse{FileID: fileID})
}
