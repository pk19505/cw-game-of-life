package gol

// Params provides the details of how to run the Game of Life and which image to load.
type Params struct {
	Turns       int
	Threads     int
	ImageWidth  int
	ImageHeight int
}

// Run starts the processing of Game of Life. It should initialise channels and goroutines.
func Run(p Params, events chan<- Event, keyPresses <-chan rune) {

	ioCommand := make(chan ioCommand)
	ioFilename := make(chan string)
	ioIdle := make(chan bool)
	ioOutput := make(chan uint8)
	ioInput := make(chan uint8)

	distributorChannels := distributorChannels{
		events,
		ioCommand,
		ioIdle,
		ioFilename,
		ioOutput,
		ioInput,
		keyPresses,
	}
	go distributor(p, distributorChannels)

	ioChannels := ioChannels{
		command:  ioCommand,
		idle:     ioIdle,
		filename: ioFilename,
		output:   ioOutput,
		input:    ioInput,
	}
	go startIo(p, ioChannels)
}
