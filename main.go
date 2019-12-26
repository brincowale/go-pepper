package main

import "./dealabs"

func main() {
	client := dealabs.New()
	client.GetHotDeals(nil)
}
