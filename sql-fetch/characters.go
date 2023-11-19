package sqlFetch

import (
	"encoding/json"
	"fmt"
	"math"
)

type Character struct {
    CitizenID string
	CharInfo CharInfo
	Money Money
	Job Job
	Gang Gang
	Vehicles []Vehicle
}

type CharInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthday string `json:"birthdate"`
	Gender int `json:"gender"`
	PhoneNumber string `json:"phone"`
	Nationality string `json:"nationality"`
	Account string `json:"account"`
	Cid int `json:"cid"`
	Backstory string `json:"backstory"`
}

type Money struct {
	Cash float64 `json:"cash"`
	Bank float64 `json:"bank"`
	Crypto float64 `json:"crypto"`
	IGC float64 `json:"igc"`
}

type Grade struct {
	Name string `json:"name"`
	Level int `json:"level"`
}

type Job struct {
	Label string `json:"label"`
	Grade Grade `json:"grade"`
	Payment float64 `json:"payment"`
	Name string `json:"name"`
	Type string `json:"type"`
	IsBoss bool `json:"isboss"`
	OnDuty bool `json:"onduty"`
}

type Gang struct {
	Label string `json:"label"`
	Grade Grade `json:"grade"`
	Name string `json:"name"`
	IsBoss bool `json:"isboss"`
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func allCharacters() ([]Character, error) {

	
    var characters []Character

    rows, err := db.Query("SELECT money, citizenid, charinfo, job, gang FROM players")
    if err != nil {
        return nil, fmt.Errorf("allCharacters: %q", err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
		var character Character
		var money string
		var charinfo string
		var job string
		var gang string
        if err := rows.Scan(&money, &character.CitizenID, &charinfo, &job, &gang); err != nil {
            return nil, fmt.Errorf("allCharacters: %q", err)
        }
		err = json.Unmarshal([]byte(charinfo), &character.CharInfo)
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = json.Unmarshal([]byte(money), &character.Money)
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = json.Unmarshal([]byte(job), &character.Job)
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = json.Unmarshal([]byte(gang), &character.Gang)
		if err != nil {
			fmt.Println("Error:", err)
		}


		character.Money.Bank = toFixed(character.Money.Bank, 2)
		character.Money.Cash = toFixed(character.Money.Cash, 2)
		if character.Gang.Label == "No Gang Affiliaton" {
			character.Gang.Label = "None"
		}

        characters = append(characters, character)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("allCharacters: %q", err)
    }
    return characters, nil
}