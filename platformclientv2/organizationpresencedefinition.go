package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationpresencedefinition
type Organizationpresencedefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// VarType - The type of definition
	VarType *string `json:"type,omitempty"`

	// LanguageLabels - The label used for the definition in each specified language
	LanguageLabels *map[string]string `json:"languageLabels,omitempty"`

	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// Deactivated
	Deactivated *bool `json:"deactivated,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organizationpresencedefinition) SetField(field string, fieldValue interface{}) {
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

func (o Organizationpresencedefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Organizationpresencedefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		LanguageLabels *map[string]string `json:"languageLabels,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Deactivated *bool `json:"deactivated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		LanguageLabels: o.LanguageLabels,
		
		SystemPresence: o.SystemPresence,
		
		DivisionId: o.DivisionId,
		
		Deactivated: o.Deactivated,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Organizationpresencedefinition) UnmarshalJSON(b []byte) error {
	var OrganizationpresencedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationpresencedefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OrganizationpresencedefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OrganizationpresencedefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := OrganizationpresencedefinitionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if LanguageLabels, ok := OrganizationpresencedefinitionMap["languageLabels"].(map[string]interface{}); ok {
		LanguageLabelsString, _ := json.Marshal(LanguageLabels)
		json.Unmarshal(LanguageLabelsString, &o.LanguageLabels)
	}
	
	if SystemPresence, ok := OrganizationpresencedefinitionMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if DivisionId, ok := OrganizationpresencedefinitionMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if Deactivated, ok := OrganizationpresencedefinitionMap["deactivated"].(bool); ok {
		o.Deactivated = &Deactivated
	}
    
	if SelfUri, ok := OrganizationpresencedefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationpresencedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
