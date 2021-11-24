package gomen

import "regexp"

// Type Const ...
const (
	Visa     = "visa"
	Master   = "master"
	Amex     = "amex"
	Diners   = "diners"
	Discover = "discover"
	JCB      = "jcb"
	Other    = "other"
)

// DetectCardType ...
func DetectCardType(number string) string {
	regVisa, _ := regexp.Compile(`^4[0-9]{12}(?:[0-9]{3})?$`)
	regMaster, _ := regexp.Compile(`^5[1-5][0-9]{14}$`)
	regAmex, _ := regexp.Compile(`^3[47][0-9]{13}$`)
	regDiners, _ := regexp.Compile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`)
	regDiscover, _ := regexp.Compile(`^6(?:011|5[0-9]{2})[0-9]{12}$`)
	regJCB, _ := regexp.Compile(`^(?:2131|1800|35\d{3})\d{11}$`)

	reg := map[string]interface{}{
		Visa:     regVisa,
		Master:   regMaster,
		Amex:     regAmex,
		Diners:   regDiners,
		Discover: regDiscover,
		JCB:      regJCB,
	}

	for t, r := range reg {
		if r.(*regexp.Regexp).MatchString(number) {
			return t
		}
	}

	return Other
}
