package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodelretrainingerror
type Predictormodelretrainingerror struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ErrorCode - Error code describing model training failure.
	ErrorCode *string `json:"errorCode,omitempty"`

	// DateOfFirstOccurrence - Date when the first retraining failure happened. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateOfFirstOccurrence *time.Time `json:"dateOfFirstOccurrence,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Predictormodelretrainingerror) SetField(field string, fieldValue interface{}) {
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

func (o Predictormodelretrainingerror) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateOfFirstOccurrence", }
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
	type Alias Predictormodelretrainingerror
	
	DateOfFirstOccurrence := new(string)
	if o.DateOfFirstOccurrence != nil {
		
		*DateOfFirstOccurrence = timeutil.Strftime(o.DateOfFirstOccurrence, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateOfFirstOccurrence = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		DateOfFirstOccurrence *string `json:"dateOfFirstOccurrence,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ErrorCode: o.ErrorCode,
		
		DateOfFirstOccurrence: DateOfFirstOccurrence,
		Alias:    (Alias)(o),
	})
}

func (o *Predictormodelretrainingerror) UnmarshalJSON(b []byte) error {
	var PredictormodelretrainingerrorMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelretrainingerrorMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PredictormodelretrainingerrorMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ErrorCode, ok := PredictormodelretrainingerrorMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if dateOfFirstOccurrenceString, ok := PredictormodelretrainingerrorMap["dateOfFirstOccurrence"].(string); ok {
		DateOfFirstOccurrence, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateOfFirstOccurrenceString)
		o.DateOfFirstOccurrence = &DateOfFirstOccurrence
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodelretrainingerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
