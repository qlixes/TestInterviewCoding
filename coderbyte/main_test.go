package coderbyte

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
