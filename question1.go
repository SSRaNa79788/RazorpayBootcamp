package main
import(
	"fmt"
)
type Matrix struct{
	row,col int
	ele [][]int
}
func (mat Matrix)setRowCol(row,col int) {
	mat.row=row
	mat.col=col
	rowSlice:=make([]int,col)
	for i:=0;i<row;i++ {
		mat.ele=append(mat.ele,rowSlice)
	}
	//fmt.Println(len(mat.ele))//printing number of rows
	//fmt.Println(len(mat.ele[0]))//printing number of columns
}
func (mat Matrix)getCol()int{
	return mat.col
}
func (mat Matrix)getRow()int{
	return mat.row
}
func (mat Matrix)setElement(i int,j int, val int){
	mat.ele[i][j]=val
}
func (mat Matrix)addMatrix(mat2 Matrix)Matrix{//Assuming that both matrix can be added
	var result Matrix
	row:=mat.row
	col:=mat.col
	result.setRowCol(row,col);
	for i:=0 ; i<row ; i++ {
		for j:=0 ; j<col ; j++ {
			result.ele[i][j]=mat.ele[i][j]+mat2.ele[i][j]
		}
	}
	return result
}
func (mat Matrix)printMatrixAsJson(){
	fmt.Println('{')
	fmt.Sprintln("    row : ",mat.row)
	fmt.Sprintln("    column : ",mat.col)
	fmt.Sprintln("    elements : ",mat.ele)
	fmt.Println('}')
}
func main(){
	fmt.Println("hello world")
	var matrix Matrix
	matrix.setRowCol(2,3)
	fmt.Println(matrix)
	//fmt.Println(len(matrix.ele))
	//fmt.Println(matrix.getRow())
	//fmt.Println(matrix.getCol())
}