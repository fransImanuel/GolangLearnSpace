package domain

import (
	"database/sql"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/errs"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct{
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll()([]Customer, *errs.AppError){
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, findAllSql)

	if err!=nil {
		logger.Error("Error While Querying customer table "+ err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err!=nil {
	// 	logger.Error("Error While Scanning customer table "+ err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected Database Error")
	// }
	return customers, nil
}

func(d CustomerRepositoryDb) ById(id string)(*Customer, *errs.AppError){
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
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

func NewCustomerRepositoryDb(client *sqlx.DB) CustomerRepositoryDb{
	return CustomerRepositoryDb{client}
}