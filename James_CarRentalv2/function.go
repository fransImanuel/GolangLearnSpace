package booking

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//Struct untuk mengambil data dari DB (harus sesuai ama DB, * untuk nullable variable hanya menyimpan alamat)
type Cars struct {
	Car_ID           int    `json:"Car_ID"`
	Name             string `json:"Name"`
	Rent_Price_Daily int32  `json:"Rent_Price_Daily"`
	Stock            int16  `json:"Stock"`
}

type Customers struct {
	Customer_ID    int      `json:"Customer_ID"`
	Name           string   `json:"Name"`
	NIK            string   `json:"NIK"`
	Phone_Number   string   `json:"Phone_Number"`
	Membership_ID  *int     `json:"Membership_ID"`
	Daily_Discount *float32 `json:"Daily_Discount,omitempty"`
}

type Driver struct {
	ID          int     `json:"Driver_ID"`
	Name        *string `json:"Name"`
	NIK         *string `json:"NIK"`
	PhoneNumber *string `json:"Phone_Number"`
	Daily       *int    `json:"Daily_Cost"`
}

type Booking struct {
	Booking_ID        int    `json:"Booking_ID"`
	Customer_ID       int    `json:"Customer_ID"`
	Car_ID            int    `json:"Car_ID"`
	Start_Time        string `json:"Start_Time"`
	End_Time          string `json:"End_Time"`
	Total_Cost        int    `json:"Total_Cost"`
	Finished          *bool  `json:"Finished"`
	Discount          *int   `json:"Discount"`
	BookType_ID       *int   `json:"BookType_ID"`
	Driver_ID         *int   `json:"Driver_ID"`
	Total_Driver_Cost *int   `json:"Total_Driver_Cost"`
}

func GetCars(car_id int) (x Cars, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
	select * from "Cars" where "Car_ID" = $1
	`

	err = db.QueryRow(queries, car_id).Scan(
		&x.Car_ID,
		&x.Name,
		&x.Rent_Price_Daily,
		&x.Stock)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return x, err
		} else {
			log.Fatal(err.Error())
		}
	}

	return x, err
}

func GetCustomers(cust_id int) (x Customers, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
	select cust.* , mem."Daily_Discount" from "Customers" cust
        left join
        "Membership" mem on mem."Membership_ID" = cust ."Membership_ID" 
        where "Customer_ID" = $1
	`

	err = db.QueryRow(queries, cust_id).Scan(
		&x.Customer_ID,
		&x.Name,
		&x.NIK,
		&x.Phone_Number,
		&x.Membership_ID,
		&x.Daily_Discount)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return x, err
		} else {
			log.Fatal(err.Error())
		}
	}

	return x, err
}

func GetDriver(driver_id int) (x Driver, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
	select * from "Driver" where "Driver_ID" = $1
	`

	err = db.QueryRow(queries, driver_id).Scan(
		&x.ID,
		&x.Name,
		&x.NIK,
		&x.PhoneNumber,
		&x.Daily)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return x, err
		} else {
			log.Fatal(err.Error())
		}
	}

	return x, err
}

func AllUsers() []Driver {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	var drivers []Driver

	queries := `
	select * from "Driver"
	`
	rows, err := db.Query(queries)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var driver Driver
		err := rows.Scan(&driver.ID, &driver.Name, &driver.NIK, &driver.PhoneNumber, &driver.Daily)

		if err != nil {
			panic(err.Error())
		}
		drivers = append(drivers, driver)
	}
	defer db.Close()

	return drivers
}

func AllBooking() (x []Booking, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
	select * from "Booking"
	`
	rows, err := db.Query(queries)

	for rows.Next() {
		var booking Booking
		err := rows.Scan(
			&booking.Booking_ID,
			&booking.Customer_ID,
			&booking.Car_ID,
			&booking.Start_Time,
			&booking.End_Time,
			&booking.Total_Cost,
			&booking.Finished,
			&booking.Discount,
			&booking.BookType_ID,
			&booking.Driver_ID,
			&booking.Total_Driver_Cost)
		if err != nil {
			panic(err.Error())
		}
		x = append(x, booking)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return x, err
}

func Book(id string) (x Booking, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
	select * from "Booking" where "Booking_ID" = $1
	`

	err = db.QueryRow(queries, id).Scan(
		&x.Booking_ID,
		&x.Customer_ID,
		&x.Car_ID,
		&x.Start_Time,
		&x.End_Time,
		&x.Total_Cost,
		&x.Finished,
		&x.Discount,
		&x.BookType_ID,
		&x.Driver_ID,
		&x.Total_Driver_Cost)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return x, err
		} else {
			log.Fatal(err.Error())
		}
	}

	return x, err
}

func InsertBooking(data Booking) (id int, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
		INSERT INTO "Booking"("Customer_ID", "Car_ID" ,"Start_Time", "End_Time", "Total_Cost", "Finished", "Discount", "BookType_ID", "Driver_ID", "Total_Driver_Cost")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING "Booking_ID";
	`

	err = db.QueryRow(queries,
		data.Customer_ID,
		data.Car_ID,
		data.Start_Time,
		data.End_Time,
		data.Total_Cost,
		data.Finished,
		data.Discount,
		data.BookType_ID,
		data.Driver_ID,
		data.Total_Driver_Cost).Scan(&id)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		} else {
			log.Fatal(err.Error())
		}
	}
	return id, err
}

