Claude helped me describe it:

Worst-case per-operation cost is O(log* n), where n is the number of elements
that have ever passed through the structure (or current size — whichever you
use should be stated explicitly, since they could differ). This follows from:
    (1) nesting depth d satisfies d(n) = O(log* n), because each level's capacity
		    grows as roughly the square (or some fixed power) of the level below
				it's threshold — i.e., the thresholds at which each level fires grow as
				a tower-like function of n, the inverse of which is log*; and
		(2) each level does O(1) work per operation, independent of d, established
				directly from the code (rebal only ever pops a fixed number of elements
				off rotin/rotout at its own level, never recursing into the contents of
				the stacks it moves).
