package main

import "goNeural/dataStructure"

func main() {
	nn := neuralNetwork.NewNetwork([]int{10, 8, 8, 8, 10})
	nn.Print()
}
