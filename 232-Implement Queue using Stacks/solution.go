package main

type stack struct {
	arr []int
}

func (s *stack) Push(x int) {
	s.arr = append(s.arr, x)
}

func (s *stack) Pop() int {
	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return r
}

func (s *stack) Size() int {
	return len(s.arr)
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

type MyQueue struct {
	stack0 stack
	stack1 stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack0: stack{
			arr: make([]int, 0),
		},
		stack1: stack{
			arr: make([]int, 0),
		},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack0.Push(x)
}

// Pop removes the element from in front of queue and returns that element
func (this *MyQueue) Pop() int {
	for !this.stack0.isEmpty() {
		this.stack1.Push(this.stack0.Pop())
	}

	r := this.stack1.Pop()
	for !this.stack1.isEmpty() {
		this.stack0.Push(this.stack1.Pop())
	}

	return r
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for !this.stack0.isEmpty() {
		this.stack1.Push(this.stack0.Pop())
	}

	r := this.stack1.Pop()

	this.stack1.Push(r)
	for !this.stack1.isEmpty() {
		this.stack0.Push(this.stack1.Pop())
	}

	return r
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack0.isEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
