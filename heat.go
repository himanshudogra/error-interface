package main

import "fmt"

type error interface {
	Error() string
}

type OverHeating float64

func (o OverHeating) Error() string {
	return fmt.Sprintf("OverHeating by %0.2f degrees", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return (OverHeating(excess))
	}
	return nil
}

func main() {
	var err error
	err = checkTemperature(402.3, 302.4)
	fmt.Println(err)
}
