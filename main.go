package main

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
)

// check sum override address 0x00000691 1681
// check  override 0x71,0x4E,0x71,0x4E

func main() {

	romBytes, _ := ioutil.ReadFile("C:\\Users\\brend\\Downloads\\NBA Live 95 (USA, Europe).md")

	// Populate the model
	modelFile, _ := ioutil.ReadFile("model.json")
	var model Model
	json.Unmarshal(modelFile, &model)

	for i := 0; i < len(model.TeamDataAddresses); i++ {

		curTeamAddress := model.TeamDataAddresses[i]

		for j := 0; j < len(model.TeamData); j++ {

			curTeamData := model.TeamData[j]

			if curTeamData.DataType == "byte" {
				curTeamData.ByteValue = romBytes[curTeamAddress+curTeamData.Offset]
				println(curTeamData.Name)
				println(curTeamData.ByteValue)
			}

			if curTeamData.DataType == "stringFixed" {
				println(curTeamData.Name)
				println(readText(curTeamAddress+curTeamData.Offset, romBytes))
			}

			if curTeamData.DataType == "stringVariable" {
				println(curTeamData.Name)
				curTeamData.ReferenceAddress = binary.BigEndian.Uint32(romBytes[curTeamAddress+curTeamData.Offset : curTeamAddress+curTeamData.Offset+4])
				println(readText(curTeamData.ReferenceAddress, romBytes))
			}
		}

		model.TeamData[0].Name = "Test"

		// write data

		println()
	}

	//byteValue, _ := ioutil.ReadFile("romAddresses.json")
	// var addresses Addresses
	// json.Unmarshal(byteValue, &addresses)

	// var teams [30]Team

	// for i := 0; i < len(addresses.TeamStats); i++ {
	// 	teams[i] = *NewTeam(addresses.TeamMenuStart+(uint32)(i*4), addresses.TeamStats[i], byteArray)
	// }

	// teams[0].scoring = 1
	// teams[0].initials = "test"
	// teams[0].name = "tbirds"
	// teams[0].Write(byteArray)

	// // write back to the file, update check sum
	// byteArray[(int)(0x00000691)] = 0x71
	// byteArray[(int)(0x00000692)] = 0x4E
	// byteArray[(int)(0x00000693)] = 0x71
	// byteArray[(int)(0x00000694)] = 0x4E

	//	ioutil.WriteFile("C:\\Users\\brend\\Downloads\\NBA Live 95 (USA, Europe) - modified.md", byteArray, 0644)

}

func readText(startingAddress uint32, rom []byte) string {
	var text string
	for i := 0; ; i++ {
		character := rom[int(startingAddress)+i]
		if character == 0 {
			return text
		}
		text += string(character)
	}
}
