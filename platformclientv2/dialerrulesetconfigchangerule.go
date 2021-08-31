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

func (o *Dialerrulesetconfigchangerule) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Order: o.Order,
		
		Category: o.Category,
		
		Actions: o.Actions,
		
		Conditions: o.Conditions,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangerule) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeruleMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeruleMap)
	if err != nil {
		return err
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
	
	if Conditions, ok := DialerrulesetconfigchangeruleMap["conditions"].([]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if AdditionalProperties, ok := DialerrulesetconfigchangeruleMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
