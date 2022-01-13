package lib

import (
	"reflect"
)

func findMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func LargestValue(array []int) int {
	highest := 0
	for index := range array {
		//Search deret
		var pattern []int
		pattern = append(pattern, array[index])
		for i := 1; i < len(array)-index; i++ {
			//Check next value
			if array[index+i] == array[index]+i {
				pattern = append(pattern, array[index+i])
			} else {
				break
			}
		}

		//Check if length pattern > 1
		if len(pattern) > 1 {
			//Search reverse deret
			array_left := array[index+len(pattern):]
			//Reverse existing pattern
			reverse_pattern := reverse(pattern)
			for j := range array_left {
				//Check if len pattern <= leftover array on the right side of the pattern
				if j+len(pattern) <= len(array_left) {
					//Check if there's matching reverse pattern on the leftover array
					if reflect.DeepEqual(reverse_pattern, array_left[j:j+len(pattern)]) {
						//Find max value from the pattern if the reverse pattern are matching
						if highest < findMax(pattern) {
							highest = findMax(pattern)
						}
					}
				}
			}
		}

	}
	return highest
}
