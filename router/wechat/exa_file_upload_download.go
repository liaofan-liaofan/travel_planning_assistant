package wechat

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	var exaFileUploadAndDownloadApi = v1.ApiGroupApp.SystemApiGroup.SysFileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile) // 上传文件
	}
}
