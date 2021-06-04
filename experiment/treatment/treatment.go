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

// Run will run the command for this treatment that will be measured.
// Returns the time in seconds it took to execute.
func (t *Treatment) Run() float64 {
	fmt.Printf("  running %s...", t.name)
	cmd := exec.Command(t.runCmd[0], t.runCmd[1:]...)
	cmd.Dir = t.path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	start := time.Now()
	err := cmd.Run()
	dur := time.Since(start)

	if err != nil {
		panic(fmt.Errorf("%s failed to run: %v", t.name, err))
	}

	fmt.Printf("%s\n", dur.String())
	return dur.Seconds()
}

// Prepare will prepare the treatment for a run to make it a clean consistent run.
// This will not be measured.
func (t *Treatment) Prepare() {
	if len(t.prepareCmd) > 0 {
		fmt.Printf("   preparing %s\n", t.name)
		cmd := exec.Command(t.prepareCmd[0], t.prepareCmd[1:]...)
		cmd.Dir = t.path
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(fmt.Errorf("%s failed to prepare: %v", t.name, err))
		}
	}
}

// String will return the name of the treatment.
func (t *Treatment) String() string {
	return t.name
}
