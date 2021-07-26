package main

import (
	f "fmt"
	m "math/rand"
	t "time"
)

func main() {
    m.Seed(t.Now().UnixNano())
	isHeistOn := true
    eludedGuards := m.Intn(100)

	if eludedGuards >= 50 {
		f.Println("Looks like you've managed to make it past the guards.  Good job, but remember, this is the first step.")
		f.Println("The heist continues.")
	} else {
		isHeistOn = false
		f.Println("Failed to elude the guards.  The heist fails. ")
		f.Println("Plan a better disguse next time?")
	}

	f.Println("Heist Status:", isHeistOn)






}