package audec

import "bytes"

type Format string

type pre struct {
	prefix []byte
	format Format
}

var prefixTable = []pre{
	{[]byte("fLaC\x00\x00\x00"), FLAC},

	{[]byte{0x49, 0x44, 0x33}, MP3}, // ID3 tagged MP3
	// This is a list of possible combinations of the sync-header,
	// audio version, layer index and protection bit of a mp3 frame-header
	{[]byte{0xFF, 0xFB}, MP3}, // version 1, layer 3, protection on
	{[]byte{0xFF, 0xFA}, MP3}, // version 1, layer 3, protection off
	{[]byte{0xFF, 0xF3}, MP3}, // version 2, layer 3, protection on
	{[]byte{0xFF, 0xF2}, MP3}, // version 2, layer 3, protection off
	{[]byte{0xFF, 0xE3}, MP3}, // version 2.5, layer 3, protection on
	{[]byte{0xFF, 0xE2}, MP3}, // version 2.5, layer 3, protection off
	// TODO: support layer 1 and 2

	{[]byte("\x4F\x67\x67\x53\x00"), OGG},
}

const (
	UnsupportedFormat Format = ""
	MP3                      = "MP3"
	FLAC                     = "FLAC"
	OGG                      = "OGG"
)

// DetectFormat takes the header of a file and tries to determine what
// kind of audio format the file contains. The returned value is a typed
// string containing the format name.
//
// Returns UnsupportedFormat if the format is not known or unsupported
func DetectFormat(h []byte) Format {
	for _, p := range prefixTable {
		l := len(p.prefix)
		if len(h) > l && bytes.Equal(h[:l], p.prefix) {
			return p.format
		}
	}

	return UnsupportedFormat
}
