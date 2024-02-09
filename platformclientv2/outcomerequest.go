package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomerequest
type Outcomerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IsActive - Whether or not the outcome is active.
	IsActive *bool `json:"isActive,omitempty"`

	// DisplayName - The display name of the outcome.
	DisplayName *string `json:"displayName,omitempty"`

	// Version - The version of the outcome.
	Version *int `json:"version,omitempty"`

	// Description - A description of the outcome.
	Description *string `json:"description,omitempty"`

	// IsPositive - Whether or not the outcome is positive.
	IsPositive *bool `json:"isPositive,omitempty"`

	// Context - The context of the outcome.
	Context *Requestcontext `json:"context,omitempty"`

	// Journey - The pattern of rules defining the filter of the outcome.
	Journey *Requestjourney `json:"journey,omitempty"`

	// AssociatedValueField - The field from the event indicating the associated value.
	AssociatedValueField *Associatedvaluefield `json:"associatedValueField,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outcomerequest) SetField(field string, fieldValue interface{}) {
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

func (o Outcomerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Outcomerequest
	
	return json.Marshal(&struct { 
		IsActive *bool `json:"isActive,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		IsPositive *bool `json:"isPositive,omitempty"`
		
		Context *Requestcontext `json:"context,omitempty"`
		
		Journey *Requestjourney `json:"journey,omitempty"`
		
		AssociatedValueField *Associatedvaluefield `json:"associatedValueField,omitempty"`
		Alias
	}{ 
		IsActive: o.IsActive,
		
		DisplayName: o.DisplayName,
		
		Version: o.Version,
		
		Description: o.Description,
		
		IsPositive: o.IsPositive,
		
		Context: o.Context,
		
		Journey: o.Journey,
		
		AssociatedValueField: o.AssociatedValueField,
		Alias:    (Alias)(o),
	})
}

func (o *Outcomerequest) UnmarshalJSON(b []byte) error {
	var OutcomerequestMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomerequestMap)
	if err != nil {
		return err
	}
	
	if IsActive, ok := OutcomerequestMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if DisplayName, ok := OutcomerequestMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Version, ok := OutcomerequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Description, ok := OutcomerequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if IsPositive, ok := OutcomerequestMap["isPositive"].(bool); ok {
		o.IsPositive = &IsPositive
	}
    
	if Context, ok := OutcomerequestMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if Journey, ok := OutcomerequestMap["journey"].(map[string]interface{}); ok {
		JourneyString, _ := json.Marshal(Journey)
		json.Unmarshal(JourneyString, &o.Journey)
	}
	
	if AssociatedValueField, ok := OutcomerequestMap["associatedValueField"].(map[string]interface{}); ok {
		AssociatedValueFieldString, _ := json.Marshal(AssociatedValueField)
		json.Unmarshal(AssociatedValueFieldString, &o.AssociatedValueField)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
