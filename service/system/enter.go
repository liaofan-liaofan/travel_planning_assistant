package system

type SysGroup struct {
	UserService
	JwtService
	OperationRecordService
	AuthorityService
	ApiService
	MenuService
	DeptService
	SystemConfigService
	CasbinService
	FileUploadAndDownloadService
}
