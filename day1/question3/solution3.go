package main

import "fmt"

type salaryCalculation interface{
	salary(int)int
}
type fulltime struct{}
type contractor struct{}
type freelancer struct{}
func (e *fulltime)salary(day int)int{
	return 500*day
}
func (e *contractor)salary(day int)int{
	return 100*day
}
func (e *freelancer)salary(hr int)int{
	return 10*hr
}
func main(){
	var sc salaryCalculation
	sc=&fulltime{}
	fmt.Println(sc.salary(30))
	sc=&contractor{}
	fmt.Println(sc.salary(30))
	sc=&freelancer{}
	fmt.Println(sc.salary(20))
}