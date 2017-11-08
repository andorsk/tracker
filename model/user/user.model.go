package user

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"tracker/proto/location"
	umodel "tracker/proto/user"
)

type UserModelInterface struct {
	ObjectModelInterface interface{}
}

func Push(db *sql.DB, u *umodel.User) error {

	statement := fmt.Sprintf("INSERT INTO users(Name, Age) VALUES ('%s', '%d')", u.Name, u.Age)
	_, err := db.Exec(statement)

	if err != nil {
		return nil
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.UserId)

	if err != nil {
		return err
	}
	return nil
}

func Update(db *sql.DB, u *umodel.User) error {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age='%d', WHERE id='%d'", u.Name, u.Age, u.UserId)
	_, err := db.Exec(statement)
	return err
}

func (u *UserModelInterface) Get(db *sql.DB, id int64) (umodel.User, error) {
	statement := fmt.Sprintf("SELECT * FROM users where UserId = %d", id)
	row, err := db.Query(statement)

	if err != nil {
		return umodel.User{}, err
	}

	defer row.Close()
	user := umodel.User{}

	for row.Next() {
		var u umodel.User
		if err := row.Scan(&u.UserId, &u.Name, &u.Age); err != nil {
			return umodel.User{}, err
		}
		user = u
	}

	return user, nil
}

func GetByUserId(db *sql.DB, val int64) (umodel.User, error) {
	strval := strconv.FormatInt(val, 10)
	return Get(db, "UserId", strval)
}

func Get(db *sql.DB, field, val string) (umodel.User, error) {

	var loc location.Location
	var uname, email, phone, fb, linkedin, site sql.NullString

	statement := fmt.Sprintf("SELECT * FROM users WHERE %s=%s", field, val)
	var user umodel.User
	fmt.Println("statement is ", statement)
	if err := db.QueryRow(statement).Scan(&user.UserId, &user.Name, &user.Age, &uname, &email, &loc, &phone, &fb, &linkedin, &site); err != nil {
		log.Panic("Failed to query with ", err, " statement was \n ", statement)
		return user, err
	}

	user.Username = uname
	user.Email = email
	user.Phonenumber = phone
	user.Facebook = fb
	user.Linkedin = linkedin
	user.Site = site

	return user, nil
}

func GetUsers(db *sql.DB, start, count int) ([]umodel.User, error) {
	statement := fmt.Sprintf("SELECT UserId,Name,Age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []umodel.User{}
	for rows.Next() {
		var u umodel.User
		if err := rows.Scan(&u.UserId, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *UserModelInterface) Delete(db *sql.DB, user *umodel.User) error {
	statement := fmt.Sprintf("DELETE FROM user WHERE id=%d", user.UserId)
	_, err := db.Exec(statement)
	return err
}
