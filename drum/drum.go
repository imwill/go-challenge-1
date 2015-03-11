// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"math"
)

func readSplice(path string, pattern *Pattern) error {
	offset := 55
	splice, error1 := ioutil.ReadFile(path)
	error2 := checkHeader(&splice)

	if error1 != nil {
		return error1
	}
	if error2 != nil {
		return error2
	}

	pattern.version = string(bytes.Trim(splice[14:33], "\x00"))
	findTempo(splice, pattern)
	findTracks(splice, &offset, pattern)

	return nil
}

func checkHeader(splice *[]byte) error {
	header := []byte("SPLICE")
	count := bytes.Count(*splice, header)

	if count > 1 {
		if index := bytes.LastIndex(*splice, header); index > 0 {
			*splice = (*splice)[0:index]
		}
	} else {
		if bytes.Index(*splice, header) != 0 {
			return errors.New("No SPLICE header found in file! Wrong file format?")
		}
	}
	return nil
}

func findTempo(spliceData []byte, pattern *Pattern) {
	bits := binary.LittleEndian.Uint32(spliceData[46:51])
	float := math.Float32frombits(bits)
	pattern.tempo = float
}

func findTracks(splice []byte, offset *int, p *Pattern) {
	if len(splice) <= *offset {
		return
	}

	track := Track{}
	track.id = int(splice[*offset-5])

	for {
		if bytes.Equal([]byte{splice[*offset]}, []byte{0x00}) {
			break
		} else if bytes.Equal([]byte{splice[*offset]}, []byte{0x01}) {
			break
		} else {
			track.name += string(splice[*offset])
			*offset++
		}
	}

	if *offset+16 > len(splice) {
		return
	}

	for i := 0; i < 16; i++ {
		if bytes.Equal([]byte{splice[*offset]}, []byte{0x00}) {
			track.steps[i] = byte(0)
		} else if bytes.Equal([]byte{splice[*offset]}, []byte{0x01}) {
			track.steps[i] = byte(1)
		}
		*offset++
	}

	p.tracks = append(p.tracks, track)

	*offset = *offset + 5
	findTracks(splice, offset, p)
}
