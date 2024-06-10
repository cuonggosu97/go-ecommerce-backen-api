package service

import "github.com/cuonggosu97/go-ecommerce-backen-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUserService() string {
	return us.userRepo.GetInfoUser()
}
