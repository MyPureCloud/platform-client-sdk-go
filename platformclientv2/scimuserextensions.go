package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimuserextensions - Genesys Cloud user extensions to SCIM RFC.
type Scimuserextensions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RoutingSkills - The list of routing skills assigned to a user. Maximum 50 skills.
	RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`

	// RoutingLanguages - The list of routing languages assigned to a user. Maximum 50 languages.
	RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`

	// ExternalIds - The list of external identifiers assigned to user. Always includes an immutable SCIM authority prefixed with \"x-pc:scimv2:v1\". ExternalIds are searchable with complex filter query parameter using 'authority' and 'value', e.g., filter=urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:externalIds[authority eq \"matchAuthName\" and value eq \"matchingExternalKeyValue\"].
	ExternalIds *[]Scimgenesysuserexternalid `json:"externalIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimuserextensions) SetField(field string, fieldValue interface{}) {
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

func (o Scimuserextensions) MarshalJSON() ([]byte, error) {
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
	type Alias Scimuserextensions
	
	return json.Marshal(&struct { 
		RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`
		
		RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`
		
		ExternalIds *[]Scimgenesysuserexternalid `json:"externalIds,omitempty"`
		Alias
	}{ 
		RoutingSkills: o.RoutingSkills,
		
		RoutingLanguages: o.RoutingLanguages,
		
		ExternalIds: o.ExternalIds,
		Alias:    (Alias)(o),
	})
}

func (o *Scimuserextensions) UnmarshalJSON(b []byte) error {
	var ScimuserextensionsMap map[string]interface{}
	err := json.Unmarshal(b, &ScimuserextensionsMap)
	if err != nil {
		return err
	}
	
	if RoutingSkills, ok := ScimuserextensionsMap["routingSkills"].([]interface{}); ok {
		RoutingSkillsString, _ := json.Marshal(RoutingSkills)
		json.Unmarshal(RoutingSkillsString, &o.RoutingSkills)
	}
	
	if RoutingLanguages, ok := ScimuserextensionsMap["routingLanguages"].([]interface{}); ok {
		RoutingLanguagesString, _ := json.Marshal(RoutingLanguages)
		json.Unmarshal(RoutingLanguagesString, &o.RoutingLanguages)
	}
	
	if ExternalIds, ok := ScimuserextensionsMap["externalIds"].([]interface{}); ok {
		ExternalIdsString, _ := json.Marshal(ExternalIds)
		json.Unmarshal(ExternalIdsString, &o.ExternalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
