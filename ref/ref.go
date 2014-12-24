package ref

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// Ref is a reference to a Bible passage
type Ref struct {
	book           Book
	chapter, verse int
}

var bookRegex = map[*regexp.Regexp]Book{
	regexp.MustCompile("(?i)^ge\\w*\\s*(\\d+)?"):      Genesis,
	regexp.MustCompile("(?i)^ex\\w*\\s*(\\d+)?"):      Exodus,
	regexp.MustCompile("(?i)^le\\w*\\s*(\\d+)?"):      Leviticus,
	regexp.MustCompile("(?i)^nu\\w*\\s*(\\d+)?"):      Numbers,
	regexp.MustCompile("(?i)^de\\w*\\s*(\\d+)?"):      Deuteronomy,
	regexp.MustCompile("(?i)^jos\\w*\\s*(\\d+)?"):     Joshua,
	regexp.MustCompile("(?i)^judg\\w*\\s*(\\d+)?"):    Judges,
	regexp.MustCompile("(?i)^ru\\w*\\s*(\\d+)?"):      Ruth,
	regexp.MustCompile("(?i)^1\\s?sa\\w*\\s*(\\d+)?"): Samuel1,
	regexp.MustCompile("(?i)^2\\s?sa\\w*\\s*(\\d+)?"): Samuel2,
	regexp.MustCompile("(?i)^1\\s?ki\\w*\\s*(\\d+)?"): Kings1,
	regexp.MustCompile("(?i)^2\\s?ki\\w*\\s*(\\d+)?"): Kings2,
	regexp.MustCompile("(?i)^1\\s?ch\\w*\\s*(\\d+)?"): Chronicles1,
	regexp.MustCompile("(?i)^2\\s?ch\\w*\\s*(\\d+)?"): Chronicles2,
	regexp.MustCompile("(?i)^ez\\w*\\s*(\\d+)?"):      Ezra,
	regexp.MustCompile("(?i)^ne\\w*\\s*(\\d+)?"):      Nehemiah,
	regexp.MustCompile("(?i)^es\\w*\\s*(\\d+)?"):      Esther,
	regexp.MustCompile("(?i)^job\\w*\\s*(\\d+)?"):     Job,
	regexp.MustCompile("(?i)^ps\\w*\\s*(\\d+)?"):      Psalm,
	regexp.MustCompile("(?i)^pr\\w*\\s*(\\d+)?"):      Proverbs,
	regexp.MustCompile("(?i)^ec\\w*\\s*(\\d+)?"):      Ecclesiastes,
	regexp.MustCompile("(?i)^so\\w*\\s*(\\d+)?"):      SongOfSolomon,
	regexp.MustCompile("(?i)^is\\w*\\s*(\\d+)?"):      Isaiah,
	regexp.MustCompile("(?i)^je\\w*\\s*(\\d+)?"):      Jeremiah,
	regexp.MustCompile("(?i)^la\\w*\\s*(\\d+)?"):      Lamentations,
	regexp.MustCompile("(?i)^ez\\w*\\s*(\\d+)?"):      Ezekiel,
	regexp.MustCompile("(?i)^da\\w*\\s*(\\d+)?"):      Daniel,
	regexp.MustCompile("(?i)^ho\\w*\\s*(\\d+)?"):      Hosea,
	regexp.MustCompile("(?i)^joe\\w*\\s*(\\d+)?"):     Joel,
	regexp.MustCompile("(?i)^am\\w*\\s*(\\d+)?"):      Amos,
	regexp.MustCompile("(?i)^ob\\w*\\s*(\\d+)?"):      Obadiah,
	regexp.MustCompile("(?i)^jon\\w*\\s*(\\d+)?"):     Jonah,
	regexp.MustCompile("(?i)^mi\\w*\\s*(\\d+)?"):      Micah,
	regexp.MustCompile("(?i)^na\\w*\\s*(\\d+)?"):      Nahum,
	regexp.MustCompile("(?i)^ha\\w*\\s*(\\d+)?"):      Habakkuk,
	regexp.MustCompile("(?i)^ze\\w*\\s*(\\d+)?"):      Zephaniah,
	regexp.MustCompile("(?i)^ha\\w*\\s*(\\d+)?"):      Haggai,
	regexp.MustCompile("(?i)^ze\\w*\\s*(\\d+)?"):      Zechariah,
	regexp.MustCompile("(?i)^mal\\w*\\s*(\\d+)?"):     Malachi,
	regexp.MustCompile("(?i)^mat\\w*\\s*(\\d+)?"):     Matthew,
	regexp.MustCompile("(?i)^mar\\w*\\s*(\\d+)?"):     Mark,
	regexp.MustCompile("(?i)^lu\\w*\\s*(\\d+)?"):      Luke,
	regexp.MustCompile("(?i)^joh\\w*\\s*(\\d+)?"):     John,
	regexp.MustCompile("(?i)^ac\\w*\\s*(\\d+)?"):      Acts,
	regexp.MustCompile("(?i)^ro\\w*\\s*(\\d+)?"):      Romans,
	regexp.MustCompile("(?i)^1\\s?co\\w*\\s*(\\d+)?"): Corinthians1,
	regexp.MustCompile("(?i)^2\\s?co\\w*\\s*(\\d+)?"): Corinthians2,
	regexp.MustCompile("(?i)^ga\\w*\\s*(\\d+)?"):      Galatians,
	regexp.MustCompile("(?i)^ep\\w*\\s*(\\d+)?"):      Ephesians,
	regexp.MustCompile("(?i)^ph\\w*\\s*(\\d+)?"):      Philippians,
	regexp.MustCompile("(?i)^co\\w*\\s*(\\d+)?"):      Colossians,
	regexp.MustCompile("(?i)^1\\s?th\\w*\\s*(\\d+)?"): Thessalonians1,
	regexp.MustCompile("(?i)^2\\s?th\\w*\\s*(\\d+)?"): Thessalonians2,
	regexp.MustCompile("(?i)^1\\s?ti\\w*\\s*(\\d+)?"): Timothy1,
	regexp.MustCompile("(?i)^2\\s?ti\\w*\\s*(\\d+)?"): Timothy2,
	regexp.MustCompile("(?i)^ti\\w*\\s*(\\d+)?"):      Titus,
	regexp.MustCompile("(?i)^ph\\w*\\s*(\\d+)?"):      Philemon,
	regexp.MustCompile("(?i)^he\\w*\\s*(\\d+)?"):      Hebrews,
	regexp.MustCompile("(?i)^ja\\w*\\s*(\\d+)?"):      James,
	regexp.MustCompile("(?i)^1\\s?pe\\w*\\s*(\\d+)?"): Peter1,
	regexp.MustCompile("(?i)^2\\s?pe\\w*\\s*(\\d+)?"): Peter2,
	regexp.MustCompile("(?i)^1\\s?jo\\w*\\s*(\\d+)?"): John1,
	regexp.MustCompile("(?i)^2\\s?jo\\w*\\s*(\\d+)?"): John2,
	regexp.MustCompile("(?i)^3\\s?jo\\w*\\s*(\\d+)?"): John3,
	regexp.MustCompile("(?i)^jude\\w*\\s*(\\d+)?"):    Jude,
	regexp.MustCompile("(?i)^re\\w*\\s*(\\d+)?"):      Revelation,
}

