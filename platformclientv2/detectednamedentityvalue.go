package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Detectednamedentityvalue
type Detectednamedentityvalue struct { 
	// Raw - The raw value of the detected named entity.
	Raw *string `json:"raw,omitempty"`


	// Resolved - The resolved value of the detected named entity.
	Resolved *string `json:"resolved,omitempty"`

}

func (o *Detectednamedentityvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Detectednamedentityvalue
	
	return json.Marshal(&struct { 
		Raw *string `json:"raw,omitempty"`
		
		Resolved *string `json:"resolved,omitempty"`
		*Alias
	}{ 
		Raw: o.Raw,
		
		Resolved: o.Resolved,
		Alias:    (*Alias)(o),
	})
}

func (o *Detectednamedentityvalue) UnmarshalJSON(b []byte) error {
	var DetectednamedentityvalueMap map[string]interface{}
	err := json.Unmarshal(b, &DetectednamedentityvalueMap)
	if err != nil {
		return err
	}
	
	if Raw, ok := DetectednamedentityvalueMap["raw"].(string); ok {
		o.Raw = &Raw
	}
    
	if Resolved, ok := DetectednamedentityvalueMap["resolved"].(string); ok {
		o.Resolved = &Resolved
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Detectednamedentityvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
