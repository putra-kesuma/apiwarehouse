package repositories

import (
	"apiwarehouse/models"
	"apiwarehouse/utils"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserRepoImp struct {
	db *sql.DB
}

func (u UserRepoImp) UpdateUser(user *models.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := `update m_user set username=?, 
				password=?,email=?, updated_at=? where id_user=?;`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return err
	}
	defer stmt.Close()

	//generate hash pass it must be function
	passHash,_ := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	pass := string(passHash)

	datenow := utils.GenDateNow()

	if _, err := stmt.Exec(user.Username,pass,user.Email,datenow,user.IdUser); err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	return tx.Commit()
}

func (u UserRepoImp) LoginUser(user *models.User) error {
	//make var for contain struct item
	dataUser := []*models.User{}
	//prepare query
	query := "SELECT username,password FROM m_user;"
	data, err := u.db.Query(query)
	//check error when exec the query
	if err != nil {
		return  err
	}
	//for get all data, use looping
	for data.Next() {
		//prepare contain for row data assign to item struct
		user := models.User{}
		//scan data
		var err = data.Scan(&user.Username,&user.Password)
		if err != nil {
			return err
		}
		dataUser = append(dataUser, &user)
	}

		var pass error
		var tempUser string
		for _, v := range dataUser {
			pass = bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(user.Password))
			tempUser= v.Username
			//ketika dapat yang sesuai break username dan password
			if pass == nil && v.Username == user.Username  {
				break
			}
		}
		fmt.Println(tempUser)
		fmt.Println(user.Username)

	if pass == nil && tempUser == user.Username  {
		return nil
	} else {
		return errors.New("uername dan password tidak sesuai")
	}

}

func (u UserRepoImp) GetAllUser() ([]*models.User, error) {
	//make var for contain struct user
	dataUser := []*models.User{}
	//prepare query
	query := "SELECT id_user,username,password,email,created_at FROM m_user"
	data, err := u.db.Query(query)
	//check error when exec the query
	if err != nil {
		return nil, err
	}
	//for get all data, use looping
	for data.Next() {
		//prepare contain for row data assign to item struct
		user := models.User{}
		//scan data
		var err = data.Scan(&user.IdUser, &user.Username, &user.Password, &user.Email,&user.CreatedAt)
		if err != nil {
			return nil, err
		}
		dataUser = append(dataUser, &user)
	}
	return dataUser, nil
}

func (u UserRepoImp) InsertUser(user *models.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := `insert into m_user (id_user,username,password,email,created_at)
								value (?,?,?,?,?);`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return err
	}
	defer stmt.Close()
	//make uuid
	uuid := utils.GenUUID()

	//generate hash pass it must be function
	passHash,_ := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	pass := string(passHash)

	datenow := utils.GenDateNow()

	if _, err := stmt.Exec(uuid,user.Username,pass,user.Email,datenow); err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	return tx.Commit()
}


func (u UserRepoImp) DeleteUser(i *int) error {
	panic("implement me")
}

func InitUserRepoImpl(db *sql.DB) UserRepository  {
	return &UserRepoImp{db}
}