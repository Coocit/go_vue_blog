package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue_blog/models/res"
)

// ImageUploadView 上传单个图片,返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)
	fmt.Println(fileHeader.Filename)
}
