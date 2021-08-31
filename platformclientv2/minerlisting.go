package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Minerlisting
type Minerlisting struct { 
	// Entities
	Entities *[]Miner `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

}

func (o *Minerlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minerlisting
	
	return json.Marshal(&struct { 
		Entities *[]Miner `json:"entities,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		NextUri: o.NextUri,
		
		SelfUri: o.SelfUri,
		
		PreviousUri: o.PreviousUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Minerlisting) UnmarshalJSON(b []byte) error {
	var MinerlistingMap map[string]interface{}
	err := json.Unmarshal(b, &MinerlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := MinerlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if NextUri, ok := MinerlistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if SelfUri, ok := MinerlistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if PreviousUri, ok := MinerlistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Minerlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
