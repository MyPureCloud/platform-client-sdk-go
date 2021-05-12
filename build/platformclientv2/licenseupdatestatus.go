package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Licenseupdatestatus
type Licenseupdatestatus struct { 
	// UserId
	UserId *string `json:"userId,omitempty"`


	// LicenseId
	LicenseId *string `json:"licenseId,omitempty"`


	// Result
	Result *string `json:"result,omitempty"`

}

// String returns a JSON representation of the model
func (o *Licenseupdatestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
