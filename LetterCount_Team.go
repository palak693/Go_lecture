package main

import "fmt"

func main() {
	arr := []string{"Palak", "Sabahat", "Reet", "Akanksha"}
	freq := make(map[string]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
		// freq[num] ++
	}

	fmt.Println("Frequency of the Array is:", freq)

	/* var int name = 0
	index := 0
	for _, name := range arr {
		index++
	}*/

	for j := 0; j <= len(arr); j++ {
		name := arr[j]
		letco := make(map[string]int)
		for _, char := range name {
			letco[string(char)]++
		}
		fmt.Println(letco)
	}
}
