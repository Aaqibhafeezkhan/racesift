# 0003: Group by the pair of code locations

**Decision:** two reports are the same race when the first frame from the user's code matches on both sides, by function, file and line, in either order.

**Why:** Go prints a report for each conflicting pair it sees, so one bug repeats across tests and goroutines with different IDs and addresses. The two code lines are what a fix changes.

**Alternatives:** grouping by memory address only works within one process, and addresses are reused. Grouping by one side only would merge different bugs that touch a common line, so that case is shown separately as "lines in more than one race".
