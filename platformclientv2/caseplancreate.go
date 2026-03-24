package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Caseplancreate
type Caseplancreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the Caseplan. Valid length between 3 and 256 characters.
	Name *string `json:"name,omitempty"`

	// DefaultDueDurationInSeconds - The default due duration in seconds for Cases created from the Caseplan. Valid range is between 1 and 31536000 seconds.
	DefaultDueDurationInSeconds *int `json:"defaultDueDurationInSeconds,omitempty"`

	// DefaultTtlSeconds - The default TTL in seconds for Cases created from the Caseplan. Valid range is between 86400 and 31536000 seconds.
	DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`

	// ReferencePrefix - The prefix of the Caseplan reference. Valid length between 2 and 8 alphanumeric characters.
	ReferencePrefix *string `json:"referencePrefix,omitempty"`

	// CustomerIntentId - The ID of the customer intent associated with this Caseplan.
	CustomerIntentId *string `json:"customerIntentId,omitempty"`

	// Description - The description of the Caseplan. Maximum length of 512 characters.
	Description *string `json:"description,omitempty"`

	// DefaultCaseOwnerId - The ID of the default owner of a Case created from the Caseplan.
	DefaultCaseOwnerId *string `json:"defaultCaseOwnerId,omitempty"`

	// DivisionId - The ID of the division the Caseplan belongs to. Use '*' for divisionless caseplans.
	DivisionId *string `json:"divisionId,omitempty"`

	// DataSchemas - The schemas that define all data for cases from this Caseplan. The schema must be defined in the TaskManagement namespace.
	DataSchemas *[]Caseplandataschema `json:"dataSchemas,omitempty"`

	// IntakeSettings - The intake format when collecting data for a case from this caseplan. There can be a maximum of 10 IntakeSettings defined for a Caseplan.
	IntakeSettings *[]Intakesetting `json:"intakeSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Caseplancreate) SetField(field string, fieldValue interface{}) {
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

func (o Caseplancreate) MarshalJSON() ([]byte, error) {
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
	type Alias Caseplancreate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DefaultDueDurationInSeconds *int `json:"defaultDueDurationInSeconds,omitempty"`
		
		DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`
		
		ReferencePrefix *string `json:"referencePrefix,omitempty"`
		
		CustomerIntentId *string `json:"customerIntentId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultCaseOwnerId *string `json:"defaultCaseOwnerId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		DataSchemas *[]Caseplandataschema `json:"dataSchemas,omitempty"`
		
		IntakeSettings *[]Intakesetting `json:"intakeSettings,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		DefaultDueDurationInSeconds: o.DefaultDueDurationInSeconds,
		
		DefaultTtlSeconds: o.DefaultTtlSeconds,
		
		ReferencePrefix: o.ReferencePrefix,
		
		CustomerIntentId: o.CustomerIntentId,
		
		Description: o.Description,
		
		DefaultCaseOwnerId: o.DefaultCaseOwnerId,
		
		DivisionId: o.DivisionId,
		
		DataSchemas: o.DataSchemas,
		
		IntakeSettings: o.IntakeSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Caseplancreate) UnmarshalJSON(b []byte) error {
	var CaseplancreateMap map[string]interface{}
	err := json.Unmarshal(b, &CaseplancreateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CaseplancreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if DefaultDueDurationInSeconds, ok := CaseplancreateMap["defaultDueDurationInSeconds"].(float64); ok {
		DefaultDueDurationInSecondsInt := int(DefaultDueDurationInSeconds)
		o.DefaultDueDurationInSeconds = &DefaultDueDurationInSecondsInt
	}
	
	if DefaultTtlSeconds, ok := CaseplancreateMap["defaultTtlSeconds"].(float64); ok {
		DefaultTtlSecondsInt := int(DefaultTtlSeconds)
		o.DefaultTtlSeconds = &DefaultTtlSecondsInt
	}
	
	if ReferencePrefix, ok := CaseplancreateMap["referencePrefix"].(string); ok {
		o.ReferencePrefix = &ReferencePrefix
	}
    
	if CustomerIntentId, ok := CaseplancreateMap["customerIntentId"].(string); ok {
		o.CustomerIntentId = &CustomerIntentId
	}
    
	if Description, ok := CaseplancreateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DefaultCaseOwnerId, ok := CaseplancreateMap["defaultCaseOwnerId"].(string); ok {
		o.DefaultCaseOwnerId = &DefaultCaseOwnerId
	}
    
	if DivisionId, ok := CaseplancreateMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if DataSchemas, ok := CaseplancreateMap["dataSchemas"].([]interface{}); ok {
		DataSchemasString, _ := json.Marshal(DataSchemas)
		json.Unmarshal(DataSchemasString, &o.DataSchemas)
	}
	
	if IntakeSettings, ok := CaseplancreateMap["intakeSettings"].([]interface{}); ok {
		IntakeSettingsString, _ := json.Marshal(IntakeSettings)
		json.Unmarshal(IntakeSettingsString, &o.IntakeSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Caseplancreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
