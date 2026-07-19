package execute

import (
	"fmt"

	"github.com/tom96da/sleepingknights/pkg"
)

func displayUsage() ExitStatus {
	fmt.Println("Usage: slk [options] [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  run <script>   Run the specified script")

	fmt.Println("Options:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")

	return ExitStatusSuccess
}

func displayVersion() ExitStatus {
	fmt.Printf("SleepingKnights  %s (%s)\n", pkg.Version, pkg.Hash)
	return ExitStatusSuccess
}

func displayArgsError(message string) ExitStatus {
	fmt.Printf("[Error] %s\n", message)
	fmt.Println()
	displayUsage()
	return ExitStatusInvalidCommandLineArgs
}

func startInteractiveMode() ExitStatus {
	displayVersion()
	return ExitStatusSuccess
}

func CommandLine(commandLineArgs []string) ExitStatus {
	if len(commandLineArgs) == 0 {
		return startInteractiveMode()
	}
	if commandLineArgs[0] == "-h" || commandLineArgs[0] == "--help" {
		return displayUsage()
	}
	if commandLineArgs[0] == "-v" || commandLineArgs[0] == "--version" {
		return displayVersion()
	}
	if commandLineArgs[0] == "run" {
		if len(commandLineArgs) < 2 {
			return displayArgsError("Missing script path")
		}
		return compileAndExecute(commandLineArgs[1])
	}
	if isScriptFile(commandLineArgs[0]) {
		return executeScript(commandLineArgs[0])
	}

	return displayArgsError(fmt.Sprintf("Unknown command: %s", commandLineArgs[0]))
}

func isScriptFile(path string) bool {
	return len(path) > 4 && path[len(path)-4:] == ".slk"
}
