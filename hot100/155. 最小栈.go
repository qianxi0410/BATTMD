package hot100

type entity struct {
	val int
	minVal int
}

type list struct {
	e entity
	next *list
}

type MinStack struct {
	*list
}


/** initialize your data structure here. */
func Constructor() MinStack {
	list := &list{
		e:    entity{
			-1,
			-1,
		},
		next: nil,
	}
	return MinStack{list}
}


func (this *MinStack) Push(val int)  {
	if this.next == nil { 	// 只有一个元素
		t := &list{e: entity{
			val:    val,
			minVal: val,
		}}
		this.next = t
	} else { // 存在元素
		t := &list{
			e:    entity{
				val: val,
				minVal: min(this.next.e.minVal, val),
			},
			next: this.next,
		}
		this.next = t
	}
}


func (this *MinStack) Pop()  {
	this.next = this.next.next
}


func (this *MinStack) Top() int {
	return this.next.e.val
}


func (this *MinStack) GetMin() int {
	return this.next.e.minVal
}

