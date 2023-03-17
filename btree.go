package btree

import (
	"hash/fnv"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int64
	Data  []byte
}

func hashData(b []byte) int64 {
	h := fnv.New64a()
	h.Write(b)
	return int64(h.Sum64())
}

func (n *TreeNode) Insert(data []byte) {
	if n == nil {
		n = &TreeNode{Value: hashData(data), Data: data}
		return
	}

	if n.Value > hashData(data) {
		if n.Left == nil {
			n.Left = &TreeNode{Value: hashData(data), Data: data}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Value: hashData(data), Data: data}
		} else {
			n.Right.Insert(data)
		}
	}
}
func (n *TreeNode) Search(data []byte) *TreeNode {
	if n == nil {
		return nil
	}

	if hashData(data) == n.Value {
		return n
	}

	if hashData(data) < n.Value {
		return n.Left.Search(data)
	}

	return n.Right.Search(data)
}

func (n *TreeNode) Delete(data []byte) *TreeNode {
	if n == nil {
		return nil
	}

	if hashData(data) < n.Value {
		n.Left = n.Left.Delete(data)
	} else if hashData(data) > n.Value {
		n.Right = n.Right.Delete(data)
	} else {
		if n.Left == nil && n.Right == nil {
			return nil
		} else if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		} else {
			successor := n.Right
			for successor.Left != nil {
				successor = successor.Left
			}
			n.Value = successor.Value
			n.Data = successor.Data
			n.Right = n.Right.Delete(successor.Data)
		}
	}

	return n
}

func (n *TreeNode) Update(data []byte, newData []byte) *TreeNode {
	if n == nil {
		return nil
	}

	if hashData(data) == n.Value {
		n.Data = newData
	} else if hashData(data) < n.Value {
		n.Left = n.Left.Update(data, newData)
	} else {
		n.Right = n.Right.Update(data, newData)
	}

	return n
}
