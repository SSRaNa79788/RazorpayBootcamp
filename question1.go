package main
import(
	"fmt"
)
type Matrix struct{
	row,col int
	ele [][]int
}
func (mat *Matrix)setRowCol(row,col int) {//verified
	mat.row=row
	mat.col=col
	for i:=0;i<row;i++ {
		rowSlice:=make([]int,col)
		mat.ele=append(mat.ele,rowSlice)
	}

	//fmt.Println(len(mat.ele))//printing number of rows
	//fmt.Println(len(mat.ele[0]))//printing number of columns
}
func (mat *Matrix)getCol()int{//verified
	return mat.col
}
func (mat *Matrix)getRow()int{//verified
	return mat.row
}
func (mat *Matrix)setElement(i int,j int, val int){//verified
	mat.ele[i][j]=val
}
func (mat *Matrix)addMatrix(mat2 Matrix)Matrix{//Assuming that both matrix can be added
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
func (mat *Matrix)printMatrixAsJson(){//verified
	fmt.Println("{")
	fmt.Print("    row : ")
	fmt.Println(mat.row)
	fmt.Print("    column : ")
	fmt.Println(mat.col)
	fmt.Print("    elements : ")
	fmt.Println(mat.ele)
	fmt.Println("}")
}
func main(){
	//fmt.Println("hello world")
	var matrix Matrix
	pt:=&matrix
	pt.setRowCol(2,3)
	//fmt.Println(matrix)
	//fmt.Println(pt.getCol())
	//fmt.Println(pt.getRow())

	//pt.setElement(1,2,23)
	//fmt.Println(pt.ele[1][2])

	//
	//var matrix2 Matrix
	//pt2:=&matrix2
	//pt2.setRowCol(2,3)
	//fmt.Println(matrix2)
	//
	//matrix.ele[0][0]=12
	//matrix2.ele[0][0]=12
	//fmt.Println(matrix)
	//
	//fmt.Println(matrix.addMatrix(matrix2))

	pt.printMatrixAsJson()
	matrix.ele[0][0]=1
	pt.printMatrixAsJson()

}