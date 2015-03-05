// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	//"regexp"
)

func readSplice(path string) *Pattern {
	var p Pattern
	splice, _ := ioutil.ReadFile(path)

	p.version = string(bytes.Trim(splice[14:33], "\x00"))
	p.tempo = getTempo(splice[46:51])
	findTracks(splice, 55, &p)

	return &p
}

func getTempo(spliceData []byte) float32 {
	bits := binary.LittleEndian.Uint32(spliceData)
	float := math.Float32frombits(bits)
	return float
}

func findTracks(splice []byte, offset int, p *Pattern) {
	//append(p.tracks, {name: "", id: "", []})
	name := ""

	fmt.Print(int(splice[offset-5]))

	for i := offset; i < len(splice); i++ {
		offset++
		if bytes.Equal([]byte{splice[i]}, []byte{0x00}) {
			//name += string(splice[i])
			break
		} else if bytes.Equal([]byte{splice[i]}, []byte{0x01}) {
			break
		} else {
			name += string(splice[i])
		}
	}

	//append(p.tracks, {name: "", id: "", []})

	fmt.Print(name + "\n")
	/*
		for i := 0; i < 16; i++ {
			if bytes.Equal([]byte{splice[offset]}, []byte{0x00}) {
				fmt.Print("-")
			} else if bytes.Equal([]byte{splice[offset]}, []byte{0x01}) {
				fmt.Print("x")
			}
			offset++
		}
	*/
	if offset+20 <= len(splice) {
		findTracks(splice, offset+20, p)
	}

}
