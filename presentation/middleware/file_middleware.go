package middleware

import (
	"io"
	"mime/multipart"
)

func GetBytesFromMultiPartFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buf, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
