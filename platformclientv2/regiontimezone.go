package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Regiontimezone
type Regiontimezone struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Offset
	Offset *int `json:"offset,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Regiontimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Regiontimezone
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Offset *int `json:"offset,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Offset: o.Offset,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Regiontimezone) UnmarshalJSON(b []byte) error {
	var RegiontimezoneMap map[string]interface{}
	err := json.Unmarshal(b, &RegiontimezoneMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RegiontimezoneMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := RegiontimezoneMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Offset, ok := RegiontimezoneMap["offset"].(float64); ok {
		OffsetInt := int(Offset)
		o.Offset = &OffsetInt
	}
	
	if SelfUri, ok := RegiontimezoneMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Regiontimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
