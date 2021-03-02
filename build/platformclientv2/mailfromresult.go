package platformclientv2
import (
	"encoding/json"
)

// Mailfromresult
type Mailfromresult struct { 
	// Status - The verification status.
	Status *string `json:"status,omitempty"`


	// Records - The list of DNS records that pertain that need to exist for verification.
	Records *[]Record `json:"records,omitempty"`


	// MailFromDomain - The custom MAIL FROM domain.
	MailFromDomain *string `json:"mailFromDomain,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mailfromresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
