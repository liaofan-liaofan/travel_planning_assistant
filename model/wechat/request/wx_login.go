package request

type WxCode struct {
	JsCode    string `json:"js_code" binding:"required"`    // 接收前端传来的code
	NickName  string `json:"nick_name" binding:"required"`  // 用户昵称
	AvatarUrl string `json:"avatar_url" binding:"required"` // 用户头像
}
