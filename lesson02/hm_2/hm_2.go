package hm_2

//目前是可以导入任意数量数进行单一类型运算的进阶I计算器

func Cal_Fac(op string) func([]int) int {
	switch op {
	case "plus":
		return func(arr []int) int {
			res := arr[0] + arr[1]
			for i := 2; i < len(arr); i++ {
				res = res + arr[i]
			}
			return res
		}

	case "minus":
		return func(arr []int) int {
			res := arr[0] - arr[1]
			for i := 2; i < len(arr); i++ {
				res = res - arr[i]
			}
			return res
		}
	case "multiply":
		return func(arr []int) int {
			res := arr[0] * arr[1]
			for i := 2; i < len(arr); i++ {
				res = res * arr[i]
			}
			return res
		}
	case "divide":
		return func(arr []int) int {
			for _, n := range arr {
				if n == 0 {
					return 0
				}
			}
			res := arr[0] / arr[1]
			for i := 2; i < len(arr); i++ {
				res = res / arr[i]
			}
			return res
		}
	}
	return nil
}
