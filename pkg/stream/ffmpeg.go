package stream

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
)

func runFfmpegCommand(options []string) <-chan error {

	cmd := exec.Command("ffmpeg", options...)

	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)

	stderr, _ := cmd.StderrPipe()
	errorScanner := bufio.NewScanner(stderr)

	done := make(chan error)

	err := cmd.Start()

	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		for errorScanner.Scan() {
			fmt.Println(errorScanner.Text())
		}

		done <- cmd.Wait()
	}()

	return done
}

func StreamVideo(source, target string, streamLoop int) error {
	inputOptions := []string{
		"-re",
	}

	if streamLoop != 0 {
		inputOptions = append(inputOptions, []string{"-stream_loop", strconv.Itoa(streamLoop)}...)
	}

	inputOptions = append(inputOptions, []string{"-i", source}...)

	outputOptions := []string{
		"-vf", "pad=ceil(iw/2)*2:ceil(ih/2)*2",
		"-c:v", "libx264",
		"-preset", "veryfast",
		"-b:v", "6000k",
		"-maxrate", "6000k",
		"-bufsize", "6000k",
		"-pix_fmt", "yuv420p",
		"-g", "50",
		"-c:a", "aac",
		"-b:a", "160k",
		"-ac", "2",
		"-ar", "44100",
		"-f", "flv",
		target,
	}

	options := []string{}

	options = append(options, inputOptions...)
	options = append(options, outputOptions...)

	return <-runFfmpegCommand(options)
}
