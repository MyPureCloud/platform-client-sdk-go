package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalsegment
type Externalsegment struct { 
	// Id - Identifier for the external segment in the system where it originates from.
	Id *string `json:"id,omitempty"`


	// Name - Name for the external segment in the system where it originates from.
	Name *string `json:"name,omitempty"`


	// Source - The external system where the segment originates from.
	Source *string `json:"source,omitempty"`

}

func (o *Externalsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalsegment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Source *string `json:"source,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Source: o.Source,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalsegment) UnmarshalJSON(b []byte) error {
	var ExternalsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalsegmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalsegmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ExternalsegmentMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Source, ok := ExternalsegmentMap["source"].(string); ok {
		o.Source = &Source
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
