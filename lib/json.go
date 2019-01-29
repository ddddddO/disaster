package lib

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

func Unmarshal(r io.Reader, apiStruct interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(b, apiStruct); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
