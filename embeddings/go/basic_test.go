package main

import (
	"fmt"
	"testing"
)

func TestWasm2WatSuccess(t *testing.T) {
	wasmer := Wasmer{}
	wasmer.New()
	result, err := wasmer.Wasm2wat([]byte{
		0, 97, 115, 109, 1, 0, 0, 0, 0,
		8, 4, 110, 97, 109, 101, 2, 1, 0,
		0, 9, 7, 108, 105, 110, 107, 105, 110,
		103, 2,
	}, 1)
	if err != nil {
		t.Error(err)
	}
	if result != "(module)\n" {
		t.Errorf("!")
	}
}

func TestWasm2WatFail(t *testing.T) {
	wasmer := Wasmer{}
	wasmer.New()
	_, err := wasmer.Wasm2wat([]byte{0}, 1)
	if err != nil {
		formatted_error := fmt.Sprint(err)
		if formatted_error != "0000000: error: unable to read uint32_t: magic\n" {
			t.Errorf("Unexpected intentional error: %s", formatted_error)
		}
	} else {
		t.Error("Wasm2wat should have errored")
	}
}
