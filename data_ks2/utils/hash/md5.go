package hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func MD5(file multipart.File) string {
	hash := md5.New()
	io.Copy(hash, file)
	hashByDate := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByDate)
}
