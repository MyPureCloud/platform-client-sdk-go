package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// ReviewedBy
	ReviewedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"reviewedBy,omitempty"`

	// ReviewedDate
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`

	// FailureReason
	FailureReason *string `json:"failureReason,omitempty"`

	// Metadata
	Metadata *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReviewedDate", }
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
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ReviewedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		Metadata *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		FailureReason: o.FailureReason,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) UnmarshalJSON(b []byte) error {
	var WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["state"].(string); ok {
		o.State = &State
	}
    
	if ReviewedBy, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if FailureReason, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    
	if Metadata, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
