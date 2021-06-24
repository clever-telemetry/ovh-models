package me

// GENERATED SDK for me API

// Representation of a Contact
type Contact struct {
	Address                             Address       `json:"address,omitempty"`
	BirthCity                           *string       `json:"birthCity,omitempty"`
	BirthCountry                        *CountryEnum  `json:"birthCountry,omitempty"`
	BirthDay                            *string       `json:"birthDay,omitempty"`
	BirthZip                            *string       `json:"birthZip,omitempty"`
	CellPhone                           *interface{}  `json:"cellPhone,omitempty"`
	CompanyNationalIdentificationNumber *string       `json:"companyNationalIdentificationNumber,omitempty"`
	Email                               string        `json:"email,omitempty"`
	Fax                                 *interface{}  `json:"fax,omitempty"`
	FirstName                           string        `json:"firstName,omitempty"`
	Gender                              *GenderEnum   `json:"gender,omitempty"`
	Id                                  int64         `json:"id,omitempty"`
	Language                            LanguageEnum  `json:"language,omitempty"`
	LastName                            string        `json:"lastName,omitempty"`
	LegalForm                           LegalFormEnum `json:"legalForm,omitempty"`
	NationalIdentificationNumber        *string       `json:"nationalIdentificationNumber,omitempty"`
	Nationality                         *CountryEnum  `json:"nationality,omitempty"`
	OrganisationName                    *string       `json:"organisationName,omitempty"`
	OrganisationType                    *string       `json:"organisationType,omitempty"`
	Phone                               *interface{}  `json:"phone,omitempty"`
	SpareEmail                          *string       `json:"spareEmail,omitempty"`
	Vat                                 *string       `json:"vat,omitempty"`
}
