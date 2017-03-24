# epdfuse [![Godoc Reference](https://img.shields.io/badge/Godoc-Reference-blue.svg)](https://godoc.org/github.com/wmarbut/go-epdfuse)
**epdfuse** is a library for interacting with the [PaPiRus epaper display](https://www.adafruit.com/products/3335) using golang and the [repaper epd fuse native library](https://github.com/repaper/gratis).

This library provides the ability to write arbitrary text and images to the display as well as clear the display.

It leverages the [goxbm](https://github.com/wmarbut/goxbm) project to convert images to the XBM format used by EPD Fuse.

