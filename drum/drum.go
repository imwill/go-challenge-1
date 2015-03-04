// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
)

func readSplice(path string) *Pattern {
	var p Pattern
	splice, _ := ioutil.ReadFile(path)

	p.version = string(bytes.Trim(splice[14:33], "\x00"))
	p.tempo = getTempo(splice[46:51])
	getTracks(splice, 55)

	return &p
}

func getTempo(spliceData []byte) float32 {
	bits := binary.LittleEndian.Uint32(spliceData)
	float := math.Float32frombits(bits)
	return float
}

func getTracks(splice []byte, offset int) {
	//append(p.tracks, {name: "", id: "", []})
	name := ""

	match, _ := regexp.Compile("[0-9A-Za-z-_ ]")
	for i := offset; i < len(splice); i++ {
		if match.MatchString(string(splice[i])) {
			name += string(splice[i])
		} else {
			fmt.Print(name)
			name = ""
		}
	}
}
