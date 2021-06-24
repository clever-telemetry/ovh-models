package me

// GENERATED SDK for me API

// Details about your OVH identifier
type Nichandle struct {
	Address                             *string           `json:"address,omitempty"`
	Area                                *string           `json:"area,omitempty"`
	BirthCity                           *string           `json:"birthCity,omitempty"`
	BirthDay                            *string           `json:"birthDay,omitempty"`
	City                                *string           `json:"city,omitempty"`
	CompanyNationalIdentificationNumber *string           `json:"companyNationalIdentificationNumber,omitempty"`
	CorporationType                     *string           `json:"corporationType,omitempty"`
	Country                             CountryEnum       `json:"country,omitempty"`
	Currency                            Currency          `json:"currency,omitempty"`
	CustomerCode                        *string           `json:"customerCode,omitempty"`
	Email                               string            `json:"email,omitempty"`
	Fax                                 *string           `json:"fax,omitempty"`
	Firstname                           *string           `json:"firstname,omitempty"`
	ItalianSDI                          *string           `json:"italianSDI,omitempty"`
	Language                            *LanguageEnum     `json:"language,omitempty"`
	Legalform                           LegalFormEnum     `json:"legalform,omitempty"`
	Name                                *string           `json:"name,omitempty"`
	NationalIdentificationNumber        *string           `json:"nationalIdentificationNumber,omitempty"`
	Nichandle                           string            `json:"nichandle,omitempty"`
	Organisation                        *string           `json:"organisation,omitempty"`
	OvhCompany                          OvhCompanyEnum    `json:"ovhCompany,omitempty"`
	OvhSubsidiary                       OvhSubsidiaryEnum `json:"ovhSubsidiary,omitempty"`
	Phone                               *string           `json:"phone,omitempty"`
	PhoneCountry                        *CountryEnum      `json:"phoneCountry,omitempty"`
	Sex                                 *GenderEnum       `json:"sex,omitempty"`
	SpareEmail                          *string           `json:"spareEmail,omitempty"`
	State                               StateEnum         `json:"state,omitempty"`
	Vat                                 *string           `json:"vat,omitempty"`
	Zip                                 *string           `json:"zip,omitempty"`
}
