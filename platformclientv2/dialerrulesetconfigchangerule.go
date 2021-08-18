package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Dialerrulesetconfigchangerule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangerule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Actions *[]Dialerrulesetconfigchangeaction `json:"actions,omitempty"`
		
		Conditions *[]Dialerrulesetconfigchangecondition `json:"conditions,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Order: u.Order,
		
		Category: u.Category,
		
		Actions: u.Actions,
		
		Conditions: u.Conditions,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
