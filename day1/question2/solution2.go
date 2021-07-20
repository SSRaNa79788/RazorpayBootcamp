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
func solve(s string,l int,r int ) *node {
	if l==r {
		cur := node(s[l],nil,nil)
		return &cur
	} else if l==r-2 {
		cur:=node(s[l+1])
		cur.left=solve(s,l,l)
		cur.right=solve(s,r,r)
		return &cur
	} else {
		//case 1: s[i+1] wil be the current treenode
		if pref()<=pref() {
			cur:=node(s[l+1])
			cur.left=solve(s,l,l)
			cur.right=solve(s,l+2,r)
		} else { //case 2: s[i+1] will be in the left part of the tree
			cur:=node(s[l+3])
			cur.left=solve(s,l,)
		}
	}
}
func main(){
	var exprsn="a+b"


}