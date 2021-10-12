package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	// Print a random silly name
	fmt.Println(randomdata.SillyName())

	// Print a male title
	fmt.Println(randomdata.Title(randomdata.Male))

	// Print a female title
	fmt.Println(randomdata.Title(randomdata.Female))

	// Print a title with random gender
	fmt.Println(randomdata.Title(randomdata.RandomGender))

	// Print a male first name
	fmt.Println(randomdata.FirstName(randomdata.Male))

	// Print a female first name
	fmt.Println(randomdata.FirstName(randomdata.Female))

	// Print a last name
	fmt.Println(randomdata.LastName())

	// Print a male name
	fmt.Println(randomdata.FullName(randomdata.Male))

	// Print a female name
	fmt.Println(randomdata.FullName(randomdata.Female))

	// Print a name with random gender
	fmt.Println(randomdata.FullName(randomdata.RandomGender))

	// Print an email
	fmt.Println(randomdata.Email())

	// Print a country with full text representation
	fmt.Println(randomdata.Country(randomdata.FullCountry))

	// Print a country using ISO 3166-1 alpha-2
	fmt.Println(randomdata.Country(randomdata.TwoCharCountry))

	// Print a country using ISO 3166-1 alpha-3
	fmt.Println(randomdata.Country(randomdata.ThreeCharCountry))

	// Print BCP 47 language tag
	fmt.Println(randomdata.Locale())

	// Print a currency using ISO 4217
	fmt.Println(randomdata.Currency())

	// Print the name of a random city
	fmt.Println(randomdata.City())

	// Print the name of a random american state
	fmt.Println(randomdata.State(randomdata.Large))

	// Print the name of a random american state using two chars
	fmt.Println(randomdata.State(randomdata.Small))

	// Print an american sounding street name
	fmt.Println(randomdata.Street())

	// Print an american sounding address
	fmt.Println(randomdata.Address())

	// Print a random number >= 10 and < 20
	fmt.Println(randomdata.Number(10, 20))

	// Print a number >= 0 and < 20
	fmt.Println(randomdata.Number(20))

	// Print a random float >= 0 and < 20 with decimal point 3
	fmt.Println(randomdata.Decimal(0, 20, 3))

	// Print a random float >= 10 and < 20
	fmt.Println(randomdata.Decimal(10, 20))

	// Print a random float >= 0 and < 20
	fmt.Println(randomdata.Decimal(20))

	// Print a bool
	fmt.Println(randomdata.Boolean())

	// Print a paragraph
	fmt.Println(randomdata.Paragraph())

	// Print a postal code
	fmt.Println(randomdata.PostalCode("SE"))

	// Print a set of 2 random numbers as a string
	fmt.Println(randomdata.StringNumber(2, "-"))

	// Print a set of 2 random 3-Digits numbers as a string
	fmt.Println(randomdata.StringNumberExt(2, "-", 3))

	// Print a random string sampled from a list of strings
	fmt.Println(randomdata.StringSample("my string 1", "my string 2", "my string 3"))

	// Print a valid random IPv4 address
	fmt.Println(randomdata.IpV4Address())

	// Print a valid random IPv6 address
	fmt.Println(randomdata.IpV6Address())

	// Print a browser's user agent string
	fmt.Println(randomdata.UserAgentString())

	// Print a day
	fmt.Println(randomdata.Day())

	// Print a month
	fmt.Println(randomdata.Month())

	// Print full date like Monday 22 Aug 2016
	fmt.Println(randomdata.FullDate())

	// Print full date <= Monday 22 Aug 2016
	fmt.Println(randomdata.FullDateInRange("2016-08-22"))

	// Print full date >= Monday 01 Aug 2016 and <= Monday 22 Aug 2016
	fmt.Println(randomdata.FullDateInRange("2016-08-01", "2016-08-22"))

	// Print phone number according to e.164
	fmt.Println(randomdata.PhoneNumber())

	// Get a complete and randomised profile of data generally used for users
	// There are many fields in the profile to use check the Profile struct definition in fullprofile.go
	profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)
	fmt.Printf("The new profile's username is: %s and password (md5): %s\n", profile.Login.Username, profile.Login.Md5)

	// Get a random country-localised street name for Great Britain
	fmt.Println(randomdata.StreetForCountry("GB"))
	// Get a random country-localised street name for USA
	fmt.Println(randomdata.StreetForCountry("US"))

	// Get a random country-localised province for Great Britain
	fmt.Println(randomdata.ProvinceForCountry("GB"))
	// Get a random country-localised province for USA
	fmt.Println(randomdata.ProvinceForCountry("US"))
}
