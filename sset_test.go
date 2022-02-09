package sset

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/intdxdt/cmp"
	"math/rand"
	"testing"
	"time"
)

func IntCompare(a, b interface{}) int {
	return a.(int) - b.(int)
}

func FloatCompare(a, b interface{}) int {
	var d = a.(float64) - b.(float64)
	if d < 0 {
		return -1
	} else if d > 0 {
		return 1
	}
	return 0
}

func StrCompare(a, b interface{}) int {
	var str, v = a.(string), b.(string)
	if str < v {
		return -1
	} else if str > v {
		return 1
	}
	return 0
}

var seed = rand.NewSource(time.Now().UnixNano())
var rnd = rand.New(seed)

func GenRandData(N int) []float64 {
	var data = make([]float64, N, N)
	for i := 0; i < N; i++ {
		data[i] = rnd.Float64()
	}
	return data
}

func GenRandIndxs(N int) []int {
	var data = make([]int, N, N)
	for i := 0; i < N; i++ {
		data[i] = rnd.Intn(N - 1)
	}
	return data
}

func TestSSet(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("SSet", func() {
		var array = []interface{}{9, 5, 3, 2, 8, 6, 4, 6, 1, 2, 3}
		var arrayF = []interface{}{9.0, 5.0, 3.0, 2.0, 8.0, 6.0, 4.0, 6.0, 1.0, 2.0, 3.0}

		var objFlt = NewSSet(FloatCompare, 2)
		for _, v := range arrayF {
			objFlt.Add(v)
		}

		g.It("should test s - common-int interface", func() {
			var st = NewSSet(IntCompare, 2)
			for _, v := range array {
				st.Add(v)
			}
			fmt.Println(st)

			var f6 = 6
			var f3 = 3

			g.Assert(st.Size() == 8).IsTrue()
			g.Assert(st.Contains(f6)).IsTrue()
			g.Assert(st.Contains(f6, f3)).IsTrue()
			g.Assert(st.Remove(f6).Contains(f6)).IsFalse()
			g.Assert(st.Size() == 7).IsTrue()
			g.Assert(st.IsEmpty()).IsFalse()

			var eachItem = make([]int, 0)
			st.ForEach(func(o interface{}, _ int) bool {
				eachItem = append(eachItem, o.(int))
				return true
			})
			g.Assert(eachItem).Eql([]int{1, 2, 3, 4, 5, 8, 9})
			var oddList = st.Filter(func(o interface{}, _ int) bool {
				return o.(int)%2 == 1
			})
			var first3oddList = make([]interface{}, 0)

			st.ForEach(func(o interface{}, _ int) bool {
				if o.(int)%2 == 1 {
					first3oddList = append(first3oddList, o.(int))
				}
				if len(first3oddList) == 3 {
					return false
				}
				return true
			})

			g.Assert(st.IndexOf(9)).Equal(len(eachItem) - 1)
			g.Assert(st.IndexOf(1)).Equal(0)
			g.Assert(st.IndexOf(6)).Equal(-1)

			g.Assert(len(oddList)).Eql(4)
			g.Assert(len(first3oddList)).Eql(3)

			var odds = make([]int, 0)
			for _, v := range oddList {
				odds = append(odds, v.(int))
			}
			g.Assert(odds).Eql([]int{1, 3, 5, 9})
			g.Assert(st.First()).Equal(1)
			g.Assert(st.Last()).Equal(9)
			g.Assert(st.PopLeft()).Equal(1)
			g.Assert(st.Pop()).Equal(9)

			g.Assert(st.First()).Equal(2)
			g.Assert(st.Last()).Equal(8)
			g.Assert(st.Remove(9).Last()).Equal(8)
			g.Assert(st.Remove(2).Last()).Equal(8)
			g.Assert(st.Remove(5).First()).Equal(3)
			g.Assert(st.Remove(4).Last()).Equal(8)
			g.Assert(st.Pop()).Equal(8)
			g.Assert(st.Remove(3).IsEmpty()).IsTrue()

			st.Empty()
			st.Add(8)
			g.Assert(st.PopLeft()).Equal(8)
			g.Assert(st.Remove(3).IsEmpty()).IsTrue()

			g.Assert(st.First()).Equal(nil)
			g.Assert(st.Last()).Equal(nil)
			g.Assert(st.PopLeft()).Equal(nil)
			g.Assert(st.Pop()).Equal(nil)
			g.Assert(st.IsEmpty()).IsTrue()
			g.Assert(st.IndexOf(8) == -1).IsTrue()
			g.Assert(st.Remove(8).IsEmpty()).IsTrue()
			g.Assert(st.Contains(8)).IsFalse()

			//print
			fmt.Println(st)
		})

		g.It("should test s - common-float interface", func() {
			var st = NewSSet(FloatCompare, 2)
			g.Assert(st.Size() == 0).IsTrue()
			for _, v := range arrayF {
				st.Add(v)
			}
			fmt.Println(st)

			var f6 = 6.0
			var f3 = 3.0
			g.Assert(st.Contains(f6)).IsTrue()
			g.Assert(st.Contains(f6, f3)).IsTrue()
			g.Assert(st.Remove(f6).Contains(f6)).IsFalse()
			g.Assert(st.IsEmpty()).IsFalse()

			var eachItem = make([]float64, 0)
			st.ForEach(func(o interface{}, _ int) bool {
				eachItem = append(eachItem, o.(float64))
				return true
			})
			g.Assert(eachItem).Eql([]float64{1, 2, 3, 4, 5, 8, 9})

			g.Assert(st.First()).Equal(1.0)
			g.Assert(st.Last()).Equal(9.0)
			g.Assert(st.PopLeft()).Equal(1.0)
			g.Assert(st.Pop()).Equal(9.0)

			g.Assert(st.First()).Equal(2.)
			g.Assert(st.Last()).Equal(8.)
			st.Empty()
			g.Assert(st.Last() == nil).IsTrue()
			g.Assert(st.PopLeft()).Equal(nil)
			g.Assert(st.Pop()).Equal(nil)
			g.Assert(st.IsEmpty()).IsTrue()

			//print
			fmt.Println(st)
		})

		g.It("should test s - common-string interface", func() {
			var st = NewSSet(StrCompare)
			st.Extend("foo", "bar", "baz", "tar", "fiz", "tau", "aww")
			fmt.Println(st)
			g.Assert(st.Size()).Equal(7)

			var s3 = "pi"
			var s6 = "tau"
			g.Assert(st.Contains(s6)).IsTrue()
			g.Assert(st.Contains(s6, s3)).IsFalse()
			g.Assert(st.Remove(s6).Contains(s6)).IsFalse()
			g.Assert(st.IsEmpty()).IsFalse()

			var eachItem = make([]interface{}, 0)
			st.ForEach(func(o interface{}, _ int) bool {
				eachItem = append(eachItem, o.(string))
				return true
			})

			fmt.Print("\nRemoved tau\n")
			fmt.Println(st)

			g.Assert(eachItem).Eql([]interface{}{"aww", "bar", "baz", "fiz", "foo", "tar"})

			g.Assert(st.First()).Equal("aww")
			g.Assert(st.PrevItem("aww") == nil).IsTrue()
			g.Assert(st.PrevItem("zzz") == nil).IsTrue()
			g.Assert(st.PrevItem("bar")).Equal("aww")
			g.Assert(st.PrevItem("tar")).Equal("foo")

			g.Assert(st.Last()).Equal("tar")
			g.Assert(st.NextItem("tar") == nil).IsTrue()
			g.Assert(st.NextItem("zzz") == nil).IsTrue()
			g.Assert(st.NextItem("bar")).Equal("baz")
			g.Assert(st.NextItem("aww")).Equal("bar")

			g.Assert(st.PopLeft()).Equal("aww")
			g.Assert(st.Pop()).Equal("tar")

			g.Assert(st.First()).Equal("bar")
			g.Assert(st.Last()).Equal("foo")
			g.Assert(st.Get(-1)).Equal("foo")
			st.Empty()
			g.Assert(st.Last() == nil).IsTrue()
			g.Assert(st.Get(-1) == nil).IsTrue()
			g.Assert(st.Get(1) == nil).IsTrue()

			g.Assert(st.NextItem("tar") == nil).IsTrue()
			g.Assert(st.PrevItem("tar") == nil).IsTrue()
			g.Assert(st.PopLeft()).Equal(nil)
			g.Assert(st.Pop()).Equal(nil)
			g.Assert(st.IsEmpty()).IsTrue()
			g.Assert(st.Size()).Equal(0)

			//print
			fmt.Println(st)
		})

	})
}

