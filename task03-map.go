package homework
//import ("fmt"
//		"sort")
func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var ky = make([]int, len(input))
	var i int
	for k:= range input{
		ky[i] = k
		i++
	}
	result = make([]string, len(input))
	sort.Ints(ky)
	for k:= range ky{
		result[k] = input[ky[k]]
	}
	return result
}

//func main(){
//	input := map[int]string{62: "a", 0: "b", 561: "c"}
//	result := sortMapValues(input)
//	fmt.Println(result)
//}
