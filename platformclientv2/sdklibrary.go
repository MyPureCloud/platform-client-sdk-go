package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sdklibrary
type Sdklibrary struct { 
	// Name - The name of the SDK.
	Name *string `json:"name,omitempty"`


	// Version - The version of the SDK.
	Version *string `json:"version,omitempty"`

}

func (o *Sdklibrary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sdklibrary
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Sdklibrary) UnmarshalJSON(b []byte) error {
	var SdklibraryMap map[string]interface{}
	err := json.Unmarshal(b, &SdklibraryMap)
	if err != nil {
		return err
	}
	
	if Name, ok := SdklibraryMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := SdklibraryMap["version"].(string); ok {
		o.Version = &Version
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sdklibrary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
