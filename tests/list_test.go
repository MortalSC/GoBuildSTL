package tests

import (
	"GOBUILDSTL/src/containers"
	"fmt"
	"testing"
)

// 测试在链表头部插入元素
func TestInsertAtHead(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtHead(10)
	list.InsertAtHead(20)

	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, but got %d", list.Size())
	}

	headValue := list.Head().GetData()
	if headValue != 20 {
		t.Errorf("Expected head value to be 20, but got %d", headValue)
	}
}

// 测试在链表尾部插入元素
func TestInsertAtTail(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtTail(10)
	list.InsertAtTail(30)

	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, but got %d", list.Size())
	}

	tailValue := list.Tail().GetData()
	if tailValue != 30 {
		t.Errorf("Expected tail value to be 30, but got %d", tailValue)
	}
}

// 测试删除头部元素错误
func TestRemoveHeadError(t *testing.T) {
	list := containers.NewList[int]()

	err := list.RemoveAtHead() // 试图删除空链表头
	if err == nil {
		t.Errorf("Expected error when removing head from empty list, but got nil")
	}
}

// 测试删除尾部元素错误
func TestRemoveTailError(t *testing.T) {
	list := containers.NewList[int]()

	err := list.RemoveAtTail() // 试图删除空链表尾
	if err == nil {
		t.Errorf("Expected error when removing tail from empty list, but got nil")
	}
}

// 测试删除指定节点错误
func TestRemoveNodeError(t *testing.T) {
	list := containers.NewList[int]()

	err := list.RemoveNode(nil) // 试图删除nil节点
	if err == nil {
		t.Errorf("Expected error when removing a nil node, but got nil")
	}
}

// 测试删除头部元素
func TestRemoveHead(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtHead(10)
	list.InsertAtHead(20)

	err := list.RemoveAtHead()

	if err != nil {
		fmt.Println("[ Error ] : ", err)
	}

	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1 after removing head, but got %d", list.Size())
	}

	headValue := list.Head().GetData()
	if headValue != 10 {
		t.Errorf("Expected head value to be 10 after removing, but got %d", headValue)
	}
}

// 测试删除尾部元素
func TestRemoveTail(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtTail(10)
	list.InsertAtTail(30)

	err := list.RemoveAtTail()

	if err != nil {
		fmt.Println("[ Error ] : ", err)
	}

	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1 after removing tail, but got %d", list.Size())
	}

	tailValue := list.Tail().GetData()
	if tailValue != 10 {
		t.Errorf("Expected tail value to be 10 after removing, but got %d", tailValue)
	}
}

// 测试遍历操作 (前向)
func TestTraverseForward(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)

	expected := []int{10, 20, 30}
	index := 0
	list.TraverseForward(func(value int) {
		if value != expected[index] {
			t.Errorf("Expected value %d at index %d, but got %d", expected[index], index, value)
		}
		index++
	})

	if index != 3 {
		t.Errorf("Expected to traverse 3 elements, but traversed %d", index)
	}
}

// 测试遍历操作 (后向)
func TestTraverseBackward(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)

	expected := []int{30, 20, 10}
	index := 0
	list.TraverseBackward(func(value int) {
		if value != expected[index] {
			t.Errorf("Expected value %d at index %d, but got %d", expected[index], index, value)
		}
		index++
	})

	if index != 3 {
		t.Errorf("Expected to traverse 3 elements, but traversed %d", index)
	}
}

// 测试删除指定节点
func TestRemoveNode(t *testing.T) {
	list := containers.NewList[int]()

	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)

	nodeToRemove := list.Head().Next() // 删除第二个节点

	err := list.RemoveNode(nodeToRemove)

	if err != nil {
		fmt.Println("[ Error ] : ", err)
	}

	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2 after removing a node, but got %d", list.Size())
	}

	expected := []int{10, 30}
	index := 0
	list.TraverseForward(func(value int) {
		if value != expected[index] {
			t.Errorf("Expected value %d at index %d, but got %d", expected[index], index, value)
		}
		index++
	})
}
