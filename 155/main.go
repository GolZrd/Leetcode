package main

func main() {

}

type MinStack struct {
	items []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		items: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
