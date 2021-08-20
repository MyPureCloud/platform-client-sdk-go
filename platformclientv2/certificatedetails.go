package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Certificatedetails - Represents the details of a parsed certificate.
type Certificatedetails struct { 
	// Issuer - Information about the issuer of the certificate.  The value of this property is a comma separated key=value format.  Each key is one of the attribute names supported by X.500.
	Issuer *string `json:"issuer,omitempty"`


	// Subject - Information about the subject of the certificate.  The value of this property is a comma separated key=value format.  Each key is one of the attribute names supported by X.500.
	Subject *string `json:"subject,omitempty"`


	// ExpirationDate - The expiration date of the certificate. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`


	// IssueDate - The issue date of the certificate. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	IssueDate *time.Time `json:"issueDate,omitempty"`


	// Expired - True if the certificate is expired, false otherwise.
	Expired *bool `json:"expired,omitempty"`


	// Valid
	Valid *bool `json:"valid,omitempty"`


	// SignatureValid
	SignatureValid *bool `json:"signatureValid,omitempty"`

}

func (u *Certificatedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Certificatedetails

	
	ExpirationDate := new(string)
	if u.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(u.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	IssueDate := new(string)
	if u.IssueDate != nil {
		
		*IssueDate = timeutil.Strftime(u.IssueDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		IssueDate = nil
	}
	

	return json.Marshal(&struct { 
		Issuer *string `json:"issuer,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		
		IssueDate *string `json:"issueDate,omitempty"`
		
		Expired *bool `json:"expired,omitempty"`
		
		Valid *bool `json:"valid,omitempty"`
		
		SignatureValid *bool `json:"signatureValid,omitempty"`
		*Alias
	}{ 
		Issuer: u.Issuer,
		
		Subject: u.Subject,
		
		ExpirationDate: ExpirationDate,
		
		IssueDate: IssueDate,
		
		Expired: u.Expired,
		
		Valid: u.Valid,
		
		SignatureValid: u.SignatureValid,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Certificatedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
