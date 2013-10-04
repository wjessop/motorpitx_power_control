package main

import (
	"github.com/mrmorphic/hwio"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	switch_pin_string string        = "GPIO8"
	sleep             time.Duration = 500 * time.Millisecond
)

func main() {
	switch_pin, pin_err := hwio.GetPinWithMode(switch_pin_string, hwio.INPUT_PULLDOWN)

	if pin_err != nil {
		log.Fatal(pin_err)
	}

	for {
		val, err := hwio.DigitalRead(switch_pin)
		if err != nil {
			log.Fatal(err)
		}

		if val == 1 {
			log.Println("Shutting down the system as due to power button press")
			cmd := exec.Command("systemctl", "poweroff")
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		time.Sleep(sleep)
	}
}
