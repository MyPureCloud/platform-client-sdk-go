package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licenseorgtoggle
type Licenseorgtoggle struct { 
	// FeatureName
	FeatureName *string `json:"featureName,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (u *Licenseorgtoggle) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseorgtoggle

	

	return json.Marshal(&struct { 
		FeatureName *string `json:"featureName,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		FeatureName: u.FeatureName,
		
		Enabled: u.Enabled,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Licenseorgtoggle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
