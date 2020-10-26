package i18n

import (
	"testing"
)

func TestLanguages(t *testing.T) {
	var tests = []struct {
		code                string
		found               bool
		expectedEnglishName string
		expectedNativeName  string
	}{
		/* 0 */ {"xy", false, "", ""},
		/* 1 */ {"de", true, "German", "Deutsch"},
		/* 2 */ {"en", true, "English", "English"},
		/* 3 */ {"fr", true, "French", "français"},
		/* 4 */ {"es", true, "Spanish", "español"},
		/* 5 */ {"zh", true, "Chinese (Simplified)", "中文(简体)"},
		/* 6 */ {"no", true, "Norwegian", "Norsk"},
		/* 7 */ {"nn", true, "Norwegian Nynorsk", "Nynorsk"},
		/* 8 */ {"nb", true, "Norwegian Bokmål", "Norsk bokmål"},
	}

	for i, f := range tests {
		l, found := Languages[f.code]
		if found != f.found {
			t.Fatalf("%d. expected language %s found flag to be %v, got %v", i, f.code, f.found, found)
		}
		if f.found {
			if l == nil {
				t.Fatalf("%d. expected language to be != nil", i)
			}
			if f.expectedEnglishName != l.EnglishName {
				t.Errorf("%d. expected EnglishName to be %v, got %v", i, f.expectedEnglishName, l.EnglishName)
			}
			if f.expectedNativeName != l.NativeName {
				t.Errorf("%d. expected NativeName to be %v, got %v", i, f.expectedNativeName, l.NativeName)
			}
		}
	}
}
