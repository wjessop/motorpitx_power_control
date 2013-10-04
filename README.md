# MotorPiTX power control

A program to listen for MotorPiTX power button presses and cleanly shutdown your Raspberry Pi.

## Building

	GOOS=linux GOARM=6 GOARCH=arm go build

then copy motorpitx_power_control to the Raspberry Pi.

## Installing

Instructions for Arch linux.

Create a a file

    /etc/systemd/power_control.service

With contents:

````
[Unit]
Description=MotorPiTX power button control

[Service]
ExecStart=/root/motorpitx_power_control

[Install]
WantedBy=multi-user.target
````

Now run:

    systemctl enable power_control
    systemctl start power_control.service

Check the process is running with ps / systemctl, and you should now be able to power down the Raspberry Pi by momentarily pressing the MotorPiTX power button.

You can test that the program is running OK by passing in the -pretend flag:

    ./motorpitx_power_control -pretend true

The program will run as normal, flashing the status LED if the button is pressed, but will not shut down the Pi. Note that the state of the power button will remain high until a total power off/on cycle, so if you start the program with -pretend true then press the power button, then subsequently start the program with no -pretend value set your system will shut down!

Installing the binary into a more correct location is left as an exercise for the user.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## TODO

* Listen for a button press rather than polling
* Raspbian/Ubuntu/Other distro install instructions in the README (I'm not going to do this myself)

## Author

* Will Jessop, @will_j, will@willj.net
