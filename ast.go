package main

import (
	"container/list"
)

type BinaryExpression struct {
	left_hand *BinaryExpression
	operand byte
	right_hand *BinaryExpression
}

type Statement interface {
	
}

type VariableDecl struct {
	name string
	expr *BinaryExpression
}

type FunctionDecl struct {
	name string
	params list.List
	body *Statement
}

type FunctionCall struct {
	name string
	params list.List
}