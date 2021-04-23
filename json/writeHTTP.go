package json

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

var Json = jsoniter.Config {
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	TagKey:                 "json",
}.Froze()

func WriteHTTP(w http.ResponseWriter, data interface{}) error {
	encoder := Default.NewEncoder(w)
	encoder.SetIndent("", "    ")

	return encoder.Encode(data)
}