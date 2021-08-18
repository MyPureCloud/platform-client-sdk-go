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


	// Certificate
	Certificate *string `json:"certificate,omitempty"`


	// Certificates
	Certificates *[]string `json:"certificates,omitempty"`


	// RelyingPartyIdentifier
	RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`


	// LogoImageData
	LogoImageData *string `json:"logoImageData,omitempty"`


	// EndpointCompression
	EndpointCompression *bool `json:"endpointCompression,omitempty"`


	// NameIdentifierFormat
	NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Genericsaml) MarshalJSON() ([]byte, error) {
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
		
		Certificate *string `json:"certificate,omitempty"`
		
		Certificates *[]string `json:"certificates,omitempty"`
		
		RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`
		
		LogoImageData *string `json:"logoImageData,omitempty"`
		
		EndpointCompression *bool `json:"endpointCompression,omitempty"`
		
		NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Disabled: u.Disabled,
		
		IssuerURI: u.IssuerURI,
		
		SsoTargetURI: u.SsoTargetURI,
		
		SloURI: u.SloURI,
		
		SloBinding: u.SloBinding,
		
		Certificate: u.Certificate,
		
		Certificates: u.Certificates,
		
		RelyingPartyIdentifier: u.RelyingPartyIdentifier,
		
		LogoImageData: u.LogoImageData,
		
		EndpointCompression: u.EndpointCompression,
		
		NameIdentifierFormat: u.NameIdentifierFormat,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Genericsaml) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
