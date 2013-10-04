package main

import (
	"flag"
	"github.com/mrmorphic/hwio"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	switch_pin_string string        = "GPIO8"
	out_pin_string    string        = "GPIO7"
	sleep             time.Duration = 500 * time.Millisecond
)

var pretend bool

func init() {
	flag.BoolVar(&pretend, "pretend", false, "Don't really shut down, just flash the LED on button press. Useful for debugging.")
	flag.Parse()
}

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
			done_chan := make(chan int)
			go func() {
				out_pin, err := hwio.GetPinWithMode(out_pin_string, hwio.OUTPUT)

				if err != nil {
					log.Fatal(err)
				}

				for i := 0; i < 10; i++ {
					hwio.DigitalWrite(out_pin, hwio.LOW)
					time.Sleep(150 * time.Millisecond)
					hwio.DigitalWrite(out_pin, hwio.HIGH)
					time.Sleep(150 * time.Millisecond)
				}
				done_chan <- 1
			}()

			log.Println("Shutting down the system due to power button press")
			if !pretend {
				cmd := exec.Command("systemctl", "poweroff")
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}

			<-done_chan
			os.Exit(0)
		}
		time.Sleep(sleep)
	}
}
