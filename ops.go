package sset

//Union - s union
func (s *SSet) Union(other *SSet) *SSet {
	u := s.Clone()
	view := other.Values()
	for _, v := range view {
		u.Add(v)
	}
	return u
}

//Intersection - s intersection
func (s *SSet) Intersection(other *SSet) *SSet {
	inter := NewSSet(s.cmp, s.load)
	view := other.Values()
	for _, v := range view {
		if s.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

//Difference- s difference
//items in s not contained in other
func (s *SSet) Difference(other *SSet) *SSet {
	diff := NewSSet(s.cmp, s.load)
	for _, v := range s.Values() {
		if !other.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}

//SymDifference - symmetric difference with between s and other
//new s with elements in either s or other but not both
func (s *SSet) SymDifference(other *SSet) *SSet {
	return s.Difference(other).Union(
		other.Difference(s),
	)
}
