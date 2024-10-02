# BuildList：List 实现说明



## 概述

> 本篇概述：使用 Go 语言实现类似于 C++ `std::list` 的双向链表数据结构。
>
> - C++ `std::list` 是一个双向链表容器，允许在常数时间内进行元素的插入和删除操作。
> - 为了模拟这一功能，我们利用 Go 语言的泛型和自动内存管理特性，构建了通用的链表容器，并提供了基本的插入、删除、遍历等操作。



## 设计目标

我们的目标是：

1. **实现通用的双向链表**，支持任意类型的元素（使用 Go 泛型 `T any`）。
2. **模拟 C++ `std::list` 的核心功能**，包括头部和尾部的高效插入与删除。
3. **支持遍历操作**，既可以从头到尾顺序遍历，也可以从尾到头逆序遍历。
4. **实现节点的迭代**，通过节点的 `Next` 和 `Prev` 方法访问相邻节点，模拟 C++ 中迭代器的行为。



## 具体实现

### 1. 数据结构

链表通过两个主要结构体实现：

- **Node[T]**：表示链表中的节点。每个节点存储一个数据元素，并且有两个指针，分别指向前驱节点和后继节点。

  ```go
  type Node[T any] struct {
      data T         // 节点的数据
      prev *Node[T]  // 前驱节点
      next *Node[T]  // 后继节点
  }
  ```

- **List[T]**：表示整个链表。它管理链表的头节点、尾节点以及链表的大小。

  ```go
  type List[T any] struct {
      head *Node[T]  // 链表的头节点
      tail *Node[T]  // 链表的尾节点
      size int       // 链表的长度
  }
  ```

### 2. 基本操作

#### - 插入操作

链表支持在头部和尾部进行元素插入：

- **InsertAtHead(data T)**：在链表头部插入一个元素。
- **InsertAtTail(data T)**：在链表尾部插入一个元素。

示例代码：

```
list.InsertAtHead(10)
list.InsertAtTail(20)
```

#### - 删除操作

支持从链表头部、尾部删除节点，也可以删除指定节点：

- **RemoveHead()**：删除链表头部节点。
- **RemoveTail()**：删除链表尾部节点。
- **RemoveNode(node \*Node[T])**：删除指定节点。

示例代码：

```go
list.RemoveHead()
list.RemoveNode(node)
```

#### - 遍历操作

为了遍历链表中的元素，我们提供了两个方法：

- **TraverseForward(func(data T))**：从头到尾顺序遍历链表。
- **TraverseBackward(func(data T))**：从尾到头逆序遍历链表。

示例代码：

```go
list.TraverseForward(func(data int) {
    fmt.Println(data)
})
```

#### - 获取链表大小

使用 **Size()** 方法可以获取链表中元素的数量：

```go
size := list.Size()
```

### 3. 节点操作

每个节点可以通过 `Next()` 方法访问下一个节点，通过 `Prev()` 方法访问前一个节点。这模拟了 C++ `std::list` 的迭代器行为。

- **GetData()**：获取节点存储的数据。
- **Next()**：获取当前节点的下一个节点。
- **Prev()**：获取当前节点的前一个节点。

示例代码：

```go
node := list.Head()
nextNode := node.Next()
prevNode := node.Prev()
```

### 4. 泛型支持

为了使链表能够存储任意类型的数据，Go 语言的泛型功能被充分利用。链表及其节点都采用了泛型参数 `T`，使得用户可以创建任何类型的链表。

示例代码：

```go
intList := containers.NewList[int]()
stringList := containers.NewList[string]()
```

### 5. 测试用例

项目中编写了全面的单元测试，确保链表的所有核心功能能够正确运行。主要测试包括：

- 插入、删除头部和尾部节点。
- 遍历操作的正确性。
- 删除特定节点的操作。

测试代码示例：

```go
list := containers.NewList[int]()
list.InsertAtHead(10)
list.InsertAtTail(20)

if list.Size() != 2 {
    t.Errorf("Expected size 2, but got %d", list.Size())
}
```



## 与 C++ STL `std::list` 的对比

- **内存管理**：C++ 中的 `std::list` 需要手动管理内存，而 Go 使用自动垃圾回收机制，这简化了开发过程，但可能带来性能上的差异。
- **迭代器**：C++ 中的 `std::list` 提供双向迭代器，而 Go 没有原生的迭代器机制。我们通过 `Next()` 和 `Prev()` 方法实现了类似的功能，便于在链表中导航。
- **泛型支持**：C++ 的 `std::list` 通过模板机制支持不同类型，而我们在 Go 中使用泛型来实现同样的效果。


## 增加错误处理
- List 的主要变化：
- RemoveHead() 和 RemoveTail()：当链表为空时，会返回一个错误，表示无法从空链表中删除元素。
- RemoveNode()：当尝试删除 nil 节点时，返回错误。
- 其余操作：如插入和遍历等操作仍然保持简洁，不需要返回错误。