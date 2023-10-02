There is no do-while loop in Go, we apply for loop instead
```
for ok := true; ok; ok = condition {
	work()
}
```

```
for {
	work()
	if !condition {
		break
	}
}
```
