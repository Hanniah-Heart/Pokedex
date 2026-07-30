package main
import "testing"
import "fmt"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
	// add more cases here
		{
			input:    "  frogs taste nice  ",
			expected: []string{"frogs", "taste", "nice"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		fmt.Printf("The actual word count is %v \n", len(actual))
		fmt.Printf("The expected word count is %v \n", len(c.expected))
		if len(actual) != len(c.expected) {
			// error and continue here
			t.Errorf("Actual number of words found does not equal Expected number of words.")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("\n %v does not match the expected word %v \n", word, expectedWord)

			}
		}
	}
}
