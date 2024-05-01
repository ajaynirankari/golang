package main

import "fmt"

func main() {

	m := make(map[int]string)

	fmt.Println(m, "len", len(m))

	m[1] = "one"
	m[2] = "two"

	fmt.Println(m, "len", len(m))

	fmt.Println("m[1] = ", m[1])
	fmt.Println("m[2] = ", m[2])
	fmt.Println("m[3] = ", m[3])

	v1 := m[1]
	fmt.Println("v1 = ", v1)
	v2 := m[2]
	fmt.Println("v2 = ", v2)

	fmt.Println("Before delete ", m)
	delete(m, 2)
	fmt.Println("After delete", m)

	clear(m)

	fmt.Println("Clear", m)

	m1 := map[int]string{11: "Eleven", 12: "Twel", 13: "Thr"}
	fmt.Println(m1)

	for k, v := range m1 {
		fmt.Println("k = ", k, ", v = ", v)
	}
	fmt.Println("---------------")

	for k, _ := range m1 {
		fmt.Println("k = ", k, ", v = ", m1[k])
	}
	fmt.Println("---------------")

	for _, v := range m1 {
		fmt.Println("v = ", v)
	}

	m2 := make(map[int]string, 15)
	fmt.Println(m2, ",len = ", len(m2))

	message := "This is test message"
	fmt.Println("Input message = ", message)
	outputMap := characterCount(message)

	for k, v := range outputMap {
		fmt.Println(k, "--->", v)
	}
}

func characterCount(message string) map[string]int {

	resultMap := make(map[string]int)

	for _, v := range message {
		count := resultMap[string(v)]

		if count == 0 {
			resultMap[string(v)] = 1
		} else {
			resultMap[string(v)] = count + 1
		}
	}

	return resultMap
}
