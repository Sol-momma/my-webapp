import React from "react";

// ファイルの型情報を定義します
interface File {
  id: string;
  name: string;
  size?: number;
}

// FileList コンポーネントが受け取る props の型を定義します
interface FileListProps {
  files: File[];
}

// FileList コンポーネント
const FileList: React.FC<FileListProps> = ({ files }) => {
  // ファイルが存在しない場合の表示
  if (files.length === 0) {
    return <p>ファイルがありません</p>;
  }

  // ファイル情報をリストとして表示
  return (
    <ul>
      {files.map((file) => (
        <li key={file.id}>
          <span>{file.name}</span>
          {file.size !== undefined && (
            <span> ({file.size.toLocaleString()} バイト)</span>
          )}
        </li>
      ))}
    </ul>
  );
};

export default FileList;
