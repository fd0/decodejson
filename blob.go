package decoder

import (
	"errors"
	"fmt"
)

// Blob is one part of a file or a tree.
type Blob struct {
	Type   BlobType
	Length uint
	ID     ID
	Offset uint
}

// BlobType specifies what a blob stored in a pack is.
type BlobType uint8

// These are the blob types that can be stored in a pack.
const (
	InvalidBlob BlobType = iota
	DataBlob
	TreeBlob
)

func (t BlobType) String() string {
	switch t {
	case DataBlob:
		return "data"
	case TreeBlob:
		return "tree"
	case InvalidBlob:
		return "invalid"
	}

	return fmt.Sprintf("<BlobType %d>", t)
}

// MarshalJSON encodes the BlobType into JSON.
func (t BlobType) MarshalJSON() ([]byte, error) {
	switch t {
	case DataBlob:
		return []byte(`"data"`), nil
	case TreeBlob:
		return []byte(`"tree"`), nil
	}

	return nil, errors.New("unknown blob type")
}

// UnmarshalJSON decodes the BlobType from JSON.
func (t *BlobType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case `"data"`:
		*t = DataBlob
	case `"tree"`:
		*t = TreeBlob
	default:
		return errors.New("unknown blob type")
	}

	return nil
}
