package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestbulkevaluatetemplateitem
type Timeoffrequestbulkevaluatetemplateitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the time-off request
	Id *string `json:"id,omitempty"`

	// User - The user associated with this time-off request
	User *Userreference `json:"user,omitempty"`

	// Status - The status of the time-off request
	Status *string `json:"status,omitempty"`

	// Substatus - The substatus of the time-off request
	Substatus *string `json:"substatus,omitempty"`

	// ProcessingStatus - The processing status of the time-off request
	ProcessingStatus *string `json:"processingStatus,omitempty"`

	// Metadata - The version metadata of the time-off request
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffrequestbulkevaluatetemplateitem) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffrequestbulkevaluatetemplateitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Timeoffrequestbulkevaluatetemplateitem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Substatus *string `json:"substatus,omitempty"`
		
		ProcessingStatus *string `json:"processingStatus,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		Status: o.Status,
		
		Substatus: o.Substatus,
		
		ProcessingStatus: o.ProcessingStatus,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffrequestbulkevaluatetemplateitem) UnmarshalJSON(b []byte) error {
	var TimeoffrequestbulkevaluatetemplateitemMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestbulkevaluatetemplateitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffrequestbulkevaluatetemplateitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if User, ok := TimeoffrequestbulkevaluatetemplateitemMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Status, ok := TimeoffrequestbulkevaluatetemplateitemMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Substatus, ok := TimeoffrequestbulkevaluatetemplateitemMap["substatus"].(string); ok {
		o.Substatus = &Substatus
	}
    
	if ProcessingStatus, ok := TimeoffrequestbulkevaluatetemplateitemMap["processingStatus"].(string); ok {
		o.ProcessingStatus = &ProcessingStatus
	}
    
	if Metadata, ok := TimeoffrequestbulkevaluatetemplateitemMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := TimeoffrequestbulkevaluatetemplateitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestbulkevaluatetemplateitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
