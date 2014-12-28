// Package flac is a small wrapper around github.com/eaburns/flac
//
// This package intents to make the API meet the io.Reader interface.
package flac

import (
	"io"

	"github.com/eaburns/flac"
)

func NewDecoder(r io.Reader) (*Decoder, error) {
	d, err := flac.NewDecoder(r)
	if err != nil {
		return nil, err
	}

	return &Decoder{MetaData: &d.MetaData, d: d}, nil
}

// Decoder wraps eaburns/flac.Decoder to follow the io.Reader
// interface.
type Decoder struct {
	// Metadata for the FLAC file
	*flac.MetaData
	// Decoder for FLAC format
	d *flac.Decoder
	// leftover is the bytes we could not return in a single Read
	leftover []byte
}

func (d *Decoder) Read(p []byte) (n int, err error) {
	var b []byte
	for {
		if len(d.leftover) > 0 {
			b = d.leftover
		} else {
			b, err = d.d.Next()
			if err != nil {
				return n, err
			}
		}

		nn := copy(p, b)
		n += nn

		if nn < len(b) {
			d.leftover = b[nn:]
			return n, nil
		}

		d.leftover = nil
		p = p[nn:]
	}
}

func (d *Decoder) Close() error {
	return nil
}
