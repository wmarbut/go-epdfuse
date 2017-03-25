# epdfuse [![Godoc Reference](https://img.shields.io/badge/Godoc-Reference-blue.svg)](https://godoc.org/github.com/wmarbut/go-epdfuse)
**epdfuse** is a library for interacting with the [PaPiRus epaper display](https://www.adafruit.com/products/3335) using golang and the [repaper epd fuse native library](https://github.com/repaper/gratis).

This library provides the ability to write arbitrary text and images to the display as well as clear the display.

It leverages the [goxbm](https://github.com/wmarbut/goxbm) project to convert images to the XBM format used by EPD Fuse.

## Install and setup

### Install Go

Go runs easily on the Raspberry PI. This stackexchange answer provides a quick overview: [http://raspberrypi.stackexchange.com/a/46828/167](http://raspberrypi.stackexchange.com/a/46828/167).

If cross compiling on a different computer, use these settings `GOARM=5 GOOS=linux GOARCH=arm`.

    GOARM=5 GOOS=linux GOARCH=arm go build

### Install Fuse Driver

The repaper fuse driver is required to use this golang library.

    # Install fuse driver
    sudo apt-get install libfuse-dev -y

    mkdir /tmp/papirus
    cd /tmp/papirus
    git clone https://github.com/repaper/gratis.git

    cd /tmp/papirus/gratis-master/PlatformWithOS
    make rpi-epd_fuse
    sudo make rpi-install
    sudo systemctl start epd-fuse.service

### Get the go-epdfuse project

Download and install the project with go

    go get github.com/wmarbut/go-epdfuse

Run an example

    cd $GOPATH/github.com/wmarbut/go-epdfuse/examples/papirus-text
    go build
    ./paprius-text "Hello World!"


