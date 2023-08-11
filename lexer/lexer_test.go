package lexer

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	md := "# Heading1\n" +
		"### Heading3\n" +
		"text\n" +
		"## Heading2\n"
	expect := []Token{
		*tokenOf(THash1),
		*tokenOf(TSpace),
		textToken("Heading1"),
		*tokenOf(TBreak),
		*tokenOf(THash3),
		*tokenOf(TSpace),
		textToken("Heading3"),
		*tokenOf(TBreak),
		textToken("text"),
		*tokenOf(TBreak),
		*tokenOf(THash2),
		*tokenOf(TSpace),
		textToken("Heading2"),
		*tokenOf(TBreak),
	}

	got := Analyze(md)

	for i := range expect {
		if !compToken(got[i], expect[i]) {
			t.Fail()
		}
	}
}

func TestAnalyze_Escape(t *testing.T) {
	md := `##\##\\`
	expect := []Token{
		*tokenOf(THash2),
		textToken("#"),
		*tokenOf(THash1),
		textToken("\\"),
	}

	got := Analyze(md)

	for i := range expect {
		if !compToken(got[i], expect[i]) {
			t.Fail()
		}
	}
}

func TestAnalyze_OrderList(t *testing.T) {
	md := "1.hoge\n12. hoge"
	expect := []Token{
		*tokenOf(TOrderList),
		textToken("hoge"),
		*tokenOf(TBreak),
		*tokenOf(TOrderList),
		*tokenOf(TSpace),
		textToken("hoge"),
	}

	got := Analyze(md)

	for i := range expect {
		if !compToken(got[i], expect[i]) {
			t.Fail()
		}
	}
}

func textToken(str string) Token {
	return Token{
		Type:  TText,
		Value: str,
	}
}

func compToken(a, b Token) bool {
	if a.Type != b.Type {
		return false
	}
	if a.Type == TText {
		return b.Type == TText && a.Value == b.Value
	}
	return a.Type == b.Type
}
