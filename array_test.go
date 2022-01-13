package main

import (
	"edenfarm-be-jiddyabd/lib"
	"fmt"
	"strings"
	"testing"
)

var (
	case_1   = []int{1, 2, 3, 8, 9, 3, 2, 1}
	case_2   = []int{1, 2, 1, 2, 2, 1}
	case_3   = []int{7, 1, 2, 9, 7, 2, 1}
	result_1 = 3
	result_2 = 2
	result_3 = 2
)

func TestCase1(t *testing.T) {
	t.Log("Case 1 : " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(case_1)), ", "), "[]"))
	t.Log("Hasil : ", lib.LargestValue(case_1))

	if lib.LargestValue(case_1) != result_1 {
		t.Error("Salah, hasil yang benar adalah " + fmt.Sprint(result_1))
	}
}

func TestCase2(t *testing.T) {
	t.Log("Case 2 : " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(case_2)), ", "), "[]"))
	t.Log("Hasil : ", lib.LargestValue(case_2))

	if lib.LargestValue(case_2) != result_2 {
		t.Error("Salah, hasil yang benar adalah " + fmt.Sprint(result_2))
	}
}

func TestCase3(t *testing.T) {
	t.Log("Case 3 : " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(case_3)), ", "), "[]"))
	t.Log("Hasil : ", lib.LargestValue(case_3))

	if lib.LargestValue(case_3) != result_3 {
		t.Error("Salah, hasil yang benar adalah " + fmt.Sprint(result_3))
	}
}
