// Example application that uses all of the available API options.
package main

import (
	"time"

	"github.com/AnishDe12020/spintron"
)

func main() {
	s := spintron.New(spintron.Options{
		Text:    "Loading",
		Delay:   time.Duration(400) * time.Millisecond,
		Padding: 2,
	})

	s.Color("red")              // Set the spinner color to red
	s.Start()                   // Start the spinner
	time.Sleep(1 * time.Second) // Run for some time to simulate work
	s.PrefixText = "Loading"
	time.Sleep(1 * time.Second)
	// s.Succeed("success")
	s.Info("info")

	// s.UpdateCharSet(spinner.CharSets[9])  // Update spinner to use a different character set
	// s.UpdateSpeed(100 * time.Millisecond) // Update the speed the spinner spins at

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
