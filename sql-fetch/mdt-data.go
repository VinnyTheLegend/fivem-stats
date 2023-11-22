package sqlFetch

import (
	"fmt"
)

type MDTData struct {
	PfpURL string
}

type Pfp struct {
	Cid string
	Url string
}

func fetchPfps(characters []Character) ([]Character, error) {


    rows, err := db.Query("SELECT cid, pfp FROM mdt_data")
    if err != nil {
        return nil, fmt.Errorf("allCharacters: %q", err)
    }
	defer rows.Close()

	for rows.Next() {
		var pfp Pfp
        if err := rows.Scan(&pfp.Cid, &pfp.Url); err != nil {
            return characters, fmt.Errorf("mdt-pfp: %q", err)
        }

        for i, character := range characters {
			if character.CitizenID == pfp.Cid {
				characters[i].MDTData.PfpURL = pfp.Url
			}
		}
    }
	return characters, err
}