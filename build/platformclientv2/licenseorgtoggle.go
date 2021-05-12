package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Licenseorgtoggle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
