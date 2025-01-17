package ast

import (
	"bytes"

	"github.com/cloudcmds/tamarin/v2/token"
)

// Program represents a complete program.
type Program struct {
	// The set of statements which comprise the program.
	statements []Node
}

func NewProgram(statements []Node) *Program {
	return &Program{statements: statements}
}

func (p *Program) Token() token.Token {
	if len(p.statements) > 0 {
		return p.statements[0].Token()
	}
	return token.Token{}
}

func (p *Program) IsExpression() bool { return false }

func (p *Program) Literal() string { return p.Token().Literal }

func (p *Program) Statements() []Node { return p.statements }

func (p *Program) First() Node {
	if len(p.statements) > 0 {
		return p.statements[0]
	}
	return nil
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}
