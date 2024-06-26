package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	testCase(t, []string{"main.go", "hello"}, " _              _   _          \n| |            | | | |         \n| |__     ___  | | | |   ___   \n|  _ \\   / _ \\ | | | |  / _ \\  \n| | | | |  __/ | | | | | (_) | \n|_| |_|  \\___| |_| |_|  \\___/  \n                               \n                               \n")
	testCase(t, []string{"main.go", "HELLO"}, " _    _   ______   _        _         ____   \n| |  | | |  ____| | |      | |       / __ \\  \n| |__| | | |__    | |      | |      | |  | | \n|  __  | |  __|   | |      | |      | |  | | \n| |  | | | |____  | |____  | |____  | |__| | \n|_|  |_| |______| |______| |______|  \\____/  \n                                             \n                                             \n")
	testCase(t, []string{"main.go", "HeLlo HuMaN"}, " _    _          _        _                 _    _           __  __           _   _  \n| |  | |        | |      | |               | |  | |         |  \\/  |         | \\ | | \n| |__| |   ___  | |      | |   ___         | |__| |  _   _  | \\  / |   __ _  |  \\| | \n|  __  |  / _ \\ | |      | |  / _ \\        |  __  | | | | | | |\\/| |  / _` | | . ` | \n| |  | | |  __/ | |____  | | | (_) |       | |  | | | |_| | | |  | | | (_| | | |\\  | \n|_|  |_|  \\___| |______| |_|  \\___/        |_|  |_|  \\__,_| |_|  |_|  \\__,_| |_| \\_| \n                                                                                     \n                                                                                     \n")
	testCase(t, []string{"main.go", "1Hello 2There"}, "     _    _          _   _                         _______   _                           \n _  | |  | |        | | | |                ____   |__   __| | |                          \n/ | | |__| |   ___  | | | |   ___         |___ \\     | |    | |__     ___   _ __    ___  \n| | |  __  |  / _ \\ | | | |  / _ \\          __) |    | |    |  _ \\   / _ \\ | '__|  / _ \\ \n| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ \n|_| |_|  |_|  \\___| |_| |_|  \\___/        |_____|    |_|    |_| |_|  \\___| |_|     \\___| \n                                                                                         \n                                                                                         \n")
	testCase(t, []string{"main.go", "\\!\" #$%&'()*+,-./"}, "__       _   _ _           _  _      _    _   __           _    __ __       _                                   __ \n\\ \\     | | ( | )        _| || |_   | |  (_) / /   ___    ( )  / / \\ \\   /\\| |/\\     _                         / / \n \\ \\    | |  V V        |_  __  _| / __)    / /   ( _ )   |/  | |   | |  \\ ` ' /   _| |_       ______         / /  \n  \\ \\   | |              _| || |_  \\__ \\   / /    / _ \\/\\     | |   | | |_     _| |_   _|     |______|       / /   \n   \\ \\  |_|             |_  __  _| (   /  / / _  | (_>  <     | |   | |  / , . \\    |_|    _            _   / /    \n    \\_\\ (_)               |_||_|    |_|  /_/ (_)  \\___/\\/     | |   | |  \\/|_|\\/          ( )          (_) /_/     \n                                                               \\_\\ /_/                    |/                       \n                                                                                                                   \n")
	testCase(t, []string{"main.go", "{|}~"}, "   __  _  __     /\\/| \n  / / | | \\ \\   |/\\/  \n | |  | |  | |        \n/ /   | |   \\ \\       \n\\ \\   | |   / /       \n | |  | |  | |        \n  \\_\\ | | /_/         \n      |_|             \n")
}

func testCase(t *testing.T, args []string, expectedOutput string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = args
	main()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if output != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, output)
	}
}
