// range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned.
package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	//range on arrays and slices provides both the index and value for each entry.
	//we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	//range on map iterates over key/value pairs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//	range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key: ", k)
	}
	//range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
