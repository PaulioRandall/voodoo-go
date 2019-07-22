package instruction

// ContextStack represents all contexts that currently exist
// at any moment in execution time. The context at the top is
// the currently active one. The second being the one that
// was active before top was created. The third being the that
// was active before the second was created, you get the
// picture. The bottom of the stack represents the context
// created for the scroll itself; the scroll finishes when
// this bottom context is dies.
type ContextStack []Context
