package sqlFetch

import (
	"fmt"
)

type Vehicle struct {
	CitizenID string
	Model string
}

func allVehicles() ([]Vehicle, error) {
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



		return vehicles, nil
}