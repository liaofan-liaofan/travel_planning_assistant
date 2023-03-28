package wechat

type RouterGroup struct {
	WxLoginRouter               // 授权登录
	ProfilePageRouter           // 个人信息页面
	TravelCirclesRouter         // 旅游圈页面
	FileUploadAndDownloadRouter // 文件上传
}
