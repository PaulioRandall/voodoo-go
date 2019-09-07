package assign

import (
	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Eval satisfies the Expr interface.
func (a assign) Eval(c ctx.Context) (value.Value, perror.Perror) {
	var e perror.Perror
	targets := len(a.dst)
	di := 0

	for si, _ := range a.src {
		if di >= targets {
			return nil, a.noTarget(si)
		}

		di, e = a.evalAssign(c, si, di)
		if e != nil {
			return nil, e
		}
	}

	if di < len(a.dst) {
		return nil, a.noSource(di)
	}

	return nil, nil
}

// evalAssign evaluates a single source expression and assigns the result to the
// target identifiers; there may be more than one target in some cases such as
// the return from functions and spells.
func (a assign) evalAssign(c ctx.Context, si, di int) (int, perror.Perror) {
	v, e := a.src[si].Eval(c)
	if e != nil {
		return 0, e
	}

	return a.doAssign(c, v, di)
}

// doAssign assigns the value to the one or many variables. The index to the
// next assignment target is returned.
func (a assign) doAssign(c ctx.Context, v value.Value, di int) (int, perror.Perror) {
	var e perror.Perror

	switch {
	case v == nil:
		t := a.dst[di].Text()
		delete(c.Vars, t)
		di++
	case a.dst[di].Kind() == token.TK_VOID:
		di++
	default:
		di, e = a.assignToID(c, v, di)
	}

	return di, e
}

// assignToID assigns a value to an identifier.
func (a assign) assignToID(c ctx.Context, new value.Value, di int) (int, perror.Perror) {
	t := a.dst[di].Text()

	switch old := c.Vars[t]; {
	case old == nil, old.SameKind(new):
		c.Vars[t] = new
	default:
		return 0, a.kindMismatch(old, new, di)
	}

	di++
	return di, nil
}

// kindMismatch returns a new Perror for when an identifier already has a value
// of a specific kind but the new value being assigned is the same or
// compatible.
func (a assign) kindMismatch(old, new value.Value, i int) perror.Perror {
	return perror.New(
		a.dst[i].Line(),
		a.dst[i].Start(),
		[]string{
			"Expression result type does not match variable storage type",
			"Variable stores type `" + old.Kind().Name() + "`",
			"New value is of type `" + new.Kind().Name() + "`",
		},
	)
}

// noSource returns a new Perror for when an identifier was declared as a
// target for assignment but no source expression exists to supply a value.
func (a assign) noSource(i int) perror.Perror {
	return perror.New(
		a.dst[i].Line(),
		a.dst[i].Start(),
		[]string{
			"No source expression for assignment target",
			"`" + a.dst[i].Text() + "` declared but no value to assign",
		},
	)
}

// noTarget returns a new Perror for when an expression was declared as a
// source for an assignment but no target identifier exists to populate the
// result with.
func (a assign) noTarget(i int) perror.Perror {
	return perror.New(
		a.src[i].Token().Line(),
		a.src[i].Token().Start(),
		[]string{
			"No target identifier for expression result",
			"`" + a.src[i].String() + "` has no assignment target",
		},
	)
}
