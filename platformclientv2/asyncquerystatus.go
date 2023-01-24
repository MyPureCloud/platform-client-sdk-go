package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncquerystatus
type Asyncquerystatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The current state of the asynchronous query
	State *string `json:"state,omitempty"`

	// ErrorMessage - The error associated with the current query, if the state is FAILED
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// ExpirationDate - The time at which results for this query will expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// SubmissionDate - The time at which the query was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SubmissionDate *time.Time `json:"submissionDate,omitempty"`

	// CompletionDate - The time at which the query completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CompletionDate *time.Time `json:"completionDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Asyncquerystatus) SetField(field string, fieldValue interface{}) {
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

func (o Asyncquerystatus) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ExpirationDate","SubmissionDate","CompletionDate", }
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
	type Alias Asyncquerystatus
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	SubmissionDate := new(string)
	if o.SubmissionDate != nil {
		
		*SubmissionDate = timeutil.Strftime(o.SubmissionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SubmissionDate = nil
	}
	
	CompletionDate := new(string)
	if o.CompletionDate != nil {
		
		*CompletionDate = timeutil.Strftime(o.CompletionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CompletionDate = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		
		SubmissionDate *string `json:"submissionDate,omitempty"`
		
		CompletionDate *string `json:"completionDate,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		ErrorMessage: o.ErrorMessage,
		
		ExpirationDate: ExpirationDate,
		
		SubmissionDate: SubmissionDate,
		
		CompletionDate: CompletionDate,
		Alias:    (Alias)(o),
	})
}

func (o *Asyncquerystatus) UnmarshalJSON(b []byte) error {
	var AsyncquerystatusMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncquerystatusMap)
	if err != nil {
		return err
	}
	
	if State, ok := AsyncquerystatusMap["state"].(string); ok {
		o.State = &State
	}
    
	if ErrorMessage, ok := AsyncquerystatusMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if expirationDateString, ok := AsyncquerystatusMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	
	if submissionDateString, ok := AsyncquerystatusMap["submissionDate"].(string); ok {
		SubmissionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", submissionDateString)
		o.SubmissionDate = &SubmissionDate
	}
	
	if completionDateString, ok := AsyncquerystatusMap["completionDate"].(string); ok {
		CompletionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", completionDateString)
		o.CompletionDate = &CompletionDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncquerystatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
