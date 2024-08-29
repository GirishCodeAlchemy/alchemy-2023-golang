package solutions

import (
	"container/list"
	"fmt"
)

type MRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key, value int
}

func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *MRUCache) Get(key int) (int, bool) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return 0, false
}

func (c *MRUCache) Put(key, value int) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}
	ele := c.list.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.list.Len() > c.capacity {
		last := c.list.Back()
		if last != nil {
			c.list.Remove(last)
			delete(c.cache, last.Value.(*entry).key)
		}
	}
}

// DisplayList prints the current state of the cache list in a detailed format
func (c *MRUCache) Display() {
	fmt.Print("Current cache list: ")
	for e := c.list.Front(); e != nil; e = e.Next() {
		entry := e.Value.(*entry)
		fmt.Printf("%d: %d -> ", entry.key, entry.value)
	}
	fmt.Println() // Move to a new line after printing the list
}

func init() {
	cache := NewMRUCache(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Display()           // Output: 2: 2 -> 1: 1 ->
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	cache.Display()           // Output: 3: 3 -> 1: 1 ->
	fmt.Println(cache.Get(2)) // returns 0 (not found)
	cache.Display()           // Output: 3: 3 -> 1: 1 ->

	cache.Put(4, 4)           // evicts key 1
	cache.Display()           // Output: 4: 4 -> 3: 3 ->
	fmt.Println(cache.Get(1)) // returns 0 (not found)
	cache.Display()           // Output: 4: 4 -> 3: 3 ->

	fmt.Println(cache.Get(3)) // returns 3
	cache.Display()           // Output: 4: 4 -> 3: 3 ->

	fmt.Println(cache.Get(4)) // returns 4
	cache.Display()           // Output: 4: 4 -> 3: 3 ->

}

// Get:
// Check if the key exists in the map.
// If it exists, move the corresponding list element to the front (indicating it was recently accessed) and return the value.
// If it doesn’t exist, return a default value (e.g., 0) and indicate the key was not found.

// Put:
// Check if the key already exists in the map.
// If it exists, update the value and move the corresponding list element to the front.
// If it doesn’t exist, add a new entry to the front of the list and the map.
// If the cache exceeds its capacity, remove the least recently used item (the one at the back of the list) and delete it from the map.

// package main

// import (
// 	"fmt"
// )

// // Node represents a node in the doubly linked list
// type Node struct {
// 	Key   int
// 	Value int
// 	Prev  *Node
// 	Next  *Node
// }

// // MRUCache represents the MRU cache
// type MRUCache struct {
// 	capacity int
// 	cache    map[int]*Node
// 	head     *Node
// 	tail     *Node
// }

// // NewMRUCache initializes a new MRU cache with the given capacity
// func NewMRUCache(capacity int) *MRUCache {
// 	head := &Node{}
// 	tail := &Node{}
// 	head.Next = tail
// 	tail.Prev = head

// 	return &MRUCache{
// 		capacity: capacity,
// 		cache:    make(map[int]*Node),
// 		head:     head,
// 		tail:     tail,
// 	}
// }

// // AddNode adds a node to the head of the list
// func (cache *MRUCache) AddNode(node *Node) {
// 	node.Prev = cache.head
// 	node.Next = cache.head.Next
// 	cache.head.Next.Prev = node
// 	cache.head.Next = node
// }

// // RemoveNode removes a node from the list
// func (cache *MRUCache) RemoveNode(node *Node) {
// 	prev := node.Prev
// 	next := node.Next
// 	prev.Next = next
// 	next.Prev = prev
// }

// // MoveToHead moves a node to the head of the list
// func (cache *MRUCache) MoveToHead(node *Node) {
// 	cache.RemoveNode(node)
// 	cache.AddNode(node)
// }

// // PopTail removes the last node (least recently used) from the list
// func (cache *MRUCache) PopTail() *Node {
// 	node := cache.tail.Prev
// 	cache.RemoveNode(node)
// 	return node
// }

// // Get retrieves a value from the cache
// func (cache *MRUCache) Get(key int) int {
// 	node, exists := cache.cache[key]
// 	if !exists {
// 		return -1
// 	}
// 	cache.MoveToHead(node)
// 	return node.Value
// }

// // Put adds a value to the cache
// func (cache *MRUCache) Put(key int, value int) {
// 	node, exists := cache.cache[key]
// 	if exists {
// 		node.Value = value
// 		cache.MoveToHead(node)
// 		return
// 	}

// 	newNode := &Node{Key: key, Value: value}
// 	cache.AddNode(newNode)
// 	cache.cache[key] = newNode

// 	if len(cache.cache) > cache.capacity {
// 		tail := cache.PopTail()
// 		delete(cache.cache, tail.Key)
// 	}
// }

// // DisplayList prints the current state of the cache list in a detailed format
// func (cache *MRUCache) DisplayList() {
// 	fmt.Print("Current cache list: ")
// 	node := cache.head.Next
// 	for node != cache.tail {
// 		fmt.Printf("%d: %d -> ", node.Key, node.Value)
// 		node = node.Next
// 	}

// 	fmt.Println() // Move to a new line after printing the list
// }

// func main() {
// 	cache := NewMRUCache(3)

// 	cache.Put(1, 1)
// 	cache.Put(2, 2)
// 	fmt.Println("Cache state after adding 1 and 2:")
// 	cache.DisplayList() // Displays the current state of the cache list

// 	fmt.Println("Get(1):", cache.Get(1)) // returns 1
// 	fmt.Println("Cache state after accessing 1:")
// 	cache.DisplayList() // Displays the current state of the cache list

// 	cache.Put(3, 3) // evicts key 2
// 	fmt.Println("Cache state after adding 3:")
// 	cache.DisplayList() // Displays the current state of the cache list

// 	cache.Put(4, 4) // evicts key 1
// 	fmt.Println("Cache state after adding 4:")
// 	cache.DisplayList()

// 	fmt.Println("Get(4):", cache.Get(4)) // returns 1
// 	cache.Put(5, 5)                      // evicts key 1
// 	fmt.Println("Cache state after adding 5:")
// 	cache.DisplayList()
// }
