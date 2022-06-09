package ffmpeg

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// Encode Splits a video file into a number of tiled .ppm image files containing the frames
func Encode(input string, output string, path string, layout string, customArgs string) (matches int, err error) {
	args := []string{"-i", input, "-y", "-vf", "tile=layout=" + layout, output}
	if customArgs != "" {
		args = append([]string{customArgs}, args...)
	}

	cmd := exec.Command(path, args...)

	// Verbosity
	cmd.Stderr = os.Stdout

	err = cmd.Start()
	if err != nil {
		return 0, err
	}

	err = cmd.Wait()
	if err != nil {
		return 0, err
	}

	// Count the number of files produced
	matchesSlice, err := filepath.Glob("out/frame*.ppm")
	if err != nil {
		return 0, err
	}

	return len(matchesSlice), err
}

// Decode Combines a number of tiled .ppm images containing frames into a video file
func Decode(input string, output string, path string, layout string, fps int, customArgs string) error {
	// if '-f image2' isn't specified before the input file, ffmpeg fails to use wildcards correctly
	// start_number is set to 0, because Fiasco starts its output files at 0
	args := []string{"-f", "image2", "-i", input, "-y", "-vf", "untile=" + layout + ",setpts=N/(" + strconv.Itoa(fps) + "*TB)", "-start_number", "0", output}
	if customArgs != "" {
		args = append([]string{customArgs}, args...)
	}

	cmd := exec.Command(path, args...)

	// Verbosity
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
