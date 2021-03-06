package testutil

import (
	"flag"
	"strings"
	"testing"
)

const (
	INTEGRATION_TEST_TYPE = "integration"
	UNIT_TEST_TYPE        = "unit"
	SSL_TEST_TYPE         = "ssl"
)

var (
	// the types of tests that should be run
	testTypes = flag.String("test.types", UNIT_TEST_TYPE, "Comma-separated list of the"+
		" types of tests to be run")
	// above, split on the comma
	testTypesParsed []string
)

func HasTestType(testType string) bool {
	// parse if necessary
	if testTypesParsed == nil {
		testTypesParsed = strings.Split(*testTypes, ",")
	}

	// skip the test if the passed-in type is not being run
	for _, typ := range testTypesParsed {
		if typ == testType {
			return true
		}
	}
	return false
}

// Skip the test if the specified type is not being run.
func VerifyTestType(t *testing.T, testType string) {
	if !HasTestType(testType) {
		t.SkipNow()
	}
}
