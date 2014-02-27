package dynaml

import (
	"github.com/cloudfoundry-incubator/spiff/yaml"
)

type ReferenceExpr struct {
	Path []string
}

func (e ReferenceExpr) Evaluate(binding Binding) (yaml.Node, bool) {
	var step yaml.Node
	var ok bool

	fromRoot := e.Path[0] == ""

	for i := 0; i < len(e.Path); i++ {
		if fromRoot {
			step, ok = binding.FindFromRoot(e.Path[1:i+1])
		} else {
			step, ok = binding.FindReference(e.Path[:i+1])
		}

		if !ok {
			return nil, false
		}

		switch step.(type) {
		case Expression:
			return e, true
		}
	}

	return step, true
}
