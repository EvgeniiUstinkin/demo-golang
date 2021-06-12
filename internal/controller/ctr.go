package controller

import (
	"github.com/booomch/demo-golang/internal/entities/contact"
	"github.com/booomch/demo-golang/internal/entities/user"
)

type Ctr struct {
	User    user.Usecase
	Contact contact.Usecase
}

func New(
	contactUsecase contact.Usecase,
	userUsecase user.Usecase,
) *Ctr {
	return &Ctr{
		Contact: contactUsecase,
		User:    userUsecase,
	}
}
