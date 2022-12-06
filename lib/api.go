package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// client is a client for the AoC service.
var client = newAocClient(getSolutionMetadata())

type aocClient struct {
	year, day, part string

	sessionCookie    string
	sessionCookieErr error
}

// newAocClient creates a new client.
func newAocClient(year, day, part string) *aocClient {
	sessionCookie, err := getSessionCookie()

	return &aocClient{
		year:             year,
		day:              day,
		part:             part,
		sessionCookie:    sessionCookie,
		sessionCookieErr: err,
	}
}

// retrieveInput retrieves the input for a given year/day.
func (a *aocClient) retrieveInput() ([]byte, error) {
	if a.sessionCookieErr != nil {
		return nil, fmt.Errorf("couldn't retrieve session cookie: %s", a.sessionCookieErr)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", a.year, a.day), nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't create get request: %s", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: a.sessionCookie})

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't send request: %s", err)
	}
	defer result.Body.Close()

	input, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read result body: %s", err)
	}

	if bytes.Contains(input, []byte("Puzzle inputs differ by user.  Please log in to get your puzzle input.")) {
		return nil, errors.New("not logged in")
	}

	return input, nil
}

// submitSolution submits the solution for the given part.
func (a *aocClient) submitSolution(solution string) (string, error) {
	if a.sessionCookieErr != nil {
		return "", fmt.Errorf("couldn't retrieve session cookie: %s", a.sessionCookieErr)
	}

	body := fmt.Sprintf("level=%s&answer=%s", a.part, solution)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", a.year, a.day), strings.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("couldn't create get request: %s", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{Name: "session", Value: a.sessionCookie})

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("couldn't send request: %s", err)
	}

	output, err := io.ReadAll(result.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read result body: %s", err)
	}

	outputStr := string(output)

	start := strings.Index(outputStr, "<article>")
	if start == -1 {
		return "", errors.New("couldn't parse response: missing <article>")
	}

	end := strings.Index(outputStr, "</article>")
	if end == -1 {
		return "", errors.New("couldn't parse response: missing </article>")
	}

	return outputStr[start+len("<article>") : end], nil
}

// getSessionCookie retrieves the AOC session cookie. It should be placed in a sessioncookie.txt file in the working directory
// of the executable.
func getSessionCookie() (string, error) {
	file, err := os.Open("sessioncookie.txt")
	if err != nil {
		return "", fmt.Errorf("couldn't open session cookie file: %s", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("couldn't read session cookie file: %s", err)
	}

	return string(data), nil
}
