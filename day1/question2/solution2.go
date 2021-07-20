package main

import (
	"bufio"
	"fmt"
	"os"
)
type node struct{
	char byte
	left,right *node
}
func (cur *node)preOrder(){
	if cur==nil {
		return
	}
	cur.preOrder(cur.left)
	fmt.Print(cur.char)
	cur.preOrder(cur.right)
}
func (cur *node)postOrder(){
	if cur==nil {
		return
	}
	cur.preOrder(cur.right)
	fmt.Print(cur.char)
	cur.preOrder(cur.left)
}
func main(){
	var exprsn="a+b"


}