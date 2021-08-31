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

func (o *Parsedcertificate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Parsedcertificate
	
	return json.Marshal(&struct { 
		CertificateDetails *[]Certificatedetails `json:"certificateDetails,omitempty"`
		*Alias
	}{ 
		CertificateDetails: o.CertificateDetails,
		Alias:    (*Alias)(o),
	})
}

func (o *Parsedcertificate) UnmarshalJSON(b []byte) error {
	var ParsedcertificateMap map[string]interface{}
	err := json.Unmarshal(b, &ParsedcertificateMap)
	if err != nil {
		return err
	}
	
	if CertificateDetails, ok := ParsedcertificateMap["certificateDetails"].([]interface{}); ok {
		CertificateDetailsString, _ := json.Marshal(CertificateDetails)
		json.Unmarshal(CertificateDetailsString, &o.CertificateDetails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Parsedcertificate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
