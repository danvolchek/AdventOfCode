package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type aocClient struct {
	metadata solutionMetadata

	sessionCookie    string
	sessionCookieErr error
}

// newAocClient creates a new client.
func newAocClient(metadata solutionMetadata) *aocClient {
	sessionCookie, err := getSessionCookie()

	return &aocClient{
		metadata:         metadata,
		sessionCookie:    sessionCookie,
		sessionCookieErr: err,
	}
}

// retrieveInput retrieves the input for a given year/day.
func (a *aocClient) retrieveInput() ([]byte, error) {
	if a.sessionCookieErr != nil {
		return nil, fmt.Errorf("couldn't retrieve session cookie: %s", a.sessionCookieErr)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", a.metadata.year, a.metadata.day), nil)
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

	if bytes.Contains(input, []byte("You don't seem to be solving the right level.  Did you already complete it?")) {
		return nil, errors.New("already solved")
	}

	if bytes.Contains(input, []byte("Please don't repeatedly request this endpoint before it unlocks!")) {
		return nil, errors.New("not unlocked yet")
	}

	return input, nil
}

// submitSolution submits the solution for the given part.
func (a *aocClient) submitSolution(solution string) (string, error) {
	solution = url.QueryEscape(solution)
	if a.sessionCookieErr != nil {
		return "", fmt.Errorf("couldn't retrieve session cookie: %s", a.sessionCookieErr)
	}

	body := fmt.Sprintf("level=%s&answer=%s", a.metadata.part, solution)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", a.metadata.year, a.metadata.day), strings.NewReader(body))
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

	display := outputStr[start+len("<article>") : end]

	if strings.Contains(display, "You don't seem to be solving the right level.  Did you already complete it?") {
		return "", errors.New("already solved")
	}

	if strings.Contains(display, "You gave an answer too recently; you have to wait after submitting an answer before trying again.") {
		return "", fmt.Errorf("rate limited: wait %s", display[strings.Index(display, "You have ")+len("You have "):strings.Index(display, " left to wait.")])
	}

	return display, nil
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
