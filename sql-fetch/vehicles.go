package sqlFetch

import (
	"fmt"
)

type Vehicle struct {
	CitizenID string
	Model string
}

func vehicleFetch(characters []Character) ([]Character, error) {
		var vehicles []Vehicle
		rows, err := db.Query("SELECT citizenid, vehicle FROM player_vehicles")
		if err != nil {
			return nil, fmt.Errorf("vehicleFetch: %q", err)
		}
		defer rows.Close()
		// Loop through rows, using Scan to assign column data to struct fields.s
		for rows.Next() {
			var vehicle Vehicle
			if err := rows.Scan(&vehicle.CitizenID, &vehicle.Model); err != nil {
				return nil, fmt.Errorf("vehicleFetch: %q", err)
			}
	
			vehicles = append(vehicles, vehicle)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("vehicleFetch: %q", err)
		}

		for c, character := range characters {
			for _, vehicle := range vehicles {
				if vehicle.CitizenID == character.CitizenID {
					characters[c].Vehicles = append(characters[c].Vehicles, vehicle)
				}
			}
		}

		return characters, nil
}