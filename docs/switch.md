1. Simple
```
switch time.Now().Weekday() {
case time.Saturday:
	fmt.Println("Today is Saturday.")
case time.Sunday:
	fmt.Println("Today is Sunday.")
default:
	fmt.Println("Today is a weekday.")
}
```
2. Execution order
First the switch expression is evaluated once, then case expressions are evaluated left-to-right and top-to-bottom
```
func Foo(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
	case Foo(4):
		fmt.Println("Second case")
	}
}
```
3. Case list
```
switch time.Now().Weekday() {
    case 1, 2, 3, 4, 5, 6:
        fmt.Println("Weekdays")
    default:
        fmt.Println("Weekend")
}
```
4. Fallthrough
A fallthrough statement transfers control to the next case.
It may be used only as the final statement in a clause.
```
switch 2 {
    case 1:
        fmt.Println("1")
        fallthrough
    case 2:
        fmt.Println("2")
        fallthrough
    case 3:
        fmt.Println("3")
}
```
5. Exit with break
If we need to break out of a surrounding loop, not the switch, we can put a label on the loop and break to that label
```
Loop:
for _, ch := range "a b\nc" {
    switch ch {
        case ' ': // skip space
            break
        case '\n': // break at newline(exit loop)
            break Loop
        default:
            fmt.Printf("%c\n", ch)
    }
}
```
