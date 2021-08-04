package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('budi','Budi')"
	_, err := db.ExecContext(ctx,query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "Select id, name FROM CUSTOMER"
	rows, err := db.QueryContext(ctx,query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}


func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "Select id, name, email, balance, rating, birth_date, married, created_at FROM CUSTOMER"
	rows, err := db.QueryContext(ctx,query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id, name string
		var  email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("=============================")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("email: ", email.String)
		}else{
			fmt.Println("email: Null")
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthDate.Valid {
			fmt.Println("birthDate: ", birthDate.Time)
		}else{
			fmt.Println("birthDate: Null")

		}
		fmt.Println("married: ", married)
		fmt.Println("createdAt: ", createdAt)
	}
}


func TestSqlInjection(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin"
	// password := "salah"
	username := "admin'; #"
	password := "salah"

	query := "Select username FROM user where username ='"+username+"' AND password = '"+password+"' limit 1"
	rows, err := db.QueryContext(ctx,query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	}else{
		fmt.Println("Gagal Login")
	}
}


func TestSqlInjectionSafe(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"
	// username := "admin'; #"
	// password := "salah"

	query := "Select username FROM user where username = ? AND password = ? limit 1"
	rows, err := db.QueryContext(ctx,query,username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	}else{
		fmt.Println("Gagal Login")
	}
}


func TestExecSqlParameter(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "ekoNEW"
	password := "ekoNEW"

	query := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx,query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}


func TestAutoIncrement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "ekoNEW"
	comment := "testcomment"

	query := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result , err := db.ExecContext(ctx,query, email, comment)
	if err != nil {
		panic(err)
	}

	InsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", InsertId)
}

func TestPrepareStatement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement , err := db.PrepareContext(ctx,query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko"+ strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke-"+ strconv.Itoa(i)

		result, err :=statement.ExecContext(ctx, email,comment)
		if err != nil {
			panic(err)
		}

		id,err:= result.LastInsertId()
		if err!=nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx,err := db.Begin()
	if err!= nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES(?,?)"

	//do transaction here
	for i := 0; i < 10; i++ {
		email := "eko"+ strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke-"+ strconv.Itoa(i)

		result, err :=tx.ExecContext(ctx, query, email,comment)
		if err != nil {
			panic(err)
		}

		id,err:= result.LastInsertId()
		if err!=nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}

	err = tx.Rollback()
	if err!=nil {
		panic(err)
	}
}