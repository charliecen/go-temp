package user

type UserServiceGroup struct {
	OperationRecordService
	UserService
	JwtService
	CasbinService
}
