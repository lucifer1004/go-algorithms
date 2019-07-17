package unionfind

import (
	"bufio"
	"fmt"
	"os"
)

// Test the implementations
func Test() {
	var (
		version, n, p, q int
		uf               UnionFind
	)
	fmt.Println("Choose the version you want to test: 1 -- 4")
	fmt.Scanln(&version)
	switch version {
	case 1:
		uf = new(IntUnionFindV1)
	case 2:
		uf = new(IntUnionFindV2)
	case 3:
		uf = new(IntUnionFindV3)
	case 4:
		uf = new(IntUnionFindV4)
	}

	fmt.Println("Please enter the number of nodes:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	uf.init(n)
	for true {
		scanner.Scan()
		if scanner.Text() == "" {
			break
		}
		fmt.Sscan(scanner.Text(), &p, &q)
		if !uf.isConnected(p, q) {
			uf.union(p, q)
			fmt.Println(p, "and", q, "are now connected.")
		} else {
			fmt.Println(p, "and", q, "are already connected.")
		}
	}
	fmt.Println("There are", uf.count(), "connected components in total.")
}
