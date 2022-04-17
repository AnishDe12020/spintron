![Peek 2022-04-16 19-13](https://user-images.githubusercontent.com/63192115/163711658-5da3f280-eb32-42de-8c02-cb33618d67b3.gif)
[![Go Reference](https://pkg.go.dev/badge/github.com/AnishDe12020/spintron.svg)](https://pkg.go.dev/github.com/AnishDe12020/spintron)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/AnishDe12020/spintron)

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

## Credits

All commits uptil [561dc95](https://github.com/AnishDe12020/spinner/commit/561dc95eeadf7fc57c2fe6ce2253f0f3361c0f75) are made by [Brian Downs](https://github.com/briandowns) and the contributors to the [original repository, briandowns/spinner](https://github.com/briandowns/spinner). The project has since been renamed to Spintron to differentiate from the original project.

This is a fork of the [original Go Spinner repository](https://github.com/briandowns/spinner) [Brian Downs](https://github.com/briandowns) licensed under [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0).
