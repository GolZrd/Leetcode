package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Решение с помощью DFS
func maxDepth(root *TreeNode) int {
	// Так как решение у нас является рекурсивным, то опишем базовый случай
	if root == nil {
		return 0
	}

	// Рекурсивно обходим левое и правое поддерево и возвращаем максимальную глубину
	// Это означает, что функция будет вызвана снова и снова, пока не достигнет
	// конечных узлов дерева (т.е. узлов, у которых нет дочерних узлов).
	// После этого мы к этой глубине мы будем прибавлять единицу
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// Решение с помощью BFS
func maxDepthBFS(root *TreeNode) int {
	// Проверяем корень на пустоту
	if root == nil {
		return 0
	}

	// Создаем очередь для обхода дерева
	queue := []*TreeNode{root}

	// Определяем переменную, которая будет обозначать глубину
	level := 0

	// Запускаем цикл пока очередь не пуста
	for len(queue) > 0 {

		// Находим длину очереди, в которой находятся узлы одного уровня
		l := len(queue)

		// Мы на каждом шаге добавляем в нашу очередь все узлы,
		// которые есть на следующем уровне дерева
		// поэтому, чтобы подсчитать глубину, нам нужно все эти узлы с предыдущего уровня удалить
		// для этого мы будем использовать цикл с ограничением по длине очереди
		for i := 0; i < l; i++ {
			// Извлекаем узел из очереди
			node := queue[0]
			// И удалем его из очереди
			queue = queue[1:]
			// проверяем, есть ли дочерние узлы, если есть, добавляем в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// После того как мы добавили новые узлы с следующего уровня
		// и удалили все старые узлы, можем увеличивать наш level на 1
		level++
	}

	return level
}
