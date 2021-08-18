package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Parsedcertificate - Represents the parsed certificate information.
type Parsedcertificate struct { 
	// CertificateDetails - The details of the certificates that were parsed correctly.
	CertificateDetails *[]Certificatedetails `json:"certificateDetails,omitempty"`

}

func (u *Parsedcertificate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Parsedcertificate

	

	return json.Marshal(&struct { 
		CertificateDetails *[]Certificatedetails `json:"certificateDetails,omitempty"`
		*Alias
	}{ 
		CertificateDetails: u.CertificateDetails,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Parsedcertificate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
