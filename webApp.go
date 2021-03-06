//Author: Alan Niemiec

//Adapted from the Minimal Example on:
//https://go-macaron.com/
package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
