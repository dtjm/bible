package ref

// Book is a book of the Bible
type Book int

const (
	nullBook Book = iota
	Genesis
	Exodus
	Leviticus
	Numbers
	Deuteronomy
	Joshua
	Judges
	Ruth
	Samuel1
	Samuel2
	Kings1
	Kings2
	Chronicles1
	Chronicles2
	Ezra
	Nehemiah
	Esther
	Job
	Psalm
	Proverbs
	Ecclesiastes
	SongOfSolomon
	Isaiah
	Jeremiah
	Lamentations
	Ezekiel
	Daniel
	Hosea
	Joel
	Amos
	Obadiah
	Jonah
	Micah
	Nahum
	Habakkuk
	Zephaniah
	Haggai
	Zechariah
	Malachi
	Matthew
	Mark
	Luke
	John
	Acts
	Romans
	Corinthians1
	Corinthians2
	Galatians
	Ephesians
	Philippians
	Colossians
	Thessalonians1
	Thessalonians2
	Timothy1
	Timothy2
	Titus
	Philemon
	Hebrews
	James
	Peter1
	Peter2
	John1
	John2
	John3
	Jude
	Revelation
)

var (
	numChapters = map[Book]int{
		Genesis:        50,
		Exodus:         40,
		Leviticus:      27,
		Numbers:        36,
		Deuteronomy:    34,
		Joshua:         24,
		Judges:         21,
		Ruth:           4,
		Samuel1:        31,
		Samuel2:        24,
		Kings1:         22,
		Kings2:         25,
		Chronicles1:    29,
		Chronicles2:    36,
		Ezra:           10,
		Nehemiah:       13,
		Esther:         10,
		Job:            42,
		Psalm:          150,
		Proverbs:       31,
		Ecclesiastes:   12,
		SongOfSolomon:  8,
		Isaiah:         66,
		Jeremiah:       52,
		Lamentations:   5,
		Ezekiel:        48,
		Daniel:         12,
		Hosea:          14,
		Joel:           3,
		Amos:           9,
		Obadiah:        1,
		Jonah:          4,
		Micah:          7,
		Nahum:          3,
		Habakkuk:       3,
		Zephaniah:      3,
		Haggai:         2,
		Zechariah:      14,
		Malachi:        4,
		Matthew:        28,
		Mark:           16,
		Luke:           24,
		John:           21,
		Acts:           28,
		Romans:         16,
		Corinthians1:   16,
		Corinthians2:   13,
		Galatians:      6,
		Ephesians:      6,
		Philippians:    4,
		Colossians:     4,
		Thessalonians1: 5,
		Thessalonians2: 3,
		Timothy1:       6,
		Timothy2:       4,
		Titus:          3,
		Philemon:       1,
		Hebrews:        13,
		James:          5,
		Peter1:         5,
		Peter2:         3,
		John1:          5,
		John2:          1,
		John3:          1,
		Jude:           1,
		Revelation:     22,
	}
)

// Next returns the next book of the Bible, wrapping around to the beginning
func (b Book) Next() Book {
	n := b + 1
	if n > Revelation {
		return Genesis
	}

	return n
}

func (b Book) String() string {
	switch b {
	case Genesis:
		return "Genesis"
	case Exodus:
		return "Exodus"
	case Leviticus:
		return "Leviticus"
	case Numbers:
		return "Numbers"
	case Deuteronomy:
		return "Deuteronomy"
	case Joshua:
		return "Joshua"
	case Judges:
		return "Judges"
	case Ruth:
		return "Ruth"
	case Samuel1:
		return "1 Samuel"
	case Samuel2:
		return "2 Samuel"
	case Kings1:
		return "1 Kings"
	case Kings2:
		return "2 Kings"
	case Chronicles1:
		return "1 Chronicles"
	case Chronicles2:
		return "2 Chronicles"
	case Ezra:
		return "Ezra"
	case Nehemiah:
		return "Nehemiah"
	case Esther:
		return "Esther"
	case Job:
		return "Job"
	case Psalm:
		return "Psalm"
	case Proverbs:
		return "Proverbs"
	case Ecclesiastes:
		return "Ecclesiastes"
	case SongOfSolomon:
		return "SongOfSolomon"
	case Isaiah:
		return "Isaiah"
	case Jeremiah:
		return "Jeremiah"
	case Lamentations:
		return "Lamentations"
	case Ezekiel:
		return "Ezekiel"
	case Daniel:
		return "Daniel"
	case Hosea:
		return "Hosea"
	case Joel:
		return "Joel"
	case Amos:
		return "Amos"
	case Obadiah:
		return "Obadiah"
	case Jonah:
		return "Jonah"
	case Micah:
		return "Micah"
	case Nahum:
		return "Nahum"
	case Habakkuk:
		return "Habakkuk"
	case Zephaniah:
		return "Zephaniah"
	case Haggai:
		return "Haggai"
	case Zechariah:
		return "Zechariah"
	case Malachi:
		return "Malachi"
	case Matthew:
		return "Matthew"
	case Mark:
		return "Mark"
	case Luke:
		return "Luke"
	case John:
		return "John"
	case Acts:
		return "Acts"
	case Romans:
		return "Romans"
	case Corinthians1:
		return "1 Corinthians"
	case Corinthians2:
		return "2 Corinthians"
	case Galatians:
		return "Galatians"
	case Ephesians:
		return "Ephesians"
	case Philippians:
		return "Philippians"
	case Colossians:
		return "Colossians"
	case Thessalonians1:
		return "1 Thessalonians"
	case Thessalonians2:
		return "2 Thessalonians"
	case Timothy1:
		return "1 Timothy"
	case Timothy2:
		return "2 Timothy"
	case Titus:
		return "Titus"
	case Philemon:
		return "Philemon"
	case Hebrews:
		return "Hebrews"
	case James:
		return "James"
	case Peter1:
		return "1 Peter"
	case Peter2:
		return "2 Peter"
	case John1:
		return "1 John"
	case John2:
		return "2 John"
	case John3:
		return "3 John"
	case Jude:
		return "Jude"
	case Revelation:
		return "Revelation"
	}

	return ""
}
