package main

import (
	"can-i-chargeeeeee/charger"
	"can-i-chargeeeeee/notify"
	"log"
	"time"
)

func main() {
	for !charger.FindCharger() {
		log.Println("No chargers currently available")
		time.Sleep(180 * time.Second) // 3 minutes
	}

	log.Println("Charger available!! HURRY!!!!")
	notify.Notify()
}
