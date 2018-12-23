package playground

import (
	"bytes"
	"fmt"
)

func example() {
	countries := []byte("Australia Canada Japan Germany India เจ้าตัวอ้วน")
	space := []byte{' '}
	splitExample := bytes.Split(countries, space)
	fmt.Println(splitExample)
	fmt.Println(space)

	dash := []byte{'-'}

	result := bytes.Join(splitExample, dash)

	fmt.Println(result)

	// way 1
	testcopy := make([]byte, len(result))
	copy(testcopy, result)

	// way 2(nice approach)
	testcopy2 := result[:]

	fmt.Println("test ", testcopy)
	fmt.Println("test2 ", testcopy2)

}
