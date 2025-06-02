package utils

import (
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

func UploadFile(file *multipart.FileHeader, uploadDir string) (string, error) {
    // Buat direktori jika belum ada
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        return "", err
    }

    // Tentukan path file
    filePath := filepath.Join(uploadDir, file.Filename)

    // Buka file sumber
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    // Buka atau buat file tujuan
    dst, err := os.Create(filePath)
    if err != nil {
        return "", err
    }
    defer dst.Close()

    // Salin isi file dari sumber ke tujuan
    if _, err := io.Copy(dst, src); err != nil {
        return "", err
    }

    return filePath, nil
}
