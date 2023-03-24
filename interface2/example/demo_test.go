package example

import "testing"

type fakeRepo struct {
}

func (fakeRepo) QueryLang(uint) Language {
	return Language{
		ID:   2,
		Name: "Java",
	}
}

func TestHandler(t *testing.T) {
	h := NewHandler(fakeRepo{})
	s := h.Do(1)

	want := "Java language"

	if s != want {
		t.Errorf("want %q but got %s", want, s)
	}
}
