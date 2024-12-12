package generator

import (
	"strings"
	"testing"
)

func TestGeneratorWordlist(t *testing.T) {
	for ii := 1; ii < 100; ii++ {
		generated := GenerateUniqueWordListID(ii)
		sz := len(strings.Split(generated, "-"))
		if sz != ii {
			t.Errorf("%d != %d", sz, ii)
		}
	}
}
