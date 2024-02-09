package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysegmentrequest
type Journeysegmentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IsActive - Whether or not the segment is active.
	IsActive *bool `json:"isActive,omitempty"`

	// DisplayName - The display name of the segment.
	DisplayName *string `json:"displayName,omitempty"`

	// Version - The version of the segment.
	Version *int `json:"version,omitempty"`

	// Description - A description of the segment.
	Description *string `json:"description,omitempty"`

	// Color - The hexadecimal color value of the segment.
	Color *string `json:"color,omitempty"`

	// Scope - The target entity that a segment applies to.
	Scope *string `json:"scope,omitempty"`

	// ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
	ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`

	// Context - The context of the segment.
	Context *Requestcontext `json:"context,omitempty"`

	// Journey - The pattern of rules defining the segment.
	Journey *Requestjourney `json:"journey,omitempty"`

	// ExternalSegment - Details of an entity corresponding to this segment in an external system.
	ExternalSegment *Requestexternalsegment `json:"externalSegment,omitempty"`

	// AssignmentExpirationDays - Time, in days, from when the segment is assigned until it is automatically unassigned.
	AssignmentExpirationDays *int `json:"assignmentExpirationDays,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeysegmentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Journeysegmentrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Journeysegmentrequest
	
	return json.Marshal(&struct { 
		IsActive *bool `json:"isActive,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Color *string `json:"color,omitempty"`
		
		Scope *string `json:"scope,omitempty"`
		
		ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`
		
		Context *Requestcontext `json:"context,omitempty"`
		
		Journey *Requestjourney `json:"journey,omitempty"`
		
		ExternalSegment *Requestexternalsegment `json:"externalSegment,omitempty"`
		
		AssignmentExpirationDays *int `json:"assignmentExpirationDays,omitempty"`
		Alias
	}{ 
		IsActive: o.IsActive,
		
		DisplayName: o.DisplayName,
		
		Version: o.Version,
		
		Description: o.Description,
		
		Color: o.Color,
		
		Scope: o.Scope,
		
		ShouldDisplayToAgent: o.ShouldDisplayToAgent,
		
		Context: o.Context,
		
		Journey: o.Journey,
		
		ExternalSegment: o.ExternalSegment,
		
		AssignmentExpirationDays: o.AssignmentExpirationDays,
		Alias:    (Alias)(o),
	})
}

func (o *Journeysegmentrequest) UnmarshalJSON(b []byte) error {
	var JourneysegmentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysegmentrequestMap)
	if err != nil {
		return err
	}
	
	if IsActive, ok := JourneysegmentrequestMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if DisplayName, ok := JourneysegmentrequestMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Version, ok := JourneysegmentrequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Description, ok := JourneysegmentrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Color, ok := JourneysegmentrequestMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if Scope, ok := JourneysegmentrequestMap["scope"].(string); ok {
		o.Scope = &Scope
	}
    
	if ShouldDisplayToAgent, ok := JourneysegmentrequestMap["shouldDisplayToAgent"].(bool); ok {
		o.ShouldDisplayToAgent = &ShouldDisplayToAgent
	}
    
	if Context, ok := JourneysegmentrequestMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if Journey, ok := JourneysegmentrequestMap["journey"].(map[string]interface{}); ok {
		JourneyString, _ := json.Marshal(Journey)
		json.Unmarshal(JourneyString, &o.Journey)
	}
	
	if ExternalSegment, ok := JourneysegmentrequestMap["externalSegment"].(map[string]interface{}); ok {
		ExternalSegmentString, _ := json.Marshal(ExternalSegment)
		json.Unmarshal(ExternalSegmentString, &o.ExternalSegment)
	}
	
	if AssignmentExpirationDays, ok := JourneysegmentrequestMap["assignmentExpirationDays"].(float64); ok {
		AssignmentExpirationDaysInt := int(AssignmentExpirationDays)
		o.AssignmentExpirationDays = &AssignmentExpirationDaysInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysegmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
