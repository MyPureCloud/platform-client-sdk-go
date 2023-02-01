package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Orphanupdaterequest
type Orphanupdaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ArchiveDate - The orphan recording's archive date. Must be greater than 1 day from now if set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ArchiveDate *time.Time `json:"archiveDate,omitempty"`

	// DeleteDate - The orphan recording's delete date. Must be greater than archiveDate and exportDate if set, otherwise one day from now. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeleteDate *time.Time `json:"deleteDate,omitempty"`

	// ExportDate - The orphan recording's export date. Must be greater than 1 day from now if set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExportDate *time.Time `json:"exportDate,omitempty"`

	// IntegrationId - IntegrationId to access AWS S3 bucket for export. This field is required if exportDate is set.
	IntegrationId *string `json:"integrationId,omitempty"`

	// ConversationId - A conversation Id that this orphan's recording is to be attached to. If not present, the conversationId will be deduced from the recording media.
	ConversationId *string `json:"conversationId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Orphanupdaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Orphanupdaterequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ArchiveDate","DeleteDate","ExportDate", }
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
	type Alias Orphanupdaterequest
	
	ArchiveDate := new(string)
	if o.ArchiveDate != nil {
		
		*ArchiveDate = timeutil.Strftime(o.ArchiveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ArchiveDate = nil
	}
	
	DeleteDate := new(string)
	if o.DeleteDate != nil {
		
		*DeleteDate = timeutil.Strftime(o.DeleteDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeleteDate = nil
	}
	
	ExportDate := new(string)
	if o.ExportDate != nil {
		
		*ExportDate = timeutil.Strftime(o.ExportDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportDate = nil
	}
	
	return json.Marshal(&struct { 
		ArchiveDate *string `json:"archiveDate,omitempty"`
		
		DeleteDate *string `json:"deleteDate,omitempty"`
		
		ExportDate *string `json:"exportDate,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		Alias
	}{ 
		ArchiveDate: ArchiveDate,
		
		DeleteDate: DeleteDate,
		
		ExportDate: ExportDate,
		
		IntegrationId: o.IntegrationId,
		
		ConversationId: o.ConversationId,
		Alias:    (Alias)(o),
	})
}

func (o *Orphanupdaterequest) UnmarshalJSON(b []byte) error {
	var OrphanupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &OrphanupdaterequestMap)
	if err != nil {
		return err
	}
	
	if archiveDateString, ok := OrphanupdaterequestMap["archiveDate"].(string); ok {
		ArchiveDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", archiveDateString)
		o.ArchiveDate = &ArchiveDate
	}
	
	if deleteDateString, ok := OrphanupdaterequestMap["deleteDate"].(string); ok {
		DeleteDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deleteDateString)
		o.DeleteDate = &DeleteDate
	}
	
	if exportDateString, ok := OrphanupdaterequestMap["exportDate"].(string); ok {
		ExportDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportDateString)
		o.ExportDate = &ExportDate
	}
	
	if IntegrationId, ok := OrphanupdaterequestMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if ConversationId, ok := OrphanupdaterequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Orphanupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
