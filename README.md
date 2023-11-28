# i3blocks-go-example

This Go code is designed to be used with i3blocks, a status bar for the i3 window manager. It provides information regarding memory and battery status by leveragin Linux command line utility. 

## Features

- Battery
- Memory

## Prerequisites

Make sure you have the following installed on your system:

- i3
- i3blocks
- Associated linux command line utility, eg. acpi

## Installation

1. Clone the repository:

    ```bash
    git clone git@github.com:boardwallfloor/i3blocks-component.git
    ```

2. Build the Go code you want:
    Example :
    ```bash
    cd i3blocks-component/battery
    go build
    ./i3_battery -d=power
    96%
    ```

3. (Optional)Move the binary to a directory in your `$PATH`:

    ```bash
    sudo mv i3_battery /usr/local/bin/
    ```

## Usage

Add the following configuration to your i3blocks configuration file (usually located at `~/.config/i3blocks/config`, if it doesn't exist simply create it yourself):

```ini
[battery-power]
command=/data/Project/i3block/battery/i3_battery -d=power
interval=5
```
