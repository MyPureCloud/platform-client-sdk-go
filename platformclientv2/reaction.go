package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reaction
type Reaction struct { 
	// Data - Parameter for this reaction. For transfer_flow, this would be the outbound flow id.
	Data *string `json:"data,omitempty"`


	// Name - Name of the parameter for this reaction. For transfer_flow, this would be the outbound flow name.
	Name *string `json:"name,omitempty"`


	// ReactionType - The reaction to take for a given call analysis result.
	ReactionType *string `json:"reactionType,omitempty"`

}

func (o *Reaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reaction
	
	return json.Marshal(&struct { 
		Data *string `json:"data,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ReactionType *string `json:"reactionType,omitempty"`
		*Alias
	}{ 
		Data: o.Data,
		
		Name: o.Name,
		
		ReactionType: o.ReactionType,
		Alias:    (*Alias)(o),
	})
}

func (o *Reaction) UnmarshalJSON(b []byte) error {
	var ReactionMap map[string]interface{}
	err := json.Unmarshal(b, &ReactionMap)
	if err != nil {
		return err
	}
	
	if Data, ok := ReactionMap["data"].(string); ok {
		o.Data = &Data
	}
	
	if Name, ok := ReactionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ReactionType, ok := ReactionMap["reactionType"].(string); ok {
		o.ReactionType = &ReactionType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
