package user

import (
	"database/sql"
	"fmt"
	umodel "tracker/proto/user"
)

type UserModelInterface struct {
	ObjectModelInterface interface{}
}

func Push(db *sql.DB, u *umodel.User) error {

	statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES ('%s', '%d')", u.Name, u.Age)
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

func Get(db *sql.DB, id int64) (umodel.User, error) {
	statement := fmt.Sprintf("SELECT is FROM user where id = %d", id)
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

func (u *UserModelInterface) GetUser(db *sql.DB, id int) error {
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", id)
	var user umodel.User
	return db.QueryRow(statement).Scan(&user.UserId, &user.Name, &user.Age)
}

func GetUsers(db *sql.DB, start, count int) ([]umodel.User, error) {
	statement := fmt.Sprintf("SELECT id,name,age FROM users LIMIT %d OFFSET %d", count, start)
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
