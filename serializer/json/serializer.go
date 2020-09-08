package json

import (
	"encoding/json"

	"github.com/lazhari/url-shortener/shortener"
	"github.com/pkg/errors"
)

// Redirect represent the redirect struct
type Redirect struct{}

// Decode decode the Redirect struct to json
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}

	return redirect, nil
}

// Encode encode the json
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
