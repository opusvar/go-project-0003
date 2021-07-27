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

	leftSafely := m.Intn(5)

	if isHeistOn{
		switch leftSafely {
		    case 0:
				isHeistOn = false
                f.Print("You got caught during the escape.  Heist failed")
			case 1:
				isHeistOn = false
				f.Print("You packed the bags too full and couldn't carry them!")
				f.Print("You had to leave the goods to make your escape.")
				f.Print("Heist failed.")
			case 2:
				isHeistOn = false
				f.Print("Pre-heist celebrations come back to haunt you.")
				f.Print("Should have hydrated better. You passed out from the strain of packing the bags.")
				f.Print("Heist failed.")
			case 3:
				isHeistOn = false
				f.Print("Cheers and success! You've escaped the property.")
				f.Print("Cheap straps on the bag broke, you had to drop the load.")
				f.Print("Heist failed.")
				f.Print("Still got to make that getaway, but this time from the boss.")
			case 4:
				f.Print("You made it out with the loot.  Congratulations.")
				f.Print("Make it to the getaway car.")
			case 5:
				f.Print("Loot in tow, no guards in sight.")
				f.Print("Heading to the getaway car.")
			default:
				f.Print("Start the getaway car!")




		}
	}
	

	f.Println("Heist Status:", isHeistOn)





}