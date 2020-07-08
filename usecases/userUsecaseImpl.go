package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"net/http"
)

type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (u UserUsecaseImpl) GetAllUser() ([]*models.User, error) {
	panic("implement me")
}

func (u UserUsecaseImpl) InsertUser(user *models.User) error {
	//if user.Username == "" || user.Password == "" {
	//	return errors.New("can't null please check")
	//} else {
		err := u.userRepo.InsertUser(user)
		if err != nil {
			return err
		}
		return nil
	//}
}

func (u UserUsecaseImpl) UpdateUser(request *http.Request) error {
	panic("implement me")
}

func (u UserUsecaseImpl) DeleteUser(i *int) error {
	panic("implement me")
}

func InitUserUsecase(userRepo repositories.UserRepository) UserUsecase{
	return &UserUsecaseImpl{userRepo }
}