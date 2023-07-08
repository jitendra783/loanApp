package service

import (
	"loanApp/pkg/db"
	v1 "loanApp/pkg/service/v1"
	v2 "loanApp/pkg/service/v2"
) 
type service struct{
	V1 v1.ServiceLayer
	V2 v2.ServiceLayer
}

type ServiceGroupLayer interface{
    GetV1Service() v1.ServiceLayer
	GetV2Service() v2.ServiceLayer
}

func NewServiceGroupObject(db db.DBLayer) ServiceGroupLayer {
	return &service{
		v1.NewServiceObject(db),
		v2.NewServiceObject(db),
	}
}

func (s *service) GetV1Service() v1.ServiceLayer {
	return s.V1
}

func (s *service) GetV2Service() v2.ServiceLayer {
	return s.V2
}