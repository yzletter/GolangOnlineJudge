package service

import "gooj-backend/repository"

type UserService struct {
	rep *repository.UserRepository
}

func NewUserService(rep *repository.UserRepository) *UserService {
	return &UserService{
		rep: rep,
	}
}
