// Copyright (c) 2021 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spintron

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/mattn/go-isatty"
)

const baseWait = 3

// syncBuffer
type syncBuffer struct {
	sync.Mutex
	bytes.Buffer
}

// Write
func (b *syncBuffer) Write(data []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	return b.Buffer.Write(data)
}

// withOutput
func withOutput(a []string, d time.Duration) (*Spinner, *syncBuffer) {
	var out syncBuffer
	s := New(Options{
		CharacterSet: a,
		Delay:        d,
	})
	s.Writer = &out
	return s, &out
}

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	for i := 0; i < len(CharSets); i++ {
		s := New(Options{
			CharacterSet: CharSets[i],
			Delay:        time.Millisecond * 100,
		})
		if reflect.TypeOf(s).String() != "*spintron.Spinner" {
			t.Errorf("New returned incorrect type kind=%d", i)
		}
	}
}

// TestStart will verify a spinner can be started
func TestStart(t *testing.T) {
	s := New(Options{
		CharacterSet: CharSets[1],
		Delay:        time.Millisecond * 100,
	})
	s.Color("red")
	s.Start()
	time.Sleep(baseWait * time.Second)
	s.Stop()
	time.Sleep(100 * time.Millisecond)
}

// TestStop will verify a spinner can be stopped
func TestStop(t *testing.T) {
	p, out := withOutput(CharSets[14], 100*time.Millisecond)
	p.Color("yellow")
	p.Start()
	time.Sleep(500 * time.Millisecond)
	p.Stop()
	// because the spinner will print an appropriate number of backspaces before stopping,
	// let it complete that sleep
	time.Sleep(100 * time.Millisecond)
	out.Lock()
	len1 := out.Len()
	out.Unlock()
	time.Sleep(300 * time.Millisecond)
	out.Lock()
	defer out.Unlock()
	len2 := out.Len()
	if len1 != len2 {
		t.Errorf("expected equal, got %v != %v", len1, len2)
	}
	p = nil
}

// TestHookFunctions will verify that hook functions works as expected
func TestHookFunctions(t *testing.T) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		t.Log("not running in a termian")
		return
	}
	s := New(Options{
		CharacterSet: CharSets[1],
		Delay:        time.Millisecond * 50,
	})
	var out syncBuffer
	s.Writer = &out
	s.PreUpdate = func(s *Spinner) {
		fmt.Fprintf(s.Writer, "pre-update")
	}
	s.PostUpdate = func(s *Spinner) {
		fmt.Fprintf(s.Writer, "post-update")
	}

	s.Start()
	s.Color("cyan")
	time.Sleep(200 * time.Millisecond)
	s.Stop()
	time.Sleep(50 * time.Millisecond)
	out.Lock()
	defer out.Unlock()
	result := out.Bytes()
	if !bytes.Contains(result, []byte("pre-update")) {
		t.Error("pre-update failed")
	}

	if !bytes.Contains(result, []byte("post-update")) {
		t.Error("post-update failed")
	}
	s = nil
}

// TestReverse will verify that the given spinner can stop and start again reversed
func TestReverse(t *testing.T) {
	a := New(Options{
		CharacterSet: CharSets[1],
		Delay:        time.Millisecond * 50,
	})
	a.Color("red")
	a.Start()
	time.Sleep(baseWait * time.Second)
	a.Reverse()
	a.Restart()
	time.Sleep(baseWait * time.Second)
	a.Reverse()
	a.Restart()
	time.Sleep(baseWait * time.Second)
	a.Stop()
	a = nil
}

// TestUpdateSpeed verifies that the delay can be updated
func TestUpdateSpeed(t *testing.T) {

	s := New(Options{
		CharacterSet: CharSets[1],
		Delay:        time.Second * 1,
	})
	delay1 := s.Delay
	s.UpdateSpeed(baseWait * time.Second)
	delay2 := s.Delay
	if delay1 == delay2 {
		t.Error("update of speed failed")
	}
	s = nil
}

// TestUpdateCharSet verifies that character sets can be updated
func TestUpdateCharSet(t *testing.T) {

	s := New(Options{
		CharacterSet: CharSets[14],
		Delay:        time.Millisecond * 50,
	})
	charSet1 := s.chars
	s.UpdateCharSet(CharSets[1])
	charSet2 := s.chars
	for i := range charSet1 {
		if charSet1[i] == charSet2[i] {
			t.Error("update of char set failed")
		}
	}
	s = nil
}

// TestGenerateNumberSequence verifies that a string slice of a spefic size is returned
func TestGenerateNumberSequence(t *testing.T) {
	elementCount := 100
	seq := GenerateNumberSequence(elementCount)
	if reflect.TypeOf(seq).String() != "[]string" {
		t.Error("received incorrect type in return from GenerateNumberSequence")
	}
	t.Log("In: ", elementCount)
	t.Log("Out: ", len(seq))
	if len(seq) != elementCount {
		t.Error("number of elements in slice doesn't match expected count")
	}
}

// TestBackspace proves that the correct number of characters are removed.
func TestBackspace(t *testing.T) {
	// Because of buffering of output and time weirdness, somethings
	// are broken for an indeterminant reason without a wait
	time.Sleep(75 * time.Millisecond)
	fmt.Println()

	s := New(Options{
		CharacterSet: CharSets[0],
		Delay:        time.Millisecond * 100,
	})
	s.Color("blue")
	s.Start()
	fmt.Print("This is on the same line as the spinner: ")
	time.Sleep(baseWait * time.Second)
	s.Stop()
}

// TestColorError tests that if an invalid color string is passed to the Color
// function, the invalid color error is returned
func TestColorError(t *testing.T) {
	s := New(Options{
		CharacterSet: CharSets[0],
		Delay:        time.Millisecond * 100,
	})

	const invalidColorName = "bluez"
	const validColorName = "green"

	if s.Color(invalidColorName) != errInvalidColor {
		t.Error("Color method did not return an error when given an invalid color.")
	}

	if s.Color(validColorName) != nil {
		t.Error("Color method did not return nil when given a valid color name.")
	}
}

/*
Benchmarks
*/

// BenchmarkNew runs a benchmark for the New() function
func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(Options{
			CharacterSet: CharSets[0],
			Delay:        time.Millisecond * 100,
		})
	}
}

func BenchmarkNewStartStop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := New(Options{
			CharacterSet: CharSets[1],
			Delay:        time.Second * 1,
		})
		s.Start()
		s.Stop()
	}
}
