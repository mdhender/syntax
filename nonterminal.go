// Copyright (c) 2023 Michael D Henderson.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package syntax

import (
	"bytes"
)

// NonTerminal is a node that implements the syntax.Node interface.
type NonTerminal struct {
	line     int
	text     []byte
	children []Node
}

// Children implements the syntax.Node interface.
func (g *NonTerminal) Children() []Node {
	return g.children
}

// Clone implements the syntax.Node interface.
func (g *NonTerminal) Clone() Node {
	return &NonTerminal{line: g.line, text: g.text, children: g.children}
}

// Copy implements the syntax.Node interface.
func (g *NonTerminal) Copy() Node {
	cp := &NonTerminal{line: g.line}
	cp.text = append(cp.text, g.text...)
	for _, child := range g.children {
		cp.children = append(cp.children, child.Copy())
	}
	return cp
}

// Eq implements the syntax.Node interface.
func (g *NonTerminal) Eq(n Node) bool {
	if n == nil {
		return false
	} else if gn, ok := n.(*NonTerminal); !ok {
		return false
	} else if g.Kind() != gn.Kind() || !bytes.Equal(g.text, gn.text) {
		return false
	} else if len(g.children) != len(gn.children) {
		return false
	} else {
		for i, child := range g.children {
			if !child.Eq(gn.children[i]) {
				return false
			}
		}
	}
	return true
}

// Kind implements the syntax.Node interface.
func (g *NonTerminal) Kind() string {
	return "green-node"
}

// Line implements the syntax.Node interface.
func (g *NonTerminal) Line() int {
	return g.line
}

// Make implements the syntax.Node interface.
func (g *NonTerminal) Make(l int, text []byte, children ...Node) Node {
	gn := NonTerminal{line: g.line, text: text, children: children}
	return gn.Copy()
}

// PartialEq implements the syntax.Node interface.
func (g *NonTerminal) PartialEq(n Node) bool {
	if n == nil {
		return false
	} else if gn, ok := n.(*NonTerminal); !ok {
		return false
	} else {
		return g.Kind() == gn.Kind()
	}
}

// Text implements the syntax.Node interface.
func (g *NonTerminal) Text() []byte {
	return g.text
}
