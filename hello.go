package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {

	byteArray, _ := ioutil.ReadFile("F:\\NBA Live 95 (USA, Europe).md")

	// ioutil.WriteFile("F:\\NBA Live 95 (USA, Europe) - Modified.md", byteArray, 0644)

	byteValue, _ := ioutil.ReadFile("romAddresses.json")
	var addresses Addresses
	json.Unmarshal(byteValue, &addresses)

	var teams [33]Team

	for i := 0; i < len(addresses.TeamStats); i++ {
		teams[i] = *NewTeam(addresses.TeamMenuStart+(uint32)(i*4), addresses.TeamStats[i], byteArray)
	}

}
