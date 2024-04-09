package main

import {

	"fmt"
	"os"
	"encoding/json"
	"sync"
	"io/ioutil"
	"path/filepath"
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

type Options struct{
	Logger
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi , err := os.Stats(path) ; os.IsNotExist(err){
		fi , err = os.Stat(path + ".json")
	}

	return 
}

func New(dir string , options *Options)(*Driver , error) {

	dir  = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		 	
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir := dir,
		mutexes: make(map[string]*sync.Mutex),
		log : opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (Database already exists)\n", dir)

		return &driver , nil
	}

	opts.Logger.Debug("Creating database at '%s'....\n", dir)


	return &driver , os.MkdirAll(dir,0755)
}



func (d *Driver) Write(collection , resource string , v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - No place to save record")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to save record - No name")
	}

	mutex := d.getORcreateMutex(collection)

	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir , resource + ".json")
	tmpPath := finalPath + ".tmp"


	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b , err := json.MarshalIndent(v , "", "\t")

	if err != nil {
		return err
	}

	b = append(b , byte('\n'))

	if err := ioutil.WriteFile(tmpPath, b , 0644); err != nil {
		return err
	}


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