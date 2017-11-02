package user

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) CreateUser(db *sql.DB) error {

	statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES ('%s', '%d')", u.Name, u.Age)
	_, err := db.Exec(statement)

	if err != nil {
		return nil
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age='%d', WHERE id='%d'", u.Name, u.Age, u.ID)
	_, err := db.Exec(statement)
	return err
}

func GetUser(db *sql.DB, id int) (User, error) {
	statement := fmt.Sprintf("SELECT is FROM user where id = %d", id)
	row, err := db.Query(statement)

	if err != nil {
		return umodel.User{}, err
	}

	defer row.Close()
	user := umodel.User{}

	for row.Next() {
		var u umodel.User
		if err := row.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return umodel.User{}, err
		}
		user = u
	}

	return user, nil
}

func (u *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
}

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
	statement := fmt.Sprintf("SELECT id,name,age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *User) DeleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM user WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}
