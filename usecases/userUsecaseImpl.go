package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"errors"
)

type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (u UserUsecaseImpl) UpdateUser(user *models.User) error {
	err := u.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecaseImpl) LoginUser(user *models.User) error {
	err := u.userRepo.LoginUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecaseImpl) GetUser() ([]*models.User, error) {
	user, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserUsecaseImpl) InsertUser(user *models.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("can't null please check")
	} else {
		err := u.userRepo.InsertUser(user)
		if err != nil {
			return err
		}
		return nil
	}
}



func (u UserUsecaseImpl) DeleteUser(i *int) error {
	panic("implement me")
}

func InitUserUsecase(userRepo repositories.UserRepository) UserUsecase{
	return &UserUsecaseImpl{userRepo }
}