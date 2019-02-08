package main

import (
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/DATA-DOG/godog/gherkin"
	"net/http"
	"os"
	"testing"
	"time"
)

var request *http.Request
var response *http.Response

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
	var err error
	response, err = http.DefaultClient.Do(request)
	if expected != response.StatusCode {
		return fmt.Errorf("expected response status: \"%d\", actual response status \"%d\"", expected, response.StatusCode)
	}
	return err
}

func theResponseBodyMUSTBe(arg1 *gherkin.DocString) error {
	return godog.ErrPending
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
