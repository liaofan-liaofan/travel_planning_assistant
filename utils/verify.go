package utils

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty(), Gt("0")}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityIds": {NotEmpty()}, "DeptId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	DeleteUserVerify       = Rules{"ID": {NotEmpty(), Gt("2")}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	DeptVerify             = Rules{"DeptName": {NotEmpty()}, "ParentID": {NotEmpty()}}
	DeleteDeptVerify       = Rules{"ID": {NotEmpty(), Gt("2")}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "DataScope": {NotEmpty()}, "Level": {NotEmpty()}}
	DeleteAuthorityVerify  = Rules{"AuthorityId": {NotEmpty(), Gt("3")}}
)
