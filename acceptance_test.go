package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func accessing(arg1 string) error {
	return godog.ErrPending
}

func theResponseStatusCodeMUSTBe(arg1 int) error {
	return godog.ErrPending
}

func theResponseBodyMUSTBe(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^accessing "([^"]*)"$`, accessing)
	s.Step(`^the response Status-Code MUST be (\d+)$`, theResponseStatusCodeMUSTBe)
	s.Step(`^the response body MUST be:$`, theResponseBodyMUSTBe)
}
