package file

import (
	"strings"
	"testing"
)

func BenchmarkParseInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(strings.NewReader(dbMiekENTNL), testzone, "stdin")
	}
}

func TestParseSerial(t *testing.T) {
	zone, err := Parse(strings.NewReader(dbMiekENTNL), testzone, "stdin")
	if err != nil {
		t.Fatalf("Failed to parse zone: %q: %s", testzone, err)
	}
	// In dbMiekENTNL the SOA serial is 1282630057.
	// If this serial is NOT that number we call this test a success. This test fails
	// once every 136 years...
	soa := zone.Apex.SOA
	if soa.Serial == 1282630057 {
		t.Fatalf("SOA serial not updated to the current Unix epoch: %d", soa.Serial)
	}
}
