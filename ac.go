package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Emit struct {
	S int
	E int
	W string
}

func newEmit(s, e int, w string) Emit {
	return Emit{s, e, w}
}

type state struct {
	children map[rune]*state
	failure  *state
	depth    int
	emits    []string

	once sync.Once
}

func newState() *state {
	return &state{children: make(map[rune]*state)}
}

func (st *state) nextState(c rune) *state {
	return st.children[c]
}

func (st *state) search(c rune) *state {
	nextState := st.nextState(c)
	for nextState == nil {
		st = st.failure

		if st == nil {
			return st
		}
		nextState = st.nextState(c)
	}

	return nextState
}

func (st *state) addChild(c rune) *state {
	nextState := st.nextState(c)
	if nextState == nil {
		nextState = newState()
		nextState.depth += 1
		st.children[c] = nextState
	}
	return nextState

}

type Trie struct {
	root *state
}

func NewTrie() *Trie {
	t := &Trie{root: newState()}
	return t
}

func (t *Trie) AddKeyword(kw string) {
	st := t.root
	for _, s := range kw {
		st = st.addChild(s)
	}
	st.emits = append(st.emits, kw)
}

func (t *Trie) Print() {
	//bfs
	queue := list.New()
	for s, c := range t.root.children {
		fmt.Printf("%c emits --> %v ", s, c.emits)
		if c.failure != nil {
			fmt.Printf("failure --> %p \n", c.failure)
		} else {
			fmt.Println()
		}
		queue.PushBack(c)
	}

	for {
		st := queue.Front()
		if st == nil {
			return
		}
		queue.Remove(st)
		for s, c := range st.Value.(*state).children {
			fmt.Printf("%c emits --> %v ", s, c.emits)
			if c.failure != nil {
				fmt.Printf("failure --> %p \n", c.failure)
			} else {
				fmt.Println()
			}
			queue.PushBack(c)
		}
	}
}

func (t *Trie) ParseText(text string) []Emit {
	t.root.once.Do(func() {
		t.buildFailureState()
	})

	st := t.root
	rs := []rune(text)
	var emits []Emit
	for pos := 0; pos < len(rs); pos++ {
		c := rs[pos]
		st = st.search(c)
		// nst := st.nextState(c)

		// for nst == nil {
		// 	st = st.failure
		// 	if st == nil {
		// 		nst = t.root
		// 		break
		// 	}
		// 	nst = st.nextState(c)
		// }
		// st = nst
		if st == nil {
			st = t.root
		}

		if len(st.emits) > 0 {
			emits = appendEmit(emits, pos, st.emits)
		}
	}

	return emits
}

func appendEmit(emits []Emit, ed int, w []string) []Emit {
	for _, s := range w {
		emits = append(emits, newEmit(ed-len([]rune(s))+1, ed, s))
	}
	return emits
}

func (t *Trie) buildFailureState() {
	queue := list.New()
	// level 1 failure to root
	for _, s := range t.root.children {
		s.failure = t.root
		queue.PushBack(s)
	}

	for {
		f := queue.Front()
		if f == nil {
			return
		}
		queue.Remove(f)
		st := f.Value.(*state)
		for c, nextSt := range st.children {
			queue.PushBack(nextSt)
			faiureState := st.failure
			nfaiureState := faiureState.nextState(c)
			for nfaiureState == nil {
				faiureState = faiureState.failure
				if faiureState == nil {
					nfaiureState = t.root
					break
				}
				nfaiureState = faiureState.nextState(c)
			}
			nextSt.failure = nfaiureState
			nextSt.emits = append(nextSt.emits, nfaiureState.emits...)
		}

	}
}
