// Package demo is an autogenerated survey ready to use with "github.com/AlecAivazis/survey/v2".
// Code generated by github.com/guumaster/surveygen, DO NOT EDIT.
// source: demo/my_awesome_survey.yaml
package demo

/*
  Example Usage:

	func main() {
		s := demo.CreateMyAwesomeSurvey()

		err := survey.Ask(s.Questions, s.Answers)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Answers: \n", s.Answers)
	}
*/

import (
	"github.com/AlecAivazis/survey/v2"
)

// SurveyMyAwesomeSurvey is an autogenerated struct for survey "MyAwesomeSurvey".
type SurveyMyAwesomeSurvey struct {
	Questions []*survey.Question
	Answers   *AnswersMyAwesomeSurvey
}

// AnswersMyAwesomeSurvey is a struct that will contain survey's answers.
type AnswersMyAwesomeSurvey struct {
	Age     int      `survey:"Age"`
	Like    bool     `survey:"Like"`
	Level   string   `survey:"level"`
	Comment string   `survey:"comment"`
	Country []string `survey:"country"`
}

// CreateMyAwesomeSurvey creates a SurveyMyAwesomeSurvey struct with .Questions and .Answers ready to use with survey.Ask()
func CreateMyAwesomeSurvey() *SurveyMyAwesomeSurvey {
	return &SurveyMyAwesomeSurvey{
		Answers: &AnswersMyAwesomeSurvey{},
		Questions: []*survey.Question{
			{
				Name: "Age",
				Prompt: &survey.Input{
					Message: "How old are you",
					Help:    "Just be honest about it",
				},
				Validate: survey.Required,
			},
			{
				Name: "Like",
				Prompt: &survey.Confirm{
					Message: "Do you like this survey",
					Help:    "It would save you some time and you know it.",
					Default: true,
				},
				Validate: survey.Required,
			}, {
				Name: "level",
				Prompt: &survey.Select{
					Message: "How much do you love it",

					Options: []string{
						"Just a bit",
						"Regular",
						"With all my heart!",
					},
				},
				Validate: survey.Required,
			}, {
				Name: "comment",
				Prompt: &survey.Multiline{
					Message: "Enter a comment",
				},
				Validate: survey.Required,
			}, {
				Name: "country",
				Prompt: &survey.MultiSelect{
					Message: "Choose some countries just for fun",

					Options: []string{
						"Afghanistan",
						"Albania",
						"Algeria",
						"American Samoa",
						"Andorra",
						"Angola",
						"Anguilla",
						"Antarctica",
						"Antigua and Barbuda",
						"Argentina",
						"Armenia",
						"Aruba",
						"Australia",
						"Austria",
						"Azerbaijan",
						"Bahamas",
						"Bahrain",
						"Bangladesh",
						"Barbados",
						"Belarus",
						"Belgium",
						"Belize",
						"Benin",
						"Bermuda",
						"Bhutan",
						"Bolivia",
						"Bosnia and Herzegovina",
						"Botswana",
						"Bouvet Island",
						"Brazil",
						"British Indian Ocean Territory",
						"Brunei Darussalam",
						"Bulgaria",
						"Burkina Faso",
						"Burundi",
						"Cambodia",
						"Cameroon",
						"Canada",
						"Cape Verde",
						"Cayman Islands",
						"Central African Republic",
						"Chad",
						"Chile",
						"China",
						"Christmas Island",
						"Cocos (Keeling) Islands",
						"Colombia",
						"Comoros",
						"Congo",
						"Congo, The Democratic Republic of the",
						"Cook Islands",
						"Costa Rica",
						"Cote D'Ivoire",
						"Croatia",
						"Cuba",
						"Cyprus",
						"Czech Republic",
						"Denmark",
						"Djibouti",
						"Dominica",
						"Dominican Republic",
						"Ecuador",
						"Egypt",
						"El Salvador",
						"Equatorial Guinea",
						"Eritrea",
						"Estonia",
						"Ethiopia",
						"Falkland Islands (Malvinas)",
						"Faroe Islands",
						"Fiji",
						"Finland",
						"France",
						"French Guiana",
						"French Polynesia",
						"French Southern Territories",
						"Gabon",
						"Gambia",
						"Georgia",
						"Germany",
						"Ghana",
						"Gibraltar",
						"Greece",
						"Greenland",
						"Grenada",
						"Guadeloupe",
						"Guam",
						"Guatemala",
						"Guernsey",
						"Guinea",
						"Guinea-Bissau",
						"Guyana",
						"Haiti",
						"Heard Island and Mcdonald Islands",
						"Holy See (Vatican City State)",
						"Honduras",
						"Hong Kong",
						"Hungary",
						"Iceland",
						"India",
						"Indonesia",
						"Iran, Islamic Republic Of",
						"Iraq",
						"Ireland",
						"Isle of Man",
						"Israel",
						"Italy",
						"Jamaica",
						"Japan",
						"Jersey",
						"Jordan",
						"Kazakhstan",
						"Kenya",
						"Kiribati",
						"Korea, Democratic People'S Republic of",
						"Korea, Republic of",
						"Kuwait",
						"Kyrgyzstan",
						"Lao People'S Democratic Republic",
						"Latvia",
						"Lebanon",
						"Lesotho",
						"Liberia",
						"Libyan Arab Jamahiriya",
						"Liechtenstein",
						"Lithuania",
						"Luxembourg",
						"Macao",
						"Macedonia, The Former Yugoslav Republic of",
						"Madagascar",
						"Malawi",
						"Malaysia",
						"Maldives",
						"Mali",
						"Malta",
						"Marshall Islands",
						"Martinique",
						"Mauritania",
						"Mauritius",
						"Mayotte",
						"Mexico",
						"Micronesia, Federated States of",
						"Moldova, Republic of",
						"Monaco",
						"Mongolia",
						"Montserrat",
						"Morocco",
						"Mozambique",
						"Myanmar",
						"Namibia",
						"Nauru",
						"Nepal",
						"Netherlands",
						"Netherlands Antilles",
						"New Caledonia",
						"New Zealand",
						"Nicaragua",
						"Niger",
						"Nigeria",
						"Niue",
						"Norfolk Island",
						"Northern Mariana Islands",
						"Norway",
						"Oman",
						"Pakistan",
						"Palau",
						"Palestinian Territory, Occupied",
						"Panama",
						"Papua New Guinea",
						"Paraguay",
						"Peru",
						"Philippines",
						"Pitcairn",
						"Poland",
						"Portugal",
						"Puerto Rico",
						"Qatar",
						"Reunion",
						"Romania",
						"Russian Federation",
						"RWANDA",
						"Saint Helena",
						"Saint Kitts and Nevis",
						"Saint Lucia",
						"Saint Pierre and Miquelon",
						"Saint Vincent and the Grenadines",
						"Samoa",
						"San Marino",
						"Sao Tome and Principe",
						"Saudi Arabia",
						"Senegal",
						"Serbia and Montenegro",
						"Seychelles",
						"Sierra Leone",
						"Singapore",
						"Slovakia",
						"Slovenia",
						"Solomon Islands",
						"Somalia",
						"South Africa",
						"South Georgia and the South Sandwich Islands",
						"Spain",
						"Sri Lanka",
						"Sudan",
						"Suriname",
						"Svalbard and Jan Mayen",
						"Swaziland",
						"Sweden",
						"Switzerland",
						"Syrian Arab Republic",
						"Taiwan, Province of China",
						"Tajikistan",
						"Tanzania, United Republic of",
						"Thailand",
						"Timor-Leste",
						"Togo",
						"Tokelau",
						"Tonga",
						"Trinidad and Tobago",
						"Tunisia",
						"Turkey",
						"Turkmenistan",
						"Turks and Caicos Islands",
						"Tuvalu",
						"Uganda",
						"Ukraine",
						"United Arab Emirates",
						"United Kingdom",
						"United States",
						"United States Minor Outlying Islands",
						"Uruguay",
						"Uzbekistan",
						"Vanuatu",
						"Venezuela",
						"Viet Nam",
						"Virgin Islands, British",
						"Virgin Islands, U.S.",
						"Wallis and Futuna",
						"Western Sahara",
						"Yemen",
						"Zambia",
						"Zimbabwe",
					},
				},
				Validate: survey.Required,
			},
		},
	}
}
