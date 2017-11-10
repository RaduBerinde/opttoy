package v3

import (
	"bytes"
)

func init() {
	registerOperator(renameOp, "rename", rename{})
}

func newRenameExpr(input *expr) *expr {
	return &expr{
		op:       renameOp,
		extra:    1,
		children: []*expr{input, nil /* filter */},
		props:    &logicalProps{},
	}
}

type rename struct{}

func (rename) kind() operatorKind {
	return relationalKind
}

func (rename) format(e *expr, buf *bytes.Buffer, level int) {
	formatRelational(e, buf, level)
	formatExprs(buf, "filters", e.filters(), level)
	formatExprs(buf, "inputs", e.inputs(), level)
}

func (rename) initKeys(e *expr, state *queryState) {
}

func (r rename) updateProps(e *expr) {
	// Rename is pass through and requires any input variables that its inputs
	// require.
	e.inputVars = 0
	for _, input := range e.inputs() {
		e.inputVars |= input.inputVars
	}

	e.props.applyFilters(e.filters())

	// TODO(peter): update keys
}

func (rename) requiredInputVars(e *expr) bitmap {
	return e.providedInputVars()
}

func (rename) equal(a, b *expr) bool {
	return true
}
