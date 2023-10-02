1. General loop
```
sum := 0
for i := 1; i < 3; i++ {
	sum += i
}
fmt.Println(sum)
```
2. While loop
```
n := 1
for n < 3 {
	n *= 2
}
fmt.Println(n)
```
3. Infinite loop
```
sum := 0
for {
	sum++
}
fmt.Println(sum)
```
4. For-each range loop
```
numbers := []string{1, 2}
for i, s := range numbers {
	fmt.Println(i, s)
}
```
5. Exit a loop
```
sum := 0
for i := 1; i < 5; i++ {
	if i%2 != 0 { // skip odd numbers
		continue
	}
	sum += i
}
fmt.Println(sum)
```