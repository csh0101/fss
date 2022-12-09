package version

import "fmt"

var (
	gitVersion   string
	gitCommit    string
	gitBranch    string
	gitTreeState string
	buildTime    string
	env          string
)

// Print prints the version attached at the compile time.
func Print() {
	fmt.Printf("%-12s : %s\n", "GitVersion", gitVersion)
	fmt.Printf("%-12s : %s\n", "GitCommit", gitCommit)
	fmt.Printf("%-12s : %s\n", "GitBranch", gitBranch)
	fmt.Printf("%-12s : %s\n", "GitTreeState", gitTreeState)
	fmt.Printf("%-12s : %s\n", "BuildTime", buildTime)
	fmt.Printf("%-12s : %s\n", "Environment", env)
}

func Version() string {
	return gitVersion
}
