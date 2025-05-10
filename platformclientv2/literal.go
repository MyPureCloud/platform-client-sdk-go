package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Literal
type Literal struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarString - A string value
	VarString *string `json:"string,omitempty"`

	// Integer - An integer value
	Integer *int `json:"integer,omitempty"`

	// Number - A decimal value
	Number *float64 `json:"number,omitempty"`

	// Date - A date value, must be in the format of yyyy-MM-dd, e.g. 2024-09-23. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`

	// Datetime - A date time value, must be in the format of yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g. 2024-10-02T01:01:01.111Z. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Datetime *time.Time `json:"datetime,omitempty"`

	// Special - A special value enum, such as Wildcard, Null, etc
	Special *string `json:"special,omitempty"`

	// Boolean - A boolean value
	Boolean *bool `json:"boolean,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Literal) SetField(field string, fieldValue interface{}) {
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

func (o Literal) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Datetime", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "Date", }

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
	type Alias Literal
	
	Date := new(string)
	if o.Date != nil {
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%d")
	} else {
		Date = nil
	}
	
	Datetime := new(string)
	if o.Datetime != nil {
		
		*Datetime = timeutil.Strftime(o.Datetime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Datetime = nil
	}
	
	return json.Marshal(&struct { 
		VarString *string `json:"string,omitempty"`
		
		Integer *int `json:"integer,omitempty"`
		
		Number *float64 `json:"number,omitempty"`
		
		Date *string `json:"date,omitempty"`
		
		Datetime *string `json:"datetime,omitempty"`
		
		Special *string `json:"special,omitempty"`
		
		Boolean *bool `json:"boolean,omitempty"`
		Alias
	}{ 
		VarString: o.VarString,
		
		Integer: o.Integer,
		
		Number: o.Number,
		
		Date: Date,
		
		Datetime: Datetime,
		
		Special: o.Special,
		
		Boolean: o.Boolean,
		Alias:    (Alias)(o),
	})
}

func (o *Literal) UnmarshalJSON(b []byte) error {
	var LiteralMap map[string]interface{}
	err := json.Unmarshal(b, &LiteralMap)
	if err != nil {
		return err
	}
	
	if VarString, ok := LiteralMap["string"].(string); ok {
		o.VarString = &VarString
	}
    
	if Integer, ok := LiteralMap["integer"].(float64); ok {
		IntegerInt := int(Integer)
		o.Integer = &IntegerInt
	}
	
	if Number, ok := LiteralMap["number"].(float64); ok {
		o.Number = &Number
	}
    
	if dateString, ok := LiteralMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02", dateString)
		o.Date = &Date
	}
	
	if datetimeString, ok := LiteralMap["datetime"].(string); ok {
		Datetime, _ := time.Parse("2006-01-02T15:04:05.999999Z", datetimeString)
		o.Datetime = &Datetime
	}
	
	if Special, ok := LiteralMap["special"].(string); ok {
		o.Special = &Special
	}
    
	if Boolean, ok := LiteralMap["boolean"].(bool); ok {
		o.Boolean = &Boolean
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Literal) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
