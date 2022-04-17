![Peek 2022-04-16 19-13](https://user-images.githubusercontent.com/63192115/163711658-5da3f280-eb32-42de-8c02-cb33618d67b3.gif)
[![Go Reference](https://pkg.go.dev/badge/github.com/AnishDe12020/spintron.svg)](https://pkg.go.dev/github.com/AnishDe12020/spintron)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/AnishDe12020/spintron)

## Features
- Start/Stop
- Customizable character sets (spinners)
- Custom spinner color, background	
- Custom spinner text
- Restarting and reversing the spinner
- Prefixed text, symbol and padding before the spinner
- Changing spinner settings while it is active
- Getting the spinner settings values
- Chain, pipe or redirect output
- Shows the seconds elasped since the spinner started (can be disabled)
- Ability to Stop and persist the spinner with custom text and a symbol
- Pre-built `Succeed`, `Fail` and `Info` functions that stops and persists the spinner with pre-defined cross-platform symbol and text

## Usage
```go
import "github.com/AnishDe12020/spintron"

func main() {
	s := spintron.New(spintron.Options{
		Text: "Working...",
	})

	s.Start() // Starts the spinner with the initial text
	// Some work that will take time. E.g. Downloading a file. Let us have a dummy
	// function for now
	err := DownloadFile()

	if err != nil {
		s.Fail("Failed to download the file") // Stops the spinner and shows a error sign with the given text
	} else {
		s.Success("Downloaded the file successfully") // Stops the spinner and shows a success sign with the given text
	}
}
```
## Options
When a new spinner is created, it can be created with a struct of options. Here are the ones available - 
### Color (string) (default: cyan)
The color of the spinner
### Text (string)
The text after the spinner
### HideCursor (bool) (defaukt: true)
Hide the cursor or not
### Symbol (string)
A symbol that will come before the prefix text
### PrefixText (string)
Text that will come before the spinner
### CharacterSet (string) (default: dots2)
Character set used for the spinner
### Writer (io.Writer) (default: color.Output)
stdOut writer
### Delay (time.Duration) (default: 100 ms)
Delay between frames in the character set
### Padding (int)
Number of chatacters in padding on the left of the spinner
### DisableElaspedSeconds (bool) (default: false)
Disable the elasped seconds timer


## Examples
### Reversing the spinner
```go
time.Sleep(time.Second * 2) // Simulate a long running process
s.Reverse() // Reverse the spinner's character set
s.Text = "I have been reversed"
time.Sleep(time.Second * 2) // Simulate a long running process
s.Stop() // Stops the spinner
```
### Updating the spinner color
```go
time.Sleep(time.Second * 2)                // Simulate a long running process
s.UpdateCharSet(spintron.CharSets["moon"]) // Update spinner to use a different character set
s.Text = "My character set has been updated"
time.Sleep(time.Second * 2) // Simulate a long running process
s.Stop()                    // Stops the spinner
```
### Updating the spinner speed
```go
time.Sleep(time.Second * 2)                         // Simulate a long running process
s.UpdateSpeed(time.Duration(50) * time.Millisecond) // Update spinner to use a different speed, here making it twice that of the default speed
s.Text = "My speed has been updated to make me faster"
time.Sleep(time.Second * 2) // Simulate a long running process
s.Stop()                    // Stops the spinner
```
### Adding padding to the spinner
```go
time.Sleep(time.Second * 2) // Simulate a long running process
s.Padding = 10              // Add a padding of 10 characters to the left of the spinner
s.Text = "My padding has been updated"
time.Sleep(time.Second * 2) // Simulate a long running process
s.Stop()                    // Stops the spinner
```
### Stop and persist the spinner
```go
time.Sleep(time.Second * 2)   // Simulate a long running process
s.StopAndPersist("👀", "Heya") // Stops the spinner and persists it with a custom symbol and text
```
### Succeeding, Failing or stopping the spinner with an info message
```go
s.Succeed("Done!") // Stops the spinner and persists it with a success sign and message
```
```go
s.Fail("Uh oh! Something went wrong!") // Stops the spinner and persists it with an error sign and message
```
```go
s.Info("Star the repo") // Stops the spinner and persists it with an info sign and message
```

## Credits

All commits uptil [561dc95](https://github.com/AnishDe12020/spinner/commit/561dc95eeadf7fc57c2fe6ce2253f0f3361c0f75) are made by [Brian Downs](https://github.com/briandowns) and the contributors to the [original repository, briandowns/spinner](https://github.com/briandowns/spinner). The project has since been renamed to Spintron to differentiate from the original project.

This is a fork of the [original Go Spinner repository](https://github.com/briandowns/spinner) [Brian Downs](https://github.com/briandowns) licensed under [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0).
