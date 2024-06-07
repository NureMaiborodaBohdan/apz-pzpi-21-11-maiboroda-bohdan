package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type AdminMysql struct {
	db *sqlx.DB
}

func NewAdminMysql(db *sqlx.DB) *AdminMysql {
	return &AdminMysql{db: db}
}

func (r *AdminMysql) GetUserByID(userID int) (AlcoSafe.User, error) {
	var user AlcoSafe.User
	query := fmt.Sprintf("SELECT UserID, Username, Email, Role FROM User WHERE UserID=?")
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}

func (r *AdminMysql) CreateUser(user AlcoSafe.User) (int, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM User WHERE email=? OR username=?")
	var count int
	err := r.db.Get(&count, checkQuery, user.Email, user.Username)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, fmt.Errorf("user with this email or username already exists")
	}

	query := fmt.Sprintf("INSERT INTO User (Username, Email, Password, Role, Name, Surname, Patronymic, CompanyID, Sex) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Role, user.Name, user.Surname, user.Patronymic, user.CompanyID, user.Sex)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return 0, fmt.Errorf("user with this email or username already exists")
		}
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
func (r *AdminMysql) GetAllUsers() ([]AlcoSafe.User, error) {
	var users []AlcoSafe.User
	query := fmt.Sprintf("SELECT UserID, Username, Email, Role, Name, Surname, Patronymic, CompanyID, Sex FROM User")
	err := r.db.Select(&users, query)
	if err != nil {
		log.Printf("Error fetching all users: %s", err)
	}
	return users, err
}
func (r *AdminMysql) UpdateUser(UserID int, input AlcoSafe.UpdateUserInput) error {
	existingUser, err := r.GetUserByID(UserID)
	if err != nil {
		return err
	}

	if input.Email != nil {
		checkEmailQuery := "SELECT UserID FROM User WHERE Email=? AND UserID != ?"
		var otherUserID int
		err := r.db.Get(&otherUserID, checkEmailQuery, *input.Email, UserID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if otherUserID != 0 {
			return errors.New("user with this email already exists")
		}

		existingUser.Email = *input.Email
	}

	if input.Username != nil {
		checkUsernameQuery := "SELECT UserID FROM User WHERE Username=? AND UserID != ?"
		var otherUserID int
		err := r.db.Get(&otherUserID, checkUsernameQuery, *input.Username, UserID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if otherUserID != 0 {
			return errors.New("user with this username already exists")
		}

		existingUser.Username = *input.Username
	}

	if input.Role != nil {
		existingUser.Role = *input.Role
	}

	if input.Name != nil {
		existingUser.Name = *input.Name
	}

	if input.Surname != nil {
		existingUser.Surname = *input.Surname
	}

	if input.Patronymic != nil {
		existingUser.Patronymic = *input.Patronymic
	}

	if input.CompanyID != nil {
		existingUser.CompanyID = *input.CompanyID
	}

	if input.Sex != nil {
		existingUser.Sex = *input.Sex
	}

	query := fmt.Sprintf("UPDATE User SET Username=?, Email=?, Password=?, Role=?, Name=?, Surname=?, Patronymic=?, CompanyID=?, Sex=? WHERE UserID=?")
	_, err = r.db.Exec(query, existingUser.Username, existingUser.Email, existingUser.Password, existingUser.Role, existingUser.Name, existingUser.Surname, existingUser.Patronymic, existingUser.CompanyID, existingUser.Sex, UserID)
	if err != nil {
		return err
	}

	return nil
}
func (r *AdminMysql) Delete(UserID int) error {
	query := fmt.Sprintf("DELETE FROM User WHERE UserID = ? ")
	_, err := r.db.Exec(query, UserID)
	return err
}
