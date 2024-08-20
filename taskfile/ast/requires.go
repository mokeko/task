package ast

import (
	"github.com/go-task/task/v3/internal/deepcopy"
	"github.com/go-task/task/v3/internal/slicesext"
)

// Requires represents a set of required variables necessary for a task to run
type Requires struct {
	Vars []string
}

func (r *Requires) DeepCopy() *Requires {
	if r == nil {
		return nil
	}

	return &Requires{
		Vars: deepcopy.Slice(r.Vars),
	}
}

func (r *Requires) Merge(other *Requires) {
	if other == nil {
		return
	}

	r.Vars = slicesext.UniqueJoin(r.Vars, other.Vars)
}
