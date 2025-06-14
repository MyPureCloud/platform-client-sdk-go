package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewjob
type Journeyviewjob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DateCreated - Timestamp of execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateCompleted - Timestamp of completion. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`

	// Status - The status of the job
	Status *string `json:"status,omitempty"`

	// JourneyView - The journey view for which the job is executed
	JourneyView *Journeyview `json:"journeyView,omitempty"`

	// DateCompletionEstimated - Timestamp for the estimated time of completion. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompletionEstimated *time.Time `json:"dateCompletionEstimated,omitempty"`

	// EstimatedCompletionMargin - Margin of error of the estimated time of completion
	EstimatedCompletionMargin *int `json:"estimatedCompletionMargin,omitempty"`

	// UserId - Id of the user who submitted the request
	UserId *string `json:"userId,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewjob) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewjob) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateCompleted","DateCompletionEstimated", }
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
	type Alias Journeyviewjob
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateCompletionEstimated := new(string)
	if o.DateCompletionEstimated != nil {
		
		*DateCompletionEstimated = timeutil.Strftime(o.DateCompletionEstimated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompletionEstimated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		JourneyView *Journeyview `json:"journeyView,omitempty"`
		
		DateCompletionEstimated *string `json:"dateCompletionEstimated,omitempty"`
		
		EstimatedCompletionMargin *int `json:"estimatedCompletionMargin,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DateCreated: DateCreated,
		
		DateCompleted: DateCompleted,
		
		Status: o.Status,
		
		JourneyView: o.JourneyView,
		
		DateCompletionEstimated: DateCompletionEstimated,
		
		EstimatedCompletionMargin: o.EstimatedCompletionMargin,
		
		UserId: o.UserId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewjob) UnmarshalJSON(b []byte) error {
	var JourneyviewjobMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyviewjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateCreatedString, ok := JourneyviewjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateCompletedString, ok := JourneyviewjobMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if Status, ok := JourneyviewjobMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if JourneyView, ok := JourneyviewjobMap["journeyView"].(map[string]interface{}); ok {
		JourneyViewString, _ := json.Marshal(JourneyView)
		json.Unmarshal(JourneyViewString, &o.JourneyView)
	}
	
	if dateCompletionEstimatedString, ok := JourneyviewjobMap["dateCompletionEstimated"].(string); ok {
		DateCompletionEstimated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletionEstimatedString)
		o.DateCompletionEstimated = &DateCompletionEstimated
	}
	
	if EstimatedCompletionMargin, ok := JourneyviewjobMap["estimatedCompletionMargin"].(float64); ok {
		EstimatedCompletionMarginInt := int(EstimatedCompletionMargin)
		o.EstimatedCompletionMargin = &EstimatedCompletionMarginInt
	}
	
	if UserId, ok := JourneyviewjobMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if SelfUri, ok := JourneyviewjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
