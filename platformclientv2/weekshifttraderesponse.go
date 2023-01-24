package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekshifttraderesponse
type Weekshifttraderesponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Trade - The shift trade details
	Trade *Shifttraderesponse `json:"trade,omitempty"`

	// MatchReview - A preview of what the schedule would look like if the shift trade is approved plus any violations
	MatchReview *Shifttradematchreviewresponse `json:"matchReview,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Weekshifttraderesponse) SetField(field string, fieldValue interface{}) {
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

func (o Weekshifttraderesponse) MarshalJSON() ([]byte, error) {
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
	type Alias Weekshifttraderesponse
	
	return json.Marshal(&struct { 
		Trade *Shifttraderesponse `json:"trade,omitempty"`
		
		MatchReview *Shifttradematchreviewresponse `json:"matchReview,omitempty"`
		Alias
	}{ 
		Trade: o.Trade,
		
		MatchReview: o.MatchReview,
		Alias:    (Alias)(o),
	})
}

func (o *Weekshifttraderesponse) UnmarshalJSON(b []byte) error {
	var WeekshifttraderesponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekshifttraderesponseMap)
	if err != nil {
		return err
	}
	
	if Trade, ok := WeekshifttraderesponseMap["trade"].(map[string]interface{}); ok {
		TradeString, _ := json.Marshal(Trade)
		json.Unmarshal(TradeString, &o.Trade)
	}
	
	if MatchReview, ok := WeekshifttraderesponseMap["matchReview"].(map[string]interface{}); ok {
		MatchReviewString, _ := json.Marshal(MatchReview)
		json.Unmarshal(MatchReviewString, &o.MatchReview)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
