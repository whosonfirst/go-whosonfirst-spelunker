package command

import (
	"fmt"
	"io"

	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

func WriteSPRResults(rsp spr.StandardPlacesResults, wr io.Writer) error {

	for _, r := range rsp.Results() {
		fmt.Fprintf(wr, "%s %s %s %s %0.6f %0.6f\n", r.Id(), r.Name(), r.Country(), r.Placetype(), r.Latitude(), r.Longitude())
	}

	return nil
}
