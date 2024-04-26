package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Node struct {
	Data Person
	Prev *Node
	Next *Node
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func NewNode(data Person) *Node {
	return &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}

func (n *Node) Hook(prev *Node, next Node, ishead bool) {
	if ishead {
		// head -> next
		n.Prev = nil
		n.Next = &next
		// next -> head
		n.Next.Prev = n
		return
	}
	// prev <- current -> next
	n.Prev = prev
	n.Next = &next

	//prev -> current
	n.Prev.Next = n
	// next -> current
	n.Next.Prev = n
}

func PrintAll(nn *Node) {

	fmt.Println("=== print all node elements ===")
	fmt.Println()
	fmt.Println("person name:", nn.Data.Name)
	fmt.Println("person age:", nn.Data.Age)

	if nn.Next == nil {
		return
	}
	PrintAll(nn.Next)
}

func CountNode(node *Node, i int) int {
	if i <= 0 {
		i = 1
	}
	if node.Next == nil {
		return i
	}

	return CountNode(node.Next, i+1)
}

func (n *Node) CountNodes(i int) int {

	if i <= 0 {
		i = 1
	}

	if n.Next == nil {
		return i
	}
	return n.CountNodes(i + 1)
}

func main() {

	person1 := Person{
		Name: "AGUS",
		Age:  20,
	}

	person2 := Person{
		Name: "Bagas",
		Age:  20,
	}

	person3 := Person{
		Name: "Bagus",
		Age:  21,
	}

	head := NewNode(person1)
	second := NewNode(person2)
	tail := NewNode(person3)

	// head
	// head -> second
	head.Hook(nil, *second, true)

	// second
	// head <- second -> tail
	second.Hook(head, *tail, false)

	fmt.Print("\n")
	fmt.Println(head.Data)
	fmt.Println(head.Next.Data)
	fmt.Print("\n")
	PrintAll(head)
	fmt.Printf("all node lenght %d \n", CountNode(head, 0))

	// print second element
	fmt.Printf(" &second : %v \n", &second)
	fmt.Printf("second -> prev : %v \n", second.Prev)
	fmt.Printf("second -> next : %v \n", second.Next)
}
