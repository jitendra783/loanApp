package v1

import (
 e "loanApp/pkg/error"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Status  bool        `json:"success"`
	Errors  []e.Error   `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
}

