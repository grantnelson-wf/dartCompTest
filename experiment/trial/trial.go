package trial

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Grant-Nelson/dartCompTest/experiment/treatment"
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

	results.WriteString("Start Time: " + t.startTime.String() + "\n")
	results.WriteString("order replicate index name seconds\n")

	// Run all the repetitions of the experiment.
	t.order = 1
	for i := 1; i <= t.repetitions; i++ {
		t.runReplicate(i, results)
	}
}

// runReplicate randomly runs the application of treatments for a replica.
func (t *Trial) runReplicate(replicate int, results *os.File) {
	fmt.Printf("replicate %d of %d\n", replicate, t.repetitions)
	repStartTime := time.Now()

	applications := t.randomizeApplicationOrder()
	for _, treatment := range applications {

		// Run cleanup
		treatment.Cleanup()

		// Run treatment
		secs := treatment.Run()

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
