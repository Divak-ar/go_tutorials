package main

import (
	"fmt"
	"reflect"
)

func main(){

	// Generics are used when we have to create a fn that can take different data types but perform the same operations on them , like calculating the items (len) in array , doesn't matter the type we have iterate over the array and return the length but when defining func we have to define the parameters type thus limitng the data type we can enter, so we have define the same func again for []int , []str, []bool data, Instead we can use Generics T to act as a place holder for any data type , []T 0> placeholder array (we also have any type) t any

	var intSlice = []int{1,2,4}
	var floatSlice = []float32{3.5,3.2,6.5}
	var boolSlice = []bool{}

	fmt.Println(sumSlice(intSlice))
	fmt.Println(sumSlice(floatSlice))
	fmt.Println(isEmpty(boolSlice))

}

// There is no special type called a class or object exist in GO. strct type in Go comes the closet to what is commonly refer as object in other programming language. Structs are used for it .

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	// perform an operation that is compatible with the data types defined here , if we did sum = strings.Toupper(v) will give an error as it is ony applicable on strings
	for _, v := range slice{
		sum += v
	}
	fmt.Println(reflect.TypeOf(sum))
	return sum
}


func isEmpty[T any](slice []T) bool{
	// type any cause len() is applicable to slices/array -> [] and array can be of any type
	return len(slice)==0
}


