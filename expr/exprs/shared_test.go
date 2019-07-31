package exprs

import (
	"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/token"
)

func dummyToken(val string) token.Token {
	return token.Token{
		Val: val,
	}
}

func dummyNumber(n string) ctx.Expression {
	num := Number{
		Number: dummyToken(n),
	}
	expr := ctx.Expression(num)
	return expr
}
