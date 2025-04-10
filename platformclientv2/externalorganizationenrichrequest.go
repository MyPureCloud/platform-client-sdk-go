package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalorganizationenrichrequest
type Externalorganizationenrichrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A user-specified tracker string, only useful in the Bulk-Enrich API. If one Bulk-Enrich operation in a request fails, the requested operation will be repeated in the Bulk API response, including this id field, allowing associating request and response operations.
	Id *string `json:"id,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writablestarrabledivision `json:"division,omitempty"`

	// MatchingIdentifiers - An ordered list of one or more Identifiers which might each be claimed by an External Organization. `action` describes what to do with any possibly matching External Organization. Identifier lookups will occur in the order specified here.
	MatchingIdentifiers *[]Externalorganizationidentifier `json:"matchingIdentifiers,omitempty"`

	// Action - The action that should be taken based on any External Organization found by `matchingIdentifiers`.
	Action *string `json:"action,omitempty"`

	// ExternalOrganization - Data to be added, either as an update to an existing External Organization or the body of a new External Organization. Omitting a field in this contract means that it will be treated as null in the `fieldRules` logic.
	ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`

	// FieldRules - Logic describing how to combine data from the submitted request with data found in the database.
	FieldRules *Enrichfieldrules `json:"fieldRules,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalorganizationenrichrequest) SetField(field string, fieldValue interface{}) {
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

func (o Externalorganizationenrichrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Externalorganizationenrichrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Writablestarrabledivision `json:"division,omitempty"`
		
		MatchingIdentifiers *[]Externalorganizationidentifier `json:"matchingIdentifiers,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`
		
		FieldRules *Enrichfieldrules `json:"fieldRules,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		MatchingIdentifiers: o.MatchingIdentifiers,
		
		Action: o.Action,
		
		ExternalOrganization: o.ExternalOrganization,
		
		FieldRules: o.FieldRules,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Externalorganizationenrichrequest) UnmarshalJSON(b []byte) error {
	var ExternalorganizationenrichrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalorganizationenrichrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalorganizationenrichrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ExternalorganizationenrichrequestMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if MatchingIdentifiers, ok := ExternalorganizationenrichrequestMap["matchingIdentifiers"].([]interface{}); ok {
		MatchingIdentifiersString, _ := json.Marshal(MatchingIdentifiers)
		json.Unmarshal(MatchingIdentifiersString, &o.MatchingIdentifiers)
	}
	
	if Action, ok := ExternalorganizationenrichrequestMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if ExternalOrganization, ok := ExternalorganizationenrichrequestMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if FieldRules, ok := ExternalorganizationenrichrequestMap["fieldRules"].(map[string]interface{}); ok {
		FieldRulesString, _ := json.Marshal(FieldRules)
		json.Unmarshal(FieldRulesString, &o.FieldRules)
	}
	
	if SelfUri, ok := ExternalorganizationenrichrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalorganizationenrichrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
