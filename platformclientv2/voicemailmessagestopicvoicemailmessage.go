package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmessagestopicvoicemailmessage
type Voicemailmessagestopicvoicemailmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Read
	Read *bool `json:"read,omitempty"`

	// AudioRecordingDurationSeconds
	AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`

	// AudioRecordingSizeBytes
	AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`

	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// CreatedDateString
	CreatedDateString *string `json:"createdDateString,omitempty"`

	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`

	// CallerName
	CallerName *string `json:"callerName,omitempty"`

	// Action
	Action *string `json:"action,omitempty"`

	// Note
	Note *string `json:"note,omitempty"`

	// Deleted
	Deleted *bool `json:"deleted,omitempty"`

	// ModifiedByUserId
	ModifiedByUserId *string `json:"modifiedByUserId,omitempty"`

	// CopiedTo
	CopiedTo *[]Voicemailmessagestopicvoicemailcopyrecord `json:"copiedTo,omitempty"`

	// CopiedFrom
	CopiedFrom *Voicemailmessagestopicvoicemailcopyrecord `json:"copiedFrom,omitempty"`

	// ModifiedDateString
	ModifiedDateString *string `json:"modifiedDateString,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailmessagestopicvoicemailmessage) SetField(field string, fieldValue interface{}) {
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

func (o Voicemailmessagestopicvoicemailmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate", }
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
	type Alias Voicemailmessagestopicvoicemailmessage
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`
		
		AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		CreatedDateString *string `json:"createdDateString,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		ModifiedByUserId *string `json:"modifiedByUserId,omitempty"`
		
		CopiedTo *[]Voicemailmessagestopicvoicemailcopyrecord `json:"copiedTo,omitempty"`
		
		CopiedFrom *Voicemailmessagestopicvoicemailcopyrecord `json:"copiedFrom,omitempty"`
		
		ModifiedDateString *string `json:"modifiedDateString,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Read: o.Read,
		
		AudioRecordingDurationSeconds: o.AudioRecordingDurationSeconds,
		
		AudioRecordingSizeBytes: o.AudioRecordingSizeBytes,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		CreatedDateString: o.CreatedDateString,
		
		CallerAddress: o.CallerAddress,
		
		CallerName: o.CallerName,
		
		Action: o.Action,
		
		Note: o.Note,
		
		Deleted: o.Deleted,
		
		ModifiedByUserId: o.ModifiedByUserId,
		
		CopiedTo: o.CopiedTo,
		
		CopiedFrom: o.CopiedFrom,
		
		ModifiedDateString: o.ModifiedDateString,
		Alias:    (Alias)(o),
	})
}

func (o *Voicemailmessagestopicvoicemailmessage) UnmarshalJSON(b []byte) error {
	var VoicemailmessagestopicvoicemailmessageMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmessagestopicvoicemailmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := VoicemailmessagestopicvoicemailmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Read, ok := VoicemailmessagestopicvoicemailmessageMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if AudioRecordingDurationSeconds, ok := VoicemailmessagestopicvoicemailmessageMap["audioRecordingDurationSeconds"].(float64); ok {
		AudioRecordingDurationSecondsInt := int(AudioRecordingDurationSeconds)
		o.AudioRecordingDurationSeconds = &AudioRecordingDurationSecondsInt
	}
	
	if AudioRecordingSizeBytes, ok := VoicemailmessagestopicvoicemailmessageMap["audioRecordingSizeBytes"].(float64); ok {
		AudioRecordingSizeBytesInt := int(AudioRecordingSizeBytes)
		o.AudioRecordingSizeBytes = &AudioRecordingSizeBytesInt
	}
	
	if createdDateString, ok := VoicemailmessagestopicvoicemailmessageMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := VoicemailmessagestopicvoicemailmessageMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if CreatedDateString, ok := VoicemailmessagestopicvoicemailmessageMap["createdDateString"].(string); ok {
		o.CreatedDateString = &CreatedDateString
	}
    
	if CallerAddress, ok := VoicemailmessagestopicvoicemailmessageMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if CallerName, ok := VoicemailmessagestopicvoicemailmessageMap["callerName"].(string); ok {
		o.CallerName = &CallerName
	}
    
	if Action, ok := VoicemailmessagestopicvoicemailmessageMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Note, ok := VoicemailmessagestopicvoicemailmessageMap["note"].(string); ok {
		o.Note = &Note
	}
    
	if Deleted, ok := VoicemailmessagestopicvoicemailmessageMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if ModifiedByUserId, ok := VoicemailmessagestopicvoicemailmessageMap["modifiedByUserId"].(string); ok {
		o.ModifiedByUserId = &ModifiedByUserId
	}
    
	if CopiedTo, ok := VoicemailmessagestopicvoicemailmessageMap["copiedTo"].([]interface{}); ok {
		CopiedToString, _ := json.Marshal(CopiedTo)
		json.Unmarshal(CopiedToString, &o.CopiedTo)
	}
	
	if CopiedFrom, ok := VoicemailmessagestopicvoicemailmessageMap["copiedFrom"].(map[string]interface{}); ok {
		CopiedFromString, _ := json.Marshal(CopiedFrom)
		json.Unmarshal(CopiedFromString, &o.CopiedFrom)
	}
	
	if ModifiedDateString, ok := VoicemailmessagestopicvoicemailmessageMap["modifiedDateString"].(string); ok {
		o.ModifiedDateString = &ModifiedDateString
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
