A struct is a typed collection of fields, useful for grouping data into records.

```
package main

import "log"

type Person struct {
	Name string
	Profession string
	DOB int
}

func (m *Person) checkAge() bool {
	return m.DOB % len(m.Name) == 0
}

func main() {
	me := Person{
		Name: "Hoan",
		Profession: "Software Engineer",
		DOB: 1990,
	}
	goodFit := me.checkAge() 
	if goodFit {
		log.Print("I am fit for the job")
	} else {
		log.Print("I am not fit for the job")
	}
}
```