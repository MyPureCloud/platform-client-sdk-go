package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchreviewresponse
type Shifttradematchreviewresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// InitiatingUser - Details for the initiatingUser side of the shift trade
	InitiatingUser *Shifttradematchreviewuserresponse `json:"initiatingUser,omitempty"`

	// ReceivingUser - Details for the receivingUser side of the shift trade
	ReceivingUser *Shifttradematchreviewuserresponse `json:"receivingUser,omitempty"`

	// Violations - Constraint violations introduced after being matched that would normally disallow a trade, but which can still be overridden by the shift trade administrator
	Violations *[]Shifttradematchviolation `json:"violations,omitempty"`

	// AdminReviewViolations - Constraint violations associated with this shift trade which require shift trade administrator review
	AdminReviewViolations *[]Shifttradematchviolation `json:"adminReviewViolations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shifttradematchreviewresponse) SetField(field string, fieldValue interface{}) {
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

func (o Shifttradematchreviewresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Shifttradematchreviewresponse
	
	return json.Marshal(&struct { 
		InitiatingUser *Shifttradematchreviewuserresponse `json:"initiatingUser,omitempty"`
		
		ReceivingUser *Shifttradematchreviewuserresponse `json:"receivingUser,omitempty"`
		
		Violations *[]Shifttradematchviolation `json:"violations,omitempty"`
		
		AdminReviewViolations *[]Shifttradematchviolation `json:"adminReviewViolations,omitempty"`
		Alias
	}{ 
		InitiatingUser: o.InitiatingUser,
		
		ReceivingUser: o.ReceivingUser,
		
		Violations: o.Violations,
		
		AdminReviewViolations: o.AdminReviewViolations,
		Alias:    (Alias)(o),
	})
}

func (o *Shifttradematchreviewresponse) UnmarshalJSON(b []byte) error {
	var ShifttradematchreviewresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradematchreviewresponseMap)
	if err != nil {
		return err
	}
	
	if InitiatingUser, ok := ShifttradematchreviewresponseMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if ReceivingUser, ok := ShifttradematchreviewresponseMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if Violations, ok := ShifttradematchreviewresponseMap["violations"].([]interface{}); ok {
		ViolationsString, _ := json.Marshal(Violations)
		json.Unmarshal(ViolationsString, &o.Violations)
	}
	
	if AdminReviewViolations, ok := ShifttradematchreviewresponseMap["adminReviewViolations"].([]interface{}); ok {
		AdminReviewViolationsString, _ := json.Marshal(AdminReviewViolations)
		json.Unmarshal(AdminReviewViolationsString, &o.AdminReviewViolations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
