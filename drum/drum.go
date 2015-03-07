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

func readSplice(path string, p *Pattern) error {
	offset := 55
	splice, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	p.version = string(bytes.Trim(splice[14:33], "\x00"))
	p.tempo = getTempo(splice[46:51])
	findTracks(splice, &offset, p)

	return nil
}

func getTempo(spliceData []byte) float32 {
	bits := binary.LittleEndian.Uint32(spliceData)
	float := math.Float32frombits(bits)
	return float
}

func findTracks(splice []byte, offset *int, p *Pattern) {
	name := ""
	track := Track{}

	track.id = int(splice[*offset-5])

	fmt.Println(*offset)

	for i := *offset; i < len(splice); i++ {
		*offset++
		fmt.Println(*offset)
		if bytes.Equal([]byte{splice[i]}, []byte{0x00}) {
			//name += string(splice[i])
			break
		} else if bytes.Equal([]byte{splice[i]}, []byte{0x01}) {
			break
		} else {
			name += string(splice[i])
		}
	}

	track.name = name
	fmt.Println("step loop")
	for i := 0; i < 16; i++ {
		if bytes.Equal([]byte{splice[*offset]}, []byte{0x00}) {
			track.steps[i] = byte(0)
		} else if bytes.Equal([]byte{splice[*offset]}, []byte{0x01}) {
			track.steps[i] = byte(1)
		}
		*offset++
	}

	p.tracks = append(p.tracks, track)

	fmt.Printf("Name: %v; Offset: %d; Splice len: %v", name, *offset, len(splice))

	if *offset <= len(splice) {
		findTracks(splice, offset, p)
	}

}
