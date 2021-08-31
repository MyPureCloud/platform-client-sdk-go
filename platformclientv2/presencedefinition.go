package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presencedefinition
type Presencedefinition struct { 
	// Id - description
	Id *string `json:"id,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Presencedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presencedefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SystemPresence: o.SystemPresence,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Presencedefinition) UnmarshalJSON(b []byte) error {
	var PresencedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &PresencedefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PresencedefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SystemPresence, ok := PresencedefinitionMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
	
	if SelfUri, ok := PresencedefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Presencedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
