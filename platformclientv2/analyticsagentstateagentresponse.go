package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsagentstateagentresponse
type Analyticsagentstateagentresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UserId - User Id - only returned if division is covered by agentStateNames permission
	UserId *string `json:"userId,omitempty"`

	// DivisionId - Division Id
	DivisionId *string `json:"divisionId,omitempty"`

	// UserName - User name - only returned if division is covered by agentStateNames permission
	UserName *string `json:"userName,omitempty"`

	// ManagerId - The user that this user reports to
	ManagerId *string `json:"managerId,omitempty"`

	// SessionCount - The count of sessions
	SessionCount *int `json:"sessionCount,omitempty"`

	// Sessions - List of sessions
	Sessions *[]Analyticsagentstateagentsessionresult `json:"sessions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsagentstateagentresponse) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsagentstateagentresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Analyticsagentstateagentresponse
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		UserName *string `json:"userName,omitempty"`
		
		ManagerId *string `json:"managerId,omitempty"`
		
		SessionCount *int `json:"sessionCount,omitempty"`
		
		Sessions *[]Analyticsagentstateagentsessionresult `json:"sessions,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		DivisionId: o.DivisionId,
		
		UserName: o.UserName,
		
		ManagerId: o.ManagerId,
		
		SessionCount: o.SessionCount,
		
		Sessions: o.Sessions,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsagentstateagentresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsagentstateagentresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsagentstateagentresponseMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := AnalyticsagentstateagentresponseMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if DivisionId, ok := AnalyticsagentstateagentresponseMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if UserName, ok := AnalyticsagentstateagentresponseMap["userName"].(string); ok {
		o.UserName = &UserName
	}
    
	if ManagerId, ok := AnalyticsagentstateagentresponseMap["managerId"].(string); ok {
		o.ManagerId = &ManagerId
	}
    
	if SessionCount, ok := AnalyticsagentstateagentresponseMap["sessionCount"].(float64); ok {
		SessionCountInt := int(SessionCount)
		o.SessionCount = &SessionCountInt
	}
	
	if Sessions, ok := AnalyticsagentstateagentresponseMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
