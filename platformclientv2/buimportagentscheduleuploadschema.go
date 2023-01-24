package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buimportagentscheduleuploadschema
type Buimportagentscheduleuploadschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UserId - The ID of the user to whom this agent schedule applies
	UserId *string `json:"userId,omitempty"`

	// WorkPlanId - The ID of the work plan for this user.  Mutually exclusive with workPlanIdsPerWeek
	WorkPlanId *Valuewrapperstring `json:"workPlanId,omitempty"`

	// WorkPlanIdsPerWeek - The IDs of the work plans per week for this user.  Mutually exclusive with workPlanId
	WorkPlanIdsPerWeek *Listwrapperstring `json:"workPlanIdsPerWeek,omitempty"`

	// Shifts - The shift definitions for this agent schedule
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`

	// FullDayTimeOffMarkers - Any full day time off markers that apply to this agent schedule
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buimportagentscheduleuploadschema) SetField(field string, fieldValue interface{}) {
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

func (o Buimportagentscheduleuploadschema) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Buimportagentscheduleuploadschema
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		WorkPlanId *Valuewrapperstring `json:"workPlanId,omitempty"`
		
		WorkPlanIdsPerWeek *Listwrapperstring `json:"workPlanIdsPerWeek,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		WorkPlanId: o.WorkPlanId,
		
		WorkPlanIdsPerWeek: o.WorkPlanIdsPerWeek,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		Alias:    (Alias)(o),
	})
}

func (o *Buimportagentscheduleuploadschema) UnmarshalJSON(b []byte) error {
	var BuimportagentscheduleuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &BuimportagentscheduleuploadschemaMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := BuimportagentscheduleuploadschemaMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if WorkPlanId, ok := BuimportagentscheduleuploadschemaMap["workPlanId"].(map[string]interface{}); ok {
		WorkPlanIdString, _ := json.Marshal(WorkPlanId)
		json.Unmarshal(WorkPlanIdString, &o.WorkPlanId)
	}
	
	if WorkPlanIdsPerWeek, ok := BuimportagentscheduleuploadschemaMap["workPlanIdsPerWeek"].(map[string]interface{}); ok {
		WorkPlanIdsPerWeekString, _ := json.Marshal(WorkPlanIdsPerWeek)
		json.Unmarshal(WorkPlanIdsPerWeekString, &o.WorkPlanIdsPerWeek)
	}
	
	if Shifts, ok := BuimportagentscheduleuploadschemaMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := BuimportagentscheduleuploadschemaMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buimportagentscheduleuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
