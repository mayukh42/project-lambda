package lambda

// reducing verbosity
type Any interface{}
type List []Any

// unary function: f(x) = y
type UnaryF func(Any) Any
type PredicateF func(Any) bool

// UnaryF :> PredicateF

// binary function: f(x, y) = z
type BinaryF func(Any, Any) Any

/** Map lambda function
 * returns a list of elements each transformed by the given function
 */
func Map(xs List, fn UnaryF) List {
    if len(xs) == 0 {
        return xs
    }
    ys := make(List, len(xs))
    for i, x := range xs {
        ys[i] = fn(x)
    }
    return ys
}

/** Reduce lambda function
 * returns accumulated value of elements by sliding through a collection and applying
 * the given function
 */
func Reduce(xs List, acc Any, fn BinaryF) Any {
    if len(xs) == 0 {
        return acc
    }
    for _, x := range xs {
        acc = fn(x, acc)
    }
    return acc
}

/** Filter lambda function
 * returns a filtered list of elements that satisfy the given predicate
 */
func Filter(xs List, fn PredicateF) List {
    if len(xs) == 0 {
        return xs
    }
    ys := make(List, 0)
    for _, x := range xs {
        if fn(x) {
            ys = append(ys, x)
        }
    }
    return ys
}
