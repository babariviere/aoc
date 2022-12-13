package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Cmp int

const (
	Lesser  Cmp = 0
	Eq          = 1
	Greater     = 2
)

type Ty int

const (
	Int  Ty = 0
	List Ty = 1
)

type PacketList []PacketData

func (p PacketList) Compare(o PacketList) Cmp {
	for i, pi := range p {
		if i >= len(o) {
			return Greater
		}
		oi := o[i]
		var cmp Cmp
		if pi.Ty() == List && oi.Ty() == List {
			cmp = pi.(PacketList).Compare(oi.(PacketList))
		} else if pi.Ty() == Int && oi.Ty() == Int {
			cmp = pi.(PacketInt).Compare(oi.(PacketInt))
		} else {
			var pl, ol PacketList
			if pi.Ty() == List {
				pl = pi.(PacketList)
			} else {
				pl = PacketList{pi}
			}

			if oi.Ty() == List {
				ol = oi.(PacketList)
			} else {
				ol = PacketList{oi}
			}

			cmp = pl.Compare(ol)
		}
		if cmp != Eq {
			return cmp
		}
	}
	if len(o) > len(p) {
		return Lesser
	}
	return Eq
}

func (p PacketList) Ty() Ty { return List }

func parseList(input string) (PacketList, int) {
	if input[0] != '[' {
		panic("invalid input")
	}
	idx := 1
	var data []PacketData
	for {
		c := input[idx]
		if c == ']' {
			idx += 1
			break
		}
		if c == '[' {
			list, innerIdx := parseList(input[idx:])
			idx += innerIdx
			data = append(data, list)
		} else {
			int, innerIdx := parseInt(input[idx:])
			idx += innerIdx
			data = append(data, int)
		}
		if input[idx] == ',' {
			idx += 1
		} else {
			idx += 1
			break
		}
	}
	return PacketList(data), idx
}

type PacketInt int

func (p PacketInt) Compare(o PacketInt) Cmp {
	if p == o {
		return Eq
	}
	if p < o {
		return Lesser
	}
	return Greater
}

func (p PacketInt) Ty() Ty { return Int }

func parseInt(input string) (PacketInt, int) {
	var len int
	for {
		c := input[len]
		if c < '0' || c > '9' {
			break
		}
		len += 1
	}
	int, _ := strconv.Atoi(input[:len])
	return PacketInt(int), len
}

type PacketData interface {
	Ty() Ty
}

func parsePacket(input string) PacketList {
	packet, _ := parseList(input)
	return packet
}

type PacketPair struct {
	a, b PacketList
}

func parsePacketPair(input string) PacketPair {
	lines := strings.Split(input, "\n")
	a := parsePacket(lines[0])
	b := parsePacket(lines[1])

	return PacketPair{
		a, b,
	}
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	packets := strings.Split(string(input), "\n\n")

	var packetsPairs []PacketPair
	for _, rawPair := range packets {
		pair := parsePacketPair(rawPair)
		packetsPairs = append(packetsPairs, pair)

	}

	var inOrder []int
	for idx, pair := range packetsPairs {
		fmt.Printf("%v\n", pair.a)
		fmt.Printf("%v\n", pair.b)
		cmp := pair.a.Compare(pair.b)
		if cmp == Lesser || cmp == Eq {
			fmt.Println("Pair", idx+1, "is in order")
			inOrder = append(inOrder, idx+1)
		} else {
			fmt.Println("Pair", idx+1, "is NOT in order")
		}
		fmt.Println()
	}

	var sum int
	for _, io := range inOrder {
		sum += io
	}
	fmt.Println("Result", sum)

	var unsortedPackets []PacketList
	packetA, packetB := parsePacket("[[2]]"), parsePacket("[[6]]")
	unsortedPackets = append(unsortedPackets, packetA, packetB)
	for _, pair := range packetsPairs {
		unsortedPackets = append(unsortedPackets, pair.a, pair.b)
	}

	sort.Slice(unsortedPackets, func(i, j int) bool {
		a := unsortedPackets[i]
		b := unsortedPackets[j]
		if a.Compare(b) == Lesser {
			return true
		}
		return false
	})

	var pai, pbi int
	for idx, packet := range unsortedPackets {
		if packet.Compare(packetA) == Eq {
			pai = idx + 1
			continue
		}
		if packet.Compare(packetB) == Eq {
			pbi = idx + 1
			break
		}
	}
	fmt.Println("Result part 2:", pai*pbi)
}
