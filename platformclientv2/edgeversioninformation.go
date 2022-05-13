package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeversioninformation
type Edgeversioninformation struct { 
	// SoftwareVersion
	SoftwareVersion *string `json:"softwareVersion,omitempty"`

}

func (o *Edgeversioninformation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeversioninformation
	
	return json.Marshal(&struct { 
		SoftwareVersion *string `json:"softwareVersion,omitempty"`
		*Alias
	}{ 
		SoftwareVersion: o.SoftwareVersion,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeversioninformation) UnmarshalJSON(b []byte) error {
	var EdgeversioninformationMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeversioninformationMap)
	if err != nil {
		return err
	}
	
	if SoftwareVersion, ok := EdgeversioninformationMap["softwareVersion"].(string); ok {
		o.SoftwareVersion = &SoftwareVersion
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeversioninformation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
