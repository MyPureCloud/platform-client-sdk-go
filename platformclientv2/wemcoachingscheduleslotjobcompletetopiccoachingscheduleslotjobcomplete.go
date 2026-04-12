package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete
type Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// AttendeeIds
	AttendeeIds *[]string `json:"attendeeIds,omitempty"`

	// FacilitatorIds
	FacilitatorIds *[]string `json:"facilitatorIds,omitempty"`

	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// SlotsType
	SlotsType *string `json:"slotsType,omitempty"`

	// Results
	Results *[]Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobresults `json:"results,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete) SetField(field string, fieldValue interface{}) {
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

func (o Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete) MarshalJSON() ([]byte, error) {
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
	type Alias Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		AttendeeIds *[]string `json:"attendeeIds,omitempty"`
		
		FacilitatorIds *[]string `json:"facilitatorIds,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		SlotsType *string `json:"slotsType,omitempty"`
		
		Results *[]Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobresults `json:"results,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		AttendeeIds: o.AttendeeIds,
		
		FacilitatorIds: o.FacilitatorIds,
		
		LengthInMinutes: o.LengthInMinutes,
		
		BusinessUnitId: o.BusinessUnitId,
		
		ActivityCodeId: o.ActivityCodeId,
		
		SlotsType: o.SlotsType,
		
		Results: o.Results,
		Alias:    (Alias)(o),
	})
}

func (o *Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete) UnmarshalJSON(b []byte) error {
	var WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap map[string]interface{}
	err := json.Unmarshal(b, &WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if AttendeeIds, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["attendeeIds"].([]interface{}); ok {
		AttendeeIdsString, _ := json.Marshal(AttendeeIds)
		json.Unmarshal(AttendeeIdsString, &o.AttendeeIds)
	}
	
	if FacilitatorIds, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["facilitatorIds"].([]interface{}); ok {
		FacilitatorIdsString, _ := json.Marshal(FacilitatorIds)
		json.Unmarshal(FacilitatorIdsString, &o.FacilitatorIds)
	}
	
	if LengthInMinutes, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if BusinessUnitId, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if ActivityCodeId, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if SlotsType, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["slotsType"].(string); ok {
		o.SlotsType = &SlotsType
	}
    
	if Results, ok := WemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcompleteMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wemcoachingscheduleslotjobcompletetopiccoachingscheduleslotjobcomplete) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
