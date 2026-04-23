package tools

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSchemaFor_StructProducesObjectSchema(t *testing.T) {
	// echoIn is defined in tool_test.go with `Message string jsonschema:"required"`.
	s := SchemaFor[echoIn]()
	if s == nil {
		t.Fatal("SchemaFor returned nil")
	}
	// invopop/jsonschema returns a schema where the top-level points at a
	// definition; we marshal-roundtrip and look for the expected substrings.
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	out := string(b)
	for _, want := range []string{`"message"`, `"required"`, `"type":"string"`} {
		if !strings.Contains(out, want) {
			t.Errorf("schema JSON missing %q\n%s", want, out)
		}
	}
}
