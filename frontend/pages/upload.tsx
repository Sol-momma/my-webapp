import React, { useState } from "react";

const UploadPage: React.FC = () => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);

  // ファイル選択時のハンドラー
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setSelectedFile(file);
      const url = URL.createObjectURL(file);
      setPreviewUrl(url);
    }
  };

  // ファイルアップロード処理
  const handleUpload = async () => {
    if (!selectedFile) return;

    const formData = new FormData();
    formData.append("file", selectedFile);

    try {
      const response = await fetch("/api/upload", {
        method: "POST",
        body: formData,
      });
      const data = await response.json();
      console.log("アップロード完了:", data);
    } catch (error) {
      console.error("アップロードエラー:", error);
    }
  };

  return (
    <div>
      <h1>ファイルアップロード</h1>
      <input type="file" onChange={handleFileChange} />
      {previewUrl && (
        <div>
          <h2>プレビュー</h2>
          <img src={previewUrl} alt="プレビュー" style={{ maxWidth: "300px", marginTop: "20px" }} />
        </div>
      )}
      <button onClick={handleUpload} disabled={!selectedFile}>
        アップロード
      </button>
    </div>
  );
};

export default UploadPage;
