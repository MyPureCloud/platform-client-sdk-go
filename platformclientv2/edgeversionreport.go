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

func (o *Edgeversionreport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeversionreport
	
	return json.Marshal(&struct { 
		OldestVersion *Edgeversioninformation `json:"oldestVersion,omitempty"`
		
		NewestVersion *Edgeversioninformation `json:"newestVersion,omitempty"`
		*Alias
	}{ 
		OldestVersion: o.OldestVersion,
		
		NewestVersion: o.NewestVersion,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeversionreport) UnmarshalJSON(b []byte) error {
	var EdgeversionreportMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeversionreportMap)
	if err != nil {
		return err
	}
	
	if OldestVersion, ok := EdgeversionreportMap["oldestVersion"].(map[string]interface{}); ok {
		OldestVersionString, _ := json.Marshal(OldestVersion)
		json.Unmarshal(OldestVersionString, &o.OldestVersion)
	}
	
	if NewestVersion, ok := EdgeversionreportMap["newestVersion"].(map[string]interface{}); ok {
		NewestVersionString, _ := json.Marshal(NewestVersion)
		json.Unmarshal(NewestVersionString, &o.NewestVersion)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeversionreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
