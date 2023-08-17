package Algorithms

import "goNeural/dataStructure"

type BackProp interface {
	bp(n NeuralNetwork.Network) NeuralNetwork.Network
}

func bp(n NeuralNetwork.Network) NeuralNetwork.Network {
	return n
}
