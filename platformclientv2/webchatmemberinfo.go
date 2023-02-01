package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatmemberinfo
type Webchatmemberinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The communicationId of this member.
	Id *string `json:"id,omitempty"`

	// DisplayName - The display name of the member.
	DisplayName *string `json:"displayName,omitempty"`

	// FirstName - The first name of the member.
	FirstName *string `json:"firstName,omitempty"`

	// LastName - The last name of the member.
	LastName *string `json:"lastName,omitempty"`

	// Email - The email address of the member.
	Email *string `json:"email,omitempty"`

	// PhoneNumber - The phone number of the member.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// AvatarImageUrl - The url to the avatar image of the member.
	AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`

	// Role - The role of the member, one of [agent, customer, acd, workflow]
	Role *string `json:"role,omitempty"`

	// JoinDate - The time the member joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	JoinDate *time.Time `json:"joinDate,omitempty"`

	// LeaveDate - The time the member left the conversation, or null if the member is still active in the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LeaveDate *time.Time `json:"leaveDate,omitempty"`

	// AuthenticatedGuest - If true, the guest member is an authenticated guest.
	AuthenticatedGuest *bool `json:"authenticatedGuest,omitempty"`

	// CustomFields - Any custom fields of information pertaining to this member.
	CustomFields *map[string]string `json:"customFields,omitempty"`

	// State - The connection state of this member.
	State *string `json:"state,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webchatmemberinfo) SetField(field string, fieldValue interface{}) {
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

func (o Webchatmemberinfo) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "JoinDate","LeaveDate", }
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
	type Alias Webchatmemberinfo
	
	JoinDate := new(string)
	if o.JoinDate != nil {
		
		*JoinDate = timeutil.Strftime(o.JoinDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		JoinDate = nil
	}
	
	LeaveDate := new(string)
	if o.LeaveDate != nil {
		
		*LeaveDate = timeutil.Strftime(o.LeaveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LeaveDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`
		
		Role *string `json:"role,omitempty"`
		
		JoinDate *string `json:"joinDate,omitempty"`
		
		LeaveDate *string `json:"leaveDate,omitempty"`
		
		AuthenticatedGuest *bool `json:"authenticatedGuest,omitempty"`
		
		CustomFields *map[string]string `json:"customFields,omitempty"`
		
		State *string `json:"state,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DisplayName: o.DisplayName,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Email: o.Email,
		
		PhoneNumber: o.PhoneNumber,
		
		AvatarImageUrl: o.AvatarImageUrl,
		
		Role: o.Role,
		
		JoinDate: JoinDate,
		
		LeaveDate: LeaveDate,
		
		AuthenticatedGuest: o.AuthenticatedGuest,
		
		CustomFields: o.CustomFields,
		
		State: o.State,
		Alias:    (Alias)(o),
	})
}

func (o *Webchatmemberinfo) UnmarshalJSON(b []byte) error {
	var WebchatmemberinfoMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatmemberinfoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatmemberinfoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DisplayName, ok := WebchatmemberinfoMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if FirstName, ok := WebchatmemberinfoMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if LastName, ok := WebchatmemberinfoMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Email, ok := WebchatmemberinfoMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if PhoneNumber, ok := WebchatmemberinfoMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if AvatarImageUrl, ok := WebchatmemberinfoMap["avatarImageUrl"].(string); ok {
		o.AvatarImageUrl = &AvatarImageUrl
	}
    
	if Role, ok := WebchatmemberinfoMap["role"].(string); ok {
		o.Role = &Role
	}
    
	if joinDateString, ok := WebchatmemberinfoMap["joinDate"].(string); ok {
		JoinDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", joinDateString)
		o.JoinDate = &JoinDate
	}
	
	if leaveDateString, ok := WebchatmemberinfoMap["leaveDate"].(string); ok {
		LeaveDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", leaveDateString)
		o.LeaveDate = &LeaveDate
	}
	
	if AuthenticatedGuest, ok := WebchatmemberinfoMap["authenticatedGuest"].(bool); ok {
		o.AuthenticatedGuest = &AuthenticatedGuest
	}
    
	if CustomFields, ok := WebchatmemberinfoMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if State, ok := WebchatmemberinfoMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatmemberinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
