package decoder

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// idSize contains the size of an ID, in bytes.
const idSize = sha256.Size

// ID references content within a repository.
type ID [idSize]byte

func (id ID) String() string {
	return hex.EncodeToString(id[:])
}

// MarshalJSON returns the JSON encoding of id.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in id.
func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	_, err = hex.Decode(id[:], []byte(s))
	if err != nil {
		return err
	}

	return nil
}
