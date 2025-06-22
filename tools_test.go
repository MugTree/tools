package tools

import (
	"testing"
)

func Test_TitleCaseUnderscoredNames(t *testing.T) {

	input := []string{"is_valid"}
	expectedOutput := "IsValid"

	for _, v := range input {

		result := underscoredToTitleCase(v)
		if result != expectedOutput {
			t.Errorf("Results incorrect, got %s, was expecting %s", result, expectedOutput)
		}

	}

}
