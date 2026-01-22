package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentreschedule
type Learningassignmentreschedule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateRecommendedForCompletion - The recommended completion date of the assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateRecommendedForCompletion *time.Time `json:"dateRecommendedForCompletion,omitempty"`

	// LengthInMinutes - The length in minutes of the assignment
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// AddToSchedule - If True, adds the assignment to their schedule
	AddToSchedule *bool `json:"addToSchedule,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningassignmentreschedule) SetField(field string, fieldValue interface{}) {
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

func (o Learningassignmentreschedule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateRecommendedForCompletion", }
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
	type Alias Learningassignmentreschedule
	
	DateRecommendedForCompletion := new(string)
	if o.DateRecommendedForCompletion != nil {
		
		*DateRecommendedForCompletion = timeutil.Strftime(o.DateRecommendedForCompletion, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateRecommendedForCompletion = nil
	}
	
	return json.Marshal(&struct { 
		DateRecommendedForCompletion *string `json:"dateRecommendedForCompletion,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		AddToSchedule *bool `json:"addToSchedule,omitempty"`
		Alias
	}{ 
		DateRecommendedForCompletion: DateRecommendedForCompletion,
		
		LengthInMinutes: o.LengthInMinutes,
		
		AddToSchedule: o.AddToSchedule,
		Alias:    (Alias)(o),
	})
}

func (o *Learningassignmentreschedule) UnmarshalJSON(b []byte) error {
	var LearningassignmentrescheduleMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentrescheduleMap)
	if err != nil {
		return err
	}
	
	if dateRecommendedForCompletionString, ok := LearningassignmentrescheduleMap["dateRecommendedForCompletion"].(string); ok {
		DateRecommendedForCompletion, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateRecommendedForCompletionString)
		o.DateRecommendedForCompletion = &DateRecommendedForCompletion
	}
	
	if LengthInMinutes, ok := LearningassignmentrescheduleMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if AddToSchedule, ok := LearningassignmentrescheduleMap["addToSchedule"].(bool); ok {
		o.AddToSchedule = &AddToSchedule
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentreschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
