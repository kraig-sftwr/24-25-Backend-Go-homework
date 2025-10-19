package hm_1

func ArrToMap(arr []int) map[int]int {
	m := make(map[int]int)

	for _, ele := range arr {
		m[ele]++
	}
	return m
}
