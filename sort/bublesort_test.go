package sort

import "testing"

func TestBubblesort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	Bubblesort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("冒泡排序失败：", values, "。应该是：12345")
	}
	t.Log("成功：", values)
}