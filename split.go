package sset

const splitFactor = 2

func (s *SSet) subOverflows(sub *subset) bool {
	return sub.size() > (s.load * 2)
}

func (s *SSet) splitSub(sub *subset) []*subset {
	var subs = make([]*subset, 0)
	var view = sub.vals()
	var n = len(view)
	var end int
	var chunk []interface{}

	var size = (sub.size() + splitFactor) / splitFactor
	var offset = sub.offset

	for i := 0; i < n; i += size {
		end = i + size
		if end > n {
			end = n
		}

		chunk = view[i:end]
		if len(chunk) > 0 {
			sb := s.allocSubset()
			sb.addVals(chunk)

			sb.offset = offset
			offset += sb.size()

			subs = append(subs, sb)
		}
	}

	return subs
}
