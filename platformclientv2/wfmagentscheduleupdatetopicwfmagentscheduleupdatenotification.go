package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification
type Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User
	User *Wfmagentscheduleupdatetopicuserreference `json:"user,omitempty"`

	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`

	// Shifts
	Shifts *[]Wfmagentscheduleupdatetopicwfmscheduleshift `json:"shifts,omitempty"`

	// FullDayTimeOffMarkers
	FullDayTimeOffMarkers *[]Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`

	// Updates
	Updates *[]Wfmagentscheduleupdatetopicwfmagentscheduleupdate `json:"updates,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
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
	type Alias Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		User *Wfmagentscheduleupdatetopicuserreference `json:"user,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Shifts *[]Wfmagentscheduleupdatetopicwfmscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Updates *[]Wfmagentscheduleupdatetopicwfmagentscheduleupdate `json:"updates,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		Updates: o.Updates,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap)
	if err != nil {
		return err
	}
	
	if User, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if startDateString, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Shifts, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if Updates, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["updates"].([]interface{}); ok {
		UpdatesString, _ := json.Marshal(Updates)
		json.Unmarshal(UpdatesString, &o.Updates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
