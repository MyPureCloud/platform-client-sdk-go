package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyassignment
type Surveyassignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SurveyForm - The survey form used for this survey.
	SurveyForm *Publishedsurveyformreference `json:"surveyForm,omitempty"`

	// Flow - The URI reference to the flow associated with this survey.
	Flow *Domainentityref `json:"flow,omitempty"`

	// InviteTimeInterval - An ISO 8601 repeated interval consisting of the number of repetitions, the start datetime, and the interval (e.g. R2/2018-03-01T13:00:00Z/P1M10DT2H30M). Total duration must not exceed 90 days.
	InviteTimeInterval *string `json:"inviteTimeInterval,omitempty"`

	// SendingUser - User together with sendingDomain used to send email, null to use no-reply
	SendingUser *string `json:"sendingUser,omitempty"`

	// SendingDomain - Validated email domain, required
	SendingDomain *string `json:"sendingDomain,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Surveyassignment) SetField(field string, fieldValue interface{}) {
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

func (o Surveyassignment) MarshalJSON() ([]byte, error) {
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
	type Alias Surveyassignment
	
	return json.Marshal(&struct { 
		SurveyForm *Publishedsurveyformreference `json:"surveyForm,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		InviteTimeInterval *string `json:"inviteTimeInterval,omitempty"`
		
		SendingUser *string `json:"sendingUser,omitempty"`
		
		SendingDomain *string `json:"sendingDomain,omitempty"`
		Alias
	}{ 
		SurveyForm: o.SurveyForm,
		
		Flow: o.Flow,
		
		InviteTimeInterval: o.InviteTimeInterval,
		
		SendingUser: o.SendingUser,
		
		SendingDomain: o.SendingDomain,
		Alias:    (Alias)(o),
	})
}

func (o *Surveyassignment) UnmarshalJSON(b []byte) error {
	var SurveyassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyassignmentMap)
	if err != nil {
		return err
	}
	
	if SurveyForm, ok := SurveyassignmentMap["surveyForm"].(map[string]interface{}); ok {
		SurveyFormString, _ := json.Marshal(SurveyForm)
		json.Unmarshal(SurveyFormString, &o.SurveyForm)
	}
	
	if Flow, ok := SurveyassignmentMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if InviteTimeInterval, ok := SurveyassignmentMap["inviteTimeInterval"].(string); ok {
		o.InviteTimeInterval = &InviteTimeInterval
	}
    
	if SendingUser, ok := SurveyassignmentMap["sendingUser"].(string); ok {
		o.SendingUser = &SendingUser
	}
    
	if SendingDomain, ok := SurveyassignmentMap["sendingDomain"].(string); ok {
		o.SendingDomain = &SendingDomain
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
