package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validationlimits
type Validationlimits struct { 
	// MinLength
	MinLength *Minlength `json:"minLength,omitempty"`


	// MaxLength
	MaxLength *Maxlength `json:"maxLength,omitempty"`


	// MinItems
	MinItems *Minlength `json:"minItems,omitempty"`


	// MaxItems
	MaxItems *Maxlength `json:"maxItems,omitempty"`


	// Minimum
	Minimum *Minlength `json:"minimum,omitempty"`


	// Maximum
	Maximum *Maxlength `json:"maximum,omitempty"`

}

func (o *Validationlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validationlimits
	
	return json.Marshal(&struct { 
		MinLength *Minlength `json:"minLength,omitempty"`
		
		MaxLength *Maxlength `json:"maxLength,omitempty"`
		
		MinItems *Minlength `json:"minItems,omitempty"`
		
		MaxItems *Maxlength `json:"maxItems,omitempty"`
		
		Minimum *Minlength `json:"minimum,omitempty"`
		
		Maximum *Maxlength `json:"maximum,omitempty"`
		*Alias
	}{ 
		MinLength: o.MinLength,
		
		MaxLength: o.MaxLength,
		
		MinItems: o.MinItems,
		
		MaxItems: o.MaxItems,
		
		Minimum: o.Minimum,
		
		Maximum: o.Maximum,
		Alias:    (*Alias)(o),
	})
}

func (o *Validationlimits) UnmarshalJSON(b []byte) error {
	var ValidationlimitsMap map[string]interface{}
	err := json.Unmarshal(b, &ValidationlimitsMap)
	if err != nil {
		return err
	}
	
	if MinLength, ok := ValidationlimitsMap["minLength"].(map[string]interface{}); ok {
		MinLengthString, _ := json.Marshal(MinLength)
		json.Unmarshal(MinLengthString, &o.MinLength)
	}
	
	if MaxLength, ok := ValidationlimitsMap["maxLength"].(map[string]interface{}); ok {
		MaxLengthString, _ := json.Marshal(MaxLength)
		json.Unmarshal(MaxLengthString, &o.MaxLength)
	}
	
	if MinItems, ok := ValidationlimitsMap["minItems"].(map[string]interface{}); ok {
		MinItemsString, _ := json.Marshal(MinItems)
		json.Unmarshal(MinItemsString, &o.MinItems)
	}
	
	if MaxItems, ok := ValidationlimitsMap["maxItems"].(map[string]interface{}); ok {
		MaxItemsString, _ := json.Marshal(MaxItems)
		json.Unmarshal(MaxItemsString, &o.MaxItems)
	}
	
	if Minimum, ok := ValidationlimitsMap["minimum"].(map[string]interface{}); ok {
		MinimumString, _ := json.Marshal(Minimum)
		json.Unmarshal(MinimumString, &o.Minimum)
	}
	
	if Maximum, ok := ValidationlimitsMap["maximum"].(map[string]interface{}); ok {
		MaximumString, _ := json.Marshal(Maximum)
		json.Unmarshal(MaximumString, &o.Maximum)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validationlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
