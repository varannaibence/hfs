package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	value byte
	zero  *Node
	one   *Node
}

type LZWBinTree struct {
	root        Node
	current     *Node
	depth       int
	maxDepth    int
	leafSum     int
	leafCount   int
	varianceSum float64
	average     float64
	deviation   float64
}

func NewLZWBinTree() *LZWBinTree {
	tree := &LZWBinTree{}
	tree.root.value = '/'
	tree.current = &tree.root
	return tree
}

func (tree *LZWBinTree) InsertBit(bit byte) {
	if bit == '0' {
		if tree.current.zero == nil {
			tree.current.zero = &Node{value: '0'}
			tree.current = &tree.root
		} else {
			tree.current = tree.current.zero
		}
		return
	}

	if tree.current.one == nil {
		tree.current.one = &Node{value: '1'}
		tree.current = &tree.root
	} else {
		tree.current = tree.current.one
	}
}

func (tree *LZWBinTree) Write(writer *bufio.Writer) error {
	tree.depth = 0
	tree.writeNode(&tree.root, writer)
	return writer.Flush()
}

func (tree *LZWBinTree) writeNode(node *Node, writer *bufio.Writer) {
	if node == nil {
		return
	}

	tree.depth++
	tree.writeNode(node.one, writer)

	for i := 0; i < tree.depth; i++ {
		fmt.Fprint(writer, "---")
	}
	fmt.Fprintf(writer, "%c(%d)\n", node.value, tree.depth-1)

	tree.writeNode(node.zero, writer)
	tree.depth--
}

func (tree *LZWBinTree) GetDepth() int {
	tree.depth = 0
	tree.maxDepth = 0
	tree.measureDepth(&tree.root)
	return tree.maxDepth - 1
}

func (tree *LZWBinTree) measureDepth(node *Node) {
	if node == nil {
		return
	}

	tree.depth++
	if tree.depth > tree.maxDepth {
		tree.maxDepth = tree.depth
	}
	tree.measureDepth(node.one)
	tree.measureDepth(node.zero)
	tree.depth--
}

func (tree *LZWBinTree) GetMean() float64 {
	tree.depth = 0
	tree.leafSum = 0
	tree.leafCount = 0
	tree.measureMean(&tree.root)
	tree.average = float64(tree.leafSum) / float64(tree.leafCount)
	return tree.average
}

func (tree *LZWBinTree) measureMean(node *Node) {
	if node == nil {
		return
	}

	tree.depth++
	tree.measureMean(node.one)
	tree.measureMean(node.zero)
	tree.depth--

	if node.one == nil && node.zero == nil {
		tree.leafCount++
		tree.leafSum += tree.depth
	}
}

func (tree *LZWBinTree) GetDeviation() float64 {
	tree.average = tree.GetMean()
	tree.varianceSum = 0.0
	tree.depth = 0
	tree.leafCount = 0
	tree.measureDeviation(&tree.root)

	if tree.leafCount-1 > 0 {
		tree.deviation = math.Sqrt(tree.varianceSum / float64(tree.leafCount-1))
	} else {
		tree.deviation = math.Sqrt(tree.varianceSum)
	}

	return tree.deviation
}

func (tree *LZWBinTree) measureDeviation(node *Node) {
	if node == nil {
		return
	}

	tree.depth++
	tree.measureDeviation(node.one)
	tree.measureDeviation(node.zero)
	tree.depth--

	if node.one == nil && node.zero == nil {
		tree.leafCount++
		difference := float64(tree.depth) - tree.average
		tree.varianceSum += difference * difference
	}
}

func usage() {
	fmt.Println("Usage: lzwtree in_file -o out_file")
}

func processInput(data []byte, tree *LZWBinTree) {
	index := 0
	for index < len(data) && data[index] != '\n' {
		index++
	}
	if index < len(data) {
		index++
	}

	commentMode := false
	for ; index < len(data); index++ {
		value := data[index]

		if value == '>' {
			commentMode = true
			continue
		}
		if value == '\n' {
			commentMode = false
			continue
		}
		if commentMode || value == 'N' {
			continue
		}

		for bitIndex := 0; bitIndex < 8; bitIndex++ {
			if value&0x80 != 0 {
				tree.InsertBit('1')
			} else {
				tree.InsertBit('0')
			}
			value <<= 1
		}
	}
}

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(-1)
	}

	inputFile := os.Args[1]
	if os.Args[2] != "-o" {
		usage()
		os.Exit(-2)
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("%s nem letezik...\n", inputFile)
		usage()
		os.Exit(-3)
	}

	outputFile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(-4)
	}
	defer outputFile.Close()

	tree := NewLZWBinTree()
	processInput(data, tree)

	writer := bufio.NewWriter(outputFile)
	tree.Write(writer)
	fmt.Fprintf(writer, "depth = %d\n", tree.GetDepth())
	fmt.Fprintf(writer, "mean = %g\n", tree.GetMean())
	fmt.Fprintf(writer, "var = %g\n", tree.GetDeviation())
	writer.Flush()
}
