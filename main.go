package main

import (
	"fmt"
	"strconv"

	"github.com/distatus/battery"
	"github.com/mqu/go-notify"
)

func getBatteryLevel() (uint8, error) {

	var batteryLevel uint8 = 0

	battery, err := battery.Get(0)
	if err != nil {
		fmt.Println("Could not get battery info!")
		return batteryLevel, err
	}

	// calculate percentage
	batteryLevel = uint8((battery.Current / battery.Full) * 100)

	return batteryLevel, err
}

func main() {

	batteryLevel, err := getBatteryLevel()
	if err != nil {
		return
	}

	notify.Init("Low Battery !")

	if batteryLevel == 20 || batteryLevel == 10 || batteryLevel == 5 {
		notif := notify.NotificationNew("Low Battery !", "Battery level: "+strconv.Itoa(int(batteryLevel))+"%", "dialog-information")
		notif.SetUrgency(notify.NOTIFY_URGENCY_CRITICAL)
		notif.Show()
	}
}
