package main

import (
	"fmt"
	"container/list"
)

type Parser struct {
	token_stream list.List
	parse_tree list.List
	token_index int
}

func (p *Parser) createParser(token_stream list.List) {
	p.token_stream = token_stream
	p.token_index = 0
}

func (p *Parser) parseStatement() {
	
}

func (p *Parser) startParsing() {
	for e := p.token_stream.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}