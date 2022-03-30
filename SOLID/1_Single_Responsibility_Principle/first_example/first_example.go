package first_example

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
In this example, Journal is a struct that holds a slice of entries.
Entries can be added or removed
*/

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// Here, the Save functionality is not a part of the Journal struct.
// Persistence functionality can be part of a whole different struct or package.
// This persistence construct could be used to save other types of data as well.
type Persistence struct {
	LineSeparator string
}

func (p *Persistence) Save(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

// func (p *Persistence) Load(filename string) { To be implemented...}
// func (p *Persistence) LoadFromWeb(url *url.URL) { To be implemented...}

func main() {
	j := Journal{}
	j.AddEntry("I laughed today")
	j.AddEntry("I ate pasta")

	p := Persistence{LineSeparator: "\n"}
	p.Save(&j, "journal.txt")
}
