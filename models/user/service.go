package user

import (
	"fmt"
	"go/api/db"
)

func GetAllUsersService(id int64) (users []User, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Sprintf("Error on connect to database")
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM user`)

	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Password)

		if err != nil {
			fmt.Sprint(err)
		}

		users = append(users, user)
	}

	return
}

func PostUserService(user User) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Sprintf("Error on connect to database")
	}
	defer conn.Close()

	sql := "INSERT INTO user (user, password, done) VALUES ($1, $2, $3) RETURNING id"

	err = conn.QueryRow(sql, user.Username, user.Password, user.Done).Scan(&id)

	return
}

func UpdateUserService(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET username=$2, passowrd=$3, done=$4 WHERE id=$1`, user.ID, user.Username, user.Password, user.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteUserService(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE id=$1`, user.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
