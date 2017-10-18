package sset

import "github.com/intdxdt/math"

func (s *SSet) Add(v interface{}) *SSet {
	var sub *subset
	if s.IsEmpty() {
		sub = s.allocSubset()
		sub.add(v)
		s.addSubset(sub)
	} else {
		view := s.list.DataView()
		idx := s.findSubsetByMax(v)
		sub = view[idx].(*subset)
		sub.add(v)
	}
	s.updateIndexAtSubset(sub)

	if s.subOverflows(sub) {
		subs := s.splitSub(sub)
		for _, sb := range subs {
			s.addSubset(sb)
		}
	}
	return s
}

//Extend SSet given list of values as params
func (s *SSet) Extend(values ...interface{}) *SSet {
	for _, v := range values {
		s.Add(v)
	}
	return s
}


//Empty SubSSet
func (s *SSet) Empty() *SSet {
	//given this snap short of the view
	//remove each item. Note : remove changes the original
	//view slice, but given this window in time, remove all subsets
	view := s.list.DataView()
	for i := 0; i < len(view); i++ {
		sub := view[i].(*subset)
		s.list.Remove(sub.max())
		sub.set.Empty()
	}
	return s
}

//Remove item from set
func (s *SSet) Remove(items ...interface{}) *SSet {
	if s.IsEmpty() {
		return s
	}

	var idx int
	var sub *subset
	var n = len(items)
	var prev_idx = n - 1
	for i := 0; i < n; i++ {
		idx = s.findSubsetByMax(items[i])
		//update prev index
		prev_idx = math.MinInt(prev_idx, idx-1)

		sub = s.list.Get(idx).(*subset)
		if sub.size() == 1 && s.cmp(sub.max(), items[i]) == 0 {
			s.list.Remove(sub)
		} else {
			sub.set.Remove(items[i])
		}
	}

	//can be -1 if idx == 0
	if prev_idx < 0 {
		prev_idx = 0
	}
	if !s.IsEmpty() {
		sub := s.list.Get(prev_idx).(*subset)
		s.updateIndexAtSubset(sub)
	}

	return s
}

//Pop item from the end of the sorted list
func (s *SSet) Pop() interface{} {
	var val interface{}
	if s.IsEmpty() {
		return val
	}
	view := s.list.DataView()
	n    := len(view)
	sub  := view[n-1].(*subset)

	if sub.size() == 1 {
		val = sub.set.Get(-1)
		s.list.Remove(sub)
	} else {
		val = sub.set.Pop()
	}

	return val
}

//PopLeft item from the beginning of the sorted list
func (s *SSet) PopLeft() interface{} {
	var val interface{}
	if s.IsEmpty() {
		return val
	}

	sub := s.list.Get(0).(*subset)

	if sub.size() == 1 {
		val = sub.set.Get(0)
		s.list.Remove(sub)
	} else {
		val = sub.set.PopLeft()
	}

	if !s.IsEmpty() {
		sub := s.list.Get(0).(*subset)
		s.updateIndexAtSubset(sub)
	}
	return val
}
