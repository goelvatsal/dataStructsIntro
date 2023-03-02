package main

type MyArray struct {
	length int
	data   []string
}

func (a *MyArray) push(item string) {
	a.data = append(a.data, item)
	a.length++
}

func (a *MyArray) pop() {
	//lastItem := a.data[a.length-1]
	a.data = a.data[:len(a.data)-1]
	a.length--
}

func (a MyArray) deleteAtIndex(index int) string {
	item := a.data[index]
	a.shiftItems(index)
	return item
}

func (a MyArray) shiftItems(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.data = a.data[:a.length-1]
	a.length--
}
