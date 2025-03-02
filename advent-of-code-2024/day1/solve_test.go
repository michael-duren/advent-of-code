package day1

import "testing"

// parseCol calls day1.parseCol with a string that contains
// numbers, checking for valid return
func TestParseColWithNumbers(t *testing.T) {
	// arragne
	col := "123    "
	// act
	val, err := parseCol(col)
	// assert
	if err != nil || val != 123 {
		t.Fatalf(`parseCol("%s") = %q, %v`, col, val, err)
	}
}

// parseCol calls day1.parseCol with a string that contains
// only whitespace, checking for error
func TestParseColWithWhitespace(t *testing.T) {
	// arragne
	col := "    "
	// act
	val, err := parseCol(col)
	// assert
	if err == nil {
		t.Fatalf(`parseCol("%s") = %q, %s`, col, val, "there should be an error for just whitespace")
	}
}

// parseCol calls day1.parseCol with a string that contains
// only whitespace, checking for error
func TestSplitRow(t *testing.T) {
	// arragne
	row := "56208   95668"

	// act
	cols, err := splitRow(row)
	// assert
	if err != nil {
		t.Fatalf(`splitRow("%s") = %q, %s`, row, cols, err)
	}
	if cols[0] != "56208" || cols[1] != "95668" {
		t.Fatalf(`splitRow("%s") = %q, %s`, row, cols, "didn't get correct values in slice")
	}
}
