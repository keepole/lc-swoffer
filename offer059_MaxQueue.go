// offer059_MaxQueue.go kee > 2021/11/14

package swoffer

// # 剑指 Offer 59 - II. 队列的最大值
// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
// 若队列为空，pop_front 和 max_value 需要返回 -1
//
// ### 示例 1：
// 输入:
// ["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
// [[],[1],[2],[],[],[]]
// 输出: [null,null,null,2,1,2]
//
// ### 示例 2：
// 输入:
// ["MaxQueue","pop_front","max_value"]
// [[],[],[]]
// 输出: [null,-1,-1]
//
// 限制：
// 1 <= push_back,pop_front,max_value的总操作数 <= 10000
// 1 <= value <= 10^5

type MaxQueue struct {
	q    []int
	m    []int
	x, y int
}

func Constructor() MaxQueue {
	return MaxQueue{x: 0, y: 0}
}

func (this *MaxQueue) Max_value() int {
	if this.y == 0 {
		return -1
	}
	return this.m[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.m) > 0 && this.m[this.y-1] < value {
		this.m = this.m[:this.y-1]
		this.y -= 1
	}
	this.m = append(this.m, value)
	this.q = append(this.q, value)
	this.x++
	this.y++
}

func (this *MaxQueue) Pop_front() (v int) {
	if this.x == 0 {
		return -1
	}
	v, this.q = this.q[0], this.q[1:]
	this.x -= 1
	if len(this.m) > 0 && this.m[0] == v {
		this.m = this.m[1:]
		this.y -= 1
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
