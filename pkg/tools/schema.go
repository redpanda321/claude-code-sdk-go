package tools

import "github.com/invopop/jsonschema"

// SchemaFor returns a JSON Schema derived from the Go struct type T.
// Uses invopop/jsonschema's reflector with default options. Tools should
// call this in their InputSchema() implementation rather than rolling
// their own reflection.
func SchemaFor[T any]() *jsonschema.Schema {
	var zero T
	return jsonschema.Reflect(&zero)
}
