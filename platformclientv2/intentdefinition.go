package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Intentdefinition
type Intentdefinition struct { 
	// Name - The name of the intent.
	Name *string `json:"name,omitempty"`


	// EntityTypeBindings - The bindings for the named entity types used in this intent.This field is mutually exclusive with entityNameReferences and entities
	EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`


	// EntityNameReferences - The references for the named entity used in this intent.This field is mutually exclusive with entityTypeBindings
	EntityNameReferences *[]string `json:"entityNameReferences,omitempty"`


	// Utterances - The utterances that act as training phrases for the intent.
	Utterances *[]Nluutterance `json:"utterances,omitempty"`

}

func (o *Intentdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Intentdefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`
		
		EntityNameReferences *[]string `json:"entityNameReferences,omitempty"`
		
		Utterances *[]Nluutterance `json:"utterances,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		EntityTypeBindings: o.EntityTypeBindings,
		
		EntityNameReferences: o.EntityNameReferences,
		
		Utterances: o.Utterances,
		Alias:    (*Alias)(o),
	})
}

func (o *Intentdefinition) UnmarshalJSON(b []byte) error {
	var IntentdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &IntentdefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := IntentdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if EntityTypeBindings, ok := IntentdefinitionMap["entityTypeBindings"].([]interface{}); ok {
		EntityTypeBindingsString, _ := json.Marshal(EntityTypeBindings)
		json.Unmarshal(EntityTypeBindingsString, &o.EntityTypeBindings)
	}
	
	if EntityNameReferences, ok := IntentdefinitionMap["entityNameReferences"].([]interface{}); ok {
		EntityNameReferencesString, _ := json.Marshal(EntityNameReferences)
		json.Unmarshal(EntityNameReferencesString, &o.EntityNameReferences)
	}
	
	if Utterances, ok := IntentdefinitionMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Intentdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
