package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Minerexecuterequest
type Minerexecuterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateStart - Start date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - End date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// UploadKey - Location of input conversations.
	UploadKey *string `json:"uploadKey,omitempty"`

	// MediaType - Media type for filtering conversations.
	MediaType *string `json:"mediaType,omitempty"`

	// ParticipantType - Type of the participant, either agent, customer or both.
	ParticipantType *string `json:"participantType,omitempty"`

	// QueueIds - List of queue IDs for filtering conversations.
	QueueIds *[]string `json:"queueIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Minerexecuterequest) SetField(field string, fieldValue interface{}) {
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

func (o Minerexecuterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Minerexecuterequest
	
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
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ParticipantType *string `json:"participantType,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		Alias
	}{ 
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		UploadKey: o.UploadKey,
		
		MediaType: o.MediaType,
		
		ParticipantType: o.ParticipantType,
		
		QueueIds: o.QueueIds,
		Alias:    (Alias)(o),
	})
}

func (o *Minerexecuterequest) UnmarshalJSON(b []byte) error {
	var MinerexecuterequestMap map[string]interface{}
	err := json.Unmarshal(b, &MinerexecuterequestMap)
	if err != nil {
		return err
	}
	
	if dateStartString, ok := MinerexecuterequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := MinerexecuterequestMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if UploadKey, ok := MinerexecuterequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if MediaType, ok := MinerexecuterequestMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ParticipantType, ok := MinerexecuterequestMap["participantType"].(string); ok {
		o.ParticipantType = &ParticipantType
	}
    
	if QueueIds, ok := MinerexecuterequestMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Minerexecuterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
