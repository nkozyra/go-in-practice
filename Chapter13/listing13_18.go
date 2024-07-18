package main

type MyTypeQueue struct { // # A
	q []MyType // # A
} // # A
func NewMyTypeQueue() *MyTypeQueue {
	return &MyTypeQueue{
		q: []MyType{},
	}
}
func (o *MyTypeQueue) Insert(v MyType) { // # B
	o.q = append(o.q, v)
}
func (o *MyTypeQueue) Remove() MyType { // # C
	if len(o.q) == 0 {
		panic("Oops.") // # D
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
