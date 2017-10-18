package sset

import "github.com/intdxdt/cmp"

func createMaxCmp(cmp cmp.Compare) cmp.Compare {
	return func(as, bs interface{}) int {
		ma, mb := as, bs
		a, ok := as.(*subset)
		if ok {
			ma = a.max()
		}

		b, ok := bs.(*subset)
		if ok {
			mb = b.max()
		}

		d := cmp(ma, mb)
		if d < 0 {
			return -1
		} else if d > 0 {
			return 1
		}
		return 0
	}
}

func offsetCmp(as, bs interface{}) int {
	var i, j int
	a, ok := as.(*subset)
	if ok {
		i = a.offset
	} else {
		i = as.(int)
	}

	b, ok := bs.(*subset)
	if ok {
		j = b.offset
	} else {
		j = bs.(int)
	}
	return i - j
}
