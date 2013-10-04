# MotorPiTX power control

A program to listen for MotorPiTX power button presses and cleanly shutdown your Raspberry Pi.

## Building

	GOOS=linux GOARM=6 GOARCH=arm go build

then copy MotorPiTX_power_button_control to the Raspberry Pi.

## Installing

Instructions for Arch linux.

Create a a file

    /etc/systemd/power_control.service

With contents:

````
[Unit]
Description=MotorPiTX power button control

[Service]
ExecStart=/root/MotorPiTX_power_button_control

[Install]
WantedBy=multi-user.target
````

Now run:

    systemctl start power_control.service

Check the process is running with ps / systemctl, and you should now be able to power down the Raspberry Pi by momentarily pressing the MotorPiTX power button.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## TODO

* Listen for a button press rather than polling

## Author

* Will Jessop, @will_j, will@willj.net