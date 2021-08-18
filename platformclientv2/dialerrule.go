package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrule
type Dialerrule struct { 
	// Id - The identifier of the rule.
	Id *string `json:"id,omitempty"`


	// Name - The name of the rule.
	Name *string `json:"name,omitempty"`


	// Order - The ranked order of the rule. Rules are processed from lowest number to highest.
	Order *int `json:"order,omitempty"`


	// Category - The category of the rule.
	Category *string `json:"category,omitempty"`


	// Conditions - A list of Conditions. All of the Conditions must evaluate to true to trigger the actions.
	Conditions *[]Condition `json:"conditions,omitempty"`


	// Actions - The list of actions to be taken if the conditions are true.
	Actions *[]Dialeraction `json:"actions,omitempty"`

}

func (u *Dialerrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Conditions *[]Condition `json:"conditions,omitempty"`
		
		Actions *[]Dialeraction `json:"actions,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Order: u.Order,
		
		Category: u.Category,
		
		Conditions: u.Conditions,
		
		Actions: u.Actions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
