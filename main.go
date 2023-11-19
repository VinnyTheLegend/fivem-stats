package main

import (
	sqlFetch "fivem-stats/sql-fetch"
	"fmt"
	"time"

)

//var charactersByFirstName []sqlFetch.Character
//var charactersByLastName []sqlFetch.Character
var charactersByBank []sqlFetch.Character
func updateData() {
	for {
		fmt.Println("upating data")
		_, _, charactersByBank = sqlFetch.FetchSortedCharacters()
		time.Sleep(1 * time.Hour)
	}
	
}

func main() {
	fmt.Println("starting data updater")
	go updateData()
	fmt.Println("starting http router")
	startRouter()
}
