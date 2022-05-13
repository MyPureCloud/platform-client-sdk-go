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

func (o *Licenseorgtoggle) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseorgtoggle
	
	return json.Marshal(&struct { 
		FeatureName *string `json:"featureName,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		FeatureName: o.FeatureName,
		
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Licenseorgtoggle) UnmarshalJSON(b []byte) error {
	var LicenseorgtoggleMap map[string]interface{}
	err := json.Unmarshal(b, &LicenseorgtoggleMap)
	if err != nil {
		return err
	}
	
	if FeatureName, ok := LicenseorgtoggleMap["featureName"].(string); ok {
		o.FeatureName = &FeatureName
	}
    
	if Enabled, ok := LicenseorgtoggleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Licenseorgtoggle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
