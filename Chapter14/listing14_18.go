package main

type MyTypeQueue struct {
	q []MyType
}

func NewMyTypeQueue() *MyTypeQueue {
	return &MyTypeQueue{
		q: []MyType{},
	}
}
func (o *MyTypeQueue) Insert(v MyType) {
	o.q = append(o.q, v)
}
func (o *MyTypeQueue) Remove() MyType {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
