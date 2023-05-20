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

// Terminal implements the syntax.Node interface.
type Terminal struct {
	line int
	text []byte
}

// Children implements the syntax.Node interface.
func (t *Terminal) Children() []Node {
	return nil
}

// Clone implements the syntax.Node interface.
func (t *Terminal) Clone() Node {
	return &Terminal{line: t.Line(), text: t.text}
}

// Copy implements the syntax.Node interface.
func (t *Terminal) Copy() Node {
	cp := &Terminal{line: t.line}
	cp.text = append(cp.text, t.text...)
	return cp
}

// Eq implements the syntax.Node interface.
func (t *Terminal) Eq(n Node) bool {
	if n == nil || t.text == nil {
		return false
	} else if tok, ok := n.(*Terminal); !ok || tok.text == nil {
		return false
	} else if t.Kind() != tok.Kind() || !bytes.Equal(t.text, tok.text) {
		return false
	}
	return true
}

// Kind implements the syntax.Node interface.
func (t *Terminal) Kind() string {
	return "token"
}

// Line implements the syntax.Node interface.
func (t *Terminal) Line() int {
	return t.line
}

// Make implements the syntax.Node interface.
func (t *Terminal) Make(l int, text []byte, _ ...Node) Node {
	tok := Terminal{line: t.line, text: text}
	return tok.Copy()
}

// PartialEq implements the syntax.Node interface.
func (t *Terminal) PartialEq(n Node) bool {
	if n == nil {
		return false
	} else if tok, ok := n.(*Terminal); !ok {
		return false
	} else if t.Kind() != tok.Kind() {
		return false
	}
	return true
}

// Text implements the syntax.Node interface.
func (t *Terminal) Text() []byte {
	return t.text
}
