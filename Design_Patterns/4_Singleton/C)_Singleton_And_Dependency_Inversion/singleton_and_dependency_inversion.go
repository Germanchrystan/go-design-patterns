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
	We really want our tests to depend on some sort of abstraction.
	So, instead of depending on the concrete Singleton database,
	we can substitute this database with something else, and just
	provide a different implementation of our population.

	So, how can we make a dummy database?
	First of all, we have to introduce some sort of abstraction,
	which has something in commmon between the real db and the dummy db.
*/
type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
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

func GetTotalPopulation(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
		//         Using interface instead of singleton struct method
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	cities := []string{"alpha", "gamma"}
	tp := GetTotalPopulation(&DummyDatabase{}, cities)

	ok := tp == 4
	fmt.Println(ok)

	// if we want to use the actual singleton db to get total population we could write:
	// tp := GetTotalPopulation(GetSingletonDatabase(), citites)
}
