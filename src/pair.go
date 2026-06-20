package Pair

import (
	"fmt"
	"reflect"

	Array "github.com/go-composites/array/src"
)

// Interface is the fixed two-element heterogeneous grouping: the natural
// key/value entry of Composition-Oriented Programming. A Pair never goes nil;
// the Null-Object variant honours the whole Interface.
type Interface interface {
	First() interface{}
	Second() interface{}
	Equal(Interface) bool
	ToArray() Array.Interface
	ToGoString() string
	IsNull() bool
}

type data struct {
	first  interface{}
	second interface{}
}

/*
New returns a Pair.Interface holding first and second.
*/
func New(first, second interface{}) Interface {
	return &data{
		first:  first,
		second: second,
	}
}

// First returns the first slot of the Pair.
func (d data) First() interface{} {
	return d.first
}

// Second returns the second slot of the Pair.
func (d data) Second() interface{} {
	return d.second
}

// Equal reports whether both slots of the receiver are deeply equal to those
// of the given Pair (compared with reflect.DeepEqual). A null Pair is never
// equal to anything.
func (d data) Equal(p Interface) bool {
	if p.IsNull() {
		return false
	}
	return reflect.DeepEqual(d.first, p.First()) &&
		reflect.DeepEqual(d.second, p.Second())
}

// ToArray returns a two-element go-composites Array holding [first, second].
func (d data) ToArray() Array.Interface {
	array := Array.New()
	array.Push(d.first)
	array.Push(d.second)
	return array
}

// ToGoString renders the Pair as "(first, second)".
func (d data) ToGoString() string {
	return fmt.Sprintf("(%v, %v)", d.first, d.second)
}

// IsNull reports that this is a real (non-null) Pair.
func (d data) IsNull() bool {
	return false
}

// null is the Null-Object variant of a Pair: a placeholder that honours the
// full Interface without ever being nil. Its slots are nil, it equals nothing,
// and it converts to the empty Array.
type null struct{}

/*
Null returns the Null-Object Pair.
*/
func Null() Interface {
	return &null{}
}

// First returns nil for the null Pair.
func (n null) First() interface{} { return nil }

// Second returns nil for the null Pair.
func (n null) Second() interface{} { return nil }

// Equal is always false: a null Pair equals nothing.
func (n null) Equal(p Interface) bool { return false }

// ToArray returns the empty Array for the null Pair.
func (n null) ToArray() Array.Interface { return Array.New() }

// ToGoString renders the null Pair as "(null)".
func (n null) ToGoString() string { return "(null)" }

// IsNull reports that this is the null Pair.
func (n null) IsNull() bool { return true }
