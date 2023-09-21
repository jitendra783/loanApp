package onboarding

import (
	dbuser "loanApp/pkg/db/onbaord"
	e "loanApp/pkg/error"
)

type GetOtpParam struct {
	Phone string `json:"phone"`
}
type VerifyOtpParam struct {
	OtpReferenceNumnber string `json:"otpReferenceNumber"`
	Otp                 string `json:"otp"`
}
type GetMpinParam struct {
	Mpin string `json:"mpin"`
}

type ApiResponse struct {
	Data    []dbuser.User `json:"data,omitempty"`
	Success bool          `json:"success"`
	Errors  []e.Error     `json:"error,omitempty"`
	Message string        `json:"message"`
}
