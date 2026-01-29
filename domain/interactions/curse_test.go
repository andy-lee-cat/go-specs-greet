package interactions_test

import (
	"testing"

	"github.com/andy-lee-cat/go-specs-greet/domain/interactions"
	"github.com/andy-lee-cat/go-specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
