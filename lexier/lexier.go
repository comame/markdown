package lexier

import "regexp"

type Token struct {
	Type  TokenType
	Value string
}

// デバッグしやすいように、iota ではなくトークンをよく表す文字列を用いる
type TokenType string

const (
	THash1 TokenType = "#"
	THash2 TokenType = "##"
	THash3 TokenType = "###"
	THash4 TokenType = "####"
	THash5 TokenType = "#####"
	THash6 TokenType = "######"

	TAsterisk1 TokenType = "*"
	TAsterisk2 TokenType = "**"

	TUnderscore1 TokenType = "_"
	TUnderscore2 TokenType = "__"

	TSquareBracketLeft  TokenType = "["
	TSquareBracketRight TokenType = "]"

	TRoundBracketLeft  TokenType = "("
	TRoundBracketRight TokenType = ")"

	TAngleBracketLeft  TokenType = "<"
	TAngleBracketRight TokenType = ">"

	TExclamation TokenType = "!"

	TBackquote1 TokenType = "`"
	TBackquote3 TokenType = "```"

	THyphen    TokenType = "-"
	TOrderList TokenType = "1."

	TPipe  TokenType = "|"
	TColon TokenType = ":"

	TSpace TokenType = " "
	TBreak TokenType = "\n"

	TText TokenType = "text"
)

func Analyze(md string) []Token {
	var tokens []Token

	var stringToken *Token

	escaped := false

	i := 0
	for {
		if i >= len(md) {
			break
		}

		// エスケープされていたら、通常の文字 TText とみなす
		if escaped {
			if stringToken == nil {
				stringToken = tokenOf(TText)
			}
			stringToken.Value = string(md[i])
			escaped = false
			i += 1
			continue
		}

		if md[i] == '\\' {
			escaped = true
			i += 1
			continue
		}

		var symbolToken *Token

		switch {
		case substr(md, i, 6) == "######":
			symbolToken = tokenOf(THash6)
			i += 6
		case substr(md, i, 5) == "#####":
			symbolToken = tokenOf(THash5)
			i += 5
		case substr(md, i, 4) == "####":
			symbolToken = tokenOf(THash4)
			i += 4
		case substr(md, i, 3) == "###":
			symbolToken = tokenOf(THash3)
			i += 3
		case substr(md, i, 2) == "##":
			symbolToken = tokenOf(THash2)
			i += 2
		case substr(md, i, 1) == "#":
			symbolToken = tokenOf(THash1)
			i += 1

		case substr(md, i, 2) == "**":
			symbolToken = tokenOf(TAsterisk2)
			i += 2
		case substr(md, i, 2) == "*":
			symbolToken = tokenOf(TAsterisk1)
			i += 1

		case substr(md, i, 2) == "__":
			symbolToken = tokenOf(TUnderscore2)
			i += 2
		case substr(md, i, 1) == "_":
			symbolToken = tokenOf(TUnderscore1)
			i += 1

		case md[i] == '[':
			symbolToken = tokenOf(TSquareBracketLeft)
			i += 1
		case md[i] == ']':
			symbolToken = tokenOf(TSquareBracketRight)
			i += 1

		case md[i] == '(':
			symbolToken = tokenOf(TRoundBracketLeft)
			i += 1
		case md[i] == ')':
			symbolToken = tokenOf(TRoundBracketRight)
			i += 1

		case md[i] == '<':
			symbolToken = tokenOf(TAngleBracketLeft)
			i += 1
		case md[i] == '>':
			symbolToken = tokenOf(TAngleBracketRight)
			i += 1

		case md[i] == '!':
			symbolToken = tokenOf(TExclamation)
			i += 1

		case substr(md, i, 3) == "```":
			symbolToken = tokenOf(TBackquote3)
			i += 3
		case substr(md, i, 1) == "`":
			symbolToken = tokenOf(TBackquote1)
			i += 1

		case md[i] == '-':
			symbolToken = tokenOf(THyphen)
			i += 1

		case md[i] == '|':
			symbolToken = tokenOf(TPipe)
			i += 1
		case md[i] == ':':
			symbolToken = tokenOf(TColon)
			i += 1

		case md[i] == ' ':
			symbolToken = tokenOf(TSpace)
			i += 1
		case md[i] == '\n':
			symbolToken = tokenOf(TBreak)
			i += 1
		}

		// 記号トークンだったら、直前の TText を閉じてから、記号トークンを配列に入れる
		if symbolToken != nil {
			if stringToken != nil {
				tokens = append(tokens, *stringToken)
				stringToken = nil
			}

			tokens = append(tokens, *symbolToken)
			continue
		}

		// 番号付きリスト TOrderList の数字部分は適当な数字が入るので、正規表現で調べる
		orderListReg := regexp.MustCompile(`^\d+\.`)
		orderListMatch := orderListReg.FindString(md[i:])
		if orderListMatch != "" {
			tokens = append(tokens, *tokenOf(TOrderList))
			i += len(orderListMatch)
			continue
		}

		// 文字列トークン TText は前回のトークンが TText だったとき、別トークンとはしない
		if stringToken == nil {
			stringToken = tokenOf(TText)
		}
		stringToken.Value += string(md[i])

		i += 1
	}

	if stringToken != nil {
		tokens = append(tokens, *stringToken)
	}

	return tokens
}

func tokenOf(t TokenType) *Token {
	return &Token{
		Type: t,
	}
}

func substr(str string, begin, length int) string {
	l := len(str)
	if l < begin {
		return ""
	}
	if l < begin+length {
		return str[begin:l]
	}
	return str[begin : begin+length]
}
