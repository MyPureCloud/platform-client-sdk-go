package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactenrichrequest
type Contactenrichrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A user-specified tracker string, only useful in the Bulk-Enrich API. If one Bulk-Enrich operation in a request fails, the requested operation will be repeated in the Bulk API response, including this id field, allowing associating of request and response operations.
	Id *string `json:"id,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writablestarrabledivision `json:"division,omitempty"`

	// MatchingIdentifiers - An ordered list of one or more Identifiers which might each be claimed by a Contact. `action` describes what to do with any possibly matching Contacts. Identifier lookups will occur in the order specified here. Between 1 and 25.
	MatchingIdentifiers *[]Contactidentifier `json:"matchingIdentifiers,omitempty"`

	// Action - The action that should be taken based on any Contacts found by `matchingIdentifiers`.
	Action *string `json:"action,omitempty"`

	// Contact - Data to be added, either as an update to an existing Contact or the body of a new Contact. Omitting a field in this contract means that it will be treated as null in the `fieldRules` logic.
	Contact *Externalcontact `json:"contact,omitempty"`

	// FieldRules - Logic describing how to combine data from the submitted request with data found in the database.
	FieldRules *Enrichfieldrules `json:"fieldRules,omitempty"`

	// Options - Additional options modifying the behavior of the API operation.
	Options *Contactenrichoptions `json:"options,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactenrichrequest) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Contactenrichrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactenrichrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Writablestarrabledivision `json:"division,omitempty"`
		
		MatchingIdentifiers *[]Contactidentifier `json:"matchingIdentifiers,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Contact *Externalcontact `json:"contact,omitempty"`
		
		FieldRules *Enrichfieldrules `json:"fieldRules,omitempty"`
		
		Options *Contactenrichoptions `json:"options,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		MatchingIdentifiers: o.MatchingIdentifiers,
		
		Action: o.Action,
		
		Contact: o.Contact,
		
		FieldRules: o.FieldRules,
		
		Options: o.Options,
		Alias:    (Alias)(o),
	})
}

func (o *Contactenrichrequest) UnmarshalJSON(b []byte) error {
	var ContactenrichrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactenrichrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContactenrichrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ContactenrichrequestMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if MatchingIdentifiers, ok := ContactenrichrequestMap["matchingIdentifiers"].([]interface{}); ok {
		MatchingIdentifiersString, _ := json.Marshal(MatchingIdentifiers)
		json.Unmarshal(MatchingIdentifiersString, &o.MatchingIdentifiers)
	}
	
	if Action, ok := ContactenrichrequestMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Contact, ok := ContactenrichrequestMap["contact"].(map[string]interface{}); ok {
		ContactString, _ := json.Marshal(Contact)
		json.Unmarshal(ContactString, &o.Contact)
	}
	
	if FieldRules, ok := ContactenrichrequestMap["fieldRules"].(map[string]interface{}); ok {
		FieldRulesString, _ := json.Marshal(FieldRules)
		json.Unmarshal(FieldRulesString, &o.FieldRules)
	}
	
	if Options, ok := ContactenrichrequestMap["options"].(map[string]interface{}); ok {
		OptionsString, _ := json.Marshal(Options)
		json.Unmarshal(OptionsString, &o.Options)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactenrichrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
