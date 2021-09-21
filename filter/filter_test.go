package filter

import (
	"testing"
)

func TestFilter_ReadFilter(t *testing.T) {
	filter := Filter{}
	filter.ReadFilter("filter")
	filter.IsLegal("a.123")
}