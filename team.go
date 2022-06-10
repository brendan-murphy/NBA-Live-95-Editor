package main

// import "encoding/binary"

// const scoringOffset byte = 0x45
// const reboundsOffset byte = 0x46
// const ballControllOffset byte = 0x47
// const defenseOffset byte = 0x48
// const overallOffset byte = 0x49

// const backgroundColorOffset byte = 0x4B
// const bannerColorOffset byte = 0x4C
// const textColorOffset byte = 0x4D

// const initialsOffset byte = 0x40
// const courtnameOffset byte = 0x38
// const courtLocationOffset byte = 0x3C
// const locationOffset byte = 0x34
// const nameOffset byte = 0x30

// type Team struct {
// 	menuAddress uint32
// 	mainAddress uint32

// 	scoring     byte
// 	rebounds    byte
// 	ballControl byte
// 	defense     byte
// 	overall     byte

// 	backgroundColor byte
// 	bannerColor     byte
// 	textColor       byte

// 	location      string
// 	name          string
// 	courtLocation string
// 	initials      string

// 	locationLength      int
// 	nameLength          int
// 	courtLocationLength int
// }

// func NewTeam(menuAddress uint32, mainAddress uint32, rom []byte) *Team {
// 	team := Team{menuAddress: menuAddress, mainAddress: mainAddress}

// 	team.scoring = rom[team.mainAddress+uint32(scoringOffset)]
// 	team.rebounds = rom[team.mainAddress+uint32(reboundsOffset)]
// 	team.ballControl = rom[team.mainAddress+uint32(ballControllOffset)]
// 	team.defense = rom[team.mainAddress+uint32(defenseOffset)]
// 	team.overall = rom[team.mainAddress+uint32(overallOffset)]
// 	//74 or 0x4A is a mystery
// 	team.backgroundColor = rom[team.mainAddress+uint32(backgroundColorOffset)]
// 	team.bannerColor = rom[team.mainAddress+uint32(bannerColorOffset)]
// 	team.textColor = rom[team.mainAddress+uint32(textColorOffset)]

// 	nameStartingAddress := binary.BigEndian.Uint32(rom[team.mainAddress+uint32(nameOffset) : team.mainAddress+48+4])
// 	team.name = readText(nameStartingAddress, rom)
// 	team.nameLength = len(team.name)

// 	locationStartingAddress := binary.BigEndian.Uint32(rom[team.mainAddress+uint32(locationOffset) : team.mainAddress+uint32(locationOffset)+4])
// 	team.location = readText(locationStartingAddress, rom)
// 	team.locationLength = len(team.location)

// 	courtStartingAddress := binary.BigEndian.Uint32(rom[team.mainAddress+uint32(courtLocationOffset) : team.mainAddress+uint32(courtLocationOffset)+4])
// 	team.courtLocation = readText(courtStartingAddress, rom)
// 	team.courtLocationLength = len(team.courtLocation)

// 	team.initials = readText(team.mainAddress+uint32(initialsOffset), rom)

// 	return &team
// }

// func readText(startingAddress uint32, rom []byte) string {
// 	var text string
// 	for i := 0; ; i++ {
// 		character := rom[int(startingAddress)+i]
// 		if character == 0 {
// 			return text
// 		}
// 		text += string(character)
// 	}
// }

// func (t *Team) Write(rom []byte) {
// 	rom[t.mainAddress+uint32(scoringOffset)] = t.scoring
// 	rom[t.mainAddress+uint32(reboundsOffset)] = t.rebounds
// 	rom[t.mainAddress+uint32(ballControllOffset)] = t.ballControl
// 	rom[t.mainAddress+uint32(defenseOffset)] = t.defense
// 	rom[t.mainAddress+uint32(overallOffset)] = t.overall
// 	rom[t.mainAddress+uint32(backgroundColorOffset)] = t.backgroundColor
// 	rom[t.mainAddress+uint32(bannerColorOffset)] = t.bannerColor
// 	rom[t.mainAddress+uint32(textColorOffset)] = t.textColor

// 	initialAddressStart := t.mainAddress + uint32(initialsOffset)
// 	for i := 0; i < 4; i++ {
// 		rom[initialAddressStart+uint32(i)] = t.initials[i]
// 	}

// 	nameStartingAddress := binary.BigEndian.Uint32(rom[t.mainAddress+uint32(nameOffset) : t.mainAddress+48+4])
// 	for i := 0; i < t.nameLength; i++ {
// 		if i < len(t.name)-1 {
// 			rom[nameStartingAddress+uint32(i)] = t.name[i]
// 		} else {
// 			rom[nameStartingAddress+uint32(i)] = 0x00
// 		}
// 	}
// }
