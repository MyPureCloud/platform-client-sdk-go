package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmessage
type Voicemailmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Conversation - The conversation that the voicemail message is associated with
	Conversation *Conversation `json:"conversation,omitempty"`

	// Read - Whether the voicemail message is marked as read
	Read *bool `json:"read,omitempty"`

	// AudioRecordingDurationSeconds - The voicemail message's audio recording duration in seconds
	AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`

	// AudioRecordingSizeBytes - The voicemail message's audio recording size in bytes
	AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`

	// CreatedDate - The date the voicemail message was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ModifiedDate - The date the voicemail message was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// DeletedDate - The date the voicemail message deleted property was set to true. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeletedDate *time.Time `json:"deletedDate,omitempty"`

	// CallerAddress - The caller address
	CallerAddress *string `json:"callerAddress,omitempty"`

	// CallerName - Optionally the name of the caller that left the voicemail message if the caller was a known user
	CallerName *string `json:"callerName,omitempty"`

	// CallerUser - Optionally the user that left the voicemail message if the caller was a known user
	CallerUser *User `json:"callerUser,omitempty"`

	// Deleted - Whether the voicemail message has been marked as deleted
	Deleted *bool `json:"deleted,omitempty"`

	// Note - An optional note
	Note *string `json:"note,omitempty"`

	// User - The user that the voicemail message belongs to or null which means the voicemail message belongs to a group or queue
	User *User `json:"user,omitempty"`

	// Group - The group that the voicemail message belongs to or null which means the voicemail message belongs to a user or queue
	Group *Group `json:"group,omitempty"`

	// Queue - The queue that the voicemail message belongs to or null which means the voicemail message belongs to a user or group
	Queue *Queue `json:"queue,omitempty"`

	// CopiedFrom - Represents where this voicemail message was copied from
	CopiedFrom *Voicemailcopyrecord `json:"copiedFrom,omitempty"`

	// CopiedTo - Represents where this voicemail has been copied to
	CopiedTo *[]Voicemailcopyrecord `json:"copiedTo,omitempty"`

	// DeleteRetentionPolicy - The retention policy for this voicemail when deleted is set to true
	DeleteRetentionPolicy *Voicemailretentionpolicy `json:"deleteRetentionPolicy,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailmessage) SetField(field string, fieldValue interface{}) {
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

func (o Voicemailmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate","DeletedDate", }
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
	type Alias Voicemailmessage
	
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
	
	DeletedDate := new(string)
	if o.DeletedDate != nil {
		
		*DeletedDate = timeutil.Strftime(o.DeletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeletedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`
		
		AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		DeletedDate *string `json:"deletedDate,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		CallerUser *User `json:"callerUser,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		CopiedFrom *Voicemailcopyrecord `json:"copiedFrom,omitempty"`
		
		CopiedTo *[]Voicemailcopyrecord `json:"copiedTo,omitempty"`
		
		DeleteRetentionPolicy *Voicemailretentionpolicy `json:"deleteRetentionPolicy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Conversation: o.Conversation,
		
		Read: o.Read,
		
		AudioRecordingDurationSeconds: o.AudioRecordingDurationSeconds,
		
		AudioRecordingSizeBytes: o.AudioRecordingSizeBytes,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		DeletedDate: DeletedDate,
		
		CallerAddress: o.CallerAddress,
		
		CallerName: o.CallerName,
		
		CallerUser: o.CallerUser,
		
		Deleted: o.Deleted,
		
		Note: o.Note,
		
		User: o.User,
		
		Group: o.Group,
		
		Queue: o.Queue,
		
		CopiedFrom: o.CopiedFrom,
		
		CopiedTo: o.CopiedTo,
		
		DeleteRetentionPolicy: o.DeleteRetentionPolicy,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Voicemailmessage) UnmarshalJSON(b []byte) error {
	var VoicemailmessageMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := VoicemailmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Conversation, ok := VoicemailmessageMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Read, ok := VoicemailmessageMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if AudioRecordingDurationSeconds, ok := VoicemailmessageMap["audioRecordingDurationSeconds"].(float64); ok {
		AudioRecordingDurationSecondsInt := int(AudioRecordingDurationSeconds)
		o.AudioRecordingDurationSeconds = &AudioRecordingDurationSecondsInt
	}
	
	if AudioRecordingSizeBytes, ok := VoicemailmessageMap["audioRecordingSizeBytes"].(float64); ok {
		AudioRecordingSizeBytesInt := int(AudioRecordingSizeBytes)
		o.AudioRecordingSizeBytes = &AudioRecordingSizeBytesInt
	}
	
	if createdDateString, ok := VoicemailmessageMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := VoicemailmessageMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if deletedDateString, ok := VoicemailmessageMap["deletedDate"].(string); ok {
		DeletedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deletedDateString)
		o.DeletedDate = &DeletedDate
	}
	
	if CallerAddress, ok := VoicemailmessageMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if CallerName, ok := VoicemailmessageMap["callerName"].(string); ok {
		o.CallerName = &CallerName
	}
    
	if CallerUser, ok := VoicemailmessageMap["callerUser"].(map[string]interface{}); ok {
		CallerUserString, _ := json.Marshal(CallerUser)
		json.Unmarshal(CallerUserString, &o.CallerUser)
	}
	
	if Deleted, ok := VoicemailmessageMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if Note, ok := VoicemailmessageMap["note"].(string); ok {
		o.Note = &Note
	}
    
	if User, ok := VoicemailmessageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Group, ok := VoicemailmessageMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Queue, ok := VoicemailmessageMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if CopiedFrom, ok := VoicemailmessageMap["copiedFrom"].(map[string]interface{}); ok {
		CopiedFromString, _ := json.Marshal(CopiedFrom)
		json.Unmarshal(CopiedFromString, &o.CopiedFrom)
	}
	
	if CopiedTo, ok := VoicemailmessageMap["copiedTo"].([]interface{}); ok {
		CopiedToString, _ := json.Marshal(CopiedTo)
		json.Unmarshal(CopiedToString, &o.CopiedTo)
	}
	
	if DeleteRetentionPolicy, ok := VoicemailmessageMap["deleteRetentionPolicy"].(map[string]interface{}); ok {
		DeleteRetentionPolicyString, _ := json.Marshal(DeleteRetentionPolicy)
		json.Unmarshal(DeleteRetentionPolicyString, &o.DeleteRetentionPolicy)
	}
	
	if SelfUri, ok := VoicemailmessageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
