/*
https://contest.yandex.ru/contest/24810/run-report/112807323/

-- ПРИНЦИП РАБОТЫ --
	Функция remove принимает указатель на корневой узел дерева и ключ для удаления. Если узел равен nil, возвращается nil, что означает, что дерево пустое или ключ не найден.
Когда ключ меньше значения текущего узла, рекурсивно вызывается remove для левого поддерева. Если ключ больше значения текущего узла, рекурсивно вызывается remove для правого поддерева.
Когда ключ равен значению текущего узла, то:
   1. Если у узла нет левого потомка, возвращается правый потомок.
   2. Если у узла нет правого потомка, возвращается левый потомок.
   3. Если у узла есть два потомка, то нахожу узел с минимальным значением в правом поддереве, его значение копирую в текущий узел, а затем этот узел удаляю из правого поддерева.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Удаление узла без потомков происходит корректно, т.к. функция возвращает nil, что означает удаление этого узла.
Удаление узла с одним потомком происходит корректно, т.к. функция возвращает ссылку на этого потомка, что бы после удаления узла перенаправить ссылку на его потомка.
Удаление узла с двумя потомками происходит корректно, т.к. я нахожу узел с минимальным значением в правом поддереве, копирую его значение в текущий узел и рекурсивно вызываю remove для этого узла с минимальным значением.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
n – количество узлов в дереве
h - высота дерева

В несбалансированным дереве, когда h равна n сложность O(n).
В сбалансированном дереве, где h равна логарифму n сложность O(log n).

Итогововя сложность O(n) в худшем случае, O(log n) в лучшем случае

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
n – количество узлов в дереве
Сложность кода по памяти в O(n)

*/

package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func remove(node *Node, key int) *Node {
	// Your code
	// “ヽ(´▽｀)ノ”

	if node == nil {
		return nil
	}

	if key < node.value {
		node.left = remove(node.left, key)
	} else if key > node.value {
		node.right = remove(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minRight := findMin(node.right)
		node.value = minRight.value
		node.right = remove(node.right, minRight.value)
	}

	return node

}

func findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}
