package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

var (
	testDeck  Deck
	testCards []Card
)

func TestMain(m *testing.M) {
	testCards = []Card{
		{"Queen", "diamond", 12},
		{"3", "spade", 3},
		{"Ace", "heart", 14},
		{"7", "club", 7},
	}
	testDeck = Deck(testCards)

	code := m.Run()
	os.Exit(code)
}

func TestArg(t *testing.T) {
	testCases := map[string]struct {
		args   []string
		result string
	}{
		"with no filepath": {
			args:   []string{"exec"},
			result: "",
		},
		"with a filepath": {
			args:   []string{"exec", "file.txt"},
			result: "file.txt",
		},
		"with an extra arg": {
			args:   []string{"exec", "file.txt", "extra"},
			result: "file.txt",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := arg(tc.args)
			if got != tc.result {
				t.Errorf("got %s want %s", got, tc.result)
			}
		})
	}
}

func TestExitWithError(t *testing.T) {
	if os.Getenv("OS_EXIT_CALLED") == "1" {
		exitWithError(errors.New("error"))
		return
	}
	subTest := exec.Command(os.Args[0], "-test.run=TestExitWithError")
	subTest.Env = append(os.Environ(), "OS_EXIT_CALLED=1")
	err := subTest.Run()
	if exitError, ok := err.(*exec.ExitError); !ok || exitError.Success() {
		t.Error("process exited with no error, wanted exit status 1")
	}
}
