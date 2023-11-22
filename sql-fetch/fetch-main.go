package sqlFetch

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"
	"os"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	fmt.Printf("dbuser: %v\n", os.Getenv("DBUSER"))
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "insanitygaming.net",
		DBName: "s30_roleplay",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

type ByField func(p1, p2 *Character) bool

func (f ByField) Sort(characters []Character) []Character {
	ps := make([]Character, len(characters))
	copy(ps, characters)
	sort.Sort(&characterSorter{characters: ps, by: f})
	return ps
}

type characterSorter struct {
	characters []Character
	by      func(p1, p2 *Character) bool
}

func (s *characterSorter) Len() int           { return len(s.characters) }
func (s *characterSorter) Swap(i, j int)      { s.characters[i], s.characters[j] = s.characters[j], s.characters[i] }
func (s *characterSorter) Less(i, j int) bool { return s.by(&s.characters[i], &s.characters[j]) }

func FetchSortedCharacters() ([]Character, []Character, []Character) {	
	var characters, charerr = allCharacters()
	if charerr != nil {
		log.Fatal(charerr)
	}
	println("Characters Loaded:", len(characters))
	vehicles, veherr := allVehicles() 
	if veherr != nil {
		log.Fatal(veherr)
	}
	println("Vehicles Loaded:", len(vehicles))

	vnew := make([]Vehicle, len(vehicles))
	copy(vnew, vehicles)
	for c, character := range characters {
		var vtemp []Vehicle
		for _, vehicle := range vnew {
			newveh := vehicle
			if vehicle.CitizenID == character.CitizenID {
				characters[c].Vehicles = append(characters[c].Vehicles, newveh)
			} else {
				vtemp = append(vtemp, newveh)
			}
		}
		vnew = vtemp
	}


	charactersByBank := ByField(func(p1, p2 *Character) bool {
		return p1.Money.Bank < p2.Money.Bank
	}).Sort(characters)

	charactersByFirstName := ByField(func(p1, p2 *Character) bool {
		name1 := p1.CharInfo.FirstName
		name2 := p2.CharInfo.FirstName
		return strings.ToLower(name1) < strings.ToLower(name2)
	}).Sort(characters)

	charactersByLastName := ByField(func(p1, p2 *Character) bool {
		name1 := p1.CharInfo.LastName
		name2 := p2.CharInfo.LastName
		return strings.ToLower(name1) < strings.ToLower(name2)
	}).Sort(characters)

	return charactersByFirstName, charactersByLastName, charactersByBank
}