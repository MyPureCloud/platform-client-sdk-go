package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification
type V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

	// LastSuccessfulSessionId
	LastSuccessfulSessionId *string `json:"lastSuccessfulSessionId,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

	// ForecastDataState
	ForecastDataState *string `json:"forecastDataState,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification) SetField(field string, fieldValue interface{}) {
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

func (o V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification) MarshalJSON() ([]byte, error) {
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
	type Alias V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification
	
	return json.Marshal(&struct { 
		SessionId *string `json:"sessionId,omitempty"`
		
		LastSuccessfulSessionId *string `json:"lastSuccessfulSessionId,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ForecastDataState *string `json:"forecastDataState,omitempty"`
		Alias
	}{ 
		SessionId: o.SessionId,
		
		LastSuccessfulSessionId: o.LastSuccessfulSessionId,
		
		State: o.State,
		
		ErrorCode: o.ErrorCode,
		
		ForecastDataState: o.ForecastDataState,
		Alias:    (Alias)(o),
	})
}

func (o *V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification) UnmarshalJSON(b []byte) error {
	var V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap)
	if err != nil {
		return err
	}
	
	if SessionId, ok := V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if LastSuccessfulSessionId, ok := V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap["lastSuccessfulSessionId"].(string); ok {
		o.LastSuccessfulSessionId = &LastSuccessfulSessionId
	}
    
	if State, ok := V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap["state"].(string); ok {
		o.State = &State
	}
    
	if ErrorCode, ok := V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ForecastDataState, ok := V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotificationMap["forecastDataState"].(string); ok {
		o.ForecastDataState = &ForecastDataState
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2wfmcontinuousforecastsessioneventcontinuousforecastsessionnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
