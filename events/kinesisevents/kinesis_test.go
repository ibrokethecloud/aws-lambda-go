// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
package kinesisevents

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events/test"
)

func TestKinesisEventMarshaling(t *testing.T) {

	// 1. read JSON from file
	inputJson := readJsonFromFile(t)

	// 2. de-serialize into Go object
	var inputEvent KinesisEvent
	if err := json.Unmarshal(inputJson, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// 3. serialize to JSON
	outputJson, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	// 4. check result
	test.AssertJsonsEqual(t, inputJson, outputJson)
}

func readJsonFromFile(t *testing.T) []byte {
	inputJson, err := ioutil.ReadFile("./test-files/kinesis-event.json")
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJson
}

func TestKinesisMarshalingMalformedJson(t *testing.T) {
	test.TestMalformedJson(t, KinesisEvent{})
}