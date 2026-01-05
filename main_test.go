package main

import "testing"

func TestAddTodo(t *testing.T) {
	todos := Todos{}
	todos.add("test todo")

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	if todos[0].Title != "test todo" {
		t.Fatalf("unexpected title: %s", todos[0].Title)
	}
}

func TestDeleteTodo(t *testing.T) {
	todos := Todos{}
	todos.add("one")
	todos.add("two")

	err := todos.delete(0)
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}
}

func TestToggleTodo(t *testing.T) {
	todos := Todos{}
	todos.add("test")

	err := todos.toggle(0)
	if err != nil {
		t.Fatal(err)
	}

	if !todos[0].Completed {
		t.Fatal("expected todo to be completed")
	}

	if todos[0].CompletedAt == nil {
		t.Fatal("expected completedAt to be set")
	}
}

func TestEditTodo(t *testing.T) {
	todos := Todos{}
	todos.add("old title")

	err := todos.edit(0, "new title")
	if err != nil {
		t.Fatal(err)
	}

	if todos[0].Title != "new title" {
		t.Fatalf("expected updated title, got %s", todos[0].Title)
	}
}

func TestInvalidIndex(t *testing.T) {
	todos := Todos{}
	err := todos.delete(5)

	if err == nil {
		t.Fatal("expected error for invalid index")
	}
}
