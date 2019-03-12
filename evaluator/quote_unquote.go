package evaluator

import (
	"github.com/tyabu12/monkey/ast"
	"github.com/tyabu12/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
