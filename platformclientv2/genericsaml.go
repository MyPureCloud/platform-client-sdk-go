package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Genericsaml
type Genericsaml struct { 
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


	// EndpointCompression
	EndpointCompression *bool `json:"endpointCompression,omitempty"`


	// NameIdentifierFormat
	NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Genericsaml) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Genericsaml
	
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
		
		EndpointCompression *bool `json:"endpointCompression,omitempty"`
		
		NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
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
		
		EndpointCompression: o.EndpointCompression,
		
		NameIdentifierFormat: o.NameIdentifierFormat,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Genericsaml) UnmarshalJSON(b []byte) error {
	var GenericsamlMap map[string]interface{}
	err := json.Unmarshal(b, &GenericsamlMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GenericsamlMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := GenericsamlMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Disabled, ok := GenericsamlMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
    
	if IssuerURI, ok := GenericsamlMap["issuerURI"].(string); ok {
		o.IssuerURI = &IssuerURI
	}
    
	if SsoTargetURI, ok := GenericsamlMap["ssoTargetURI"].(string); ok {
		o.SsoTargetURI = &SsoTargetURI
	}
    
	if SloURI, ok := GenericsamlMap["sloURI"].(string); ok {
		o.SloURI = &SloURI
	}
    
	if SloBinding, ok := GenericsamlMap["sloBinding"].(string); ok {
		o.SloBinding = &SloBinding
	}
    
	if RelyingPartyIdentifier, ok := GenericsamlMap["relyingPartyIdentifier"].(string); ok {
		o.RelyingPartyIdentifier = &RelyingPartyIdentifier
	}
    
	if Certificate, ok := GenericsamlMap["certificate"].(string); ok {
		o.Certificate = &Certificate
	}
    
	if Certificates, ok := GenericsamlMap["certificates"].([]interface{}); ok {
		CertificatesString, _ := json.Marshal(Certificates)
		json.Unmarshal(CertificatesString, &o.Certificates)
	}
	
	if LogoImageData, ok := GenericsamlMap["logoImageData"].(string); ok {
		o.LogoImageData = &LogoImageData
	}
    
	if EndpointCompression, ok := GenericsamlMap["endpointCompression"].(bool); ok {
		o.EndpointCompression = &EndpointCompression
	}
    
	if NameIdentifierFormat, ok := GenericsamlMap["nameIdentifierFormat"].(string); ok {
		o.NameIdentifierFormat = &NameIdentifierFormat
	}
    
	if SelfUri, ok := GenericsamlMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Genericsaml) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
