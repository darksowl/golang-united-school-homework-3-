package homework
//import ("fmt")
func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	var n float32
	for _, s := range input {
		if s!=0 {
			sum += s
			n++
		}
	}
	return sum/n
}

//func main(){
//	var input = [15]float32{1,2,3,4,5,6}
//	var av float32
//	av = average(input)
//	fmt.Print(av)
//}