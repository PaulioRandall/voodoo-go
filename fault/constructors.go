package fault

// Bug returns a new Fault with a message formatted
// to present a bug with this program.
func Bug(m string) Fault {
	return stdFault{
		msg:     m,
		errType: DevBug,
	}
}

// Paren returns a new Fault with a message formatted
// to present an error with a parenthesis within a scroll.
func Paren(m string) Fault {
	return stdFault{
		msg:     m,
		errType: Parenthesis,
	}
}

// Num returns a new Fault with a message formatted
// to present an error with a literal number within
// a scroll.
func Num(m string) Fault {
	return stdFault{
		msg:     m,
		errType: Number,
	}
}

// Func returns a new Fault with a message formatted
// to present an error with a function within a scroll.
func Func(m string) Fault {
	return stdFault{
		msg:     m,
		errType: Function,
	}
}

// Str returns a new Fault with a message formatted
// to present an error with a string literal within
// a scroll.
func Str(m string) Fault {
	return stdFault{
		msg:     m,
		errType: String,
	}
}

// Sym returns a new Fault with a message formatted
// to present an error with a symbol within
// a scroll. E.g. operators etc
func Sym(m string) Fault {
	return stdFault{
		msg:     m,
		errType: Symbol,
	}
}
