package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Addshifttradejobrequest
type Addshifttradejobrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// InitiatingShift - The shift that the initiating user wants to give up in this trade
	InitiatingShift *Initiatingshiftrequestitem `json:"initiatingShift,omitempty"`

	// AcceptableIntervals - Time frames when the initiating user is willing to accept a shift in exchange. Empty means giving up the shift without taking on another one
	AcceptableIntervals *[]Requireddaterange `json:"acceptableIntervals,omitempty"`

	// Target - Optional shift trade target, can be used for example for direct user to user trade
	Target *Shifttradetargetrequestitem `json:"target,omitempty"`

	// ExpirationDate - When this shift trade will expire. Date time is represented as an ISO-8601 string
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Addshifttradejobrequest) SetField(field string, fieldValue interface{}) {
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

func (o Addshifttradejobrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ExpirationDate", }
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
	type Alias Addshifttradejobrequest
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	return json.Marshal(&struct { 
		InitiatingShift *Initiatingshiftrequestitem `json:"initiatingShift,omitempty"`
		
		AcceptableIntervals *[]Requireddaterange `json:"acceptableIntervals,omitempty"`
		
		Target *Shifttradetargetrequestitem `json:"target,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		Alias
	}{ 
		InitiatingShift: o.InitiatingShift,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		Target: o.Target,
		
		ExpirationDate: ExpirationDate,
		Alias:    (Alias)(o),
	})
}

func (o *Addshifttradejobrequest) UnmarshalJSON(b []byte) error {
	var AddshifttradejobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddshifttradejobrequestMap)
	if err != nil {
		return err
	}
	
	if InitiatingShift, ok := AddshifttradejobrequestMap["initiatingShift"].(map[string]interface{}); ok {
		InitiatingShiftString, _ := json.Marshal(InitiatingShift)
		json.Unmarshal(InitiatingShiftString, &o.InitiatingShift)
	}
	
	if AcceptableIntervals, ok := AddshifttradejobrequestMap["acceptableIntervals"].([]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if Target, ok := AddshifttradejobrequestMap["target"].(map[string]interface{}); ok {
		TargetString, _ := json.Marshal(Target)
		json.Unmarshal(TargetString, &o.Target)
	}
	
	if expirationDateString, ok := AddshifttradejobrequestMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Addshifttradejobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
