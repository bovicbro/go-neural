package main

import (
	"goNeural/algorithms"
	"goNeural/dataStructure"
)

func main() {
	nn := NeuralNetwork.NewNetwork([]int{10, 8, 8, 8, 10})
	nn.Print()
	nn = Algorithms.FeedForward(nn)
}
