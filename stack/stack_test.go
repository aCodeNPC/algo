package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type opval[T any] struct {
	op  string
	val T
}

func TestSequentialEqual(t *testing.T) {
	tests := []struct {
		name    string
		cap     int
		ov      []opval[int]
		compare []int
		want    bool
	}{
		{"test1", -1, nil, nil, true},
		{"test2", -1, nil, []int{}, false},
		{"test3", 1, nil, nil, false},
		{"test4", 2, []opval[int]{{"Push", 1}, {"Push", 2}}, []int{2}, false},
		{"test5", 2, []opval[int]{{"Push", 1}, {"Push", 2}}, []int{2, 1}, false},
		{"test6", 2, []opval[int]{{"Push", 1}, {"Push", 2}}, []int{1, 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSequentialStack[int](tt.cap)
			for _, ov := range tt.ov {
				if ov.op == "Push" {
					s.Push(ov.val)
				} else {
					s.Pop()
				}
			}

			assert.Equal(t, tt.want, s.Equal(tt.compare))
		})
	}

	strtests := []struct {
		name    string
		cap     int
		ov      []opval[string]
		compare []string
		want    bool
	}{
		{"strtest1", -1, nil, nil, true},
		{"strtest2", -1, nil, []string{}, false},
		{"strtest3", 1, nil, nil, false},
		{"strtest4", 2, []opval[string]{{"Push", "1"}, {"Push", "2"}}, []string{"2"}, false},
		{"strtest5", 2, []opval[string]{{"Push", "1"}, {"Push", "2"}}, []string{"2", "1"}, false},
		{"strtest6", 2, []opval[string]{{"Push", "1"}, {"Push", "2"}}, []string{"1", "2"}, true},
	}

	for _, tt := range strtests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSequentialStack[string](tt.cap)
			for _, ov := range tt.ov {
				if ov.op == "Push" {
					s.Push(ov.val)
				} else {
					s.Pop()
				}
			}

			assert.Equal(t, tt.want, s.Equal(tt.compare))
		})
	}
}

func TestSequentialStack(t *testing.T) {
	s := NewSequentialStack[int](3)
	s.Push(1)
	s.Push(2)
	c := s.Count()
	assert.Equal(t, 2, c)

	v, ok := s.Pop()
	assert.Equal(t, 2, v)
	assert.Equal(t, true, ok)
	c = s.Count()
	assert.Equal(t, 1, c)

	v, ok = s.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, true, ok)
	c = s.Count()
	assert.Equal(t, 0, c)

	v, ok = s.Pop()
	assert.Equal(t, 0, v)
	assert.Equal(t, false, ok)
	c = s.Count()
	assert.Equal(t, 0, c)

	s.Push(10)
	s.Push(2)
	s.Pop()
	s.Push(3)
	s.Pop()
	c = s.Count()
	assert.Equal(t, 1, c)
	v, ok = s.Pop()
	assert.Equal(t, 10, v)
	assert.Equal(t, true, ok)
	c = s.Count()
	assert.Equal(t, 0, c)
	t.Log("empty chain stack", s)

	tests := []struct {
		name string
		cap  int
		ov   []opval[int]
		want []int
	}{
		{"test1", 2, []opval[int]{}, []int{}},
		{"test2", 2, []opval[int]{{"Push", 1}, {"Push", 2}}, []int{1, 2}},
		{"test3", 2, []opval[int]{{"Push", 1}, {"Push", 2}, {"Push", 3}}, []int{1, 2}},
		{"test4", 2, []opval[int]{{"Push", 1}, {op: "Pop"}, {op: "Pop"}, {"Push", 2}, {"Push", 3}}, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSequentialStack[int](tt.cap)
			for _, ov := range tt.ov {
				if ov.op == "Push" {
					s.Push(ov.val)
				} else {
					s.Pop()
				}
			}

			assert.Equal(t, true, s.Equal(tt.want))
		})
	}

	strtests := []struct {
		name string
		cap  int
		ov   []opval[string]
		want []string
	}{
		{"strtest1", 2, []opval[string]{}, []string{}},
		{"strtest2", 2, []opval[string]{{"Push", "1"}, {"Push", "2"}}, []string{"1", "2"}},
		{"strtest3", 2, []opval[string]{{"Push", "1"}, {"Push", "2"}, {"Push", "3"}}, []string{"1", "2"}},
		{"strtest4", 2, []opval[string]{{"Push", "1"}, {op: "Pop"}, {op: "Pop"}, {"Push", "2"}, {"Push", "3"}}, []string{"2", "3"}},
	}

	for _, tt := range strtests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSequentialStack[string](tt.cap)
			for _, ov := range tt.ov {
				if ov.op == "Push" {
					s.Push(ov.val)
				} else {
					s.Pop()
				}
			}

			assert.Equal(t, true, s.Equal(tt.want))
		})
	}
}

