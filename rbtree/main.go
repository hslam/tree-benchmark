package main

import (
	"flag"
	"fmt"
	"github.com/hslam/rbtree"
	"time"
)

var total int

func init() {
	flag.IntVar(&total, "total", 10000000, "-total=1000000")
	flag.Parse()
}

func main() {
	if total < 1 {
		return
	}
	Insert()
	Insert()
	Insert()
	Search()
	Search()
	Search()
	Delete()
	Delete()
	Delete()
	Iterate()
	Iterate()
	Iterate()
}

func Insert() {
	tree := rbtree.New()
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Insert(rbtree.Int(i))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Insert:\ttotal %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Search() {
	tree := rbtree.New()
	for i := 0; i < total; i++ {
		tree.Insert(rbtree.Int(i))
	}
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Search(rbtree.Int(i))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Search:\ttotal %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Delete() {
	tree := rbtree.New()
	for i := 0; i < total; i++ {
		tree.Insert(rbtree.Int(i))
	}
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Delete(rbtree.Int(i))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Delete:\ttotal %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Iterate() {
	tree := rbtree.New()
	for i := 0; i < total; i++ {
		tree.Insert(rbtree.Int(i))
	}
	start := time.Now()
	iter := tree.Min()
	for iter != nil {
		iter.Item()
		iter = iter.Next()
	}
	d := time.Now().Sub(start)
	fmt.Printf("Iterate:\ttotal %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}
