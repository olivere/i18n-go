package i18n

import (
	"strings"
)

// Address is a generic structure for a location.
type Address struct {
	StreetAddress   string
	ExtendedAddress string
	Locality        string
	PostalCode      string
	Region          string
	Country         string
}

// FormattedParts returns an array of strings that, when combined,
// form a country-specific address. If no rules can be found for the given
// address, the German address format will be used.
func (a *Address) FormattedParts() []string {
	rule, found := formatRules[strings.ToUpper(a.Country)]
	if !found {
		rule = formatRules["DE"]
	}
	return a.formatWithRule(rule)
}

// formatWithRule generates an array of strings according to the given rule.
func (a *Address) formatWithRule(r []string) []string {
	parts := make([]string, 0)
	for _, line := range r {
		s := line
		s = strings.Replace(s, "{StreetAddress}", a.StreetAddress, -1)
		s = strings.Replace(s, "{ExtendedAddress}", a.ExtendedAddress, -1)
		s = strings.Replace(s, "{POBox}", "", -1)
		s = strings.Replace(s, "{PostalCode}", a.PostalCode, -1)
		s = strings.Replace(s, "{Locality}", a.Locality, -1)
		s = strings.Replace(s, "{Region}", a.Region, -1)
		if strings.Contains(s, "{-Region}") {
			if a.Region != "" {
				s = strings.Replace(s, "{-Region}", "-"+a.Region, -1)
			} else {
				s = strings.Replace(s, "{-Region}", "", -1)
			}
		}
		if strings.Contains(s, "{, Region}") {
			if a.Region != "" {
				s = strings.Replace(s, "{, Region}", ", "+a.Region, -1)
			} else {
				s = strings.Replace(s, "{, Region}", "", -1)
			}
		}
		if a.Country != "" {
			if t, found := Territories[strings.ToUpper(a.Country)]; found {
				s = strings.Replace(s, "{Country}", t.EnglishName, -1)
			} else {
				s = strings.Replace(s, "{Country}", a.Country, -1)
			}
		} else {
			s = strings.Replace(s, "{Country}", "", -1)
		}
		parts = append(parts, s)
	}

	// Trim and remove empty parts
	cleanedParts := make([]string, 0)
	for _, part := range parts {
		s := strings.TrimSpace(part)
		if s != "" {
			cleanedParts = append(cleanedParts, s)
		}
	}

	return cleanedParts
}

// formatRules specifies the country specific formatting rules for countries.
// See e.g. http://bitboost.com/ref/international-address-formats.html
// for international address formats.
var formatRules = map[string][]string{
	"AT": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "A-{PostalCode} {Locality}", "{Country}"},
	"AR": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"AU": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality} {Region} {PostalCode}", "{Country}"},
	"BE": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "B-{PostalCode} {Locality}", "{Country}"},
	"BR": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality}{-Region}", "{PostalCode}", "{Country}"},
	"CA": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality} {Region} {PostalCode}", "{Country}"},
	"CH": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality} {Region}", "{Country}"},
	"CL": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"CN": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Region}", "{Country}"},
	"CZ": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"DE": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"DK": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "DK-{PostalCode} {Locality}", "{Country}"},
	"EE": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"ES": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"FI": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "FI-{PostalCode} {Locality}", "{Country}"},
	"FR": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"HK": []string{"{StreetAddress}", "{ExtendedAddress}", "{Region}", "{Locality}", "{Country}"},
	"GB": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality}", "{Region}", "{PostalCode}", "{Country}"},
	"IN": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality} {PostalCode}", "{Country}"},
	"IL": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"IT": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality} {Region}", "{Country}"},
	"JP": []string{"{StreetAddress}", "{ExtendedAddress}", "{PostalCode} {Locality}", "{Region}", "{Country}"},
	"KR": []string{"{StreetAddress}", "{ExtendedAddress}", "{Region}", "{Locality} {PostalCode}", "{Country}"},
	"LU": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "L-{PostalCode} {Locality}", "{Country}"},
	"MY": []string{"{StreetAddress}", "{ExtendedAddress}", "{PostalCode} {Locality}", "{Region}", "{Country}"},
	"MX": []string{"{StreetAddress}", "{ExtendedAddress}", "{PostalCode} {Locality}{, Region}", "{Country}"},
	"NL": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"NZ": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality} {PostalCode}", "{Country}"},
	"NO": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "NO-{PostalCode} {Locality}", "{Country}"},
	"PL": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"PT": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Country}"},
	"RO": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{PostalCode} {Locality}", "{Region}", "{Country}"},
	"RU": []string{"{StreetAddress}", "{ExtendedAddress}", "{Region}", "{Locality}", "{PostalCode}", "{Country}"},
	"SA": []string{"{StreetAddress}", "{ExtendedAddress}", "{Region}", "{Locality}", "{PostalCode}", "{Country}"},
	"SE": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "SE-{PostalCode} {Locality}", "{Country}"},
	"SG": []string{"{StreetAddress}", "{ExtendedAddress}", "{Locality} {PostalCode}", "{Country}"},
	"US": []string{"{StreetAddress}", "{ExtendedAddress}", "{POBox}", "{Locality} {Region} {PostalCode}", "{Country}"},
}
