package main

import "encoding/binary"

type Team struct {
	menuAddress uint32
	mainAddress uint32

	scoring     byte
	rebounds    byte
	ballControl byte
	defense     byte
	overall     byte

	backgroundColor byte
	bannerColor     byte
	textColor       byte

	location      string
	name          string
	courtLocation string
	initials      string
}

func NewTeam(menuAddress uint32, mainAddress uint32, rom []byte) *Team {
	team := Team{menuAddress: menuAddress, mainAddress: mainAddress}

	team.scoring = rom[team.mainAddress+69]
	team.rebounds = rom[team.mainAddress+70]
	team.ballControl = rom[team.mainAddress+71]
	team.defense = rom[team.mainAddress+72]
	team.overall = rom[team.mainAddress+73]
	//74 or 0x41 is a mystery
	team.backgroundColor = rom[team.mainAddress+75]
	team.bannerColor = rom[team.mainAddress+76]
	team.textColor = rom[team.mainAddress+77]

	nameStartingAddress := binary.BigEndian.Uint32(rom[team.mainAddress+48 : team.mainAddress+48+4])
	team.name = readText(nameStartingAddress, rom)

	locationStartingAddress := binary.BigEndian.Uint32(rom[team.mainAddress+52 : team.mainAddress+52+4])
	team.location = readText(locationStartingAddress, rom)

	courtStartingAddress := binary.BigEndian.Uint32(rom[team.mainAddress+60 : team.mainAddress+60+4])
	team.courtLocation = readText(courtStartingAddress, rom)

	team.initials = readText(team.mainAddress+64, rom)

	return &team
}

func readText(startingAddress uint32, rom []byte) string {
	var text string
	for i := 0; ; i++ {
		character := rom[(int)(startingAddress)+i]
		if character == 0 {
			return text
		}
		text += string(character)
	}
}
