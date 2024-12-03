package main

func main() {

}

// Решение с помощью двойного связного списка

// Определяем узел Node двухсвязного списка
type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

// Определяем структуру нашего LRU кэша
type LRUCache struct {
	cap   int           // Здесь будут такие поля как cap - capacity (вместимость)
	cache map[int]*Node // Мапа для хранения ссылки на узел (node) в связном списке.
	left  *Node         // Ссылка на левый узел, то есть первый узел в связном списке. Это также узел, который давно не использовался, поэтому он будет удаляться
	right *Node         // Ссылка на правый узел, то есть последний узел в связном списке. Также можно сказать, что он недавно использовался
}

// capacity - количество элементов, которые мы будем хранить в кэше.
// Определяется при вызове функции и возвращается экземпляр LRUCache
func Constructor(capacity int) LRUCache {
	// Определяем сткрутуру нашего LRUCache
	lru := LRUCache{
		cap:   capacity,            // указываем передаенное capacity
		cache: make(map[int]*Node), // инициализируем мапу
		left:  &Node{},             // Определяем для поля left пустую ноду
		right: &Node{},             // Определяем для поля right пустую ноду
	}
	// Так как у нас будет двусвязный список,
	// нужно заполнить указатели на следующий и предыдущий узел
	lru.left.next = lru.right
	lru.right.prev = lru.left
	// возвращаем экземпляр LRUCache
	return lru
}

// Определим дополнительную функцию для удаления узла из двухсвязного списка
func (this *LRUCache) remove(node *Node) {
	// Сохраняем указатели на предыдущий и следующий узлы, у ноды, которую будем удалять (который был передан в функцию)
	prev, next := node.prev, node.next
	// Теперь просто как бы вырезаем этот узел
	// Сначала для узла, который был предыдущим (prev) меняем ссылку на следующий (next)
	// то есть на тот узел, который шел после удаляемого узла
	prev.next = next
	// Также меняем указатель на предыдущий узел у ноды, которая шла после удаляемого (next)
	next.prev = prev
	// Таким образом мы как бы пропустили удаляемый узел и теперь у нас нет ссылок на него
	// Мы могли бы это написать следующим образом
	// prev.next = next.next
	// next.prev = prev
}

// Теперь определим функцию для добавления узла в двухсвязный список
func (this *LRUCache) insert(node *Node) {
	// Сохраняем указатели на предыдущий и следующий узлы у ноды right (то есть последний узел в связном списке)
	prev, next := this.right.prev, this.right
	// right у нас указывает на пустую ноду, поэтому нода с элементом находится на this.right.prev
	prev.next = node
	next.prev = node
	node.prev = prev
	node.next = next

	// По сути у нас остается такой же указатель на правую ноду в нашем LRUCache
	// мы его не меняем и не перезаписываем, мы просто при добавлении нового его двигаем вправо
	// а вот указатель prev у этого узла, мы меняем каждый раз при добавлении нового узла
}

// Проверяем, существует ли значение по этому ключу, если есть, то возвращаем его, иначе -1
func (this *LRUCache) Get(key int) int {
	// Проверяем в мапе cache нашей структуре LRUCache, есть ли значение по переданному ключу
	if node, ok := this.cache[key]; ok {
		// Если значение есть, тогда мы должны сначала удалить его из нашего двусвязного списка
		this.remove(node)
		// После удаления, мы его добавляем в наш список, но уже в конец, он должен быть с правой стороны
		// Так как этот узел недавно вызывался, он не может оставаться слева на прежнем месте
		// Поэтому мы его двигаем вправо
		this.insert(node)
		// После этого возвращаем значение этого узла
		return node.value
	}
	// Если такого ключа нет в мапе cache, то возвращаем -1
	return -1
}

// Функция для помещения пары ключ значение в наш кэш
// Если ключ уже существует, то мы обновляем его значение (value) и сдвигаем его вправо, если нет, то помещаем новую пару
// Если количество ключей превышает capacity, то нам нужно удалить самый старый элемент
func (this *LRUCache) Put(key int, value int) {
	// Проверяем, существует ли уже в мапе cache, переданный ключ (key)
	if node, ok := this.cache[key]; ok {
		// Если существует, то мы удаляем его из двусвязного списка с помощью функции remove
		this.remove(node)
		// Также нам нужно удалить этот ключ из мапы cache
		delete(this.cache, key)
	}
	// Теперь создаем новый узел с переданными ключом (key) и значением (value)
	node := &Node{key: key, value: value}
	// Добавляем созданный узел (node) в мапу по переданному ключу (key)
	this.cache[key] = node
	// Добавляем созданный узел (node) в двусвязный список. По сути мы его добавляем в конец.
	this.insert(node)
	// Сравниваем количество ключей в нашей мапе cache с capacity.
	// Если количество ключей превышает capacity, то нам нужно удалить самый старый элемент
	if len(this.cache) > this.cap {
		// Получаем первый узел в нашем двусвязном списке
		lru := this.left.next
		// Удаляем этот узел
		this.remove(lru)
		// Также удаляем его из мапы cache
		delete(this.cache, lru.key)
	}
}