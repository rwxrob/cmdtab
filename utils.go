package cmdtab

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Found returns true if the given path was absolutely found to exist on
// the system. A false return value means either the file does not
// exists or it was not able to determine if it exists or not. WARNING:
// do not use this function if a definitive check for the non-existence
// of a file is required since the possible indeterminate error state is
// a possibility. These checks are also not atomic on many file systems
// so avoid this usage for pseudo-semaphore designs and depend on file
// locks.
func Found(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ArgsFromStdin converts each line of standard input into a slice of
// strings suitable for using as arguments. If fields is passed any
// occurance of {n} will be replaced with the appropriate field in order
// with n beginning at 1.
func ArgsFromStdin(fields ...string) []string {
	args := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for n, v := range fields {
			lk := fmt.Sprintf("{%v}", n+1)
			line = strings.ReplaceAll(line, lk, v)
		}
		args = append(args, line)
	}
	return args
}
