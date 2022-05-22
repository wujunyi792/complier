package util

import (
	"bufio"
	"compiler/lexer/define"
	"container/list"
	"io"
)

type Stream struct {
	scanner    *bufio.Scanner
	queueCache *list.List
	endToken   string
	isEnd      bool
	line       int
	column     int
}

func NewStream(r io.Reader, et string) *Stream {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	return &Stream{scanner: s, queueCache: list.New(), endToken: et, isEnd: false, line: 1, column: 1}
}

func (s *Stream) GetLine() int {
	return s.line
}

func (s *Stream) GetColumn() int {
	return s.column
}

func (s *Stream) Next() string {
	char := ""
	if s.queueCache.Len() != 0 {
		e := s.queueCache.Front()
		char = s.queueCache.Remove(e).(string)
	} else if s.scanner.Scan() {
		char = s.scanner.Text()
	} else {
		s.isEnd = true

		char = s.endToken
	}
	if define.IsNewLine(char) {
		s.line += 1
		s.column = 0
	}
	s.column += 1
	return char
}

func (s *Stream) HasNext() bool {
	if s.queueCache.Len() != 0 {
		return true
	}

	if s.scanner.Scan() {
		s.queueCache.PushBack(s.scanner.Text())
		return true
	}

	if !s.isEnd {
		return true
	}

	return false
}

func (s *Stream) Peek() string {
	if s.queueCache.Len() != 0 {
		return s.queueCache.Front().Value.(string)
	}

	if s.scanner.Scan() {
		e := s.scanner.Text()
		s.queueCache.PushBack(e)
		return e
	}

	return s.endToken
}

func (s *Stream) PutBack(e string) {
	s.queueCache.PushFront(e)
	s.column -= 1
}
