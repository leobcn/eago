package eago

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDNA(t *testing.T) {
	// rand seed set to 0
	rand.Seed(0)

	// create a new binary-coded DNA
	d := NewDNA(10)
	fmt.Printf("DNA size: %d\n", d.Size())
	fmt.Printf("DNA gene: %s\n", d.Gene())

	// rand seed set to 1
	rand.Seed(1)

	// create a new DNA
	d = NewDNA(10)
	fmt.Printf("DNA size: %d\n", d.Size())
	fmt.Printf("DNA gene: %s\n", d.Gene())

	// bit flip mutation
	mrate := 0.3
	d.Mutate(mrate)
	fmt.Printf("Mutation rate: %f\n", mrate)
	fmt.Printf("Mutated DNA gene: %s\n", d.gene)

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
