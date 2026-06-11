package main



type Node struct{
	Key int
	Value int
	Next *Node
	Prev *Node
}


type LRUcache struct{
	Capacity int
	cache map[int]*Node
	Head *Node
	Tail *Node
}

func NewLruCache(capicity int) *LRUcache{
	head := &Node{}
	tail := &Node{}

	head.Next=tail
	tail.Prev=head

	return &LRUcache{
		Capacity:capacity,
		cache : make(map[int]*Node)
		Head :head
		Tail :tail
	}
}

func (l *LRUcache) removeNode(node *Node){

	node.Prev.Next=node.Next
	node.Next.Prev=node.Prev


}

func ( ; *LRUcache) removeTail(node *Node){
	tail :=l.Tail.Prev

	l.removeNode(tail)

	delete(l.cache,tail.Key)
}



func (l *LRUcache) addToFront(node *Node){

	if node.Prev != nil || node.Next !=nil{
		l.removeNode(node)
	}

	node.Prev = l.Head
	node.Next= l.Head.Next

	l.Head.Next.Prev=node

	l.Head.Next=node

}


