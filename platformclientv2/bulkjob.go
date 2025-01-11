package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkjob
type Bulkjob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// State - The bulk job state.
	State *string `json:"state,omitempty"`

	// Action - The bulk job action. This determines what the bulk job does, for example, terminate workitems.
	Action *string `json:"action,omitempty"`

	// TotalCount - Total count of items to be processed in the bulk job.
	TotalCount *int `json:"totalCount,omitempty"`

	// SuccessfulCount - Count of successfully processed items in the bulk job.
	SuccessfulCount *int `json:"successfulCount,omitempty"`

	// FailedCount - Count of failed processed items in the bulk job.
	FailedCount *int `json:"failedCount,omitempty"`

	// DateStarted - The bulk job start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// DateFinished - The bulk job finished date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateFinished *time.Time `json:"dateFinished,omitempty"`

	// DateModified - The bulk job modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bulkjob) SetField(field string, fieldValue interface{}) {
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

func (o Bulkjob) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStarted","DateFinished","DateModified", }
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
	type Alias Bulkjob
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateFinished := new(string)
	if o.DateFinished != nil {
		
		*DateFinished = timeutil.Strftime(o.DateFinished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateFinished = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		
		SuccessfulCount *int `json:"successfulCount,omitempty"`
		
		FailedCount *int `json:"failedCount,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateFinished *string `json:"dateFinished,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Action: o.Action,
		
		TotalCount: o.TotalCount,
		
		SuccessfulCount: o.SuccessfulCount,
		
		FailedCount: o.FailedCount,
		
		DateStarted: DateStarted,
		
		DateFinished: DateFinished,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Bulkjob) UnmarshalJSON(b []byte) error {
	var BulkjobMap map[string]interface{}
	err := json.Unmarshal(b, &BulkjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := BulkjobMap["state"].(string); ok {
		o.State = &State
	}
    
	if Action, ok := BulkjobMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if TotalCount, ok := BulkjobMap["totalCount"].(float64); ok {
		TotalCountInt := int(TotalCount)
		o.TotalCount = &TotalCountInt
	}
	
	if SuccessfulCount, ok := BulkjobMap["successfulCount"].(float64); ok {
		SuccessfulCountInt := int(SuccessfulCount)
		o.SuccessfulCount = &SuccessfulCountInt
	}
	
	if FailedCount, ok := BulkjobMap["failedCount"].(float64); ok {
		FailedCountInt := int(FailedCount)
		o.FailedCount = &FailedCountInt
	}
	
	if dateStartedString, ok := BulkjobMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if dateFinishedString, ok := BulkjobMap["dateFinished"].(string); ok {
		DateFinished, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateFinishedString)
		o.DateFinished = &DateFinished
	}
	
	if dateModifiedString, ok := BulkjobMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := BulkjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
