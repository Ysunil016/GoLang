package main

func main() {

}

// Node ...
type Node struct {
	val   int
	count int
}

// StockSpanner ...
type StockSpanner struct {
	stack []Node
}

// Constructor ...
func Constructor() StockSpanner {
	return StockSpanner{stack: []Node{}}
}

// Next ...
func (sS *StockSpanner) Next(price int) int {
	count := 1
	for len(sS.stack) != 0 && sS.stack[len(sS.stack)-1].val <= price {
		count += sS.stack[len(sS.stack)-1].count
		sS.stack = sS.stack[0 : len(sS.stack)-1]
	}
	sS.stack = append(sS.stack, Node{price, count})
	return count
}
