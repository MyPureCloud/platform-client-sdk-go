package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Detectednamedentity
type Detectednamedentity struct { 
	// Name - The name of the detected named entity.
	Name *string `json:"name,omitempty"`


	// EntityType - The type of the detected named entity.
	EntityType *string `json:"entityType,omitempty"`


	// Probability - The probability of the detected named entity.
	Probability *float64 `json:"probability,omitempty"`


	// Value - The value of the detected named entity.
	Value *Detectednamedentityvalue `json:"value,omitempty"`

}

func (u *Detectednamedentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Detectednamedentity

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		Probability *float64 `json:"probability,omitempty"`
		
		Value *Detectednamedentityvalue `json:"value,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		EntityType: u.EntityType,
		
		Probability: u.Probability,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Detectednamedentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
