package main

func getLetterNodes() *TreeNode {
	return &TreeNode{
		Val: "f",
		Left: &TreeNode{
			Val: "b",
			Left: &TreeNode{
				Val: "a",
			},
			Right: &TreeNode{
				Val: "d",
				Left: &TreeNode{
					Val: "c",
				},
				Right: &TreeNode{
					Val: "e",
				},
			},
		},
		Right: &TreeNode{
			Val: "g",
			Right: &TreeNode{
				Val: "i",
				Left: &TreeNode{
					Val: "h",
				},
			},
		},
	}
}

func getFirst() *TreeNode {
	return &TreeNode{
		Val: "1",
		Left: &TreeNode{
			Val: "4",
			Left: &TreeNode{
				Val: "2",
			},
		},
		Right: &TreeNode{
			Val: "3",
		},
	}
}

func getSecond() *TreeNode {
	return &TreeNode{
		Val: "2",
		Left: &TreeNode{
			Val: "3",
			Left: &TreeNode{
				Val: "1",
			},
		},
	}
}

func getThird() *TreeNode {
	return &TreeNode{
		Val: "3",
		Left: &TreeNode{
			Val: "1",
			Right: &TreeNode{
				Val: "2",
			},
		},
	}
}

func getForth() *TreeNode {
	return &TreeNode{
		Val: "2",
		Left: &TreeNode{
			Val: "3",
			Left: &TreeNode{
				Val: "1",
			},
		},
	}
}

