package evaluator

import (
	"context"

	"github.com/cloudcmds/tamarin/ast"
	"github.com/cloudcmds/tamarin/object"
	"github.com/cloudcmds/tamarin/scope"
)

func (e *Evaluator) evalHashLiteral(
	ctx context.Context,
	node *ast.HashLiteral,
	s *scope.Scope,
) object.Object {
	m := object.NewMap(make(map[string]interface{}, len(node.Pairs)))
	for keyNode, valueNode := range node.Pairs {
		key := e.Evaluate(ctx, keyNode, s)
		if isError(key) {
			return key
		}
		keyStr, err := object.AsString(key)
		if err != nil {
			return err
		}
		value := e.Evaluate(ctx, valueNode, s)
		if isError(value) {
			return value
		}
		m.Set(keyStr, value)
	}
	return m
}