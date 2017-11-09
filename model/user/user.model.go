package user

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	umodel "tracker/proto/user"

	gorp "gopkg.in/gorp.v2"
)

type UserModelInterface struct {
	ObjectModelInterface interface{}
}

type User struct {
}

func Push(db *sql.DB, u *umodel.User) error {

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.AddTableWithName(umodel.User{}, "users")
	fmt.Println("U is ", u)
	err := dbmap.Insert(u)
	//statement := fmt.Sprintf("INSERT INTO users(Name, Age) VALUES  ('%s', '%d')", u.Name, u.Age)
	//_, err := db.Exec(statement)

	if err != nil {
		fmt.Println("Failed to insert")
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

	var uname, loc, email, phone, fb, linkedin, site sql.NullString

	statement := fmt.Sprintf("SELECT * FROM users WHERE %s=%s", field, val)
	var user umodel.User

	if err := db.QueryRow(statement).Scan(&user.UserId, &user.Name, &user.Age, &uname, &email, &loc, &phone, &fb, &linkedin, &site); err != nil {
		return user, err
	}

	addIfNotNil(user.Location, loc)
	addIfNotNil(user.Username, uname)
	addIfNotNil(user.UserId, user.UserId)
	addIfNotNil(user.Username, uname)
	addIfNotNil(user.Email, email)
	addIfNotNil(user.Phonenumber, phone)
	addIfNotNil(user.Facebook, fb)
	addIfNotNil(user.Linkedin, linkedin)
	addIfNotNil(user.Site, site)
	return user, nil
}

func addIfNotNil(place, vv interface{}) {
	scan := reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	uses := !reflect.PtrTo(reflect.TypeOf(vv)).Implements(scan)
	if uses {
		place = vv
	}
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
