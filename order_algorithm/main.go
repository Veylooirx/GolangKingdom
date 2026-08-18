package main

import "fmt"

func main() {

	numbers := []int{12, 233, 53, 62, 85, -9, 123} 

	fmt.Println(ordenar(numbers))
}

func ordenar(numbers[] int ) [] int  {

	ordened := make([]int, 0, len(numbers))

	value := minValue(numbers)
	newSlice := delete(numbers, value)
	ordened = append(ordened, value)

	for { 
		if len(ordened) < len(numbers)  { 
			value = minValue(newSlice)
			ordened = append(ordened, value)
			newSlice = delete(newSlice, value)

		}else {
			break	
		}
	}
	return ordened

}


func minValue(numbers []int) int { 

	value := numbers[0]
	for _, num := range numbers {
		if value > num {  
			value = num
		}
	}
	return value
}

func delete(numbers []int, value int) []int {

    slice := make([]int, len(numbers))

	 j := 0
    for _, num := range numbers {
        if num != value  {
            slice[j] = num 
			j++
        }
    }
    return slice[:j]
}
