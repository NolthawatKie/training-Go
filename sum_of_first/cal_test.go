package cal

import "testing"

func Testsum_of_first(t *testing.T) {
	let := 4
	expect := 10
	get := sum_of_first(let)
	if get != expect {
		t.Errorf("given %d expect %d but get %d", let, expect, get)
	}

}
