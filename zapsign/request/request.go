package request

import (
	"errors"
)

// CreateNewModelOptions  is create new model
type Model struct {
	TemplateID            string  `json:"template_id"`             //TemplateID template id for substituition
	SignerName            string  `json:"signer_name"`             //SignerName signer name
	ExternalID            string  `json:"external_id"`             //ExternalID for systems referencies
	SendAutomaticEmail    bool    `json:"send_automatic_email"`    // SendAutomaticEmail flag for send email notification
	SendAutomaticWhatsapp bool    `json:"send_automatic_whatsapp"` // SendAutomaticEmail flag for send whatsapp notification
	PhoneCountry          string  `json:"phone_country"`
	PhoneNumber           string  `json:"phone_number"`
	Data                  []Datum `json:"data"` // Data is key value for include text dinamiclly into document
}

// Docs is zapsign document request body
type Docs struct {
	Name              string   `json:"name"`                // Name is name to docs
	Base64Pdf         string   `json:"base64_pdf"`          // Base64Pdf is file to upload
	BrandPrimaryColor string   `json:"brand_primary_color"` // BrandPrimaryColor primary color for documentation
	Observers         []string `json:"observers"`
	Signers           []Signer `json:"signers"`
}

func (d *Docs) AddOnserver(o string) *Docs {
	d.Observers = append(d.Observers, o)
	return &Docs{}
}

func (d *Docs) AddSigners(s Signer) *Docs {
	d.Signers = append(d.Signers, s)
	return &Docs{}
}

// NewDocs create a new instance from Docs
func NewDocs(name, base64Pdf, brandPrimaryColor string) (*Docs, error) {

	if name == "" {
		return nil, errors.New("name is not should be null")
	}
	if base64Pdf == "" {
		return nil, errors.New("base64pdf is not should be null")
	}

	return &Docs{
		Name:              name,
		Base64Pdf:         base64Pdf,
		BrandPrimaryColor: brandPrimaryColor,
	}, nil
}

// Attachment is a representation to new doc attachment
type Attachment struct {
	Name      string `json:"name"`
	Base64PDF string `json:"base64_pdf"`
}

// Datum variables for substituition in document
type Datum struct {
	Key   string `json:"de"`
	Value string `json:"para"`
}

// Signers document signer
type Signer struct {
	Name                  string `json:"name"`
	Email                 string `json:"email,omitempty"`
	SendAutomaticEmail    bool   `json:"send_automatic_email,omitempty"`
	SendAutomaticWhatsapp bool   `json:"send_automatic_whatsapp,omitempty"`
	CustomMessage         string `json:"custom_message,omitempty"`
	LockEmail             bool   `json:"lock_email,omitempty"`
	LockPhone             bool   `json:"lock_phone,omitempty"`
	PhoneCountry          string `json:"phone_country,omitempty"`
	PhoneNumber           string `json:"phone_number,omitempty"`
	AuthMode              string `json:"auth_mode,omitempty"`
}

type Webhook struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}
