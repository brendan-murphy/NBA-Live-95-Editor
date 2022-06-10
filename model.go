package main

type RomValue struct {
	Name             string
	DataType         string
	Offset           uint32
	Length           int
	ByteValue        byte
	StringValue      string
	ReferenceAddress uint32
}

type Model struct {
	TeamMenuAddressStart uint32
	TeamDataAddresses    []uint32
	TeamData             []RomValue
}
