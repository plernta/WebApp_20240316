// Object Oriented Programming (OOP)
package log

// class
type Counter struct {
	name  string
	value int
}

// methods
func (c *counter) Increment() {
	c.value++
}

func main() {
	counter1 := Counter{name: "First", value: 0}
	log.printf("Counter %s has value %d", counter1.name, counter1.value)
	counter1.Increment()
	log.printf("Counter %s has value %d", counter1.name, counter1.value)
}
