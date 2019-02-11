package main

import (
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/christianhujer/assert"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"
)

var request *http.Request
var response *http.Response

var variables map[string]string


var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Paths:  []string{"features"},
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func startBackendFromTest() {
	go func() { StartBackendEphimeral() }()
	time.Sleep(1 * time.Second)
}

func accessing(endpoint string) error {
	startBackendFromTest()
	fmt.Fprintf(os.Stderr, "Accessing http://"+SocketAddress+endpoint+"\n")
	var err error
	request, err = http.NewRequest(HttpGET, "http://"+SocketAddress+endpoint, nil)
	return err
}

func theResponseStatusCodeMUSTBe(expected int) error {
	if err := sendRequestIfNotSent(); err != nil {
		return err
	}
	if expected != response.StatusCode {
		return fmt.Errorf("expected response status: \"%d\", actual response status \"%d\"", expected, response.StatusCode)
	}
	return nil
}

func theResponseBodyMUSTBe(expectedBodyPattern *gherkin.DocString) error {
	if err := sendRequestIfNotSent(); err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// Warning
	// Go totally sucks at encoding.
	// Go only supports UTF-8.
	// We're silently assuming UTF-8 here and not supporting anything else in the tests.
	return equalsWithCaptureAndReplace(string(bodyBytes), expectedBodyPattern.Content)
}

func replace(pattern string) string {
	for name, value := range variables {
		pattern = regexp.MustCompile(`\$\([<]`+name+`\)`).ReplaceAllString(pattern, value)
	}
	return pattern
}

func equalsWithCaptureAndReplace(input string, pattern string) error {
	pattern = regexp.MustCompile(`\(\?<`).ReplaceAllString(pattern, `(?P<`)
	pattern = regexp.MustCompile(`\$\(>([^()]+)\)`).ReplaceAllString(pattern, `(?P<$1>.*?)`)
	regex := regexp.MustCompile(`^` + pattern + `$`)
	match := regex.FindStringSubmatch(input)
	if match != nil {
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				variables[name] = match[i]
			}
		}
	}
	pattern = replace(pattern)
	return assert.True(nil, regexp.MustCompile(pattern).MatchString(input))
}

func sendRequestIfNotSent() error {
	if request != nil {
		//_, _ = fmt.Fprintf(os.Stderr, "Accessing %v\n", request.URL)
		var err error
		response, err = http.DefaultClient.Do(request)
		request = nil
		return err
	}
	return nil
}


func featureContext(s *godog.Suite) {
	s.Step(`^accessing "([^"]*)"$`, accessing)
	s.Step(`^the response Status-Code MUST be (\d+)$`, theResponseStatusCodeMUSTBe)
	s.Step(`^the response body MUST be:$`, theResponseBodyMUSTBe)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()
	opt.Tags = "~@future"
	status := godog.RunWithOptions("acceptance", func(s *godog.Suite) {
		featureContext(s)
	}, opt)
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
