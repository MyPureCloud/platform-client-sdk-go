package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsresolution
type Analyticsresolution struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`

	// QueueId - The ID of the last queue on which the conversation was handled.
	QueueId *string `json:"queueId,omitempty"`

	// UserId - The ID of the last user who handled the conversation.
	UserId *string `json:"userId,omitempty"`

	// NNextContactAvoided
	NNextContactAvoided *int `json:"nNextContactAvoided,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsresolution) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsresolution) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime", }
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
	type Alias Analyticsresolution
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		NNextContactAvoided *int `json:"nNextContactAvoided,omitempty"`
		Alias
	}{ 
		EventTime: EventTime,
		
		QueueId: o.QueueId,
		
		UserId: o.UserId,
		
		NNextContactAvoided: o.NNextContactAvoided,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsresolution) UnmarshalJSON(b []byte) error {
	var AnalyticsresolutionMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsresolutionMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := AnalyticsresolutionMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if QueueId, ok := AnalyticsresolutionMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if UserId, ok := AnalyticsresolutionMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if NNextContactAvoided, ok := AnalyticsresolutionMap["nNextContactAvoided"].(float64); ok {
		NNextContactAvoidedInt := int(NNextContactAvoided)
		o.NNextContactAvoided = &NNextContactAvoidedInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsresolution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
