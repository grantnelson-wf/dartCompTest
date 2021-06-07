package treatment

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

const (
	// pathSep is the seperator between paths as a string.
	pathSep = string(os.PathSeparator)
)

// Treatment defines how to run an experimental unit.
type Treatment struct {

	// index is used to quickly lookup and add to data for this treatment whilst
	// the treatments are ran in randomized order. Randomizing the order prevents
	// an additional temporal variable being unexpectedly being added to the experiment.
	index int

	// name is the indication of what the treatment is doing.
	name string

	// path is the folder to run the command within.
	path string

	// prepareCmd is the command to run to prepare the treatment.
	// This command will not be measured.
	// If empty then nothing will be run.
	prepareCmd []string

	// runCmd is the command to run and will be measured.
	runCmd []string
}

// New will construct a new treatment that can be run as an experimental unit.
func New(index int) *Treatment {
	return &Treatment{
		index:      index,
		name:       `Unnamed`,
		path:       `.`,
		prepareCmd: nil,
		runCmd:     nil,
	}
}

// Index gets the index this treatment has created with.
func (t *Treatment) Index() int {
	return t.index
}

// Name set the indication of what the treatment is doing.
// To make your life simpler when processing the results, the name shouldn't contain spaces.
// Returns the receiver so that these calls can be chained together.
func (t *Treatment) Name(name string) *Treatment {
	t.name = name
	return t
}

// Path sets the folder to run the command within.
// The multiple parts will be joined by a path separator.
// Returns the receiver so that these calls can be chained together.
func (t *Treatment) Path(parts ...string) *Treatment {
	t.path = path.Join(parts...)
	return t
}

// RunCommand sets the command to run and will be measured.
// Returns the receiver so that these calls can be chained together.
func (t *Treatment) RunCommand(cmd string, args ...string) *Treatment {
	t.runCmd = append([]string{cmd}, args...)
	return t
}

// PrepareCommand is the command to run to prepare the treatment.
// This command will not be measured. If not set then nothing will be run.
// Returns the receiver so that these calls can be chained together.
func (t *Treatment) PrepareCommand(cmd string, args ...string) *Treatment {
	t.prepareCmd = append([]string{cmd}, args...)
	return t
}

// Prepare will prepare the treatment for a run to make it a clean consistent run.
// This will not be measured.
func (t *Treatment) Prepare(timeout time.Duration) {
	if len(t.prepareCmd) > 0 {
		fmt.Printf("   preparing %s\n", t.name)
		_, err := t.performCmd(t.prepareCmd, timeout)
		if err != nil {
			panic(fmt.Errorf("%s failed to prepare: %v", t.name, err))
		}
	}
}

// Run will run the command for this treatment that will be measured.
// Returns the time in seconds it took to execute.
func (t *Treatment) Run(timeout time.Duration) float64 {
	fmt.Printf("  running %s...\n", t.name)
	result, err := t.performCmd(t.runCmd, timeout)
	if err != nil {
		panic(fmt.Errorf("%s failed to run: %v", t.name, err))
	}
	fmt.Printf("  running %s...%f sec\n", t.name, result)
	return result
}

// asyncCmd is the asynchronous part of the performCmd method which is to run
// in its own go-routine and use the given channels to return results.
func (t *Treatment) asyncCmd(cmd *exec.Cmd, resultCh chan float64, errCh chan error) {
	defer func() {
		if r := recover(); r != nil {
			errCh <- fmt.Errorf("%v", r)
		}
	}()

	start := time.Now()
	err := cmd.Run()
	dur := time.Since(start)

	if err != nil {
		errCh <- err
	}
	resultCh <- dur.Seconds()
}

// performCmd performs the given command with the given time out.
// This will return the amount of time it took to run the command or an error.
// This will block until the command has finished being run to keep treatments from overlapping.
func (t *Treatment) performCmd(args []string, timeout time.Duration) (float64, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = t.path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	errCh := make(chan error)
	resultCh := make(chan float64)
	timeoutCh := time.After(timeout)
	go t.asyncCmd(cmd, resultCh, errCh)

	select {
	case <-timeoutCh:
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return 0.0, fmt.Errorf("timed out")
	case err := <-errCh:
		return 0.0, err
	case result := <-resultCh:
		return result, nil
	}
}

// String will return the name of the treatment.
func (t *Treatment) String() string {
	return t.name
}
