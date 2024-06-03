package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
            lang: "fr",
            want: "Bonjour le monde",
        },
        "Akkadian, not supported": {
            lang: "akk",
            want: `unsupported language: "akk"`,
        },
        "Greek": {
            lang: "el",
            want: "Χαίρετε Κόσμε",
        },
        "Hebrew": {
            lang: "he",
            want: "שלום עולם",
        },
        "Urdu": {
            lang: "ur",
            want: "ہیلو دنیا",
        },
        "Vietnamese": {
            lang: "vi",
            want: "Xin chào Thế Giới",
        },
        "Empty": {
            lang: "",
            want: `unsupported language: ""`,
        },
	}
    
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("Expected: %q, received: %q", tc.want, got)
			}
		})
	}
}