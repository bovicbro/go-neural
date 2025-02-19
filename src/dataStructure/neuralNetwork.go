package NeuralNetwork

import (
	"math/rand"
	"strconv"
)

type Network struct {
	inputLayer   []*Node
	outputLayer  []*Node
	hiddenLayers [][]*Node
}

type Node struct {
	value       float64
	connections *[]Connection
}

type Connection struct {
	weight float64
	to     *Node
	from   *Node
	// Pointers here, are we doing it correctly? remeber, Go will deep copy objects by default when passed to functions
}

func NewNetwork(size []int) Network {
	r := rand.New(rand.NewSource(99))
	inputLayer := make([]*Node, size[0])
	hiddenLayers := make([][]*Node, len(size)-2)
	outputLayer := make([]*Node, size[len(size)-1])

	for i := range outputLayer {
		outputLayer[i] = &Node{r.Float64(), nil}
	}

	for i := len(hiddenLayers) - 1; i >= 0; i-- {
		hiddenLayers[i] = make([]*Node, size[i+1])
		for j := range hiddenLayers[i] {
			if i == len(hiddenLayers)-1 {
				hiddenLayers[i][j] = &Node{r.Float64(), constructConnectionsToLayer(outputLayer, hiddenLayers[i][j])}
			} else {
				hiddenLayers[i][j] = &Node{r.Float64(), constructConnectionsToLayer(hiddenLayers[i+1], hiddenLayers[i][j])}
			}
		}
	}
	// Acceptable nestled loops

	for i := range inputLayer {
		inputLayer[i] = &Node{r.Float64(), constructConnectionsToLayer(hiddenLayers[0], inputLayer[i])}
	}

	return Network{
		inputLayer:   inputLayer,
		outputLayer:  outputLayer,
		hiddenLayers: hiddenLayers,
	}
}

func constructConnectionsToLayer(l []*Node, from *Node) *[]Connection {
	connections := make([]Connection, len(l))
	for i, n := range l {
		connections[i] = Connection{rand.Float64(), n, from}
	}
	return &connections
}

func (n *Network) Print() {
	println("Input Layer:  " + layerToString(n.inputLayer))
	for _, l := range n.hiddenLayers {
		println("Hidden Layer: " + layerToString(l))
	}
	println("outputLayer:  " + layerToString(n.outputLayer))
}

func layerToString(l []*Node) string {
	result := ""
	for _, n := range l {
		result += nodeToString(n) + " "
	}
	return result
	// Ugly empty strings
}

func nodeToString(n *Node) string {
	return strconv.FormatFloat(n.value, 'f', 2, 64)
}

func ArrayMap[A, B any](f func(A) B, a []A) []B {
	b := make([]B, len(a))
	for i, a := range a {
		b[i] = f(a)
	}
	return b
}

func ArrayFilter[A any](f func(A) bool, a []A) []A {
	b := make([]A, 0)
	for _, a := range a {
		if f(a) {
			b = append(b, a)
		}
	}
	return b
}
