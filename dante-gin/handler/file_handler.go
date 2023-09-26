package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const dst = "/Users/dante/Documents/Project/go-world/src/dante-gin/"

func UploadMultiFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		// 上传文件至指定目录
		err := c.SaveUploadedFile(file, dst+file.Filename)
		if err != nil {
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

/**

curl -X POST http://localhost:8900/upload \                                                                                      dante@dandingdeMacBook-Pro
  -F "upload[]=@/Users/dante/Documents/Project/go-world/src/dante-gin/handler/file1.txt" \
  -F "upload[]=@/Users/dante/Documents/Project/go-world/src/dante-gin/handler/file2.csv" \
  -H "Content-Type: multipart/form-data"

*/
