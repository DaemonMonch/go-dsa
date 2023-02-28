package godsa

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if ok, v := s.Pop(); ok {
		if v.(int) != 3 {
			t.Error(v.(int))
		}
		return
	}
	if ok, v := s.Pop(); ok {
		if v.(int) != 2 {
			t.Error(v.(int))
		}
		return
	}
	if ok, v := s.Pop(); ok {
		if v.(int) != 1 {
			t.Error(v.(int))
		}
		return
	}
	if ok, v := s.Pop(); ok {
		t.Error(v.(int))
		return
	}
	t.Error(s.Len())
}
