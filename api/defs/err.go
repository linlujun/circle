package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	//请求解析错误
	ErrorRequestBodyParseFailed = ErrResponse{
		HttpSC: 400, Error: Err{
			Error: "Request body is not correct", ErrorCode: "001"},
	}
	//用户验证不通过
	ErrorNotAuthUser = ErrResponse{
		HttpSC: 401, Error: Err{
			Error: "User authentication failed.", ErrorCode: "002"},
	}
	//DB错误
	ErrorDBError = ErrResponse{
		HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"},
	}
	//内部错误
	ErrorInternalFaults = ErrResponse{
		HttpSC: 500, Error: Err{
			Error:     "Internal service error",
			ErrorCode: "004"},
	}
	//用户角色权限错误
	ErrorRoleFaults = ErrResponse{
		HttpSC: 500, Error: Err{
			Error:     "Role privileges error",
			ErrorCode: "005"},
	}
)
