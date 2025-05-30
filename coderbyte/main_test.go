package main

import "testing"

func TestLongestWordCase1(t *testing.T) {
	got := LongestWord("hello world")
	want := "hello"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase2(t *testing.T) {
	got := LongestWord("this is some sort of sentence")
	want := "sentence"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase3(t *testing.T) {
	got := LongestWord("longest word!!")
	want := "longest"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase4(t *testing.T) {
	got := LongestWord("coderbyte")
	want := "coderbyte"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase5(t *testing.T) {
	got := LongestWord("a beautiful sentence^&!")
	want := "beautiful"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase6(t *testing.T) {
	got := LongestWord("oxford press")
	want := "oxford"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase7(t *testing.T) {
	got := LongestWord("123456789 98765432")
	want := "123456789"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase8(t *testing.T) {
	got := LongestWord("letter after letter!!")
	want := "letter"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase9(t *testing.T) {
	got := LongestWord("a b c dee")
	want := "dee"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLongestWordCase10(t *testing.T) {
	got := LongestWord("a confusing /:sentence:/[ this is not!!!!!!!~")
	want := "confusing"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFirstFactorial01(t *testing.T) {
	got := FirstFactorial(1)
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial02(t *testing.T) {
	got := FirstFactorial(2)
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial03(t *testing.T) {
	got := FirstFactorial(3)
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial04(t *testing.T) {
	got := FirstFactorial(4)
	want := 24

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial05(t *testing.T) {
	got := FirstFactorial(5)
	want := 120

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial06(t *testing.T) {
	got := FirstFactorial(6)
	want := 720

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial07(t *testing.T) {
	got := FirstFactorial(7)
	want := 5040

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial08(t *testing.T) {
	got := FirstFactorial(8)
	want := 40320

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial09(t *testing.T) {
	got := FirstFactorial(9)
	want := 362880

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstFactorial10(t *testing.T) {
	got := FirstFactorial(10)
	want := 3628800

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstReserve01(t *testing.T) {
	got := FirstReverse("I Love Coding")
	want := "gnidoC evoL I"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve02(t *testing.T) {
	got := FirstReverse("h333llLo")
	want := "oLll333h"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve03(t *testing.T) {
	got := FirstReverse("Yo0")
	want := "0oY"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve04(t *testing.T) {
	got := FirstReverse("thisiscool")
	want := "loocsisiht"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve05(t *testing.T) {
	got := FirstReverse("commacomma!")
	want := "!ammocammoc"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve06(t *testing.T) {
	got := FirstReverse("123456789")
	want := "987654321"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve07(t *testing.T) {
	got := FirstReverse("lettersz!23z")
	want := "z32!zsrettel"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve08(t *testing.T) {
	got := FirstReverse("aq")
	want := "qa"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFirstReserve09(t *testing.T) {
	got := FirstReverse("b")
	want := "b"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
