package json

import (
	"errors"
	"github.com/json-iterator/go"
)

var DB = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	TagKey:                 "db",
}.Froze()

var Default = jsoniter.ConfigCompatibleWithStandardLibrary

func Unmarshal(src interface{}, out interface{}) error {
	data, ok := src.([]byte)
	if !ok {
		return errors.New("invalid type")
	}

	return DB.Unmarshal(data, out)
}