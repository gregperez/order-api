package leetcode

// ListNode Definición de un nodo en una lista enlazada.
type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
}

// AddTwoNumbers suma dos números representados por listas enlazadas.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummyHead.Next
}

// ConvertDoublyLinkedListToArray - Convertir doble lista enlazada a arreglo
func ConvertDoublyLinkedListToArray(head *DoublyListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
