package hm_2

//目前只有功能测试
import "testing"

func TestCal_Fac(t *testing.T) {
	mid := Cal_Fac("divide")
	arr := []int{24, 0, 3}
	res := mid(arr)
	exp := 0

	if res != exp {
		t.Errorf("Cal_Fac(arr) = %d; wanted %d", res, exp)
	}
}
