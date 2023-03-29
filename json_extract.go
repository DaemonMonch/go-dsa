package dsl

import (
	"errors"
	"fmt"
	"io"
	"unicode"
)

type JsonExtractor interface {
	FindKey(k string) JsonExtractor

	Val() []interface{}
}

type Ctx struct {
	path []string
}

type JEcallback interface {
	OnScalar(ctx Ctx, k string, v interface{}) bool
}

type cb struct{}

func (cb) OnScalar(ctx Ctx, k string, v interface{}) bool {
	fmt.Println(ctx.path, k, v)
	return true
}

type JE struct {
	curOffset int
	curKey    string

	cb JEcallback
	rs []rune
}

func (je *JE) parse(k string, json string) {

	rs := []rune(json)
	je.cb = cb{}
	for {
		err := je.skipBlank(rs)
		if err != nil {
			return
		}
		c := rs[je.curOffset]
		// fmt.Printf("%c \n", c)
		if c == rune('{') {
			je.parseObj(rs)
		} else if c == rune('}') {
			je.curOffset++
			// ctx.path = ctx.path[0:len()]
		} else if c == rune('[') {
			je.parseArray(rs)
		} else if c == rune(']') {
			//
			je.curOffset++
		} else if c == rune(',') {
			je.curOffset++
		} else {
			err = je.parseKey(rs)
			if err != nil {
				return
			}
			je.curOffset++
		}
	}
}

func (je *JE) parseKey(rs []rune) error {
	if je.curKey == "" {
		return errors.New("invalid format")
	}

	err := je.skipBlank(rs)
	if err != nil {
		return fmt.Errorf("key err offset %d ", je.curOffset)
	}

	c := rs[je.curOffset]
	if c != rune('"') {
		return fmt.Errorf("key err offset %d ", je.curOffset)
	}

	je.curOffset++
	stk := je.curOffset
	err = je.find(rs, rune('"'))
	if err != nil {
		return fmt.Errorf("key err offset %d ", je.curOffset)
	}
	keyend := je.curOffset
	err = je.find(rs, rune(':'))

	if err != nil {
		return fmt.Errorf("key err offset %d ", je.curOffset)
	}

	je.curKey = string(rs[stk:keyend])
	// fmt.Println("key ", je.curKey, " start ", stk, "end ", keyend)
	// je.keys = append(je.keys, je.curKey)
	je.curOffset++
	return je.parseVal(rs)
}

func (je *JE) parseVal(rs []rune) error {
	err := je.skipBlank(rs)
	if err != nil {
		return fmt.Errorf("value err offset %d ", je.curOffset)
	}
	c := rs[je.curOffset]
	if c == rune('"') { //str val
		// st := je.curOffset
		je.curOffset++
		err := je.find(rs, rune('"'))
		if err != nil {
			return fmt.Errorf("value err offset %d ", je.curOffset)
		}
		je.curOffset++
	} else if unicode.IsDigit(c) || c == rune('t') || c == rune('f') || c == rune('n') {
		err := je.find(rs, rune(','), rune('}'))
		if err != nil {
			return fmt.Errorf("value err offset %d ", je.curOffset)
		}
	} else {

	}

	return nil
}
func (je *JE) parseObj(rs []rune) error {
	je.curOffset++
	if je.curKey == "" {
		je.curKey = "$"
	}
	return je.parseKey(rs)
}

func (je *JE) parseArray(rs []rune) error {
	je.curOffset++
	return nil
}

func (je *JE) find(rs []rune, cs ...rune) error {
	if je.curOffset >= len(rs) {
		return io.EOF
	}
	for {
		for _, c := range cs {
			if rs[je.curOffset] == c {
				return nil
			}
		}

		je.curOffset++
		if je.curOffset >= len(rs) {
			return io.EOF
		}
	}
}

func (je *JE) skipBlank(rs []rune) error {
	if je.curOffset >= len(rs) {
		return io.EOF
	}
	for rs[je.curOffset] == rune(' ') || rs[je.curOffset] == rune('\r') || rs[je.curOffset] == rune('\n') || rs[je.curOffset] == rune('\t') {
		je.curOffset++
		if je.curOffset >= len(rs) {
			return io.EOF
		}
	}
	return nil
}
