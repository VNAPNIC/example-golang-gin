package code

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400
	TokenInvalid  = 401
	UnknownError  = 900

	ErrorAuthCheckTokenFail     = 20001
	ErrorAuthCheckTokenTimeout  = 20002
	ErrorAuthToken              = 20003
	ErrorAuth                   = 20004
	ErrorUserPasswordInvalid    = 20005
	AuthTokenInBlockList        = 20006
	ErrorUserOldPasswordInvalid = 20007
	ErrorFailedAddNewUser       = 20008
	ErrorFailedAddNewRole       = 20009
)

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "fail",
	UnknownError:                "unknown mistake",
	InvalidParams:               "wrong request parameter",
	TokenInvalid:                "The Token parameter is invalid or does not exist",
	ErrorAuthCheckTokenFail:     "Token authentication failed",
	ErrorAuthCheckTokenTimeout:  "Token has timed out",
	ErrorAuthToken:              "Token generation failed",
	ErrorAuth:                   "Token error",
	ErrorUserPasswordInvalid:    "The corresponding username or password is incorrect",
	AuthTokenInBlockList:        "The Token already exists in blockList",
	ErrorUserOldPasswordInvalid: "The corresponding user's original password is incorrect",
	ErrorFailedAddNewUser:       "Failed to add new user",
	ErrorFailedAddNewRole:       "Failed to add new role",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
