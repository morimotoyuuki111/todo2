package model

import (
        "strconv"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)

// DBの情報
const username = "root"
const password = "password"
const schema = "go"

// 使用するカラムを定義
type User struct {
	id int
	name string
}

// 会員の一覧を取得するメソッド
func GetUserList()(map[int]string) {
	// DBとの接続
	db, err := sql.Open("mysql", username + ":" + password + "@tcp(127.0.0.1:3306)/" + schema)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// クエリを実行
	rows, err := db.Query("select id, name from user order by id")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// メソッドの呼び出し元に返す配列を定義
	list := make(map[int]string)

	// 一行ずつアクセスして配列に格納
	for rows.Next() {
		var user User

		err := rows.Scan(&user.id, &user.name)
		if err != nil {
            panic(err)
		}

		list[user.id] = user.name
	}

	// 無事に終了していることを確認
	err = rows.Err()
        if err != nil {
            panic(err)
        }

	return list
}

// 会員の情報を取得するメソッド
func GetUserData(id string)(map[string]string) {
	// DBとの接続
	db, err := sql.Open("mysql", username + ":" + password + "@tcp(127.0.0.1:3306)/" + schema)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// プリペアードステートメントを利用してクエリを実行
	rows, err := db.Query("select id, name from user where id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	data := make(map[string]string)

	// 一行ずつアクセスして配列に格納
	for rows.Next() {
		var user User

		err := rows.Scan(&user.id, &user.name)
		if err != nil {
            panic(err)
		}

		data["id"] = strconv.Itoa(user.id)
		data["name"] = user.name
	}

	// 正常にに終了していることを確認
	err = rows.Err()
        if err != nil {
            panic(err)
        }

	return data
}

// 会員の情報を更新するメソッド
func EditUserData(id string, name string) {
	// DBとの接続
	db, err := sql.Open("mysql", username + ":" + password + "@tcp(127.0.0.1:3306)/" + schema)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// クエリを定義
	update, err := db.Prepare("update user set name = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()

	// プリペアードステートメントを利用してクエリを実行
	_, err = update.Exec(name, id)
	if err != nil {
		panic(err.Error())
	}
}

// 新しく会員を追加するメソッド
func AddUserData(name string) {
	// DBとの接続
	db, err := sql.Open("mysql", username + ":" + password + "@tcp(127.0.0.1:3306)/" + schema)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// クエリを定義
	insert, err := db.Prepare("insert user(name) values(?)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	// プリペアードステートメントを利用してクエリを実行
	_, err = insert.Exec(name)
	if err != nil {
		panic(err.Error())
	}
}