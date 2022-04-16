// Example application that uses all of the available API options.
package main

import (
	"time"

	"github.com/AnishDe12020/spintron"
)

func main() {
	s := spintron.New(spintron.Options{
		Text:                  "Loading",
		Delay:                 time.Duration(100) * time.Millisecond,
		DisableElaspedSeconds: false,
	})

	s.Start() // Start the spinner
	time.Sleep(2 * time.Second)
	s.Text = "New character set 👀"             // Change the text
	s.UpdateCharSet(spintron.CharSets["pong"]) // Update spinner to use a different character set
	time.Sleep(2 * time.Second)
	s.Text = "Made the spinner faster :)"
	s.UpdateSpeed(time.Duration(50) * time.Millisecond) // Update spinner to use a different speed
	time.Sleep(2 * time.Second)
	s.Text = "We have gone back to the default character set and speed and changed the color of the spinner"
	s.UpdateCharSet(spintron.CharSets["dots2"])
	s.Delay = time.Duration(100) * time.Millisecond
	err := s.Color("red") // Update spinner to use a different color
	if err != nil {
		s.Fail("Uh oh! Something went wrong!")
	}
	time.Sleep(4 * time.Second)
	s.Text = "The spinner has been reversed"
	s.Reverse()
	time.Sleep(2 * time.Second)
	s.Unicorn("Hope you liked the demo!")

	// color.Green("\n:)")

	// if err := s.Color("yellow"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.Start()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("red"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[20])
	// s.Reverse()
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("blue"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[3])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("cyan"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[28])
	// s.Reverse()
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("green"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[25])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("magenta"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[32])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("white"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[31])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// s.Stop() // Stop the spinner

	// s.UpdateCharSet(spinner.CharSets[39])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// s.Stop() // Stop the spinner

	// println("")
}
