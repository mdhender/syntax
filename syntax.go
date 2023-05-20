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

// Node defines the implementation of a syntax node.
type Node interface {
	// Children returns a slice containing the children of the Node.
	Children() []Node

	// Clone returns a shallow copy of the Node.
	Clone() Node

	// Copy returns a deep copy of the Node.
	Copy() Node

	// Eq does not use Line number.
	Eq(n Node) bool

	// Line returns the line number the Node was created on.
	Line() int

	// Kind returns the type of Node.
	Kind() string

	// Make is a Node factory.
	Make(int, []byte, ...Node) Node

	// PartialEq does a comparison?
	PartialEq(n Node) bool

	// Text returns the text of the Node.
	Text() []byte
}
