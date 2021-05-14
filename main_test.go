package main

import (
	"testing"
)

func TestPermutate(t *testing.T) {

	domains := []string{
		"google.com",
		"microsoft.com",
		"apple.com",
		"tesla.com",
		"infected-database.org",
		"stackoverflow.com",
		"golang.org",
		"oracle.com",
		"github.com",
		"youtube.com",
	}
	permutationList := []string{"mail", "vpn", "admin", "www", "remote"}

	out := make(chan string, 1)

	go func() {
		defer close(out)
		for _, domain := range domains {
			Permutate(out, &domain, &permutationList)
		}
	}()

	count := 0

	for range out {
		count++
	}

	expected := len(domains) * len(permutationList)

	if expected != count {
		t.Errorf("expected %d, got %d", expected, count)
	}

}
