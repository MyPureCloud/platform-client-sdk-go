package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowoutcome - Flow Outcome data related to a bot flow which is exiting gracefully.
type Textbotflowoutcome struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OutcomeId - The Flow Outcome ID.
	OutcomeId *string `json:"outcomeId,omitempty"`

	// OutcomeValue - The value of the FlowOutcome.
	OutcomeValue *string `json:"outcomeValue,omitempty"`

	// DateStart - The timestamp for when the Flow Outcome began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The timestamp for when the Flow Outcome finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// Milestones - The Flow Milestones for the Flow Outcome.
	Milestones *[]Textbotflowmilestone `json:"milestones,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotflowoutcome) SetField(field string, fieldValue interface{}) {
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

func (o Textbotflowoutcome) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd", }
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
	type Alias Textbotflowoutcome
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		OutcomeId *string `json:"outcomeId,omitempty"`
		
		OutcomeValue *string `json:"outcomeValue,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Milestones *[]Textbotflowmilestone `json:"milestones,omitempty"`
		Alias
	}{ 
		OutcomeId: o.OutcomeId,
		
		OutcomeValue: o.OutcomeValue,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		Milestones: o.Milestones,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotflowoutcome) UnmarshalJSON(b []byte) error {
	var TextbotflowoutcomeMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowoutcomeMap)
	if err != nil {
		return err
	}
	
	if OutcomeId, ok := TextbotflowoutcomeMap["outcomeId"].(string); ok {
		o.OutcomeId = &OutcomeId
	}
    
	if OutcomeValue, ok := TextbotflowoutcomeMap["outcomeValue"].(string); ok {
		o.OutcomeValue = &OutcomeValue
	}
    
	if dateStartString, ok := TextbotflowoutcomeMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := TextbotflowoutcomeMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if Milestones, ok := TextbotflowoutcomeMap["milestones"].([]interface{}); ok {
		MilestonesString, _ := json.Marshal(Milestones)
		json.Unmarshal(MilestonesString, &o.Milestones)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
