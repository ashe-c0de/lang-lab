func main() {
	// 创建map
	myMap := make(map[string]any)
	// var cache = map[string]any{}
	
	// 写
	myMap["a"] = 1
	myMap["b"] = 2
	myMap["c"] = 3

	// 读
	v1, ok := myMap["c"]
	if ok {
		fmt.Println(v1)
	}

	v2, ok := myMap["d"]
	if ok {
		fmt.Println(v2)
	}

	// 删
	delete(myMap, "c")
	delete(myMap, "d")

	// 遍历
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
