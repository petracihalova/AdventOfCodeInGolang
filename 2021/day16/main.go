package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Puzzle 1
	binaryInput := GetPuzzleInputAsBinary()
	a := Packet{}
	a.Init(binaryInput)
	result := a.SumOfVersions(0)
	fmt.Println("Puzzle 1 =", result)

	//Puzzle 2
	result = int(a.Value)
	fmt.Println("Puzzle 2 =", result)
}

type Packet struct {
	Version 		 int64
	TypeId 			 int64
	LengthTypeId	 int64
	Length			 int64
	Value 			 int64
	BinaryInput 	 string
	SubPackets		 []Packet
	SubPacketsBinary string
	Rest			 string
}

func (p *Packet) Init(binaryInput string) {
	p.BinaryInput = binaryInput
	p.Version, _ = strconv.ParseInt(binaryInput[:3], 2, 64)
	p.TypeId, _ = strconv.ParseInt(binaryInput[3:6], 2, 64)

	if p.TypeId == 4 {
		p.Value, p.Rest = p.GetLiteralValue(binaryInput[6:])
		return
	}
	num, _ := strconv.Atoi(string(binaryInput[6]))
	p.LengthTypeId = int64(num)
	switch p.LengthTypeId {
	case 0:
		p.Length, _ = strconv.ParseInt(binaryInput[7:22], 2, 64)
		p.SubPacketsBinary = binaryInput[22 : 22 + p.Length]
		p.Rest = binaryInput[22 + p.Length : ]
		if !strings.Contains(p.Rest, "1") {
			p.Rest = ""
		}
		for {
			newPacket := Packet{}
			newPacket.Init(p.SubPacketsBinary)
			p.SubPackets = append(p.SubPackets, newPacket)
			p.SubPacketsBinary = newPacket.Rest
			if p.SubPacketsBinary == "" || !strings.Contains(p.SubPacketsBinary, "1") {
				break
			}
		}
	case 1:
		p.Length, _ = strconv.ParseInt(binaryInput[7:18], 2, 64)
		p.SubPacketsBinary = binaryInput[18:]

		for i := 0; i < int(p.Length); i++ {
			newPacket := Packet{}
			newPacket.Init(p.SubPacketsBinary)
			p.SubPackets = append(p.SubPackets, newPacket)
			p.SubPacketsBinary = newPacket.Rest
		}
		p.Rest = p.SubPacketsBinary


	default:
		panic("something wrong")
	}

	p.CalculateValue()


}

func (p *Packet) PrintPacket() {
	fmt.Printf("verion: %d, type ID: %d, length type ID: %d, value: %d, subpackets: %d \n", p.Version, p.TypeId, p.LengthTypeId, p.Value, len(p.SubPackets))
}

func (p *Packet) GetLiteralValue(input string) (int64, string) {
	var literalValueBinary string
	counter := len(input) / 5
	for i := 0; i < counter; i++ {
		temp := input[:5]
		input = input[5:]
		literalValueBinary += temp[1:]
		if string(temp[0]) == "0" {
			literalValue, _ := strconv.ParseInt(literalValueBinary, 2, 64)
			return literalValue, input
		}
	}
	return 0, ""
}

func (p *Packet) CalculateValue() {
	switch p.TypeId {
	case 0:
		for _, subPacket := range p.SubPackets {
			p.Value += subPacket.Value
		}
	case 1:
		p.Value = 1
		for _, subPacket := range p.SubPackets {
			p.Value *= subPacket.Value
		}
	case 2:
		p.Value = p.SubPackets[0].Value
		for _, subPacket := range p.SubPackets {
			if p.Value > subPacket.Value {
				p.Value = subPacket.Value
			}
		}
	case 3:
		p.Value = p.SubPackets[0].Value
		for _, subPacket := range p.SubPackets {
			if p.Value < subPacket.Value {
				p.Value = subPacket.Value
			}
		}
	case 5:
		if p.SubPackets[0].Value > p.SubPackets[1].Value {
			p.Value = 1
		}
	case 6:
		if p.SubPackets[0].Value < p.SubPackets[1].Value {
			p.Value = 1
		}
	case 7:
		if p.SubPackets[0].Value == p.SubPackets[1].Value {
			p.Value = 1
		}
	default:
		panic(fmt.Sprintf("unexpected type ID = %d", p.TypeId))
	}

}

func (p *Packet) SumOfVersions(result int) int {
	result += int(p.Version)
	for _, subPacket := range p.SubPackets {
		result = subPacket.SumOfVersions(result)
	}
	return result
}

func GetPuzzleInputAsBinary() string {
	pwd, _ := os.Getwd()
	dataBytes, _ := os.ReadFile(pwd + "/2021/day16/input.txt")
	var result string
	for _, item := range dataBytes {
		resultAsHex, _ := strconv.ParseUint(string(item), 16, 64)
		result += fmt.Sprintf("%04b", resultAsHex)
	}
	return result
}
