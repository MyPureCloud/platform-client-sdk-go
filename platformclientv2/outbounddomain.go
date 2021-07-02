package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Outbounddomain
type Outbounddomain struct { 
	// Id - Unique Id of the domain such as: example.com
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CnameVerificationResult - CNAME registration Status
	CnameVerificationResult *Verificationresult `json:"cnameVerificationResult,omitempty"`


	// DkimVerificationResult - DKIM registration Status
	DkimVerificationResult *Verificationresult `json:"dkimVerificationResult,omitempty"`


	// FromEmail - The email that is used to display sender informations
	FromEmail *Emailaddress `json:"fromEmail,omitempty"`


	// ReplyToEmail - The email that is used if replies are expected
	ReplyToEmail *Emailaddress `json:"replyToEmail,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outbounddomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
