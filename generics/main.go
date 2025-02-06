package main
import "fmt"
func main(){
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}
	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}
	fmt.Printf("Non-Generic sums: %v and %v \n", SumInts(ints), SumFloats(floats))

	fmt.Printf("Generic sums: %v and %v \n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}
// Non Generics
func SumInts(m map[string]int64) int64{
	var s int64

	for _, v := range m{
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64{

	var s float64
	
	for _, v := range m{
		s += v
	}
	return s
}

// Generic

/* 
	Type parameters are written in [] prior to the function arguments ()
	V can be either int or float
	K is a comparable constraint. It allows for any types that can be compared with comparison operators such as == or !=
	
	GO required map keys to be comparable
	
*/
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V{
	var s V

	for  _, v := range m{
		s += v
	}
	return s


}
