package mock

import "github.com/jakha/lc/pkg"


func GetLetterNodes() *pkg.TreeNode {
	return &pkg.TreeNode{
		Val: "f",
		Left: &pkg.TreeNode{
			Val: "b",
			Left: &pkg.TreeNode{
				Val: "a",
			},
			Right: &pkg.TreeNode{
				Val: "d",
				Left: &pkg.TreeNode{
					Val: "c",
				},
				Right: &pkg.TreeNode{
					Val: "e",
				},
			},
		},
		Right: &pkg.TreeNode{
			Val: "g",
			Right: &pkg.TreeNode{
				Val: "i",
				Left: &pkg.TreeNode{
					Val: "h",
				},
			},
		},
	}
}

func getFirst() *pkg.TreeNode {
	return &pkg.TreeNode{
		Val: "1",
		Left: &pkg.TreeNode{
			Val: "4",
			Left: &pkg.TreeNode{
				Val: "2",
			},
		},
		Right: &pkg.TreeNode{
			Val: "3",
		},
	}
}

func getSecond() *pkg.TreeNode {
	return &pkg.TreeNode{
		Val: "2",
		Left: &pkg.TreeNode{
			Val: "3",
			Left: &pkg.TreeNode{
				Val: "1",
			},
		},
	}
}

func getThird() *pkg.TreeNode {
	return &pkg.TreeNode{
		Val: "3",
		Left: &pkg.TreeNode{
			Val: "1",
			Right: &pkg.TreeNode{
				Val: "2",
			},
		},
	}
}

func getForth() *pkg.TreeNode {
	return &pkg.TreeNode{
		Val: "2",
		Left: &pkg.TreeNode{
			Val: "3",
			Left: &pkg.TreeNode{
				Val: "1",
			},
		},
	}
}

