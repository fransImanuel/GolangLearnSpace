package domain

import (
	"database/sql"
	"time"

	"github.com/fransimanuel/RestfulApiwithHexagonalArch/errs"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct{
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll()([]Customer, *errs.AppError){
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err!=nil {
		logger.Error("Error While Querying customer table "+ err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}

	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err!=nil {
		logger.Error("Error While Scanning customer table "+ err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}
	return customers, nil
}

func(d CustomerRepositoryDb) ById(id string)(*Customer, *errs.AppError){
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err!=nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer Not Found")
			return nil, errs.NewNotFoundError("Customer Not Found")
		}else{
			// log.Println("Error While Scanning customer table "+ err.Error())
			logger.Error("Unexpected Database Error")
			return nil, errs.NewUnexpectedError("Unexpected Database Error")
		}
	}

	return &c, nil
 
}

func NewCustomerRepositoryDb() CustomerRepositoryDb{
	client, err := sql.Open("mysql", "root@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}