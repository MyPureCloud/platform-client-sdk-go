package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmailboxinfo
type Voicemailmailboxinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UsageSizeBytes - The total number of bytes for all voicemail message audio recordings
	UsageSizeBytes *int `json:"usageSizeBytes,omitempty"`

	// TotalCount - The total number of voicemail messages
	TotalCount *int `json:"totalCount,omitempty"`

	// UnreadCount - The total number of voicemail messages marked as unread
	UnreadCount *int `json:"unreadCount,omitempty"`

	// DeletedCount - The total number of voicemail messages marked as deleted
	DeletedCount *int `json:"deletedCount,omitempty"`

	// CreatedDate - The date of the oldest voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ModifiedDate - The date of the most recent voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// NewestUnreadDate - The date of the most recent unread voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	NewestUnreadDate *time.Time `json:"newestUnreadDate,omitempty"`

	// OldestUnreadDate - The date of the most oldest unread voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	OldestUnreadDate *time.Time `json:"oldestUnreadDate,omitempty"`

	// NewestReadDate - The date of the most recent read voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	NewestReadDate *time.Time `json:"newestReadDate,omitempty"`

	// OldestReadDate - The date of the most oldest read voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	OldestReadDate *time.Time `json:"oldestReadDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailmailboxinfo) SetField(field string, fieldValue interface{}) {
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

func (o Voicemailmailboxinfo) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate","NewestUnreadDate","OldestUnreadDate","NewestReadDate","OldestReadDate", }
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
	type Alias Voicemailmailboxinfo
	
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
	
	NewestUnreadDate := new(string)
	if o.NewestUnreadDate != nil {
		
		*NewestUnreadDate = timeutil.Strftime(o.NewestUnreadDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		NewestUnreadDate = nil
	}
	
	OldestUnreadDate := new(string)
	if o.OldestUnreadDate != nil {
		
		*OldestUnreadDate = timeutil.Strftime(o.OldestUnreadDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OldestUnreadDate = nil
	}
	
	NewestReadDate := new(string)
	if o.NewestReadDate != nil {
		
		*NewestReadDate = timeutil.Strftime(o.NewestReadDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		NewestReadDate = nil
	}
	
	OldestReadDate := new(string)
	if o.OldestReadDate != nil {
		
		*OldestReadDate = timeutil.Strftime(o.OldestReadDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OldestReadDate = nil
	}
	
	return json.Marshal(&struct { 
		UsageSizeBytes *int `json:"usageSizeBytes,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		DeletedCount *int `json:"deletedCount,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		NewestUnreadDate *string `json:"newestUnreadDate,omitempty"`
		
		OldestUnreadDate *string `json:"oldestUnreadDate,omitempty"`
		
		NewestReadDate *string `json:"newestReadDate,omitempty"`
		
		OldestReadDate *string `json:"oldestReadDate,omitempty"`
		Alias
	}{ 
		UsageSizeBytes: o.UsageSizeBytes,
		
		TotalCount: o.TotalCount,
		
		UnreadCount: o.UnreadCount,
		
		DeletedCount: o.DeletedCount,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		NewestUnreadDate: NewestUnreadDate,
		
		OldestUnreadDate: OldestUnreadDate,
		
		NewestReadDate: NewestReadDate,
		
		OldestReadDate: OldestReadDate,
		Alias:    (Alias)(o),
	})
}

func (o *Voicemailmailboxinfo) UnmarshalJSON(b []byte) error {
	var VoicemailmailboxinfoMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmailboxinfoMap)
	if err != nil {
		return err
	}
	
	if UsageSizeBytes, ok := VoicemailmailboxinfoMap["usageSizeBytes"].(float64); ok {
		UsageSizeBytesInt := int(UsageSizeBytes)
		o.UsageSizeBytes = &UsageSizeBytesInt
	}
	
	if TotalCount, ok := VoicemailmailboxinfoMap["totalCount"].(float64); ok {
		TotalCountInt := int(TotalCount)
		o.TotalCount = &TotalCountInt
	}
	
	if UnreadCount, ok := VoicemailmailboxinfoMap["unreadCount"].(float64); ok {
		UnreadCountInt := int(UnreadCount)
		o.UnreadCount = &UnreadCountInt
	}
	
	if DeletedCount, ok := VoicemailmailboxinfoMap["deletedCount"].(float64); ok {
		DeletedCountInt := int(DeletedCount)
		o.DeletedCount = &DeletedCountInt
	}
	
	if createdDateString, ok := VoicemailmailboxinfoMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := VoicemailmailboxinfoMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if newestUnreadDateString, ok := VoicemailmailboxinfoMap["newestUnreadDate"].(string); ok {
		NewestUnreadDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", newestUnreadDateString)
		o.NewestUnreadDate = &NewestUnreadDate
	}
	
	if oldestUnreadDateString, ok := VoicemailmailboxinfoMap["oldestUnreadDate"].(string); ok {
		OldestUnreadDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", oldestUnreadDateString)
		o.OldestUnreadDate = &OldestUnreadDate
	}
	
	if newestReadDateString, ok := VoicemailmailboxinfoMap["newestReadDate"].(string); ok {
		NewestReadDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", newestReadDateString)
		o.NewestReadDate = &NewestReadDate
	}
	
	if oldestReadDateString, ok := VoicemailmailboxinfoMap["oldestReadDate"].(string); ok {
		OldestReadDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", oldestReadDateString)
		o.OldestReadDate = &OldestReadDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmailboxinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
