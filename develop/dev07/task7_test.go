package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testOr struct {
	channels []<-chan interface{}
}

var testOrs = []testOr{
	{[]<-chan interface{}{MakeChan(2 * time.Hour), MakeChan(5 * time.Minute), MakeChan(1 * time.Second), MakeChan(1 * time.Hour), MakeChan(1 * time.Minute)}},
}

func TestOr(t *testing.T) {
	for _, test := range testOrs {
		start := time.Now()
		<-Union(test.channels...)
		assert.Less(t, time.Since(start)-time.Second, time.Millisecond*50)
	}
}

func TestMakeChan(t *testing.T) {
	c := MakeChan(500 * time.Millisecond)

	select {
	case <-c:
		// test passed
	case <-time.After(1 * time.Second):
		t.Errorf("MakeChan took too long to complete")
	}
}
