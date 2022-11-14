package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataactionconditionsettings
type Dataactionconditionsettings struct { 
	// DataActionId - The Data Action Id to use for this condition.
	DataActionId *string `json:"dataActionId,omitempty"`


	// ContactIdField - The input field from the data action that the contactId will be passed into.
	ContactIdField *string `json:"contactIdField,omitempty"`


	// DataNotFoundResolution - The result of this condition if the data action returns a result indicating there was no data.
	DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`


	// Predicates - A list of predicates defining the comparisons to use for this condition.
	Predicates *[]Digitaldataactionconditionpredicate `json:"predicates,omitempty"`


	// ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields.
	ContactColumnToDataActionFieldMappings *[]Dataactioncontactcolumnfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`

}

func (o *Dataactionconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dataactionconditionsettings
	
	return json.Marshal(&struct { 
		DataActionId *string `json:"dataActionId,omitempty"`
		
		ContactIdField *string `json:"contactIdField,omitempty"`
		
		DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`
		
		Predicates *[]Digitaldataactionconditionpredicate `json:"predicates,omitempty"`
		
		ContactColumnToDataActionFieldMappings *[]Dataactioncontactcolumnfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`
		*Alias
	}{ 
		DataActionId: o.DataActionId,
		
		ContactIdField: o.ContactIdField,
		
		DataNotFoundResolution: o.DataNotFoundResolution,
		
		Predicates: o.Predicates,
		
		ContactColumnToDataActionFieldMappings: o.ContactColumnToDataActionFieldMappings,
		Alias:    (*Alias)(o),
	})
}

func (o *Dataactionconditionsettings) UnmarshalJSON(b []byte) error {
	var DataactionconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DataactionconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if DataActionId, ok := DataactionconditionsettingsMap["dataActionId"].(string); ok {
		o.DataActionId = &DataActionId
	}
    
	if ContactIdField, ok := DataactionconditionsettingsMap["contactIdField"].(string); ok {
		o.ContactIdField = &ContactIdField
	}
    
	if DataNotFoundResolution, ok := DataactionconditionsettingsMap["dataNotFoundResolution"].(bool); ok {
		o.DataNotFoundResolution = &DataNotFoundResolution
	}
    
	if Predicates, ok := DataactionconditionsettingsMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	
	if ContactColumnToDataActionFieldMappings, ok := DataactionconditionsettingsMap["contactColumnToDataActionFieldMappings"].([]interface{}); ok {
		ContactColumnToDataActionFieldMappingsString, _ := json.Marshal(ContactColumnToDataActionFieldMappings)
		json.Unmarshal(ContactColumnToDataActionFieldMappingsString, &o.ContactColumnToDataActionFieldMappings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dataactionconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
