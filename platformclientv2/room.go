package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Room
type Room struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The jid of the room if adhoc, the id of the group for group rooms
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// DateCreated - Room's created time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// RoomType - The type of room
	RoomType *string `json:"roomType,omitempty"`

	// Description - Room's description
	Description *string `json:"description,omitempty"`

	// Subject - Room's subject
	Subject *string `json:"subject,omitempty"`

	// ParticipantLimit - Room's size limit
	ParticipantLimit *int `json:"participantLimit,omitempty"`

	// Owners - Room's owners
	Owners *[]Userreference `json:"owners,omitempty"`

	// PinnedMessages - Room's pinned messages
	PinnedMessages *[]Addressableentityref `json:"pinnedMessages,omitempty"`

	// Jid - The jid of the room
	Jid *string `json:"jid,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Room) SetField(field string, fieldValue interface{}) {
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

func (o Room) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Room
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		RoomType *string `json:"roomType,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		ParticipantLimit *int `json:"participantLimit,omitempty"`
		
		Owners *[]Userreference `json:"owners,omitempty"`
		
		PinnedMessages *[]Addressableentityref `json:"pinnedMessages,omitempty"`
		
		Jid *string `json:"jid,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		RoomType: o.RoomType,
		
		Description: o.Description,
		
		Subject: o.Subject,
		
		ParticipantLimit: o.ParticipantLimit,
		
		Owners: o.Owners,
		
		PinnedMessages: o.PinnedMessages,
		
		Jid: o.Jid,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Room) UnmarshalJSON(b []byte) error {
	var RoomMap map[string]interface{}
	err := json.Unmarshal(b, &RoomMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RoomMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := RoomMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := RoomMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if RoomType, ok := RoomMap["roomType"].(string); ok {
		o.RoomType = &RoomType
	}
    
	if Description, ok := RoomMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Subject, ok := RoomMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if ParticipantLimit, ok := RoomMap["participantLimit"].(float64); ok {
		ParticipantLimitInt := int(ParticipantLimit)
		o.ParticipantLimit = &ParticipantLimitInt
	}
	
	if Owners, ok := RoomMap["owners"].([]interface{}); ok {
		OwnersString, _ := json.Marshal(Owners)
		json.Unmarshal(OwnersString, &o.Owners)
	}
	
	if PinnedMessages, ok := RoomMap["pinnedMessages"].([]interface{}); ok {
		PinnedMessagesString, _ := json.Marshal(PinnedMessages)
		json.Unmarshal(PinnedMessagesString, &o.PinnedMessages)
	}
	
	if Jid, ok := RoomMap["jid"].(string); ok {
		o.Jid = &Jid
	}
    
	if SelfUri, ok := RoomMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Room) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
