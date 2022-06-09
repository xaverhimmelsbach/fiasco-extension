package fiasco

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// Encode Encodes several files matching with the input pattern in fiasco
func Encode(inputTemplate string, outputTemplate string, threads int, matches int, path string, customArgs string) error {

	// Calculate parameters for threaded execution
	filesPerThread := matches / threads
	rest := matches % threads
	start := 1
	end := filesPerThread

	var wg sync.WaitGroup

	for i := 1; i <= threads; i++ {
		// Add 1 of the rest frames to current thread if there are any remaining
		if rest > 0 {
			end++
			rest--
		}

		// Insert start and end into input file name
		input := fmt.Sprintf(inputTemplate, start, end)

		// Insert index into output file name
		outputSplit := strings.Split(outputTemplate, ".")
		outputSplit = append(outputSplit[:len(outputSplit)-1], fmt.Sprintf("%03d", i), outputSplit[len(outputSplit)-1])
		output := strings.Join(outputSplit, ".")

		wg.Add(1)

		go func(wg *sync.WaitGroup, input string, output string) {
			// Report thread as done after execution
			defer wg.Done()

			// Only encode to I-Frames for now, as the default pattern causes crashes while decoding
			args := []string{"-V", "2", "-q", "100", "-i", input, "-o", output, "--pattern=I"}
			if customArgs != "" {
				args = append([]string{customArgs}, args...)
			}

			cmd := exec.Command(path, args...)

			// Verbosity
			cmd.Stderr = os.Stdout

			err := cmd.Start()
			if err != nil {
				panic(err)
			}

			_ = cmd.Wait()
		}(&wg, input, output)

		// Update start and end
		start = end + 1
		end = start + filesPerThread - 1

	}

	// Wait for all threads to finish
	wg.Wait()
	return nil
}

// Decode Decodes a fiasco file
func Decode(inputTemplate string, outputTemplate string, threads int, path string, customArgs string) error {

	var wg sync.WaitGroup

	for i := 1; i <= threads; i++ {

		// Insert index into input file name
		inputSplit := strings.Split(inputTemplate, ".")
		inputSplit = append(inputSplit[:len(inputSplit)-1], fmt.Sprintf("%03d", i), inputSplit[len(inputSplit)-1])
		input := strings.Join(inputSplit, ".")

		// Insert index into output file name
		outputSplit := strings.Split(outputTemplate, ".")
		outputSplit = append(outputSplit[:len(outputSplit)-1], fmt.Sprintf("%03d", i), outputSplit[len(outputSplit)-1])
		output := strings.Join(outputSplit, ".")

		wg.Add(1)

		go func(wg *sync.WaitGroup, input string, output string) {
			// Report thread as done after execution
			defer wg.Done()

			args := []string{"-o", output, input}
			if customArgs != "" {
				args = append([]string{customArgs}, args...)
			}

			cmd := exec.Command(path, args...)

			// Verbosity
			cmd.Stderr = os.Stdout

			err := cmd.Start()
			if err != nil {
				panic(err)
			}

			_ = cmd.Wait()
		}(&wg, input, output)
	}

	// Wait for all threads to finish
	wg.Wait()
	return nil
}
