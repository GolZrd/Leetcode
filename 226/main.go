package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Используя поиск в глубину (DFS)
func invertTree(root *TreeNode) *TreeNode {
	// Если корень дерева пустой (nil), то возвращаем nil
	if root == nil {
		return nil
	}

	// Меняем местами левое и правое поддерево
	root.Left, root.Right = root.Right, root.Left

	// Рекурсивно обходим левое и правое поддерево
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// Используя поиск в ширину (BFS)
func invertTreeBFS(root *TreeNode) *TreeNode {
	// проверка корня на пустоту
	if root == nil {
		return nil
	}

	// Создаем очередь для обхода дерева
	queue := []*TreeNode{root}

	// Запускаем цикл пока очередь не пуста
	for len(queue) > 0 {
		// Извлекаем узел из очереди
		node := queue[0]
		// И удаляем его. Очередь работает по принципу FIFO. Поэтому удаляем первый элемент.
		queue = queue[1:]

		// Меняем местами левое и правое поддерево
		node.Left, node.Right = node.Right, node.Left

		// Добавляем левое и правое поддерево в очередь
		// Проверяем что поддерево не пустое
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