func TestChainStack(t *testing.T) {
	s := NewChainStack()
	s.Push(1)
	s.Push(2)
	c := s.Count()
	assert.Equal(t, 2, c)

	v, ok := s.Pop()
	assert.Equal(t, 2, v)
	assert.Equal(t, true, ok)
	c = s.Count()
	assert.Equal(t, 1, c)

	v, ok = s.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, true, ok)
	c = s.Count()
	assert.Equal(t, 0, c)

	v, ok = s.Pop()
	assert.Equal(t, 0, v)
	assert.Equal(t, false, ok)
	c = s.Count()
	assert.Equal(t, 0, c)

	s.Push(10)
	s.Push(2)
	s.Pop()
	s.Push(3)
	s.Pop()
	c = s.Count()
	assert.Equal(t, 1, c)
	v, ok = s.Pop()
	assert.Equal(t, 10, v)
	assert.Equal(t, true, ok)
	c = s.Count()
	assert.Equal(t, 0, c)
	t.Log("empty chain stack", s)
}

func TestExpEvaluationSeq(t *testing.T) {
	var tests = []struct {
		name  string
		exp   string
		wantV int
		wantE error
	}{
		{"test1", "", 0, nil},
		{"test2", "1+2", 3, nil},
		{"test3", "1+2*3/2-1", 3, nil},
		{"test4", "1+2*3/2-1*2+4%3", 3, nil},
		{"test5", "1+2*3/2-1*2+4%3+", 0, errInvalidExp},
		{"test6", "3*2+2-5*6", -22, nil},
		{"test7", "12*200", 2400, nil},
		{"test8", "12(200+2", 0, errInvalidExp},
		{"test9", "12/0", 0, errDivisionByZero},
		{"test10", "12%0", 0, errDivisionByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, e := ExpEvaluationSeq(tt.exp)
			t.Log(v, e)
			assert.Equal(t, tt.wantE, e)
			assert.Equal(t, tt.wantV, v)
		})
	}
}

func TestBracketMatch(t *testing.T) {
	tests := []struct {
		name string
		exp  string
		want bool
	}{
		// {"test1", "", true},
		// {"test2", "abc1234sdk", true},
		// {"test3", "[]", true},
		{"test4", "{}", true},
		{"test5", "()", true},
		{"test6", "()[]{}", true},
		{"test7", "([{()}])", true},
		{"test8", "([{()}])[[{{([])}}]]", true},
		{"test9", "([{()}])[[{{([])}}]][", false},
		{"test10", "(", false},
		{"test11", "a)c", false},
		{"test12", "ddf(1345", false},
		{"test13", "a[c", false},
		{"test14", "ddf]1345", false},
		{"test15", "a{c", false},
		{"test16", "ddf}1345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BracketMatch(tt.exp)
			assert.Equal(t, tt.want, r)
		})
	}
}

func TestBrowserFB(t *testing.T) {
	tests := []struct {
		name string
		ops  []webOp
		end1 []string
		end2 []string
	}{
		{"test1", []webOp{{}}, []string{"首页"}, []string{}},
		{"test2", []webOp{{"n", "web1"}}, []string{"首页", "web1"}, []string{}},
		{"test3", []webOp{{"n", "web1"}, {op: "b"}}, []string{"首页"}, []string{"web1"}},
		{"test4", []webOp{{"n", "web1"}, {op: "f"}}, []string{"首页", "web1"}, []string{}},
		{"test5", []webOp{{"n", "web1"}, {op: "b"}, {op: "f"}}, []string{"首页", "web1"}, []string{}},
		{"test6", []webOp{{"n", "web1"}, {"n", "web2"}, {"n", "web3"}, {op: "f"}, {op: "b"}, {op: "b"}, {op: "b"}}, []string{"首页"}, []string{"web3", "web2", "web1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r1, r2 := BrowserFB(tt.ops)
			assert.Equal(t, true, r1.Equal(tt.end1))
			assert.Equal(t, true, r2.Equal(tt.end2))
		})
	}
}
