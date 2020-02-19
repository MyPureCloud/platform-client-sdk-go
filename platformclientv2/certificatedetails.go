package platformclientv2
import (
	"time"
	"encoding/json"
)

// Certificatedetails - Represents the details of a parsed certificate.
type Certificatedetails struct { 
	// Issuer - Information about the issuer of the certificate.  The value of this property is a comma separated key=value format.  Each key is one of the attribute names supported by X.500.
	Issuer *string `json:"issuer,omitempty"`


	// Subject - Information about the subject of the certificate.  The value of this property is a comma separated key=value format.  Each key is one of the attribute names supported by X.500.
	Subject *string `json:"subject,omitempty"`


	// ExpirationDate - The expiration date of the certificate. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`


	// IssueDate - The issue date of the certificate. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	IssueDate *time.Time `json:"issueDate,omitempty"`


	// Expired - True if the certificate is expired, false otherwise.
	Expired *bool `json:"expired,omitempty"`


	// SignatureValid
	SignatureValid *bool `json:"signatureValid,omitempty"`


	// Valid
	Valid *bool `json:"valid,omitempty"`

}

// String returns a JSON representation of the model
func (o *Certificatedetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
