package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingusernotificationtopiccoachingusernotification
type Wemcoachingusernotificationtopiccoachingusernotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// MarkedAsRead
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`

	// ActionType
	ActionType *string `json:"actionType,omitempty"`

	// Relationship
	Relationship *string `json:"relationship,omitempty"`

	// Appointment
	Appointment *Wemcoachingusernotificationtopiccoachingappointmentreference `json:"appointment,omitempty"`

	// DateStart
	DateStart *time.Time `json:"dateStart,omitempty"`

	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wemcoachingusernotificationtopiccoachingusernotification) SetField(field string, fieldValue interface{}) {
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

func (o Wemcoachingusernotificationtopiccoachingusernotification) MarshalJSON() ([]byte, error) {
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
	type Alias Wemcoachingusernotificationtopiccoachingusernotification
	
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
		
		Appointment *Wemcoachingusernotificationtopiccoachingappointmentreference `json:"appointment,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MarkedAsRead: o.MarkedAsRead,
		
		ActionType: o.ActionType,
		
		Relationship: o.Relationship,
		
		Appointment: o.Appointment,
		
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Status: o.Status,
		Alias:    (Alias)(o),
	})
}

func (o *Wemcoachingusernotificationtopiccoachingusernotification) UnmarshalJSON(b []byte) error {
	var WemcoachingusernotificationtopiccoachingusernotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WemcoachingusernotificationtopiccoachingusernotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MarkedAsRead, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
    
	if ActionType, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["actionType"].(string); ok {
		o.ActionType = &ActionType
	}
    
	if Relationship, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["relationship"].(string); ok {
		o.Relationship = &Relationship
	}
    
	if Appointment, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["appointment"].(map[string]interface{}); ok {
		AppointmentString, _ := json.Marshal(Appointment)
		json.Unmarshal(AppointmentString, &o.Appointment)
	}
	
	if dateStartString, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Status, ok := WemcoachingusernotificationtopiccoachingusernotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wemcoachingusernotificationtopiccoachingusernotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
