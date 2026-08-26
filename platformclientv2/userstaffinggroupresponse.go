package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstaffinggroupresponse
type Userstaffinggroupresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - Effective start date of the user assignment in ISO-8601 format or empty value. Empty value means no limit on start-date.
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - Effective end date of the user assignment in ISO-8601 format or empty value. Empty value means no limit on end-date.
	EndDate *time.Time `json:"endDate,omitempty"`

	// User - The user associated with the staffing group
	User *Userreference `json:"user,omitempty"`

	// StaffingGroup - The staffing group associated with the user
	StaffingGroup *Staffinggroupreference `json:"staffingGroup,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userstaffinggroupresponse) SetField(field string, fieldValue interface{}) {
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

func (o Userstaffinggroupresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "StartDate","EndDate", }

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
	type Alias Userstaffinggroupresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%d")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		StaffingGroup *Staffinggroupreference `json:"staffingGroup,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		User: o.User,
		
		StaffingGroup: o.StaffingGroup,
		Alias:    (Alias)(o),
	})
}

func (o *Userstaffinggroupresponse) UnmarshalJSON(b []byte) error {
	var UserstaffinggroupresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UserstaffinggroupresponseMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := UserstaffinggroupresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := UserstaffinggroupresponseMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02", endDateString)
		o.EndDate = &EndDate
	}
	
	if User, ok := UserstaffinggroupresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if StaffingGroup, ok := UserstaffinggroupresponseMap["staffingGroup"].(map[string]interface{}); ok {
		StaffingGroupString, _ := json.Marshal(StaffingGroup)
		json.Unmarshal(StaffingGroupString, &o.StaffingGroup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstaffinggroupresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
