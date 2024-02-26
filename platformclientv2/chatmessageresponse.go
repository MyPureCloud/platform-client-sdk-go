package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmessageresponse
type Chatmessageresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of the message
	Id *string `json:"id,omitempty"`

	// DateCreated - Message's created time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Message's last updated time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ToJid - Jid of message's recipient (roomJid or userJid)
	ToJid *string `json:"toJid,omitempty"`

	// Jid - Jid of message's sender (userJid)
	Jid *string `json:"jid,omitempty"`

	// Body - Message's body
	Body *string `json:"body,omitempty"`

	// Mentions - Message's mentions
	Mentions *map[string]string `json:"mentions,omitempty"`

	// Edited - If message was edited
	Edited *bool `json:"edited,omitempty"`

	// AttachmentDeleted - If message's attachment was deleted
	AttachmentDeleted *bool `json:"attachmentDeleted,omitempty"`

	// FileUri - URI of file attachment
	FileUri *string `json:"fileUri,omitempty"`

	// Thread - The id for a thread this message corresponds to
	Thread *Entity `json:"thread,omitempty"`

	// ParentThread - Parent thread id for thread replies
	ParentThread *Entity `json:"parentThread,omitempty"`

	// User - The user who sent the message
	User *Addressableentityref `json:"user,omitempty"`

	// ToUser - The receiving user of the message
	ToUser *Addressableentityref `json:"toUser,omitempty"`

	// Reactions - The emoji reactions to this message
	Reactions *[]Chatreaction `json:"reactions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Chatmessageresponse) SetField(field string, fieldValue interface{}) {
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

func (o Chatmessageresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Chatmessageresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ToJid *string `json:"toJid,omitempty"`
		
		Jid *string `json:"jid,omitempty"`
		
		Body *string `json:"body,omitempty"`
		
		Mentions *map[string]string `json:"mentions,omitempty"`
		
		Edited *bool `json:"edited,omitempty"`
		
		AttachmentDeleted *bool `json:"attachmentDeleted,omitempty"`
		
		FileUri *string `json:"fileUri,omitempty"`
		
		Thread *Entity `json:"thread,omitempty"`
		
		ParentThread *Entity `json:"parentThread,omitempty"`
		
		User *Addressableentityref `json:"user,omitempty"`
		
		ToUser *Addressableentityref `json:"toUser,omitempty"`
		
		Reactions *[]Chatreaction `json:"reactions,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ToJid: o.ToJid,
		
		Jid: o.Jid,
		
		Body: o.Body,
		
		Mentions: o.Mentions,
		
		Edited: o.Edited,
		
		AttachmentDeleted: o.AttachmentDeleted,
		
		FileUri: o.FileUri,
		
		Thread: o.Thread,
		
		ParentThread: o.ParentThread,
		
		User: o.User,
		
		ToUser: o.ToUser,
		
		Reactions: o.Reactions,
		Alias:    (Alias)(o),
	})
}

func (o *Chatmessageresponse) UnmarshalJSON(b []byte) error {
	var ChatmessageresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ChatmessageresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ChatmessageresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateCreatedString, ok := ChatmessageresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := ChatmessageresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ToJid, ok := ChatmessageresponseMap["toJid"].(string); ok {
		o.ToJid = &ToJid
	}
    
	if Jid, ok := ChatmessageresponseMap["jid"].(string); ok {
		o.Jid = &Jid
	}
    
	if Body, ok := ChatmessageresponseMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if Mentions, ok := ChatmessageresponseMap["mentions"].(map[string]interface{}); ok {
		MentionsString, _ := json.Marshal(Mentions)
		json.Unmarshal(MentionsString, &o.Mentions)
	}
	
	if Edited, ok := ChatmessageresponseMap["edited"].(bool); ok {
		o.Edited = &Edited
	}
    
	if AttachmentDeleted, ok := ChatmessageresponseMap["attachmentDeleted"].(bool); ok {
		o.AttachmentDeleted = &AttachmentDeleted
	}
    
	if FileUri, ok := ChatmessageresponseMap["fileUri"].(string); ok {
		o.FileUri = &FileUri
	}
    
	if Thread, ok := ChatmessageresponseMap["thread"].(map[string]interface{}); ok {
		ThreadString, _ := json.Marshal(Thread)
		json.Unmarshal(ThreadString, &o.Thread)
	}
	
	if ParentThread, ok := ChatmessageresponseMap["parentThread"].(map[string]interface{}); ok {
		ParentThreadString, _ := json.Marshal(ParentThread)
		json.Unmarshal(ParentThreadString, &o.ParentThread)
	}
	
	if User, ok := ChatmessageresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if ToUser, ok := ChatmessageresponseMap["toUser"].(map[string]interface{}); ok {
		ToUserString, _ := json.Marshal(ToUser)
		json.Unmarshal(ToUserString, &o.ToUser)
	}
	
	if Reactions, ok := ChatmessageresponseMap["reactions"].([]interface{}); ok {
		ReactionsString, _ := json.Marshal(Reactions)
		json.Unmarshal(ReactionsString, &o.Reactions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatmessageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
