package operation

// TODO: Just wanted to get the idea down somewhere.
// TODO: Do we need a Moment?
// TODO: How will it be used?
// TODO: Could it be used to store meta-information about
// TODO: the running scroll?
// TODO: Could a Moment just be a single immutable context
// TODO: in an ever-changing series of moments? Might be
// TODO: inuitive to think to think about execution in this way.
//
// OFF-TOPIC: Mis-spelled `moment` as `mement`. I'm now
// defining a mement as a single meme in the context of
// an evolving meme.
//
// Moment represents an instant in time bewteen instruction
// executions. An instruction execution is considered a zero
// time operation. Of course it's no where near but by
// allowing me to think about the execution of instructions
// as indivisable units of activity, I can reason about
// the behaviour of it all in a manner that fits in my head.
type Moment interface {
}
