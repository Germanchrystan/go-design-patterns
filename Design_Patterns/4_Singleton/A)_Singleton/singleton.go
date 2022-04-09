package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int // Name of capital : Population
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

/*
	We need to make only one instance of the singletonDatabase struct available.
	We will declare this variable called "once", which is going to be a type sync.Once
	sync.Once is a struct that ensures something gets calleds only once

	Thread safety: We don't want two threads to start initializing the struct at the same time.
	- One option is using sync.Once.
	- The other is using the package level "init" function
*/

var once sync.Once

/*
	We are also declaring a variable instance of type singletonDatabase pointer.

	Laziness feature: It means that we only construct the database, or read it from a memory
	whenever someone asks for it, not before.
	This can not be guaranteed by the "init" package level function, but it can be done
	using things like sync.Once inside our own function


*/
var instance *singletonDatabase

/*
	Then we create a function that will return the only instance
*/
func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData(".\\capitals.txt")
		db := singletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul = ", pop)
}

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}
