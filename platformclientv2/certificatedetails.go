package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Certificatedetails - Represents the details of a parsed certificate.
type Certificatedetails struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Certificatedetails) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Certificatedetails) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ExpirationDate","IssueDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
	}{ 
		Issuer: o.Issuer,
		
		Subject: o.Subject,
		
		ExpirationDate: ExpirationDate,
		
		IssueDate: IssueDate,
		
		Expired: o.Expired,
		
		Valid: o.Valid,
		
		SignatureValid: o.SignatureValid,
		Alias:    (Alias)(o),
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
