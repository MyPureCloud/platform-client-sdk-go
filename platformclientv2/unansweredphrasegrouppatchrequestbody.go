package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredphrasegrouppatchrequestbody
type Unansweredphrasegrouppatchrequestbody struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PhraseAssociations - List of phrases and documents to be linked
	PhraseAssociations *[]Phraseassociations `json:"phraseAssociations,omitempty"`

	// DateStart - The start date to be used for filtering phrases. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The end date to be used for filtering phrases. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Unansweredphrasegrouppatchrequestbody) SetField(field string, fieldValue interface{}) {
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

func (o Unansweredphrasegrouppatchrequestbody) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateStart","DateEnd", }

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
	type Alias Unansweredphrasegrouppatchrequestbody
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%d")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		PhraseAssociations *[]Phraseassociations `json:"phraseAssociations,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		Alias
	}{ 
		PhraseAssociations: o.PhraseAssociations,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		Alias:    (Alias)(o),
	})
}

func (o *Unansweredphrasegrouppatchrequestbody) UnmarshalJSON(b []byte) error {
	var UnansweredphrasegrouppatchrequestbodyMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredphrasegrouppatchrequestbodyMap)
	if err != nil {
		return err
	}
	
	if PhraseAssociations, ok := UnansweredphrasegrouppatchrequestbodyMap["phraseAssociations"].([]interface{}); ok {
		PhraseAssociationsString, _ := json.Marshal(PhraseAssociations)
		json.Unmarshal(PhraseAssociationsString, &o.PhraseAssociations)
	}
	
	if dateStartString, ok := UnansweredphrasegrouppatchrequestbodyMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := UnansweredphrasegrouppatchrequestbodyMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredphrasegrouppatchrequestbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
