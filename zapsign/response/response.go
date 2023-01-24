package response

type Docs struct {
	OpenID       int64       `json:"open_id"`
	Token        string      `json:"token"`
	Status       string      `json:"status"`
	Name         string      `json:"name"`
	OriginalFile string      `json:"original_file"`
	SignedFile   interface{} `json:"signed_file"`
	CreatedAt    string      `json:"created_at"`
	LastUpdateAt string      `json:"last_update_at"`
	Signers      []Signer    `json:"signers"`
	Answers      []Answer    `json:"answers"`
}

type Attachment struct {
	OpenID       int64       `json:"open_id"`
	Token        string      `json:"token"`
	Name         string      `json:"name"`
	OriginalFile string      `json:"original_file"`
	SignedFile   interface{} `json:"signed_file"`
}

type Answer struct {
	Variable string `json:"variable"`
	Value    string `json:"value"`
}

type Signer struct {
	Token        string      `json:"token"`
	SignURL      string      `json:"sign_url"`
	Status       string      `json:"status"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	PhoneCountry string      `json:"phone_country"`
	PhoneNumber  string      `json:"phone_number"`
	TimesViewed  int64       `json:"times_viewed"`
	LastViewAt   interface{} `json:"last_view_at"`
	SignedAt     interface{} `json:"signed_at"`
}
