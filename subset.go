package sset

import "github.com/intdxdt/subsset"

func (s *SSet) addSubset(sub *subset) *SSet {
	s.list.Add(sub)
	return s
}

func (s *SSet) allocSubset() *subset {
	sub := subsset.NewSubSSet(s.cmp, subN)
	return &subset{set: sub, offset: 0}
}

//----------------------------------------------------------
type subset struct {
	set    *subsset.SubSSet
	offset int
}

func (sb *subset) clone() *subset {
	return &subset{set: sb.set.Clone(), offset: sb.offset}
}

func (sb *subset) max() interface{} {
	return sb.set.Get(-1)
}

func (sb *subset) add(v interface{}) {
	sb.set.Add(v)
}

func (sb *subset) size() int {
	return sb.set.Size()
}

func (sb *subset) vals() []interface{} {
	return sb.set.DataView()
}

func (sb *subset) addVals(vals []interface{}) {
	for _, v := range vals {
		sb.add(v)
	}
}
