package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Verificationresult
type Verificationresult struct { 
	// Status - The verification status.
	Status *string `json:"status,omitempty"`


	// Records - The list of DNS records that pertain that need to exist for verification.
	Records *[]Record `json:"records,omitempty"`

}

// String returns a JSON representation of the model
func (o *Verificationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
