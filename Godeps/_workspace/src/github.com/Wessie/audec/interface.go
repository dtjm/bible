package audec

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var UnsupportedFormatError = fmt.Errorf("unsupported format")

// format to decoder mapping, can be added to with RegisterDecoder
var decoderReg = map[Format]DecoderFunc{}

type DecoderFunc func(io.Reader) (Decoder, error)

type Decoder interface {
	io.ReadCloser
}

func RegisterDecoder(form Format, fn DecoderFunc) {
	if _, ok := decoderReg[form]; ok {
		panic("double registration for decoder of format: " + string(form))
	}

	decoderReg[form] = fn
}

// GetDecoder returns the DecoderFunc associated with format f. Returns nil
// when no decoder is associated with f.
//
// Generally you want to be using OpenFile or FromReader instead of GetDecoder.
func GetDecoder(f Format) DecoderFunc {
	return decoderReg[f]
}

// OpenFile is a small helper to open the filename given and passing it to
// FromReader. Acts the same as FromReader with the exception of a possible
// error returned from os.Open being returned.
func OpenFile(filename string) (Decoder, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return FromReader(f)
}

func FromReader(r io.Reader) (Decoder, error) {
	r = bufio.NewReader(r)

	// We will have to read the first few bytes to test for the format we expect
	var buf bytes.Buffer
	io.CopyN(&buf, r, 4096)

	form := DetectFormat(buf.Bytes())
	if form == UnsupportedFormat {
		return nil, UnsupportedFormatError
	}

	// Now we have a format, and have to find a decoder for it
	fn := GetDecoder(form)
	if fn == nil {
		return nil, UnsupportedFormatError
	}

	// Fix the reader to be whole again, as if we never read from it
	r = io.MultiReader(&buf, r)

	return fn(r)
}
