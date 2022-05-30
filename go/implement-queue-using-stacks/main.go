package implementqueueusingstacks

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	length := len(this.data)

	if length == 0 {
		return -1
	}

	v := this.data[0]
	this.data = this.data[1:]

	return v
}

func (this *MyQueue) Peek() int {
	length := len(this.data)

	if length == 0 {
		return -1
	}

	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}
