package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangerule
type Dialerrulesetconfigchangerule struct { 
	// Conditions - The list of rule conditions; all must evaluate to true to trigger the rule actions
	Conditions *[]Dialerrulesetconfigchangecondition `json:"conditions,omitempty"`


	// Id - The identifier of the rule
	Id *string `json:"id,omitempty"`


	// Name - The name of the rule
	Name *string `json:"name,omitempty"`


	// Order - The ranked order of the rule; rules are processed from lowest number to highest
	Order *int `json:"order,omitempty"`


	// Category - The category of the rule
	Category *string `json:"category,omitempty"`


	// Actions - The list of rule actions to be taken if the conditions are true
	Actions *[]Dialerrulesetconfigchangeaction `json:"actions,omitempty"`

}

func (o *Dialerrulesetconfigchangerule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangerule
	
	return json.Marshal(&struct { 
		Conditions *[]Dialerrulesetconfigchangecondition `json:"conditions,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Actions *[]Dialerrulesetconfigchangeaction `json:"actions,omitempty"`
		*Alias
	}{ 
		Conditions: o.Conditions,
		
		Id: o.Id,
		
		Name: o.Name,
		
		Order: o.Order,
		
		Category: o.Category,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangerule) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeruleMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeruleMap)
	if err != nil {
		return err
	}
	
	if Conditions, ok := DialerrulesetconfigchangeruleMap["conditions"].([]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Id, ok := DialerrulesetconfigchangeruleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialerrulesetconfigchangeruleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Order, ok := DialerrulesetconfigchangeruleMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if Category, ok := DialerrulesetconfigchangeruleMap["category"].(string); ok {
		o.Category = &Category
	}
	
	if Actions, ok := DialerrulesetconfigchangeruleMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
