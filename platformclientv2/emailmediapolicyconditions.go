package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailmediapolicyconditions
type Emailmediapolicyconditions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ForUsers - List of users to apply this policy to. Each user object can include the 'id' field containing the user's unique identifier. Example: [{\"id\":\"<userId>\"}].
	ForUsers *[]Policyuser `json:"forUsers,omitempty"`

	// DateRanges
	DateRanges *[]string `json:"dateRanges,omitempty"`

	// ForQueues
	ForQueues *[]Queue `json:"forQueues,omitempty"`

	// WrapupCodes
	WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`

	// Languages
	Languages *[]Language `json:"languages,omitempty"`

	// TimeAllowed
	TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`

	// Teams - Teams to match conversations against
	Teams *[]Team `json:"teams,omitempty"`

	// CustomerParticipation
	CustomerParticipation *string `json:"customerParticipation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Emailmediapolicyconditions) SetField(field string, fieldValue interface{}) {
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

func (o Emailmediapolicyconditions) MarshalJSON() ([]byte, error) {
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
	type Alias Emailmediapolicyconditions
	
	return json.Marshal(&struct { 
		ForUsers *[]Policyuser `json:"forUsers,omitempty"`
		
		DateRanges *[]string `json:"dateRanges,omitempty"`
		
		ForQueues *[]Queue `json:"forQueues,omitempty"`
		
		WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`
		
		Languages *[]Language `json:"languages,omitempty"`
		
		TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`
		
		Teams *[]Team `json:"teams,omitempty"`
		
		CustomerParticipation *string `json:"customerParticipation,omitempty"`
		Alias
	}{ 
		ForUsers: o.ForUsers,
		
		DateRanges: o.DateRanges,
		
		ForQueues: o.ForQueues,
		
		WrapupCodes: o.WrapupCodes,
		
		Languages: o.Languages,
		
		TimeAllowed: o.TimeAllowed,
		
		Teams: o.Teams,
		
		CustomerParticipation: o.CustomerParticipation,
		Alias:    (Alias)(o),
	})
}

func (o *Emailmediapolicyconditions) UnmarshalJSON(b []byte) error {
	var EmailmediapolicyconditionsMap map[string]interface{}
	err := json.Unmarshal(b, &EmailmediapolicyconditionsMap)
	if err != nil {
		return err
	}
	
	if ForUsers, ok := EmailmediapolicyconditionsMap["forUsers"].([]interface{}); ok {
		ForUsersString, _ := json.Marshal(ForUsers)
		json.Unmarshal(ForUsersString, &o.ForUsers)
	}
	
	if DateRanges, ok := EmailmediapolicyconditionsMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	
	if ForQueues, ok := EmailmediapolicyconditionsMap["forQueues"].([]interface{}); ok {
		ForQueuesString, _ := json.Marshal(ForQueues)
		json.Unmarshal(ForQueuesString, &o.ForQueues)
	}
	
	if WrapupCodes, ok := EmailmediapolicyconditionsMap["wrapupCodes"].([]interface{}); ok {
		WrapupCodesString, _ := json.Marshal(WrapupCodes)
		json.Unmarshal(WrapupCodesString, &o.WrapupCodes)
	}
	
	if Languages, ok := EmailmediapolicyconditionsMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if TimeAllowed, ok := EmailmediapolicyconditionsMap["timeAllowed"].(map[string]interface{}); ok {
		TimeAllowedString, _ := json.Marshal(TimeAllowed)
		json.Unmarshal(TimeAllowedString, &o.TimeAllowed)
	}
	
	if Teams, ok := EmailmediapolicyconditionsMap["teams"].([]interface{}); ok {
		TeamsString, _ := json.Marshal(Teams)
		json.Unmarshal(TeamsString, &o.Teams)
	}
	
	if CustomerParticipation, ok := EmailmediapolicyconditionsMap["customerParticipation"].(string); ok {
		o.CustomerParticipation = &CustomerParticipation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailmediapolicyconditions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
