package ne

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDNAFloat64(t *testing.T) {
	// rand seed set to 2
	rand.Seed(2)

	// create a new DNAFloat64
	df := NewDNAFloat64(10, 6.0, 2.0)
	fmt.Printf("DNA size: %d\n", df.Size())
	fmt.Printf("DNA gene: %f\n", df.Gene())

	// rand seed set to 3
	rand.Seed(3)

	df = NewDNAFloat64(10, 0.0, 1.0)
	fmt.Printf("DNA size: %d\n", df.Size())
	fmt.Printf("DNA gene: %f\n", df.Gene())

	// gaussian convolution mutation
	mrate = 0.5
	df.Mutate(mrate)
	fmt.Printf("Mutation rate: %f\n", mrate)
	fmt.Printf("Mutated DNA gene: %f\n", df.gene)

}
