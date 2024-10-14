/*
(c)Nokia 2024
Author: Maciej Norberciak
companion code for "Unit testing in Go" talk

Test Dive Conference, Krakow, 17.10.2024
*/

package main

import (
	"testing"
)

func TestEnglish(t *testing.T) {
	language := "english"
	name := ""
	expected := "Hello"
	actual := HelloName(language, name)

	if expected != actual {
		t.Errorf("Error, got: %s, expected %s.", actual, expected)
	}
}

/*
func TestEmptyLanguageEmptyName(t *testing.T) {

	language := ""
	name := ""
	expected := " "
	actual := HelloName(language, name)
	if expected != actual {
		t.Errorf("Error, got: %s, expected %s.", actual, expected)
	}

}

func TestSpanish(t *testing.T) {
	language := "spanish"
	name := "Pablo"
	expected := "Hola Pablo"
	actual := HelloName(language, name)

	if expected != actual {
		t.Errorf("Error, got: %s, expected %s.", actual, expected)
	}
}

func TestPolish(t *testing.T) {

	language := "polish"
	name := "Wierzchosław"

	if testing.Short() == false {

		expected := "Cześć Wierzchosław"
		actual := HelloName(language, name)

		if expected != actual {
			t.Errorf("Error, got: %s, expected %s.", actual, expected)
		}
	} else {
		t.Skip()
	}
}

func TestHelloNameTable(t *testing.T) {
	testCases := []struct {
		language string
		name     string
		expected string
	}{
		{"english", "Alice", "Hello Alice"},
		{"spanish", "Pablo", "Hola Pablo"},
		{"polish", "Wierzchosław", "Cześć Wierzchosław"},
		{"german", "Hans", "Hallo  Hans"},
		{"polish", "", "Cześć "},
		//{"japanese", "武", "今日は 武"},
		{"", "Frank", " Frank"},     // Test case for empty language
		{"french", "Gina", " Gina"}, // Test case for unsupported language

	}

	for _, tc := range testCases {
		actual := HelloName(tc.language, tc.name)
		if actual != tc.expected {
			t.Errorf("Error, got: %s, expected %s.", actual, tc.expected)
		}
	}

}

func TestHelloNameTableSub(t *testing.T) {
	testCases := []struct {
		language string
		name     string
		expected string
	}{
		{"english", "Alice", "Hello Alice"},
		{"spanish", "Bob", "Hola Bob"},
		{"polish", "Wierzchosław", "Cześć Wierzchosław"},
		{"german", "Hans", "Hallo  Hans"},
		{"polish", "", "Cześć "},
		{"", "Frank", " Frank"},     // Test case for empty language
		{"french", "Gina", " Gina"}, // Test case for unsupported language
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s+%s", tc.name, tc.language), func(t *testing.T) {
			t.Parallel()
			actual := HelloName(tc.language, tc.name)
			if actual != tc.expected {
				t.Errorf("Error, got: %s, expected %s.", actual, tc.expected)
			}
		})
	}
}

/*
//this is a sample benchmark
//you can run it with
//go test -benchmem -bench .
//to see all the information
//if you don't comment out the rest of the suite you need
//to add -run=^$

func BenchmarkHelloName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloName("", "")
	}

}

/*
//this is how you can set up the test fixture for
//individual test funcitons and tear it down
//on test function exit
//Cleanup works just like defer, but only
//for test functions
//
//run this test function and notice the oder
//of messages on the console
//

func TestDummy(t *testing.T) {
	t.Logf("Test fixture setup\n")
	t.Cleanup(func() {
		t.Logf("Cleaning up the fixture\n")
	})

	t.Logf("Actual test code goes here\n")
}

//this function is implicitly called on startup
//of the test suite (~file with tests)
//you can initialize your fixture here
//no way to tear it down in the similar manner though
//
//run any test with this function uncommented and notice
//the messagess on the console

func TestMain(m *testing.M) {
	log.Println("Put your setup here")
	exitVal := m.Run()
	log.Println("Put your teardown here")

	os.Exit(exitVal)
}

*/
