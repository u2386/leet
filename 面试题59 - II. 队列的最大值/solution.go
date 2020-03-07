package main

type MaxQueue struct {
	q0 []int
	q1 []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		q0: []int{},
		q1: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q0) == 0 {
		return -1
	}
	return this.q1[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q0 = append(this.q0, value)
	for len(this.q1) != 0 && value >= this.q1[len(this.q1)-1] {
		this.q1 = this.q1[:len(this.q1)-1]
	}
	this.q1 = append(this.q1, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q0) == 0 {
		return -1
	}
	var v int
	v, this.q0 = this.q0[0], this.q0[1:]
	if v == this.q1[0] {
		this.q1 = this.q1[1:]
	}
	return v
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
