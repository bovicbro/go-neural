package FeedForward

import "goNeural/dataStructure"

type FeedForward interface {
	ff(n neuralNetwork.Network) neuralNetwork.Network
}

func ff(n neuralNetwork.Network) neuralNetwork.Network {
	return n
}
