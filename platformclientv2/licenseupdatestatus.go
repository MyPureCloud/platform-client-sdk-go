package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Licenseupdatestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseupdatestatus

	

	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		Result *string `json:"result,omitempty"`
		*Alias
	}{ 
		UserId: u.UserId,
		
		LicenseId: u.LicenseId,
		
		Result: u.Result,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Licenseupdatestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
