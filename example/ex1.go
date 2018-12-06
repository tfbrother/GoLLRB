package main

import (
	"fmt"
	"github.com/tfbrother/GoLLRB/llrb"
)

func main() {
	tree := llrb.New()
	tree.ReplaceOrInsert(llrb.Int(1))
	tree.ReplaceOrInsert(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(3))
	tree.ReplaceOrInsert(llrb.Int(4))
	tree.DeleteMin()
	//tree.Delete(llrb.Int(4))
	fmt.Println(tree.Max())
	fmt.Println(tree.Min())
}
