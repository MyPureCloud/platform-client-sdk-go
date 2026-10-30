package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Selfschedulemanagementunitsettingsrequest
type Selfschedulemanagementunitsettingsrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether agent self-scheduling is enabled for this management unit
	Enabled *bool `json:"enabled,omitempty"`

	// CheckWorkPlanConstraints - Whether self-scheduling is further restricted by work plan constraints
	CheckWorkPlanConstraints *bool `json:"checkWorkPlanConstraints,omitempty"`

	// MinimumMinutesBeforeEarliestScheduleChange - The minimum number of minutes before the new activity start time to which the agent is allowed to self-schedule
	MinimumMinutesBeforeEarliestScheduleChange *int `json:"minimumMinutesBeforeEarliestScheduleChange,omitempty"`

	// MovableActivityCategories - Movable activity categories enabled for self-scheduling
	MovableActivityCategories *Setwrapperactivitycategory `json:"movableActivityCategories,omitempty"`

	// MovableActivityCodeIds - Movable activity codes enabled for self-scheduling
	MovableActivityCodeIds *Setwrapperstring `json:"movableActivityCodeIds,omitempty"`

	// Metadata - Version metadata for this management unit's self-scheduling settings
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Selfschedulemanagementunitsettingsrequest) SetField(field string, fieldValue interface{}) {
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

func (o Selfschedulemanagementunitsettingsrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Selfschedulemanagementunitsettingsrequest
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		CheckWorkPlanConstraints *bool `json:"checkWorkPlanConstraints,omitempty"`
		
		MinimumMinutesBeforeEarliestScheduleChange *int `json:"minimumMinutesBeforeEarliestScheduleChange,omitempty"`
		
		MovableActivityCategories *Setwrapperactivitycategory `json:"movableActivityCategories,omitempty"`
		
		MovableActivityCodeIds *Setwrapperstring `json:"movableActivityCodeIds,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		CheckWorkPlanConstraints: o.CheckWorkPlanConstraints,
		
		MinimumMinutesBeforeEarliestScheduleChange: o.MinimumMinutesBeforeEarliestScheduleChange,
		
		MovableActivityCategories: o.MovableActivityCategories,
		
		MovableActivityCodeIds: o.MovableActivityCodeIds,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Selfschedulemanagementunitsettingsrequest) UnmarshalJSON(b []byte) error {
	var SelfschedulemanagementunitsettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SelfschedulemanagementunitsettingsrequestMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SelfschedulemanagementunitsettingsrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if CheckWorkPlanConstraints, ok := SelfschedulemanagementunitsettingsrequestMap["checkWorkPlanConstraints"].(bool); ok {
		o.CheckWorkPlanConstraints = &CheckWorkPlanConstraints
	}
    
	if MinimumMinutesBeforeEarliestScheduleChange, ok := SelfschedulemanagementunitsettingsrequestMap["minimumMinutesBeforeEarliestScheduleChange"].(float64); ok {
		MinimumMinutesBeforeEarliestScheduleChangeInt := int(MinimumMinutesBeforeEarliestScheduleChange)
		o.MinimumMinutesBeforeEarliestScheduleChange = &MinimumMinutesBeforeEarliestScheduleChangeInt
	}
	
	if MovableActivityCategories, ok := SelfschedulemanagementunitsettingsrequestMap["movableActivityCategories"].(map[string]interface{}); ok {
		MovableActivityCategoriesString, _ := json.Marshal(MovableActivityCategories)
		json.Unmarshal(MovableActivityCategoriesString, &o.MovableActivityCategories)
	}
	
	if MovableActivityCodeIds, ok := SelfschedulemanagementunitsettingsrequestMap["movableActivityCodeIds"].(map[string]interface{}); ok {
		MovableActivityCodeIdsString, _ := json.Marshal(MovableActivityCodeIds)
		json.Unmarshal(MovableActivityCodeIdsString, &o.MovableActivityCodeIds)
	}
	
	if Metadata, ok := SelfschedulemanagementunitsettingsrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Selfschedulemanagementunitsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
