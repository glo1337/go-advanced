package user

type UserAuthByPhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type UserAuthByPhoneResponse struct {
	SessionId string `json:"session_id"`
}

type UserVerifyAuthRequest struct {
	SessionId string `json:"session_id" validate:"required"`
	Code      uint   `json:"code" validate:"required"`
}

type UserVerifyAuthResponse struct {
	Token string `json:"token"`
}
