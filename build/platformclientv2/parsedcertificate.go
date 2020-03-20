package platformclientv2
import (
	"encoding/json"
)

// Parsedcertificate - Represents the parsed certificate information.
type Parsedcertificate struct { 
	// CertificateDetails - The details of the certificates that were parsed correctly.
	CertificateDetails *[]Certificatedetails `json:"certificateDetails,omitempty"`

}

// String returns a JSON representation of the model
func (o *Parsedcertificate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
