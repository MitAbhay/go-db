package main

import {

	"fmt"
	"os"
	"encoding/json"
}

type Address struct {
	City string
	State string
	Country string
	PinCode json.Number
}

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}

func main() {

	dir := "./"d

	db , err := New(dir , nil)
	if(err != nil) fmt.Println("Error",err)

	empolyees := []User{
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
		{"John" , "24" , "9987654" ,"oracle" , Address{"banglore" , "karnataka" , "India" , "342342"}},
	}


	for _, employee := range employees {
		db.Write("users" , employee.Name , User{
			Name : employee.Name,
			Age : employee.Age,
			Contact : employee.Contact,
			Company : employee.Company,
			Address : employee.Address,
		})
	}

	records , err := db.ReadALL("users")
	if(err != nil) fmt.Println("Error" , err)

	fmt.Println("Records" , records)


	allUsers := []User{}

	for _, record := range records{

		employeeFound := User{}
		

	}

	
}