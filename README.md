# low_battery_alert
little tool to pop a message when your laptop battery is low

### Compil

- go get "github.com/distatus/battery"
- go get "github.com/mqu/go-notify"
- make

### Usage

- put the program in ~/.config/i3/script or where you whant
- add this line in your i3 config file `exec /path/to/i3_battery_alert`

You will be alerted when battery reach 20, 10 or 5 percent

![2018-11-22-004300_1366x768_scrot](https://user-images.githubusercontent.com/9384676/48859082-b7b72580-edef-11e8-9c27-4d2b59f5cff8.png)
