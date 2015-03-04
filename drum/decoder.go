package drum

import "fmt"

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	//p := &Pattern{}
	p := readSplice(path)

	return p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	version string
	tempo   float32
	tracks  []struct {
		id    int
		name  string
		steps [16]byte
	}
}

func (p Pattern) String() string {
	output := fmt.Sprintf("Saved with HW Version: %v\n", p.version)
	output += fmt.Sprintf("Tempo: %v\n", p.tempo)
	return output
}
