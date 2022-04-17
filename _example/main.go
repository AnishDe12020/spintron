// Example application that uses all of the available API options.
package main

import (
	"time"

	"github.com/AnishDe12020/spintron"
)

func main() {
	s := spintron.New(spintron.Options{
		Text: "Loading",
	})

	s.Start()                                                                                                // Start the spinner
	time.Sleep(2 * time.Second)                                                                              // Simulate a long running process
	s.Text = "New character set 👀"                                                                           // Change the text
	s.UpdateCharSet(spintron.CharSets["pong"])                                                               // Update spinner to use a different character set
	time.Sleep(2 * time.Second)                                                                              // Simulate a long running process
	s.Text = "Made the spinner faster :)"                                                                    // Updates spinner text
	s.UpdateSpeed(time.Duration(50) * time.Millisecond)                                                      // Update spinner to use a different speed
	time.Sleep(2 * time.Second)                                                                              // Simulate a long running process
	s.Text = "We have gone back to the default character set and speed and changed the color of the spinner" // Updates spinner text
	s.UpdateCharSet(spintron.CharSets["dots2"])                                                              // Update spinner to use a different character set
	s.Delay = time.Duration(100) * time.Millisecond                                                          // Update spinner to use a different delay
	err := s.Color("red")                                                                                    // Update spinner to use a different color
	if err != nil {
		s.Fail("Uh oh! Something went wrong!") // Stops the spinner and persists it with an error sign and message
	}
	time.Sleep(4 * time.Second)              // Simulate a long running process
	s.Text = "The spinner has been reversed" // Updates spinner text
	s.Reverse()                              // Reverses the spinner character set
	time.Sleep(2 * time.Second)              // Simulate a long running process
	s.Unicorn("Hope you liked the demo!")    // Stops the spinner and persists it with a unicorn sign and message

}
