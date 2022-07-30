/*
- Problem: https://leetcode.com/problems/min-stack/
*/

type MinStack struct {
    mainStack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.mainStack = append(this.mainStack, val)
    if len(this.minStack) > 0 && this.minStack[len(this.minStack)-1] < val {
        val = this.minStack[len(this.minStack)-1]
    }
    this.minStack = append(this.minStack, val)
}


func (this *MinStack) Pop()  {
    this.mainStack = this.mainStack[:len(this.mainStack) - 1]
    this.minStack = this.minStack[:len(this.minStack) - 1]
}


func (this *MinStack) Top() int {
    return this.mainStack[len(this.mainStack) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
