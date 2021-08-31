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

func (o *Certificatedetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Certificatedetails
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	IssueDate := new(string)
	if o.IssueDate != nil {
		
		*IssueDate = timeutil.Strftime(o.IssueDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Issuer: o.Issuer,
		
		Subject: o.Subject,
		
		ExpirationDate: ExpirationDate,
		
		IssueDate: IssueDate,
		
		Expired: o.Expired,
		
		Valid: o.Valid,
		
		SignatureValid: o.SignatureValid,
		Alias:    (*Alias)(o),
	})
}

func (o *Certificatedetails) UnmarshalJSON(b []byte) error {
	var CertificatedetailsMap map[string]interface{}
	err := json.Unmarshal(b, &CertificatedetailsMap)
	if err != nil {
		return err
	}
	
	if Issuer, ok := CertificatedetailsMap["issuer"].(string); ok {
		o.Issuer = &Issuer
	}
	
	if Subject, ok := CertificatedetailsMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if expirationDateString, ok := CertificatedetailsMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	
	if issueDateString, ok := CertificatedetailsMap["issueDate"].(string); ok {
		IssueDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", issueDateString)
		o.IssueDate = &IssueDate
	}
	
	if Expired, ok := CertificatedetailsMap["expired"].(bool); ok {
		o.Expired = &Expired
	}
	
	if Valid, ok := CertificatedetailsMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
	
	if SignatureValid, ok := CertificatedetailsMap["signatureValid"].(bool); ok {
		o.SignatureValid = &SignatureValid
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Certificatedetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
