package booking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "35.187.248.198"
	port     = 5432
	user     = "postgres"
	password = "d3v3l0p8015"
	dbname   = "CenturyNet"
)

var (
	a bool = true
	b bool = false
)

type Message struct {
	Message    string `json:"message"`
	Booking_ID int    `json:"Booking_ID,omitempty"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	drivers := AllUsers()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(drivers)
}

func GetAllBooking(w http.ResponseWriter, r *http.Request) {
	bookings, err := AllBooking()

	if err != nil {
		var message Message
		message.Message = "Task Failed !!"
		json.NewEncoder(w).Encode(message)
		return
	}

	w.Header().Set("Content-Type", "applicaton/json; charset=utf-8")
	json.NewEncoder(w).Encode(bookings)
}

func GetBooking(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := mux.Vars(r)
	id := data["id"]

	booking, err := Book(id)
	if err != nil {
		var message Message
		message.Message = fmt.Sprintf("No data was found with Booking_ID = %s", id)
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(booking)
}

func AddBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var message Message
	decoder := json.NewDecoder(r.Body)
	var body Booking
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err.Error())
	}

	layout := "2006-01-02"
	start, _ := time.Parse(layout, body.Start_Time)
	end, _ := time.Parse(layout, body.End_Time)

	if start.Before(time.Now()) {
		fmt.Fprintf(w, "Start date is invalid !! %s", body.Start_Time)
		return
	}

	if end.Before(start) {
		fmt.Fprintf(w, "Date is invalid !!")
		return
	}

	if end.After(time.Now()) {
		body.Finished = &b
	} else {
		body.Finished = &a
	}

	get_cars, err := GetCars(body.Car_ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	get_driver, err := GetDriver(*body.Driver_ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	get_customer, err := GetCustomers(body.Customer_ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	diffdays := end.Sub(start).Hours() / 24

	body.Total_Cost = int(diffdays * float64(get_cars.Rent_Price_Daily))
	total_driver_cost := int(diffdays * float64(*get_driver.Daily))
	body.Total_Driver_Cost = &total_driver_cost
	incentive := int(float32(body.Total_Cost) * 0.05)

	if get_customer.Membership_ID != nil {
		discount := int(*get_customer.Daily_Discount * float32(body.Total_Cost))
		body.Discount = &discount
	}

	if *body.BookType_ID == 1 {
		body.Total_Driver_Cost = nil
		body.Driver_ID = nil
		incentive = 0
		if get_cars.Stock > 0 {
			book_id, err := InsertBooking(body)
			if err != nil {
				message.Message = "Failed to insert data to Booking !!"
				json.NewEncoder(w).Encode(message)
				return
			}
			var new_stock int = int(get_cars.Stock - 1)
			AddDriverIncentive(book_id, int32(incentive))
			if !*body.Finished {
				UpdateCarStock(body.Car_ID, int(new_stock))
			}
			message.Message = "Successfully inserted new Booking data !!"
			message.Booking_ID = book_id
			json.NewEncoder(w).Encode(message)
		} else {
			message.Message = "Failed to insert data to Booking !!"
			json.NewEncoder(w).Encode(message)
			return
		}
	} else {
		var new_stock int = int(get_cars.Stock - 1)
		check := ErrDriverTaken(start, end, *body.Driver_ID, 0)
		if !check && get_cars.Stock > 0 {
			book_id, err := InsertBooking(body)
			if err != nil {
				log.Fatal(err.Error())
			}
			AddDriverIncentive(book_id, int32(incentive))
			if !*body.Finished {
				UpdateCarStock(body.Car_ID, int(new_stock))
			}
			message.Message = "Successfully inserted new Booking data !!"
			message.Booking_ID = book_id
			json.NewEncoder(w).Encode(message)
		} else {
			if check {
				message.Message = "Driver is Unavailable"
				json.NewEncoder(w).Encode(message)
			} else {
				message.Message = "Car is Unavailable"
				json.NewEncoder(w).Encode(message)
			}
		}
	}
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	book_id := data["id"]
	err := DelBooking(book_id)
	if err != nil {
		log.Fatal(err.Error())
	}
	var message Message
	message.Message = "Data Booking Successfully Deleted"
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(message)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := mux.Vars(r)
	book_id, _ := strconv.Atoi(data["id"])
	decoder := json.NewDecoder(r.Body)
	var body Booking
	var message Message
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err.Error())
	}

	layout := "2006-01-02"
	start, _ := time.Parse(layout, body.Start_Time)
	end, _ := time.Parse(layout, body.End_Time)

	if start.Before(time.Now()) {
		fmt.Fprintf(w, "Start date is invalid !! %s", body.Start_Time)
		return
	}

	if end.Before(start) {
		fmt.Fprintf(w, "Date is invalid !!")
		return
	}

	if end.After(time.Now()) {
		body.Finished = &b
	} else {
		body.Finished = &a
	}

	get_cars, err := GetCars(body.Car_ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	get_driver, err := GetDriver(*body.Driver_ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	get_customer, err := GetCustomers(body.Customer_ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	get_oldbook, err := Book(strconv.Itoa(book_id))
	if err != nil {
		log.Fatal(err.Error())
	}

	diffdays := end.Sub(start).Hours() / 24

	body.Total_Cost = int(diffdays * float64(get_cars.Rent_Price_Daily))
	total_driver_cost := int(diffdays * float64(*get_driver.Daily))
	body.Total_Driver_Cost = &total_driver_cost
	incentive := int(float32(body.Total_Cost) * 0.05)

	if get_customer.Membership_ID != nil {
		discount := int(*get_customer.Daily_Discount * float32(body.Total_Cost))
		body.Discount = &discount
	}

	if *body.BookType_ID == 1 {
		body.Total_Driver_Cost = nil
		body.Driver_ID = nil
		incentive = 0
		flag := 0
		stock := get_cars.Stock
		if get_oldbook.Car_ID == body.Car_ID && stock == 0 {
			flag++
			stock++
		}
		if stock > 0 {
			err := UpdateBook(body, strconv.Itoa(book_id))
			if err != nil {
				message.Message = "Failed to insert data to Booking !!"
				json.NewEncoder(w).Encode(message)
				return
			}
			var new_stock int
			UpdateDriverIncentive(book_id, int32(incentive))
			if !*body.Finished && *get_oldbook.Finished {
				if flag == 0 {
					new_stock = int(stock) - 1
				} else {
					new_stock = 0
				}
			} else if *body.Finished && !*get_oldbook.Finished {
				new_stock = int(stock) + 1
			} else {
				new_stock = int(get_cars.Stock)
			}
			UpdateCarStock(body.Car_ID, int(new_stock))
			if body.Car_ID != get_oldbook.Car_ID {
				get_cars, err = GetCars(get_oldbook.Car_ID)
				if err != nil {
					log.Fatal(err.Error())
				}
				new_stock = int(get_cars.Stock) + 1
				UpdateCarStock(get_oldbook.Car_ID, new_stock)
			}
			message.Message = "Successfully Updated Booking data !!"
			json.NewEncoder(w).Encode(message)
		} else {
			message.Message = "Failed to Update data to Booking !!"
			json.NewEncoder(w).Encode(message)
			return
		}
	} else {
		flag := 0
		stock := get_cars.Stock
		if get_oldbook.Car_ID == body.Car_ID && stock == 0 {
			flag++
			stock++
		}
		var new_stock int = int(stock - 1)
		check := ErrDriverTaken(start, end, *body.Driver_ID, book_id)
		if !check && stock > 0 {
			err = UpdateBook(body, strconv.Itoa(book_id))
			if err != nil {
				log.Fatal(err.Error())
			}
			UpdateDriverIncentive(book_id, int32(incentive))
			if !*body.Finished && *get_oldbook.Finished {
				if flag == 0 {
					new_stock = int(stock) - 1
				} else {
					new_stock = 0
				}
			} else if *body.Finished && !*get_oldbook.Finished {
				new_stock = int(stock) + 1
			} else {
				new_stock = int(get_cars.Stock)
			}
			UpdateCarStock(body.Car_ID, int(new_stock))
			if body.Car_ID != get_oldbook.Car_ID {
				get_cars, err = GetCars(get_oldbook.Car_ID)
				if err != nil {
					log.Fatal(err.Error())
				}
				new_stock = int(get_cars.Stock) + 1
				UpdateCarStock(get_oldbook.Car_ID, new_stock)
			}
			message.Message = "Successfully Updated Booking data !!"
			message.Booking_ID = book_id
			json.NewEncoder(w).Encode(message)
		} else {
			if check {
				message.Message = "Driver is Unavailable"
				json.NewEncoder(w).Encode(message)
			} else {
				message.Message = "Car is Unavailable"
				json.NewEncoder(w).Encode(message)
			}
		}
	}
}
