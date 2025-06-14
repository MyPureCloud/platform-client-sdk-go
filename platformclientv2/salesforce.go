package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Salesforce
type Salesforce struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Disabled
	Disabled *bool `json:"disabled,omitempty"`

	// IssuerURI
	IssuerURI *string `json:"issuerURI,omitempty"`

	// SsoTargetURI
	SsoTargetURI *string `json:"ssoTargetURI,omitempty"`

	// SloURI
	SloURI *string `json:"sloURI,omitempty"`

	// SloBinding
	SloBinding *string `json:"sloBinding,omitempty"`

	// RelyingPartyIdentifier
	RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`

	// Certificate
	Certificate *string `json:"certificate,omitempty"`

	// Certificates
	Certificates *[]string `json:"certificates,omitempty"`

	// LogoImageData
	LogoImageData *string `json:"logoImageData,omitempty"`

	// NameIdentifierFormat
	NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`

	// SsoBinding
	SsoBinding *string `json:"ssoBinding,omitempty"`

	// SignAuthnRequests
	SignAuthnRequests *bool `json:"signAuthnRequests,omitempty"`

	// ProviderName
	ProviderName *string `json:"providerName,omitempty"`

	// DisplayOnLogin
	DisplayOnLogin *bool `json:"displayOnLogin,omitempty"`

	// MetadataURL
	MetadataURL *string `json:"metadataURL,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Salesforce) SetField(field string, fieldValue interface{}) {
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

func (o Salesforce) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Salesforce
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		IssuerURI *string `json:"issuerURI,omitempty"`
		
		SsoTargetURI *string `json:"ssoTargetURI,omitempty"`
		
		SloURI *string `json:"sloURI,omitempty"`
		
		SloBinding *string `json:"sloBinding,omitempty"`
		
		RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`
		
		Certificate *string `json:"certificate,omitempty"`
		
		Certificates *[]string `json:"certificates,omitempty"`
		
		LogoImageData *string `json:"logoImageData,omitempty"`
		
		NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`
		
		SsoBinding *string `json:"ssoBinding,omitempty"`
		
		SignAuthnRequests *bool `json:"signAuthnRequests,omitempty"`
		
		ProviderName *string `json:"providerName,omitempty"`
		
		DisplayOnLogin *bool `json:"displayOnLogin,omitempty"`
		
		MetadataURL *string `json:"metadataURL,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Disabled: o.Disabled,
		
		IssuerURI: o.IssuerURI,
		
		SsoTargetURI: o.SsoTargetURI,
		
		SloURI: o.SloURI,
		
		SloBinding: o.SloBinding,
		
		RelyingPartyIdentifier: o.RelyingPartyIdentifier,
		
		Certificate: o.Certificate,
		
		Certificates: o.Certificates,
		
		LogoImageData: o.LogoImageData,
		
		NameIdentifierFormat: o.NameIdentifierFormat,
		
		SsoBinding: o.SsoBinding,
		
		SignAuthnRequests: o.SignAuthnRequests,
		
		ProviderName: o.ProviderName,
		
		DisplayOnLogin: o.DisplayOnLogin,
		
		MetadataURL: o.MetadataURL,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Salesforce) UnmarshalJSON(b []byte) error {
	var SalesforceMap map[string]interface{}
	err := json.Unmarshal(b, &SalesforceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SalesforceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SalesforceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Disabled, ok := SalesforceMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
    
	if IssuerURI, ok := SalesforceMap["issuerURI"].(string); ok {
		o.IssuerURI = &IssuerURI
	}
    
	if SsoTargetURI, ok := SalesforceMap["ssoTargetURI"].(string); ok {
		o.SsoTargetURI = &SsoTargetURI
	}
    
	if SloURI, ok := SalesforceMap["sloURI"].(string); ok {
		o.SloURI = &SloURI
	}
    
	if SloBinding, ok := SalesforceMap["sloBinding"].(string); ok {
		o.SloBinding = &SloBinding
	}
    
	if RelyingPartyIdentifier, ok := SalesforceMap["relyingPartyIdentifier"].(string); ok {
		o.RelyingPartyIdentifier = &RelyingPartyIdentifier
	}
    
	if Certificate, ok := SalesforceMap["certificate"].(string); ok {
		o.Certificate = &Certificate
	}
    
	if Certificates, ok := SalesforceMap["certificates"].([]interface{}); ok {
		CertificatesString, _ := json.Marshal(Certificates)
		json.Unmarshal(CertificatesString, &o.Certificates)
	}
	
	if LogoImageData, ok := SalesforceMap["logoImageData"].(string); ok {
		o.LogoImageData = &LogoImageData
	}
    
	if NameIdentifierFormat, ok := SalesforceMap["nameIdentifierFormat"].(string); ok {
		o.NameIdentifierFormat = &NameIdentifierFormat
	}
    
	if SsoBinding, ok := SalesforceMap["ssoBinding"].(string); ok {
		o.SsoBinding = &SsoBinding
	}
    
	if SignAuthnRequests, ok := SalesforceMap["signAuthnRequests"].(bool); ok {
		o.SignAuthnRequests = &SignAuthnRequests
	}
    
	if ProviderName, ok := SalesforceMap["providerName"].(string); ok {
		o.ProviderName = &ProviderName
	}
    
	if DisplayOnLogin, ok := SalesforceMap["displayOnLogin"].(bool); ok {
		o.DisplayOnLogin = &DisplayOnLogin
	}
    
	if MetadataURL, ok := SalesforceMap["metadataURL"].(string); ok {
		o.MetadataURL = &MetadataURL
	}
    
	if SelfUri, ok := SalesforceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Salesforce) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
