package v3

var setApply = [numOperators]operator{
	innerJoinOp: innerJoinApplyOp,
	leftJoinOp:  leftJoinApplyOp,
	rightJoinOp: rightJoinApplyOp,
	fullJoinOp:  fullJoinApplyOp,
	semiJoinOp:  semiJoinApplyOp,
	antiJoinOp:  antiJoinApplyOp,
}

var clearApply = [numOperators]operator{
	innerJoinApplyOp: innerJoinOp,
	leftJoinApplyOp:  leftJoinOp,
	rightJoinApplyOp: rightJoinOp,
	fullJoinApplyOp:  fullJoinOp,
	semiJoinApplyOp:  semiJoinOp,
	antiJoinApplyOp:  antiJoinOp,
}

var hasApply = [numOperators]bool{
	innerJoinApplyOp: true,
	leftJoinApplyOp:  true,
	rightJoinApplyOp: true,
	fullJoinApplyOp:  true,
	semiJoinApplyOp:  true,
	antiJoinApplyOp:  true,
}

func init() {
	for i := range setApply {
		if setApply[i] == unknownOp {
			setApply[i] = operator(i)
		}
	}
	for i := range clearApply {
		if clearApply[i] == unknownOp {
			clearApply[i] = operator(i)
		}
	}

	// TODO(peter): using the join operator here is a temporary hack.
	registerOperator(innerJoinApplyOp, "inner join (apply)", join{})
	registerOperator(leftJoinApplyOp, "left join (apply)", join{})
	registerOperator(rightJoinApplyOp, "right join (apply)", join{})
	registerOperator(fullJoinApplyOp, "full join (apply)", join{})
	registerOperator(semiJoinApplyOp, "semi-join (apply)", join{})
	registerOperator(antiJoinApplyOp, "anti-join (apply)", join{})
}
