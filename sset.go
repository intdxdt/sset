package sset

import (
	"fmt"
	"bytes"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/math"
	"github.com/intdxdt/subsset"
)

const subN = 32
const Load = 1000

//SSet type
type SSet struct {
	cmp     func(a, b interface{}) int
	list    *subsset.SubSSet
	maxes   []interface{}
	offsets []int
	load    int
}

//NewSSet Sorted Set
func NewSSet(cmp cmp.Compare, load ...int) *SSet {
	var ldN = Load
	if len(load) > 0 {
		ldN = load[0]
	}

	var maxCmp = createMaxCmp(cmp)
	var list = subsset.NewSubSSet(maxCmp)

	return &SSet{
		cmp:  cmp,
		list: list,
		load: math.MinInt(ldN, Load),
	}
}

//Clone SSet
func (s *SSet) Clone() *SSet {
	clone := NewSSet(s.cmp, s.load)

	view := s.list.DataView()
	for _, sb := range view {
		sub := sb.(*subset)
		clone.addSubset(sub.clone())
	}
	return clone
}

func (s *SSet) Size() int {
	var n = 0
	if s.IsEmpty() {
		return n
	}
	var view = s.list.DataView()
	var sub  = view[len(view)-1].(*subset)
	n = sub.offset + sub.size()
	return n
}

func (s *SSet) IsEmpty() bool {
	return s.list.IsEmpty()
}

//First item in s
func (s *SSet) First() interface{} {
	if !s.IsEmpty() {
		sub := s.list.Get(0).(*subset)
		return sub.set.First()
	}
	return nil
}

//Last Item in s
func (s *SSet) Last() interface{} {
	if !s.IsEmpty() {
		sub := s.list.Get(-1).(*subset)
		return sub.set.Last()
	}
	return nil
}

//Get value at given index in O(lgN)
func (s *SSet) Get(index int) interface{} {
	if index < 0 {
		index += s.Size()
	}

	idx := s.findSubsetByIndex(index)
	if idx >= 0 {
		sub := s.list.Get(idx).(*subset)
		i, j := sub.offset, sub.offset+sub.size()
		if index >= i && index < j {
			idx = index - i //i=offset
			val := sub.set.Get(idx)
			return val
		}
	}
	return nil
}

//Contains item for the presence of a value in the Array - O(2lgN)
func (s *SSet) Contains(items ...interface{}) bool {
	if s.IsEmpty() {
		return false
	}
	view := s.list.DataView()

	var idx int
	var sub *subset
	var bln = true
	var n = len(items)

	for i := 0; bln && i < n; i++ {
		idx = s.findSubsetByMax(items[i])
		sub = view[idx].(*subset)
		bln = sub.set.Contains(items[i])
	}
	return bln
}

//IndexOf item for the presence of a value in the Array - O(2lgN)
func (s *SSet) IndexOf(item interface{}) int {
	idx := -1
	if s.IsEmpty() {
		return idx
	}
	view := s.list.DataView()
	idx = s.findSubsetByMax(item)
	sub := view[idx].(*subset)
	idx = sub.set.IndexOf(item)
	if idx >= 0 {
		idx += sub.offset
	}
	return idx
}

//Values of the set
func (s *SSet) Values() []interface{} {
	vals := make([]interface{}, 0)
	view := s.list.DataView()
	for i := 0; i < len(view); i++ {
		sub := view[i].(*subset)
		vals = append(vals, sub.vals()...)
	}
	return vals
}

//NextItem gets next given item in the sorted set
func (s *SSet) NextItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	idx := s.IndexOf(v)
	n := s.Size() - 1

	var prev interface{} = nil
	if idx >= 0 && idx < n {
		prev = s.Get(idx + 1)
	}
	return prev
}

//PrevItem gets previous given item in the sorted s
func (s *SSet) PrevItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	idx := s.IndexOf(v)
	n := s.Size() - 1

	var prev interface{} = nil
	if idx > 0 && idx <= n {
		prev = s.Get(idx - 1)
	}
	return prev
}

//Loop through items in the queue with a callback
// if callback returns bool. Break looping with callback
// return as false
func (s *SSet) ForEach(fn func(interface{}, int) bool) {
	vals := s.Values()
	for i, v := range vals {
		if !fn(v, i) {
			break
		}
	}
}

//Filters items based on predicate : func (item Item, i int) bool
func (s *SSet) Filter(fn func(interface{}, int) bool) []interface{} {
	var items = make([]interface{}, 0)
	s.ForEach(func(v interface{}, i int) bool {
		if fn(v, i) {
			items = append(items, v)
		}
		return true
	})
	return items
}

func (s *SSet) String() string {
	var buffer bytes.Buffer
	view := s.Values()
	n := len(view) - 1

	buffer.WriteString("[")
	for i, o := range view {
		token := fmt.Sprintf("%v", o)
		if i < n {
			token += ", "
		}
		buffer.WriteString(token)
	}
	buffer.WriteString("]")
	return buffer.String()
}

//Values of the set
//func (s *SSet) DebugValues() {
//	view := s.list.DataView()
//	for i := 0; i < len(view); i++ {
//		sub := view[i].(*subset)
//		fmt.Println(fmt.Sprintf("%v : %v", i, sub.set.String()))
//	}
//	fmt.Println(strings.Repeat("-", 80))
//}
