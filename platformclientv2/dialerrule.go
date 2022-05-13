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

func (o *Dialerrule) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Order: o.Order,
		
		Category: o.Category,
		
		Conditions: o.Conditions,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrule) UnmarshalJSON(b []byte) error {
	var DialerruleMap map[string]interface{}
	err := json.Unmarshal(b, &DialerruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialerruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Order, ok := DialerruleMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if Category, ok := DialerruleMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Conditions, ok := DialerruleMap["conditions"].([]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Actions, ok := DialerruleMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