func TestSSet_LargeSet(t *testing.T) {
	g := goblin.Goblin(t)
	var N = 1000000
	var BenchData = GenRandData(N)
	var BenchIdxs = GenRandIndxs(1000)
	var lst = NewSSet(cmp.F64)

	g.Describe("Large SSet", func() {
		g.Before(func() {
			fmt.Println("\t >> building large set...")
			for i := 0; i < len(BenchData); i++ {
				lst.Add(BenchData[i])
			}
		})
		g.It("should test build and index", func() {
			for i := 0; i < len(BenchIdxs); i++ {
				idx := lst.IndexOf(BenchData[BenchIdxs[i]])
				g.Assert(idx >= 0).IsTrue()
			}
			for i := 0; i < len(BenchData); i++ {
				if rnd.Float64() > 0.5 {
					lst.Pop()
				} else {
					lst.PopLeft()
				}
			}
			g.Assert(lst.IsEmpty()).IsTrue()
		})
	})
}

func TestSSet_Set(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SSet - Set Opt", func() {
		var s1 = NewSSet(IntCompare)
		var s2 = NewSSet(IntCompare)
		var s3 = NewSSet(IntCompare)

		s1.Add(1)
		s1.Add(0)
		s1.Add(5)
		s1.Add(4)
		s1.Add(7)
		s1.Add(10)
		s1.Add(13)

		//===
		s2.Add(4)
		s2.Add(7)
		s2.Add(9)
		s2.Add(10)
		s2.Add(20)
		s2.Add(17)
		s2.Add(91)
		//===
		s3.Add(21)
		s3.Add(11)
		s3.Add(12)
		s3.Add(41)
		s3.Add(92)

		union_set := s1.Union(s2)
		intertree := s1.Intersection(s2)
		d1tree := s1.Difference(s2)
		d2tree := s2.Difference(s1)
		symtree := s2.SymDifference(s1)
		uset := make([]interface{}, 0)
		for _, v := range []int{0, 1, 4, 5, 7, 9, 10, 13, 17, 20, 91} {
			uset = append(uset, v)
		}
		g.Assert(union_set.Values()).Eql(uset)
		g.Assert(intertree.Values()).Eql([]interface{}{4, 7, 10})
		g.Assert(d1tree.Values()).Eql([]interface{}{0, 1, 5, 13})
		g.Assert(d2tree.Values()).Eql([]interface{}{9, 17, 20, 91})
		g.Assert(symtree.Values()).Eql([]interface{}{
			0, 1, 5, 9, 13, 17, 20, 91,
		})
	})
}
