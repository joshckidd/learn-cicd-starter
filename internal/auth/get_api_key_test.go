package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		input    http.Header
		expected string
		err      error
	}{
		{
			input: http.Header{
				"Authorization": []string{"ApiKey abc125"},
			},
			expected: "abc125",
			err:      nil,
		},
		{
			input: http.Header{
				"content-length": []string{"0"},
			},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
	}

	for _, c := range cases {
		actual, err := GetAPIKey(c.input)
		fmt.Println(c.input.Get("Authorization"))

		if actual != c.expected {
			t.Errorf("Expected key %s, got key %s", c.expected, actual)
		}

		if err != nil && c.err == nil {
			t.Errorf("Expected nil error, got %s", err.Error())
		} else if err == nil && c.err != nil {
			t.Errorf("Expercted error %s, got nil", c.err.Error())
		} else if err != nil && c.err != nil && err.Error() != c.err.Error() {
			t.Errorf("Expected error %s, got error %s", c.err.Error(), err.Error())
		}
	}

}
