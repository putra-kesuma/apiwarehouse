package repositories

import (
	"apiwarehouse/models"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserRepoImp struct {
	db *sql.DB
}

func (u UserRepoImp) GetAllUser() ([]*models.User, error) {
	panic("implement me")
}

func (u UserRepoImp) InsertUser(user *models.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := `insert into m_user (username,password)
								value (?,?);`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return err
	}
	defer stmt.Close()
	passHash,_ := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if _, err := stmt.Exec(user.Username,passHash); err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	return tx.Commit()
}

func (u UserRepoImp) UpdateUser(request *http.Request) error {
	panic("implement me")
}

func (u UserRepoImp) DeleteUser(i *int) error {
	panic("implement me")
}

func InitUserRepoImpl(db *sql.DB) UserRepository  {
	return &UserRepoImp{db}
}