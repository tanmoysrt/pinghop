package main

func main() {
	// start server
	go pingHopServer()
	// Start CLI
	pingHopCli()
}
