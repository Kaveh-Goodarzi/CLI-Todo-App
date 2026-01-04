package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	// todos.add("Buy Milk")
	// todos.add("Buy bread")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)
}
