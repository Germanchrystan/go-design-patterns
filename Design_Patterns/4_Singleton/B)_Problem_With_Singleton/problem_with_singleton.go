package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

/*
	So why is it the singleton a bad idea?
	The problem is that a singleton quite often breaks the dependency inversion principle,
	which we talked about at the beginning of the course.

	Let's go back to the simpleton example from before
*/

type singletonDatabase struct {
	capitals map[string]int // Name of capital : Population
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

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

/*
	Now, imagine you now want to get the total population of several cities.

*/
func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
		//         DIP violation
	}
	return result
}

/*
	Seems like everything is right, doesn't it?
	The problem is that the test is dependent upon data from a real life database.
	In real life software engineering, you almost never test against a live database (DUH!).

	Database can change at any time, and those changes can affect the outcome of a test.

	There is also a performance consideration, because even though we just want to test this
		"ok := tp == (17500000 + 17400000)"
	we want to unit test this particular function. But since we are going into a database,
	the unit test transform into an integration test.
*/

func main() {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	// Testing using real life database
	// What would happen if someone updated the numbers on the database?
	ok := tp == (17500000 + 17400000)
	fmt.Println(ok)
}

/*
	One of the ideas of the dependency inversion principle is that instead of depending
	on concrete implementations, we want to depend upon abstractions, which typically
	implies depending upon an interface
*/
