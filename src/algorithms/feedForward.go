package Algorithms

import (
	"goNeural/dataStructure"
	"math"
)

func FeedForward(nw NeuralNetwork.Network) NeuralNetwork.Network {

	return nw
}

func nodeFeedForward(c *NeuralNetwork.Connection) {

}

func activationFunction(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
