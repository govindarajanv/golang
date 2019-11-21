package main

import "fmt"

type car struct {
	gas_pedal      uint16 //min 0 to max 65535
	brake_pedal    uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh  float64
}

func main() {
	a_car := car{gas_pedal: 22000,
		brake_pedal:    0,
		steering_wheel: 1000,
		top_speed_kmh:  225.0}

	// a_car := car{22000,0,1000,225.0}
	fmt.Println(a_car.gas_pedal)
}
