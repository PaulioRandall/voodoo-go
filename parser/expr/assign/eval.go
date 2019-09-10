package assign

import (
	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Eval satisfies the Expr interface.
func (a assign) Eval(c ctx.Context) (value.Value, perror.Perror) {

	values, e := a.evalSrcExprs(c)
	if e != nil {
		return nil, e
	}

	return nil, a.doAssignments(c, values)
}

// evalSrcExprs evaluates the source expressions.
func (a assign) evalSrcExprs(c ctx.Context) ([]value.Value, perror.Perror) {
	targets := len(a.dst)
	values := []value.Value{}

	for i, src := range a.src {
		if i >= targets {
			return nil, a.noTarget(i)
		}

		v, e := src.Eval(c)
		if e != nil {
			return nil, e
		}

		values = append(values, v)
	}

	return values, nil
}

// doAssignments assigns the values to their identifiers.
func (a assign) doAssignments(c ctx.Context, values []value.Value) perror.Perror {
	var e perror.Perror
	targets := len(a.dst)
	idIndex := 0

	for i, v := range values {
		if idIndex >= targets {
			return a.noTarget(i)
		}

		idIndex, e = a.doAssign(c, v, idIndex)
		if e != nil {
			return e
		}
	}

	if idIndex < len(a.dst) {
		return a.noSource(idIndex)
	}

	return nil
}

// doAssign assigns the value to the one or many variables. The index to the
// next assignment target is returned.
func (a assign) doAssign(c ctx.Context, v value.Value, idIndex int) (int, perror.Perror) {

	switch {
	case v == nil:
		t := a.dst[idIndex].Text()
		delete(c.Vars, t)
		idIndex++
	case a.dst[idIndex].Kind() == token.TK_VOID:
		idIndex++
	default:
		return a.assignToID(c, v, idIndex)
	}

	return idIndex, nil
}

// assignToID assigns a value to an identifier.
func (a assign) assignToID(c ctx.Context, new value.Value, idIndex int) (int, perror.Perror) {
	t := a.dst[idIndex].Text()

	switch old := c.Vars[t]; {
	case old == nil, old.SameKind(new):
		c.Vars[t] = new
	default:
		return 0, a.kindMismatch(old, new, idIndex)
	}

	idIndex++
	return idIndex, nil
}

// kindMismatch returns a new Perror for when an identifier already has a value
// of a specific kind but the new value being assigned is the same or
// compatible.
func (a assign) kindMismatch(old, new value.Value, i int) perror.Perror {
	return perror.New(
		a.dst[i].Line(),
		[]int{
			a.src[i].Token().Start(),
			a.dst[i].Start(),
		},
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
		[]int{
			a.dst[i].Start(),
		},
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
		[]int{
			a.src[i].Token().Start(),
		},
		[]string{
			"No target identifier for expression result",
			"`" + a.src[i].String() + "` has no assignment target",
		},
	)
}
