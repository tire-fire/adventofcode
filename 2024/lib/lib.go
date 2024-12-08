package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func ReadInput() ([]string, error) {
	day, err := InferDayFromDirectory()
	if err != nil {
		log.Printf("Error figuring out the day: %v", err)
		return nil, err
	}

	filename, err := GetInputFile(day)
	if err != nil {
		log.Printf("Error getting the input filename: %v", err)
		return nil, err
	}

	return ReadLines(filename)
}

// ReadLines reads the content of a file and returns a slice of strings, each representing a line.
func ReadLines(filename string) ([]string, error) {
	log.Printf("Reading file: %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %v", err)
		return nil, err
	}

	return lines, nil
}

// GetSessionCookie retrieves the session cookie from session.txt.
func GetSessionCookie() (string, error) {
	lines, err := ReadLines(filepath.Join("..", "lib", "session.txt"))
	if err != nil {
		return "", fmt.Errorf("failed to read session.txt: %w", err)
	}
	if len(lines) == 0 {
		return "", fmt.Errorf("session.txt is empty")
	}
	return lines[0], nil
}

// DownloadInputFile downloads the input file for the given day.
func DownloadInputFile(day int, sessionCookie string) (string, error) {
	inputURL := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Cookie", "session="+sessionCookie)
	req.Header.Set("User-Agent", "AoC Downloader (Golang)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch input file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	filename := filepath.Join(".", "input")
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write input file: %w", err)
	}

	log.Printf("Input file downloaded and saved to: %s", filename)
	return filename, nil
}

// GetInputFile ensures the input file exists in the current directory or downloads it if necessary.
func GetInputFile(day int) (string, error) {
	filename := filepath.Join(".", "input")
	if _, err := os.Stat(filename); err == nil {
		log.Printf("Input file already exists: %s", filename)
		return filename, nil
	}

	log.Printf("Input file not found; attempting to download...")
	sessionCookie, err := GetSessionCookie()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve session cookie: %w", err)
	}

	return DownloadInputFile(day, sessionCookie)
}

// InferDayFromDirectory infers the day based on the current directory.
func InferDayFromDirectory() (int, error) {
	dir, err := os.Getwd()
	if err != nil {
		return 0, fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Extract the last directory name (expected to be the day number)
	_, dayStr := filepath.Split(dir)
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, fmt.Errorf("failed to infer day from directory: %s", dayStr)
	}

	return day, nil
}

