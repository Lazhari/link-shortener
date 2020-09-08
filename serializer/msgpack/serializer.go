package msgpack

import (
	"github.com/lazhari/url-shortener/shortener"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

// Redirect represent the redirect structure
type Redirect struct{}

// Decode decodes the message pack
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode encodes a Redirect struct to message pack
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}

	return rawMsg, err
}
