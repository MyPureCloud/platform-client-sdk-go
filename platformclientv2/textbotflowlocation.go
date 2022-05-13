package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlocation - Describes a flow location.
type Textbotflowlocation struct { 
	// ActionName - The name of the action that was active when the event of interest happened.
	ActionName *string `json:"actionName,omitempty"`


	// ActionNumber - The number of the action that was active when the event of interest happened.
	ActionNumber *int `json:"actionNumber,omitempty"`


	// SequenceName - The name of the state or task which was active when the event of interest happened.
	SequenceName *string `json:"sequenceName,omitempty"`

}

func (o *Textbotflowlocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotflowlocation
	
	return json.Marshal(&struct { 
		ActionName *string `json:"actionName,omitempty"`
		
		ActionNumber *int `json:"actionNumber,omitempty"`
		
		SequenceName *string `json:"sequenceName,omitempty"`
		*Alias
	}{ 
		ActionName: o.ActionName,
		
		ActionNumber: o.ActionNumber,
		
		SequenceName: o.SequenceName,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotflowlocation) UnmarshalJSON(b []byte) error {
	var TextbotflowlocationMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowlocationMap)
	if err != nil {
		return err
	}
	
	if ActionName, ok := TextbotflowlocationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
    
	if ActionNumber, ok := TextbotflowlocationMap["actionNumber"].(float64); ok {
		ActionNumberInt := int(ActionNumber)
		o.ActionNumber = &ActionNumberInt
	}
	
	if SequenceName, ok := TextbotflowlocationMap["sequenceName"].(string); ok {
		o.SequenceName = &SequenceName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
