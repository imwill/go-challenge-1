// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	"io/ioutil"
	"math"
	//"regexp"
)

func readSplice(path string, p *Pattern) error {
	offset := 55
	splice, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	p.version = string(bytes.Trim(splice[14:33], "\x00"))
	p.tempo = findTempo(splice[46:51])
	findTracks(splice, &offset, p)

	return nil
}

func findTempo(spliceData []byte) float32 {
	bits := binary.LittleEndian.Uint32(spliceData)
	float := math.Float32frombits(bits)
	return float
}

func findTracks(splice []byte, offset *int, p *Pattern) {
	if *offset > len(splice) {
		return
	}

	name := ""
	track := Track{}
	track.id = int(splice[*offset-5])

	if string(splice[*offset-5:*offset+1]) == "SPLICE" {
		return
	}

	for {
		if bytes.Equal([]byte{splice[*offset]}, []byte{0x00}) {
			break
		} else if bytes.Equal([]byte{splice[*offset]}, []byte{0x01}) {
			break
		} else {
			name += string(splice[*offset])
			*offset++
		}
	}

	if *offset+16 > len(splice) {
		return
	}

	for i := 0; i < 16; i++ {
		if bytes.Equal([]byte{splice[*offset]}, []byte{0x00}) {
			//fmt.Printf("00 Name: %v, Offset: %d\n", name, *offset)
			track.steps[i] = byte(0)
		} else if bytes.Equal([]byte{splice[*offset]}, []byte{0x01}) {
			//fmt.Printf("01 Name: %v, Offset: %d\n", name, *offset)
			track.steps[i] = byte(1)
		}
		*offset++
	}

	track.name = name
	p.tracks = append(p.tracks, track)

	*offset = *offset + 5
	findTracks(splice, offset, p)
}
