package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulequeryresponse
type Buagentschedulequeryresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User - The user to whom this agent schedule applies
	User *Userreference `json:"user,omitempty"`

	// Shifts - The shift definitions for this agent schedule
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`

	// FullDayTimeOffMarkers - Full day time off markers which apply to this agent schedule
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`

	// WorkPlan - The work plan for this user
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`

	// WorkPlansPerWeek - The work plans per week for this user from the work plan rotation. Null values in the list denotes that user is not part of any work plan for that week
	WorkPlansPerWeek *[]Workplanreference `json:"workPlansPerWeek,omitempty"`

	// Metadata - Versioned entity metadata for this agent schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buagentschedulequeryresponse) SetField(field string, fieldValue interface{}) {
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

func (o Buagentschedulequeryresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Buagentschedulequeryresponse
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		
		WorkPlansPerWeek *[]Workplanreference `json:"workPlansPerWeek,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		WorkPlan: o.WorkPlan,
		
		WorkPlansPerWeek: o.WorkPlansPerWeek,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Buagentschedulequeryresponse) UnmarshalJSON(b []byte) error {
	var BuagentschedulequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulequeryresponseMap)
	if err != nil {
		return err
	}
	
	if User, ok := BuagentschedulequeryresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Shifts, ok := BuagentschedulequeryresponseMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := BuagentschedulequeryresponseMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if WorkPlan, ok := BuagentschedulequeryresponseMap["workPlan"].(map[string]interface{}); ok {
		WorkPlanString, _ := json.Marshal(WorkPlan)
		json.Unmarshal(WorkPlanString, &o.WorkPlan)
	}
	
	if WorkPlansPerWeek, ok := BuagentschedulequeryresponseMap["workPlansPerWeek"].([]interface{}); ok {
		WorkPlansPerWeekString, _ := json.Marshal(WorkPlansPerWeek)
		json.Unmarshal(WorkPlansPerWeekString, &o.WorkPlansPerWeek)
	}
	
	if Metadata, ok := BuagentschedulequeryresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
