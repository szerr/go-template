package er

// 声明内置异常,错误码用固定编码,避免代码更改时导致变动
var (
	// 系统状态
	Ok                 = NewErr(0, "Ok", "ok", Error, true)                                  // 没有错误
	Cancelled          = NewErr(1, "Cancelled", "cancelled", Error, true)                    // 操作被调用者取消
	Unknown            = NewErr(2, "Unknown", "unknown", Error, true)                        // 未知错误（服务端错误）
	InvalidArgument    = NewErr(3, "InvalidArgument", "invalid_argument", Error, true)       // 无效的参数，参数错误
	DeadlineExceeded   = NewErr(4, "DeadlineExceeded", "deadline_exceeded", Error, true)     // 超时
	NotFound           = NewErr(5, "NotFound", "not_found", Error, true)                     // 未找到
	AlreadyExists      = NewErr(6, "AlreadyExists", "already_exists", Error, true)           // 已存在
	PermissionDenied   = NewErr(7, "PermissionDenied", "permission_denied", Error, true)     // 没有权限
	ResourceExhausted  = NewErr(8, "ResourceExhausted", "resource_exhausted", Error, true)   // 资源已耗尽
	FailedPrecondition = NewErr(9, "FailedPrecondition", "failed_precondition", Error, true) // 操作被拒绝
	Aborted            = NewErr(10, "Aborted", "aborted", Error, true)                       // 终止
	OutOfRange         = NewErr(11, "OutOfRange", "out_of_range", Error, true)               // 超出范围
	Unimplemented      = NewErr(12, "Unimplemented", "unimplemented", Error, true)           // 未实现或不支持的操作
	Internal           = NewErr(13, "Internal", "internal", Error, true)                     // 内部错误，为严重错误保留
	Unavailable        = NewErr(14, "Unavailable", "unavailable", Error, true)               // 不可用
	DataLoss           = NewErr(15, "DataLoss", "data_loss", Error, true)                    // 数据丢失
	Unauthenticated    = NewErr(16, "Unauthenticated", "unauthenticated", Error, true)       // 该请求没有用于该操作的有效身份验证凭据。（未登录）

	// 内部错误
	UnsupportedOption           = NewErr(100, "UnsupportedOption", "unsupported option", Error, false)                    // 不支持的选项
	IncorrectValueConversionErr = NewErr(101, "IncorrectValueConversionErr", "Incorrect value conversion.", Error, false) // 不正确的数值转换
	InvalidToken                = NewErr(102, "InvalidToken", "invalid token", Error, false)                              // 令牌无效
	ConfigError                 = NewErr(103, "configError", "config error", Error, false)                                // 配置文件错误
	RowsAffectedErr             = NewErr(201, "RowsAffectedErr", "rows affected not match", Error, false)                 // 数据库操作，rows 数量与预期不同
	FirstErr                    = NewErr(202, "FirstErr", "first error", Error, false)
	FindErr                     = NewErr(203, "FindErr", "find error", Error, false)
	DuplicateEntry              = NewErr(204, "DuplicateKey", "duplicate key", Info, false) // 重复的 key

	// 业务错误
	WrongUserNameOrPassword = NewErr(300, "WrongUserNameOrPassword", "wrong user name or password", Error, true) // 用户名或密码错误
)
