package main

/*
	Nowadays, machine learning is all the rage, and part of machine learning is the use of neural networks.
	So let's see if we can model a very simple neuron network using the Go programming language.
*/

type Neuron struct {
	In, Out []*Neuron // Incoming and outcoming connections
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

/*
	Now, imagine the situation becomes more complicated. Imagine it is not convenient to work with
	the individual neurons, but instead is preferable to work with neuron layers
	(a collection of neurons all stored together).
*/

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

/*
	We want to be able to have a method for connecting
	neurons to neurons, neurons to layers, layers to neurons and layers to layers.
*/

type NeuronInterface interface {
	Iter() []*Neuron
}

// We have implemented this interface in both a collection object and a scalar object
func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n} // In a way, we get a scalar object to masquerade as if it were a collection
}

func Connect(left, right NeuronInterface) { // Both Neuron and NeuronLayer implement NeuronInterface
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			// This ConnectTo method is actually defined on a neuron,
			// so this interconnects one neuron with another.
			l.ConnectTo(r)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron2)
	Connect(layer1, layer2)
}
