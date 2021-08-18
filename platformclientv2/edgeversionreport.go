package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeversionreport
type Edgeversionreport struct { 
	// OldestVersion
	OldestVersion *Edgeversioninformation `json:"oldestVersion,omitempty"`


	// NewestVersion
	NewestVersion *Edgeversioninformation `json:"newestVersion,omitempty"`

}

func (u *Edgeversionreport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeversionreport

	

	return json.Marshal(&struct { 
		OldestVersion *Edgeversioninformation `json:"oldestVersion,omitempty"`
		
		NewestVersion *Edgeversioninformation `json:"newestVersion,omitempty"`
		*Alias
	}{ 
		OldestVersion: u.OldestVersion,
		
		NewestVersion: u.NewestVersion,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgeversionreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
