package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatitem
type Chatitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Open - If the chat is open
	Open *bool `json:"open,omitempty"`

	// Favorite - The favorite entity for the chat, only appears if the chat is favorited
	Favorite *Chatfavorite `json:"favorite,omitempty"`

	// Images - Avatar images for the chat
	Images *[]Image `json:"images,omitempty"`

	// DateLastMessage - The date of the last message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastMessage *time.Time `json:"dateLastMessage,omitempty"`

	// DateClosed - The date the chat was closed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateClosed *time.Time `json:"dateClosed,omitempty"`

	// User - The other 1on1 user
	User *Chatuserref `json:"user,omitempty"`

	// Room - The room of the chat
	Room *Room `json:"room,omitempty"`

	// ChatType - The type of chat
	ChatType *string `json:"chatType,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Chatitem) SetField(field string, fieldValue interface{}) {
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

func (o Chatitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastMessage","DateClosed", }
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
	type Alias Chatitem
	
	DateLastMessage := new(string)
	if o.DateLastMessage != nil {
		
		*DateLastMessage = timeutil.Strftime(o.DateLastMessage, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastMessage = nil
	}
	
	DateClosed := new(string)
	if o.DateClosed != nil {
		
		*DateClosed = timeutil.Strftime(o.DateClosed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateClosed = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Open *bool `json:"open,omitempty"`
		
		Favorite *Chatfavorite `json:"favorite,omitempty"`
		
		Images *[]Image `json:"images,omitempty"`
		
		DateLastMessage *string `json:"dateLastMessage,omitempty"`
		
		DateClosed *string `json:"dateClosed,omitempty"`
		
		User *Chatuserref `json:"user,omitempty"`
		
		Room *Room `json:"room,omitempty"`
		
		ChatType *string `json:"chatType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Open: o.Open,
		
		Favorite: o.Favorite,
		
		Images: o.Images,
		
		DateLastMessage: DateLastMessage,
		
		DateClosed: DateClosed,
		
		User: o.User,
		
		Room: o.Room,
		
		ChatType: o.ChatType,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Chatitem) UnmarshalJSON(b []byte) error {
	var ChatitemMap map[string]interface{}
	err := json.Unmarshal(b, &ChatitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ChatitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ChatitemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Open, ok := ChatitemMap["open"].(bool); ok {
		o.Open = &Open
	}
    
	if Favorite, ok := ChatitemMap["favorite"].(map[string]interface{}); ok {
		FavoriteString, _ := json.Marshal(Favorite)
		json.Unmarshal(FavoriteString, &o.Favorite)
	}
	
	if Images, ok := ChatitemMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if dateLastMessageString, ok := ChatitemMap["dateLastMessage"].(string); ok {
		DateLastMessage, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastMessageString)
		o.DateLastMessage = &DateLastMessage
	}
	
	if dateClosedString, ok := ChatitemMap["dateClosed"].(string); ok {
		DateClosed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateClosedString)
		o.DateClosed = &DateClosed
	}
	
	if User, ok := ChatitemMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Room, ok := ChatitemMap["room"].(map[string]interface{}); ok {
		RoomString, _ := json.Marshal(Room)
		json.Unmarshal(RoomString, &o.Room)
	}
	
	if ChatType, ok := ChatitemMap["chatType"].(string); ok {
		o.ChatType = &ChatType
	}
    
	if SelfUri, ok := ChatitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Chatitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
