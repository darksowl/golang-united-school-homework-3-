package homework
//import ("fmt")
func reverse(input []int64) (result []int64) {
	//Place your code here
	l := len(input)
	c := cap(input)
	result = make([]int64,l,c)
	for i:=0; i<l; i++ {
		result[l-1-i] = input[i]
	}
	return result
}
//func main() {
//	input := []int64{1,2,5,15}
//	result := reverse(input)
//	fmt.Println(result)
//}
