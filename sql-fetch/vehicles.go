package sqlFetch

import (
	"fmt"
)

type Vehicle struct {
	CitizenID string
	Model string
	Plate string
	Garage *string
	Fuel int
	Engine float64
	Body float64
	State int
	Location *string
}

func allVehicles() ([]Vehicle, error) {
	var vehicles []Vehicle
	rows, err := db.Query("SELECT citizenid, vehicle, plate, garage, fuel, engine, body, state, location FROM player_vehicles")
	if err != nil {
		return nil, fmt.Errorf("vehicleFetch: %q", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.s
	for rows.Next() {
		var vehicle Vehicle
		if err := rows.Scan(&vehicle.CitizenID, &vehicle.Model, &vehicle.Plate, &vehicle.Garage, &vehicle.Fuel, &vehicle.Engine, &vehicle.Body, &vehicle.State, &vehicle.Location); err != nil {
			return nil, fmt.Errorf("vehicleFetch: %q", err)
		}

		vehicles = append(vehicles, vehicle)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("vehicleFetch: %q", err)
	}



	return vehicles, nil
}