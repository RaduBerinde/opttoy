package v3

import (
	"bytes"
)

func init() {
	registerOperator(projectOp, "project", project{})
}

type project struct{}

func (project) kind() operatorKind {
	return relationalKind
}

func (project) format(e *expr, buf *bytes.Buffer, level int) {
	formatRelational(e, buf, level)
	formatExprs(buf, "projections", e.projections(), level)
	formatExprs(buf, "filters", e.filters(), level)
	formatExprs(buf, "inputs", e.inputs(), level)
}

func (p project) updateProps(e *expr) {
	e.inputVars = p.requiredInputVars(e)
	e.inputVars &^= (e.props.outputVars() | e.providedInputVars())
	for _, input := range e.inputs() {
		e.inputVars |= input.inputVars
	}

	// TODO(peter): update expr.props.
}

func (project) requiredInputVars(e *expr) bitmap {
	var v bitmap
	for _, filter := range e.filters() {
		v |= filter.inputVars
	}
	for _, project := range e.projections() {
		v |= project.inputVars
	}
	return v
}
