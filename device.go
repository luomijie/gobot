package gobot

import (
	"errors"
	"fmt"
	"log"
)

// JSONDevice is a JSON representation of a Gobot Device.
type JSONDevice struct {
	Name       string   `json:"name"`
	Driver     string   `json:"driver"`
	Connection string   `json:"connection"`
	Commands   []string `json:"commands"`
}

type Device Driver

type devices []Device

// Len returns devices length
func (d *devices) Len() int {
	return len(*d)
}

// Each calls `f` function each device
func (d *devices) Each(f func(Device)) {
	for _, device := range *d {
		f(device)
	}
}

// Start starts all the devices.
func (d *devices) Start() (errs []error) {
	log.Println("Starting devices...")
	for _, device := range *d {
		info := "Starting device " + device.Name()
		if device.Pin() != "" {
			info = info + " on pin " + device.Pin()
		}
		log.Println(info + "...")
		if errs = device.Start(); len(errs) > 0 {
			for i, err := range errs {
				errs[i] = errors.New(fmt.Sprintf("Device %q: %v", device.Name(), err))
			}
			return
		}
	}
	return
}

// Halt stop all the devices.
func (d *devices) Halt() (errs []error) {
	for _, device := range *d {
		if derrs := device.Halt(); len(derrs) > 0 {
			for i, err := range derrs {
				derrs[i] = errors.New(fmt.Sprintf("Device %q: %v", device.Name(), err))
			}
			errs = append(errs, derrs...)
		}
	}
	return
}
