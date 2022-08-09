package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsecureattributes
type Conversationsecureattributes struct { 
	// Attributes - The map of attribute keys to values.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// Version - The version used to detect conflicting updates when using PUT. Not used for PATCH.
	Version *int `json:"version,omitempty"`

}

func (o *Conversationsecureattributes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationsecureattributes
	
	return json.Marshal(&struct { 
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Attributes: o.Attributes,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationsecureattributes) UnmarshalJSON(b []byte) error {
	var ConversationsecureattributesMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsecureattributesMap)
	if err != nil {
		return err
	}
	
	if Attributes, ok := ConversationsecureattributesMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Version, ok := ConversationsecureattributesMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsecureattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