// Parse takes a passage reference and returns a Ref object
func Parse(s string) (*Ref, error) {
	book := nullBook
	chapter := 0
	var err error
	for re, b := range bookRegex {
		if matches := re.FindStringSubmatch(s); len(matches) > 0 {
			log.Printf("got matches for %q: %q", re.String(), matches)
			book = b
			if len(matches) > 1 && matches[1] != "" {
				chapter, err = strconv.Atoi(matches[1])
				if err != nil {
					return &Ref{}, err
				}
			}
		}
	}

	if book == nullBook {
		return &Ref{}, fmt.Errorf("Error parsing ref string: %q", s)
	}

	return &Ref{
		book:    book,
		chapter: chapter,
	}, nil
}

// Book returns the Book
func (r *Ref) Book() Book {
	return r.book
}

func (r *Ref) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(r.book.String())

	if r.chapter > 0 {
		buf.WriteString(fmt.Sprintf(" %d", r.chapter))
	}

	if r.verse > 0 {
		buf.WriteString(fmt.Sprintf(":%d", r.verse))
	}

	return buf.String()
}

// NextChapter returns the next chapter for a given reference
func (r *Ref) NextChapter() *Ref {
	nextRef := Ref{
		book:    r.book,
		chapter: r.chapter + 1,
	}

	if nextRef.chapter > numChapters[r.book] {
		nextRef.book = r.book.Next()
		nextRef.chapter = 1
	}

	return &nextRef
}
