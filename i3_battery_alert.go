package main

import (
	"fmt"
	"strconv"

	"github.com/distatus/battery"
	"github.com/mqu/go-notify"
)

type notifAlert struct {
	twentyPercent bool
	tenPercent    bool
	fivePercent   bool
}

func getBatteryLevel(battery *battery.Battery) uint8 {

	var batteryLevel uint8 = 0

	// calculate percentage
	batteryLevel = uint8((battery.Current / battery.Full) * 100)

	return batteryLevel
}

func resetVar(alert notifAlert, batteryLevel uint8) {

	if batteryLevel > 14 {
		alert.twentyPercent = false
	}
	if batteryLevel > 10 {
		alert.tenPercent = false
	}
	if batteryLevel > 5 {
		alert.fivePercent = false
	}
}

func setNotification(alert *notifAlert, batteryLevel uint8) (notif *notify.NotifyNotification) {

	if batteryLevel < 14 && alert.twentyPercent == false {
		notif = notify.NotificationNew("Low Battery !", "Battery level: "+strconv.Itoa(int(batteryLevel))+"%", "dialog-information")
		notif.SetUrgency(notify.NOTIFY_URGENCY_CRITICAL)
		alert.twentyPercent = true
	} else if batteryLevel < 10 && alert.tenPercent == false {
		notif = notify.NotificationNew("Low Battery !", "Battery level: "+strconv.Itoa(int(batteryLevel))+"%", "dialog-information")
		notif.SetUrgency(notify.NOTIFY_URGENCY_CRITICAL)
		alert.tenPercent = true
	} else if batteryLevel < 5 && alert.fivePercent == false {
		notif = notify.NotificationNew("Low Battery !", "Battery level: "+strconv.Itoa(int(batteryLevel))+"%", "dialog-information")
		notif.SetUrgency(notify.NOTIFY_URGENCY_CRITICAL)
		alert.fivePercent = true
	}

	return
}

func main() {

	var (
		notif *notify.NotifyNotification = nil
		alert notifAlert
	)

	notify.Init("Low Battery !")

	for {
		// Get battery info
		batt, err := battery.Get(0)
		if err != nil {
			fmt.Println("Could not get battery info!")
			return
		}

		if batt.State.String() == "Discharging" {
			// Get battery level in percent
			batteryLevel := getBatteryLevel(batt)
			if err != nil {
				return
			}

			// Reset var if battery level charge up to limit (20, 10 or 5)
			resetVar(alert, batteryLevel)

			// set a popup
			notif = setNotification(&alert, batteryLevel)

			// Show popup
			if notif != nil {
				notif.Show()
				notif = nil
			}

			//time.Sleep(60 * time.Second)
		}
	}
}
