package hm_1

import (
	"reflect"
	"testing"
)

func TestArrToMap(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2, 3, 4, 5, 5}
	res := ArrToMap(arr)
	exp := map[int]int{1: 2, 2: 3, 3: 1, 4: 1, 5: 2}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("res = %v, expect %v", res, exp)
	}

}
