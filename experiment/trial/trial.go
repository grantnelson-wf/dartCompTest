package trial

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/grantnelson-wf/dartCompTest/experiment/treatment"
)

const (
	// resultPrecision is the precision for the result values.
	resultPrecision = 5
)

type (

	// Trial is an object for running an experiment.
	Trial struct {

		// order is a monotonically increasing treatment order number.
		// This is used to account for temporal variance because it
		// indicates the order for each treatment is run in all replicas.
		order int

		// startTime is the time the experiment started at.
		startTime time.Time

		// prepTimeout is the amount of time to let preparation run before cancelling it
		// and stopping the experiment.
		prepTimeout time.Duration

		// runTimeout is the amount of time to let a treatment run before cancelling it
		// and stopping the experiment.
		runTimeout time.Duration

		// prepWait is the amount of time to wait before running prepare. This normally is
		// zero unless something is holding onto files that need to be removed while cleaning.
		prepWait time.Duration

		// runWait is the amount of unmeasured time to wait after prepare is run before
		// the test is run. Useful for letting a process "cool down" if there are background
		// tasks that need to finish shutting down between runs.
		runWait time.Duration

		// repetitions is the number of times to run the experiment.
		repetitions int

		// resultFile is the result file to write the duration of the application to.
		resultFile string

		// treatments are the full list of treatments in order of addition.
		// The treatment's index will match the index of that treatment in this list.
		treatments []*treatment.Treatment
	}
)

// New constructs a new experiment trial.
// The [repetitions] is the number of times to run each treatment.
// The [resultFile] is the location the results of the experiment are written.
func New(repetitions int, resultFile string) *Trial {
	return &Trial{
		order:       -1,
		startTime:   time.Now(),
		prepTimeout: time.Minute,
		runTimeout:  time.Minute,
		prepWait:    0,
		runWait:     0,
		repetitions: repetitions,
		resultFile:  resultFile,
		treatments:  nil,
	}
}

// AddTreatment adds a treatment that will be run in this experiment.
// The treatment is returned so that it can be configured.
func (t *Trial) AddTreatment() *treatment.Treatment {
	treatment := treatment.New(len(t.treatments))
	t.treatments = append(t.treatments, treatment)
	return treatment
}

// SetTimeouts will set how long each command is given before it times out.
func (t *Trial) SetTimeouts(prepTimeout, runTimeout time.Duration) {
	t.prepTimeout = prepTimeout
	t.runTimeout = runTimeout
}

// SetWaits will set how long to wait, a "cool down" period, before running a command.
func (t *Trial) SetWaits(prepWait, runWait time.Duration) {
	t.prepWait = prepWait
	t.runWait = runWait
}

// Run runs the full experiment.
func (t *Trial) Run() {
	t.startTime = time.Now()

	// Seed the random number generator with current nanoseconds.
	rand.Seed(t.startTime.UTC().UnixNano())

	// Create the results output file.
	results, err := os.Create(t.resultFile)
	if err != nil {
		panic(err)
	}
	defer results.Close()

	// Write data column names to the file.
	results.WriteString("order replicate index name seconds\n")

	// Run all the repetitions of the experiment.
	t.order = 1
	for i := 1; i <= t.repetitions; i++ {
		t.runReplicate(i, results)
	}
}

// wait will pause execution for the given "cool down" time.
func (t *Trial) wait(dur time.Duration) {
	if dur > 0 {
		fmt.Printf("  waiting %f secs\n", dur.Seconds())
		time.Sleep(dur)
	}
}

// runReplicate randomly runs the application of treatments for a replica.
func (t *Trial) runReplicate(replicate int, results *os.File) {
	fmt.Printf("replicate %d of %d\n", replicate, t.repetitions)
	repStartTime := time.Now()

	applications := t.randomizeApplicationOrder()
	for _, treatment := range applications {

		// Prepare the treatment.
		t.wait(t.prepWait)
		treatment.Prepare(t.prepTimeout)

		// Run treatment
		t.wait(t.runWait)
		secs := treatment.Run(t.runTimeout)

		// Write results to the result file.
		result := fmt.Sprintf("%d %d %d %s %.*f\n",
			t.order, replicate, treatment.Index(), treatment.String(), resultPrecision, secs)
		results.WriteString(result)
		t.order++
	}

	// sync the results file so if something happens now
	results.Sync()

	// Display estimated expected amount of time for experiment to finish.
	repDur := time.Since(repStartTime)
	remaining := time.Duration(float64(t.repetitions-replicate) * float64(time.Since(t.startTime)) / float64(replicate))
	fmt.Printf("  took %s, about %s remaining\n", repDur.String(), remaining.String())
}

// randomizeApplicationOrder gets the randomized order of the treatments and the paired order.
// The point of randomizing the order is to remove (at minimum reduce) any temporal/order variable being added into variance,
// otherwise for example a cold start could cause the first run to always seem slower when infact any first run are slower.
func (t *Trial) randomizeApplicationOrder() []*treatment.Treatment {
	length := len(t.treatments)
	randomized := make([]*treatment.Treatment, length)
	copy(randomized, t.treatments)

	rand.Shuffle(length, func(i, j int) {
		randomized[i], randomized[j] = randomized[j], randomized[i]
	})
	return randomized
}
