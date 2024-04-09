package main

import {

	"fmt"
	"os"
	"encoding/json"
	"sync"
	"github.com/jcelliott/lumber"
}

const version = "1.0.0"

type{
	Logger interface{
		Fatal(string , ...interface{})
		Error(string , ...interface{})
		Warn(string , ...interface{})
		Info(string , ...interface{})
		Debug(string , ...interface{})
		Trace(string , ...interface{})

	}

	Driver struct{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
}

type Options{
	Logger
}

func New()(){

}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {
	
}

func (d *Driver) ReadAll() () {
	
}

func (d *Driver) Delete() error {
	
}

func (d *Driver) getORcreateMutex() *sync.Mutex {
	
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

		if err := json.Unmarshal([]byte(record), &employeeFound); err != nil {
			fmt.Println("Error" , err)
		}

		allUsers = append(allUsers, employeeFound)
		
	}

	fmt.Println("AllUsers" , allUsers)


	// if err := db.Delete("user" , "abhay") ; err != nil {
	// 	fmt.Println("Error" ,err)
	// }

	// if err := db.Delete("user" , ""); err != nil {
	// 	fmt.Println("Error" , err)
	// }



	
}