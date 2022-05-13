package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Certificate - Represents a certificate to parse.
type Certificate struct { 
	// Certificate - The certificate to parse.
	Certificate *string `json:"certificate,omitempty"`

}

func (o *Certificate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Certificate
	
	return json.Marshal(&struct { 
		Certificate *string `json:"certificate,omitempty"`
		*Alias
	}{ 
		Certificate: o.Certificate,
		Alias:    (*Alias)(o),
	})
}

func (o *Certificate) UnmarshalJSON(b []byte) error {
	var CertificateMap map[string]interface{}
	err := json.Unmarshal(b, &CertificateMap)
	if err != nil {
		return err
	}
	
	if Certificate, ok := CertificateMap["certificate"].(string); ok {
		o.Certificate = &Certificate
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Certificate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
