package algorithms

type MyLinkedList struct {
	Prev  *MyLinkedList
	Next  *MyLinkedList
	Value int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this.GetHead()
	curIdx := 0
	for {
		if curIdx == index || cur.Next == nil {
			return cur.Value
			break
		}
		curIdx++
		cur = cur.Next
	}
	return 0
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.GetHead()
	head.Prev = &MyLinkedList{
		Next:  head,
		Value: val,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Next == nil {
		this.Next = &MyLinkedList{
			Prev:  this,
			Value: val,
		}

	} else {
		next := this.Next
		for {
			if next == nil {
				next.Next = &MyLinkedList{
					Prev:  next,
					Value: val,
				}
				break
			}
			next = next.Next
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newList := &MyLinkedList{
		Value: val,
	}

	cur := this.GetHead()
	curIdx := 0
	for {
		if curIdx == index || cur.Next == nil {
			newList.Prev, newList.Next = cur.Prev, cur
			break
		}
		curIdx++
		cur = cur.Next
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	head := this.GetHead()
	curIdx := 0
	cur := head
	for {
		if curIdx == index || cur.Next == nil {
			if cur.Next != nil {
				cur.Next.Prev, cur.Prev.Next = cur.Prev, cur.Next
			} else {
				cur.Prev.Next = cur.Next
			}
			break
		}
		curIdx++
		cur = cur.Next
	}
}

func (this *MyLinkedList) GetHead() *MyLinkedList {
	if this.Prev == nil || this.Next == nil {
		return this
	}
	prev := this.Prev
	for {
		if prev.Prev != nil {
			prev = prev.Prev
		} else {
			return prev.Next
		}

	}

	return this
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
