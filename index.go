package sset

import "github.com/intdxdt/algor"

func (s *SSet) updateIndexAtSubset(sub *subset) *SSet {
	var prev *subset
	var view = s.list.DataView()
	var idx = s.list.IndexOf(sub)

	if idx > 0 {
		prev = view[idx-1].(*subset)
		sub.offset = prev.offset + prev.size()
	} else if idx == 0 {
		sub.offset = 0
	}

	n := sub.offset + sub.size()
	for i := idx + 1; i < len(view); i++ {
		sub = view[i].(*subset)
		sub.offset = n
		n += sub.size()
	}
	return s
}

//func (s *SSet) update_indexx() *SSet {
//	n := 0
//	view := s.list.DataView()
//	for i := 0; i < s.list.Size(); i++ {
//		sub := view[i].(*subset)
//		sub.offset = n
//		n += sub.size()
//	}
//	return s
//}

func (s *SSet) findSubsetByIndex(index int) int {
	view := s.list.DataView()
	idx := algor.BS(view, index, offsetCmp, 0, len(view)-1)
	if idx < 0 {
		idx = -idx - 2
	}
	return idx
}

func (s *SSet) findSubsetByMax(v interface{}) int {
	var view   = s.list.DataView()
	var maxCmp = createMaxCmp(s.cmp)
	var n      = len(view) - 1

	idx := algor.BS(view, v, maxCmp, 0, n)
	if idx < 0 {
		idx = -idx - 1
	}

	if idx > n {
		idx = n
	}
	return idx
}
