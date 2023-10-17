1. If
```
if ... {
	...
}
```
2. If with initial value
```
if x := f(); x <= y {
	return x
}
```
3. If Else
```
if ... {
	...
} else {
	...
}
```
4. Nested if
```
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```