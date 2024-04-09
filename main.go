package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

func New(dir string, options *Options) (*Driver, error) {

	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (Database already exists)\n", dir)

		return &driver, nil
	}

	opts.Logger.Debug("Creating database at '%s'....\n", dir)

	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
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
	finalPath := filepath.Join(dir, resource+".json")
	tmpPath := finalPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")

	if err != nil {
		return err
	}

	b = append(b, byte('\n'))

	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, finalPath)
}

func (d *Driver) Read(collection, resource string, v interface{}) error {

	if collection == "" {
		return fmt.Errorf("missing collection - unable to read")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to read record - No name")
	}

	record := filepath.Join(d.dir, collection, resource)

	if _, err := stat(record); err != nil {
		return err
	}

	b, err := ioutil.ReadFile(record + ".json")

	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {

	if collection == "" {
		return nil, fmt.Errorf("missing collection - unable to read")
	}

	dir := filepath.Join(d.dir, collection)

	if _, err := stat(dir); err != nil {
		return nil, err
	}

	files, _ := ioutil.ReadDir(dir)

	var records []string

	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		records = append(records, string(b))
	}
	return records, nil
}

func (d *Driver) Delete(collection, resource string) error {

	path := path.Join(collection, resource)
	mutex := d.getORcreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("Unable to find file", path)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}

	return nil
}

func (d *Driver) getORcreateMutex(collection string) *sync.Mutex {

	d.mutex.Lock()
	defer d.mutex.Unlock()

	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}

type Address struct {
	City    string
	State   string
	Country string
	PinCode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {

	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "24", "9987654", "oracle", Address{"banglore", "karnataka", "India", "342342"}},
		{"Abhay", "24", "9983454", "Unplaced", Address{"new york", "US", "US", "342342"}},
		{"Parth", "24", "9934654", "oracle", Address{"banglore", "karnataka", "India", "342342"}},
		{"Ishaan", "24", "9347654", "un", Address{"chandigarh", "chandigarh", "India", "342342"}},
		{"Arpit", "24", "99876454", "unv", Address{"banglore", "karnataka", "India", "342342"}},
		{"Aman", "24", "9987654", "beyond", Address{"chandigarh", "chandigarh", "India", "342342"}},
		{"Sagar", "24", "9934654", "octro", Address{"banglore", "karnataka", "India", "342342"}},
	}

	for _, employee := range employees {
		db.Write("users", employee.Name, User{
			Name:    employee.Name,
			Age:     employee.Age,
			Contact: employee.Contact,
			Company: employee.Company,
			Address: employee.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Records", records)

	allUsers := []User{}

	for _, record := range records {

		employeeFound := User{}

		if err := json.Unmarshal([]byte(record), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}

		allUsers = append(allUsers, employeeFound)

	}

	fmt.Println("AllUsers", allUsers)

	// if err := db.Delete("users", "Abhay"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	if err := db.Delete("users", ""); err != nil {
		fmt.Println("Error", err)
	}
}
