package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangerule
type Dialerrulesetconfigchangerule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Order
	Order *int `json:"order,omitempty"`


	// Category
	Category *string `json:"category,omitempty"`


	// Actions
	Actions *[]Dialerrulesetconfigchangeaction `json:"actions,omitempty"`


	// Conditions
	Conditions *[]Dialerrulesetconfigchangecondition `json:"conditions,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
