package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanSpell(t *testing.T, in string, exp, expErr *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanSpell(r)
	token.AssertToken(t, exp, tk)
	token.AssertToken(t, expErr, errTk)
}

func dummySpellToken(end int, s string) token.Token {
	return token.DummyToken(0, 0, end, s, token.TT_SPELL)
}

func TestScanSpell_1(t *testing.T) {
	in := `@Println`
	exp := dummySpellToken(8, `@Println`)
	doTestScanSpell(t, in, &exp, nil)
}

func TestScanSpell_2(t *testing.T) {
	in := `@a__12__xy__`
	exp := dummySpellToken(12, `@a__12__xy__`)
	doTestScanSpell(t, in, &exp, nil)
}

func TestScanSpell_3(t *testing.T) {
	in := `@Println(msg)`
	exp := dummySpellToken(8, `@Println`)
	doTestScanSpell(t, in, &exp, nil)
}

func TestScanSpell_4(t *testing.T) {
	in := `@2`
	expErr := token.ErrDummyToken(0, 0, 1)
	doTestScanSpell(t, in, nil, &expErr)
}
