package solutions

import (
	"fmt"
	"strings"
)

type Node struct {
	prev *Node
	next *Node
	value rune
}
type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}
func (d *DoubleLinkedList) AppendValue(value rune) *Node {
	var n *Node
	if d.size == 0 {
		// First element
		n = &Node{value: value}
		d.head = n
		d.tail = n
	} else {
		n = &Node{value: value}
		// Attach the new node to the tail
		n.prev = d.tail
		// Connect it the other way too
		d.tail.next = n
		// Make the new node the new tail
		d.tail = n
	}
	d.size++
	return n
}
func (d *DoubleLinkedList) RemoveNode(n *Node) rune {
	if n == d.head {
		d.head = n.next
		if d.head != nil {
			d.head.prev = nil
		}
	} else if n == d.tail {
		d.tail = d.tail.prev
	} else {
		n.next.prev = n.prev
		n.prev.next = n.next
	}
	d.size--
	return n.value
}
func (d *DoubleLinkedList) ToArray() []string {
	var arr []string = make([]string, 0)
	var n *Node = d.head
	for n != nil {
		arr = append(arr, string(n.value))
		n = n.next
	}
	return arr
}
func (d *DoubleLinkedList) SPrint() string {
	var arr []string = d.ToArray()
	return strings.Join(arr, "")
}

func ReactionTable() map[rune]rune {
	// Map A-Z to a-z, and vice versa
	var m map[rune]rune = make(map[rune]rune)

	// Loop through ascii ordinal values of the range [65,90],
	// i.e,, A-Z.
	for i:= 65; i <= 90; i++ {
		var upper rune = rune(i)
		var lower rune = rune(i + 32)
		m[upper] = lower
		m[lower] = upper
	}
	return m
}

func CollapseReactions(dll *DoubleLinkedList) {
	if dll.size <= 2 {
		// There couldn't possibly be anything left to eliminate
		return
	}

	var rt map[rune]rune = ReactionTable()

	var current *Node
	var next *Node

	current = dll.head  // pos 0
	next = current.next // pos 1
	var remove1 *Node
	var remove2 *Node

	for current != nil && next != nil {
		var removed bool = false
		if current == dll.head {
			// Head of the list:
			if current.value == rt[next.value] {
				// Remove current and next
				remove1 = current
				remove2 = next

				// current and next become the next two
				next = next.next.next
				current = next.prev

				dll.RemoveNode(remove1)
				dll.RemoveNode(remove2)
				removed = true
			}
		} else if current.value == rt[next.value] {
			// Normal case: not at head or tail of this
			if current.value == rt[next.value] {
				// Keep track of which ones we're removing
				remove1 = current
				remove2 = next

				// Before removing, get the nodes directly before and
				// after these two
				current = current.prev
				next = next.next

				dll.RemoveNode(remove1)
				dll.RemoveNode(remove2)
				removed = true
			}
		}
		if !removed {
			next = next.next
			current = current.next
		}
	}
}

func loadList(input string) *DoubleLinkedList {
	var dll *DoubleLinkedList = &DoubleLinkedList{}
	for _, r := range input {
		if r != '\n' {
			dll.AppendValue(r)
		}
	}
	return dll
}

func Day5Part1(input string) string {
	var dll *DoubleLinkedList = loadList(input)
	CollapseReactions(dll)
	fmt.Println(dll.SPrint())
	return fmt.Sprintf("Units remaining: %d", dll.size)
}

func Day5Part2(input string) string {
	var dll *DoubleLinkedList
	var minLength int = len(input)

	for i := 65; i <= 90; i++ {
		var inputCopy string = input
		// Convert from ascii code to string of single char
		// Remove all instances of thos chars
		inputCopy = strings.Replace(inputCopy, string(i), "", -1)
		inputCopy = strings.Replace(inputCopy, string(i + 32), "", -1)
		dll = loadList(inputCopy)
		CollapseReactions(dll)
		if dll.size < minLength {
			minLength = dll.size
		}
	}
	return fmt.Sprintf("Min length of polymer: %d", minLength)
}
