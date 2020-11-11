package main

import (
	"flag"
	"fmt"
	"github.com/hslam/btree"
	"time"
)

var total int
var degree int

func init() {
	flag.IntVar(&total, "total", 10000000, "-total=1000000")
	flag.IntVar(&degree, "degree", 1000, "-degree=256")
	flag.Parse()
}

func main() {
	if total < 1 || degree < 2 {
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
	tree := btree.New(degree)
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(i))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Insert:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Search() {
	tree := btree.New(degree)
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(i))
	}
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Search(btree.Int(i))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Search:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Delete() {
	tree := btree.New(degree)
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(i))
	}
	start := time.Now()
	for i := 0; i < total; i++ {
		tree.Delete(btree.Int(i))
	}
	d := time.Now().Sub(start)
	fmt.Printf("Delete:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}

func Iterate() {
	tree := btree.New(degree)
	for i := 0; i < total; i++ {
		tree.Insert(btree.Int(i))
	}
	start := time.Now()
	iter := tree.Min().MinIterator()
	for iter != nil {
		iter.Item()
		iter = iter.Next()
	}
	d := time.Now().Sub(start)
	fmt.Printf("Iterate:\ttotal %d,\tdegree %d,\ttime %v,\t%d ns/op,\t%d items/s\n", total, degree, d, int(d.Nanoseconds())/total, int(float64(total)*1e9/float64(d.Nanoseconds())))
}