func UpdateCarStock(car_id int, new_stock int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
		update "Cars"
		set "Stock" = $2
		where "Car_ID" = $1
	`

	rows, _ := db.Exec(queries, car_id, new_stock)

	defer db.Close()

	count, err := rows.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	if (err != nil) || count <= 0 {
		if err == sql.ErrNoRows {
			panic(err.Error())
		} else {
			log.Fatal(err.Error())
		}
	}
}

func AddDriverIncentive(book_id int, incentive int32) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
		INSERT INTO "Driver_Incentive"("Booking_ID", "Incentive")
		VALUES ($1, $2);
	`

	_, err = db.Exec(queries, book_id, incentive)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			panic(err.Error())
		} else {
			log.Fatal(err.Error())
		}
	}
}

func UpdateDriverIncentive(book_id int, incentive int32) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
		update "Driver_Incentive"
		set "Incentive" = $2
		where "Booking_ID" = $1
	`

	_, err = db.Exec(queries, book_id, incentive)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			panic(err.Error())
		} else {
			log.Fatal(err.Error())
		}
	}
}

func ErrDriverTaken(start time.Time, end time.Time, driver_id int, book_id int) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	var queries string
	if book_id != 0 {
		queries = `
		select * from "Booking" b where 
		($1 between "Start_Time" and "End_Time"
		or $2 between "Start_Time" and "End_Time"
		or "Start_Time" between $1 and $2
		or "End_Time" between $1 and $2) and "Driver_ID" = $3 and "Finished" = false and "Booking_ID" != $4
		`
		rows, err := db.Query(queries, start, end, driver_id, book_id)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		return rows.Next()
	}

	queries = `
	select * from "Booking" b where 
	($1 between "Start_Time" and "End_Time"
	or $2 between "Start_Time" and "End_Time"
	or "Start_Time" between $1 and $2
	or "End_Time" between $1 and $2) and "Driver_ID" = $3 and "Finished" = false
	`
	rows, err := db.Query(queries, start, end, driver_id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return rows.Next()
}

func DelBooking(book_id string) (err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
		delete from "Booking" 
		where "Booking_ID" = $1
	`

	_, err = db.Exec(queries, book_id)

	defer db.Close()

	return err
}

func UpdateBook(data Booking, book_id string) (err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	queries := `
		update "Booking" 
		set 
		"Customer_ID" = $1,
		"Car_ID" = $2,
		"Start_Time" = $3,
		"End_Time" = $4,
		"Total_Cost" = $5,
		"Finished" = $6,
		"Discount" = $7,
		"BookType_ID" = $8,
		"Driver_ID" = $9,
		"Total_Driver_Cost" = $10
		where "Booking_ID" = $11
	`

	_, err = db.Exec(queries,
		data.Customer_ID,
		data.Car_ID,
		data.Start_Time,
		data.End_Time,
		data.Total_Cost,
		data.Finished,
		data.Discount,
		data.BookType_ID,
		data.Driver_ID,
		data.Total_Driver_Cost,
		book_id)

	defer db.Close()

	return err
}
