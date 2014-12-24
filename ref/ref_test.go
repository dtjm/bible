package ref

import (
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		s string
		r Ref
	}{
		{"gen 1", Ref{book: Genesis, chapter: 1}},
		{"1timo", Ref{book: Timothy1}},
		{"1 John 3", Ref{book: John1, chapter: 3}},
	}

	for _, c := range cases {
		pr, err := Parse(c.s)
		if err != nil {
			t.Errorf("Parse error: %q", err)
			continue
		}

		if c.r != pr {
			t.Errorf("parsed ref: %+v, expected %+v", pr, c.r)
		}
	}
}

func TestBookNext(t *testing.T) {
	cases := []struct {
		book, next Book
	}{
		{Genesis, Exodus},
		{Revelation, Genesis},
		{Malachi, Matthew},
	}

	for _, c := range cases {
		if c.book.Next() != c.next {
			t.Errorf("(%s).Next -> %s, wanted %s ",
				c.book.String(), c.book.Next().String(), c.next.String())
		}
	}
}

func TestChapterNext(t *testing.T) {
	cases := []struct {
		ref, next Ref
	}{
		{
			Ref{book: Genesis, chapter: 1},
			Ref{book: Genesis, chapter: 2},
		},
		{
			Ref{book: Genesis, chapter: 50},
			Ref{book: Exodus, chapter: 1},
		},
	}

	for _, c := range cases {
		nextRef := c.ref.NextChapter()
		if nextRef != c.next {
			t.Errorf("(%v).NextChapter -> %v, wanted %v",
				c.ref, nextRef, c.next)
		}
	}
}
