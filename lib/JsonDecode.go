package lib

import (
	"encoding/json"
	"io"
)

func JsonDecode(in io.Reader, out any) error {
	if err := json.NewDecoder(in).Decode(out); err != nil {
		return err
	}
	return nil
}
