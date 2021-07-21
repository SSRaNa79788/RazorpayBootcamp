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
func preOrder(cur *node){
	if cur==nil {
		return
	}
	preOrder(cur.left)
	fmt.Print(cur.char)
	preOrder(cur.right)
}
func postOrder(cur *node){
	if cur==nil {
		return
	}
	postOrder(cur.right)
	fmt.Print(cur.char)
	postOrder(cur.left)
}
func pref (c byte)int{
	if c=='*' || c=='/'{
		return 2
	} else{
		return 1
	}
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
		if pref(s[l+1])<=pref(s[l+3]) {
			cur:=node(s[l+1])
			cur.left=solve(s,l,l)
			cur.right=solve(s,l+2,r)
			return cur
		} else { //case 2: s[i+1] will be in the left part of the tree
			cur:=node(s[l+3])
			cur.left=solve(s,l,l+2)
			cur.right=solve(s,l+3,r)
			return cur
		}
	}
}
func main(){
	var exprsn="a+b"
	root:=solve(exprsn,0,len(exprsn)-1))
	preOrder(root)
	fmt.Println(" ")
	postOrder(root)

}