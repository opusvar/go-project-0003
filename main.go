package main

import (
	f "fmt"
	m "math/rand"
	t "time"
)

func main() {
    m.Seed(t.Now().UnixNano()) // determines the seed
	isHeistOn := true  // declares heist status
    eludedGuards := m.Intn(100) // determines if the robbers eludes the guards

	// logic to detmine if the robbers elude the guards
	if eludedGuards >= 50 {
		f.Print("Looks like you've managed to make it past the guards. \n")  
		f.Print("Good job, but remember, this is the first step. \n")
		f.Print("The heist continues.\n")
		f.Println("Still got that safe to crack... \n")
	} else {
		isHeistOn = false
		f.Println("Failed to elude the guards.  The heist fails. ")
		f.Println("Plan a better disguse next time?")
	}

    openedVault := m.Intn(100) // determines if the robbers open the vault

    if isHeistOn && openedVault >= 70 {
        f.Println("Vault Open, Grab and GO!")
		f.Println("The heist continues.")
		f.Println("Still got make the escape.. \n")
	} else if isHeistOn && openedVault < 70 {
        isHeistOn = false
		f.Println("Safecracker failed to open the vault.  You missed your shot!")
		f.Println("The boss doesn't tolerate failure, better go on vacation.")
	} 
	

	f.Println("Heist Status:", isHeistOn)





}