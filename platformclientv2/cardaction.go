package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cardaction - A card action that a user can take.
type Cardaction struct { 
	// VarType - Describes the type of action.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a button.
	Payload *string `json:"payload,omitempty"`


	// Url - The location of the image file associated with action.
	Url *string `json:"url,omitempty"`


	// IsSelected - Indicates if the card option is selected by end customer.
	IsSelected *bool `json:"isSelected,omitempty"`

}

func (o *Cardaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cardaction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		IsSelected *bool `json:"isSelected,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		
		Url: o.Url,
		
		IsSelected: o.IsSelected,
		Alias:    (*Alias)(o),
	})
}

func (o *Cardaction) UnmarshalJSON(b []byte) error {
	var CardactionMap map[string]interface{}
	err := json.Unmarshal(b, &CardactionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := CardactionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := CardactionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := CardactionMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Url, ok := CardactionMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if IsSelected, ok := CardactionMap["isSelected"].(bool); ok {
		o.IsSelected = &IsSelected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Cardaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
