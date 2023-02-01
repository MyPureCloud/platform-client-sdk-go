package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Activitycode
type Activitycode struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// Name - The name of the activity code. Default activity codes will be created with an empty name
	Name *string `json:"name,omitempty"`

	// IsActive - Whether this activity code is active or has been deleted
	IsActive *bool `json:"isActive,omitempty"`

	// IsDefault - Whether this is a default activity code
	IsDefault *bool `json:"isDefault,omitempty"`

	// Category - The activity code's category.
	Category *string `json:"category,omitempty"`

	// LengthInMinutes - The default length of the activity in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// CountsAsPaidTime - Whether an agent is paid while performing this activity
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

	// CountsAsWorkTime - Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
	CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`

	// AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
	AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`

	// Metadata - Version metadata for the associated management unit's list of activity codes
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Activitycode) SetField(field string, fieldValue interface{}) {
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

func (o Activitycode) MarshalJSON() ([]byte, error) {
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
	type Alias Activitycode
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CountsAsWorkTime *bool `json:"countsAsWorkTime,omitempty"`
		
		AgentTimeOffSelectable *bool `json:"agentTimeOffSelectable,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		Name: o.Name,
		
		IsActive: o.IsActive,
		
		IsDefault: o.IsDefault,
		
		Category: o.Category,
		
		LengthInMinutes: o.LengthInMinutes,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		CountsAsWorkTime: o.CountsAsWorkTime,
		
		AgentTimeOffSelectable: o.AgentTimeOffSelectable,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Activitycode) UnmarshalJSON(b []byte) error {
	var ActivitycodeMap map[string]interface{}
	err := json.Unmarshal(b, &ActivitycodeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActivitycodeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := ActivitycodeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Name, ok := ActivitycodeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if IsActive, ok := ActivitycodeMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if IsDefault, ok := ActivitycodeMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    
	if Category, ok := ActivitycodeMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if LengthInMinutes, ok := ActivitycodeMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if CountsAsPaidTime, ok := ActivitycodeMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if CountsAsWorkTime, ok := ActivitycodeMap["countsAsWorkTime"].(bool); ok {
		o.CountsAsWorkTime = &CountsAsWorkTime
	}
    
	if AgentTimeOffSelectable, ok := ActivitycodeMap["agentTimeOffSelectable"].(bool); ok {
		o.AgentTimeOffSelectable = &AgentTimeOffSelectable
	}
    
	if Metadata, ok := ActivitycodeMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Activitycode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
