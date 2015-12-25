package dynaml

import (
	"github.com/cloudfoundry-incubator/spiff/yaml"
//	"github.com/cloudfoundry-incubator/spiff/debug"
)

type MergeExpr struct {
	Path []string
}

func (e MergeExpr) Evaluate(binding Binding) (yaml.Node, bool) {
	return binding.FindInStubs(e.Path)
}

func (e MergeExpr) String() string {
	return "merge"
}
