package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingnotification
type Coachingnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the appointment for this notification.
	Name *string `json:"name,omitempty"`

	// MarkedAsRead - Indicates if notification is read or unread
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`

	// ActionType - Action causing the notification.
	ActionType *string `json:"actionType,omitempty"`

	// Relationship - The relationship of this user to this notification's appointment
	Relationship *string `json:"relationship,omitempty"`

	// DateStart - The start time of the appointment relating to this notification. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// LengthInMinutes - The duration of the appointment on this notification
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// Status - The status of the appointment for this notification
	Status *string `json:"status,omitempty"`

	// User - The user of this notification
	User *Userreference `json:"user,omitempty"`

	// Appointment - The appointment
	Appointment *Coachingappointmentresponse `json:"appointment,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Coachingnotification) SetField(field string, fieldValue interface{}) {
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

func (o Coachingnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart", }
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
	type Alias Coachingnotification
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		ActionType *string `json:"actionType,omitempty"`
		
		Relationship *string `json:"relationship,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Appointment *Coachingappointmentresponse `json:"appointment,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MarkedAsRead: o.MarkedAsRead,
		
		ActionType: o.ActionType,
		
		Relationship: o.Relationship,
		
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Status: o.Status,
		
		User: o.User,
		
		Appointment: o.Appointment,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Coachingnotification) UnmarshalJSON(b []byte) error {
	var CoachingnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CoachingnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CoachingnotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MarkedAsRead, ok := CoachingnotificationMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
    
	if ActionType, ok := CoachingnotificationMap["actionType"].(string); ok {
		o.ActionType = &ActionType
	}
    
	if Relationship, ok := CoachingnotificationMap["relationship"].(string); ok {
		o.Relationship = &Relationship
	}
    
	if dateStartString, ok := CoachingnotificationMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := CoachingnotificationMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Status, ok := CoachingnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if User, ok := CoachingnotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Appointment, ok := CoachingnotificationMap["appointment"].(map[string]interface{}); ok {
		AppointmentString, _ := json.Marshal(Appointment)
		json.Unmarshal(AppointmentString, &o.Appointment)
	}
	
	if SelfUri, ok := CoachingnotificationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
