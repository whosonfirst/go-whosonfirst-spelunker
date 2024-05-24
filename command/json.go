package command

import (
	"encoding/json"
	"io"
)

func WriteJSON(body interface{}, wr io.Writer) error {

	enc := json.NewEncoder(wr)
	return enc.Encode(body)
}
