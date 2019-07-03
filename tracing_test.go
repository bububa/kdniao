package kdniao

import (
	"fmt"
	"testing"
)

const (
	TEST_BUESINESS_ID = "1549289"
	TEST_API_KEY      = "5bd3c986-078e-4d91-8f55-7d5116783880"
)

func TestInstanceSearch(t *testing.T) {
	client := NewClient(TEST_BUESINESS_ID, TEST_API_KEY)
	resp, err := InstantSearch(client, "SF", "074463202889", "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%#v\n", resp)
}
