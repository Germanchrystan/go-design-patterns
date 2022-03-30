package second_example

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

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

//func (j *Journal) RemoveEntry(index int) { To be implemented...}

// Breaking SRP with adding functions that deal with another concern.
/*
Separation of concerns: different concerns or different problems that the system solves have
to reside in different constructs, or whether attached to different structures or packages.
*/
// God Object is basically when you take every functionality into a single construct.

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// func (j *Journal) Load(filename string) { To be implemented...}
// func (j *Journal) LoadFromWeb(url *url.URL) { To be implemented...}

/*
What we are doing here is breaking the SRP, because the responsibility of the journal is to do with the management of the entries.
The responsibility of the journal is not to deal with persistence. Persistence can be handled by a separate component, whetherit is a separate package,
or a separate struct that has methods related to persistence.
*/

func main() {

}
