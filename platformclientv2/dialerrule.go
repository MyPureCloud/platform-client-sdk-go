package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Dialerrule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
