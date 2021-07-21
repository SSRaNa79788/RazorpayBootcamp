package main

import (
	"fmt"
)
type node struct{
	char string
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
	postOrder(cur.left)
	fmt.Print(cur.char)
}
func pref (c string)int{
	if c=="*" || c=="/"{
		return 2
	} else{
		return 1
	}
}
func buildTree(s string,l int,r int ) *node {
	if l==r {
		cur := &node{char:string([]rune(s)[l])}
		return cur
	} else if l==r-2 {
		cur:=&node{char:string([]rune(s)[l+1])}
		cur.left=buildTree(s,l,l)
		cur.right=buildTree(s,r,r)
		return cur
	} else {
		//case 1: s[i+1] wil be the current treenode
		if pref(string([]rune(s)[l+1]))<=pref(string([]rune(s)[l+3])) {
			cur:=&node{char:string([]rune(s)[l+1])}
			cur.left=buildTree(s,l,l)
			cur.right=buildTree(s,l+2,r)
			return cur
		} else { //case 2: s[i+1] will be in the left part of the tree
			cur:=&node{char:string([]rune(s)[l+3])}
			cur.left=buildTree(s,l,l+2)
			cur.right=buildTree(s,l+4,r)
			return cur
		}
	}
}
func main(){
	exprsn:="a*b-c"
	//fmt.Println(string([]rune(exprsn)[1]))

	root:=buildTree(exprsn,0,len(exprsn)-1)
	fmt.Println("The Preorder traversal is :")
	preOrder(root)
	fmt.Println("\nThe Postorder traversal is :")
	postOrder(root)
	fmt.Println("")

}