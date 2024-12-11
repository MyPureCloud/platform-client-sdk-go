package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentdatepicker - DateTimePicker content object.
type Conversationcontentdatepicker struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`

	// Subtitle - Text to show in the description.
	Subtitle *string `json:"subtitle,omitempty"`

	// ImageUrl - URL of an image
	ImageUrl *string `json:"imageUrl,omitempty"`

	// DateMinimum - The minimum Date Enabled in the datepicker calendar, format: ISO 8601.
	DateMinimum *time.Time `json:"dateMinimum,omitempty"`

	// DateMaximum - The maximum Date Enabled in the datepicker calendar, format: ISO 8601.
	DateMaximum *time.Time `json:"dateMaximum,omitempty"`

	// AvailableTimes - An array of available times objects.
	AvailableTimes *[]Conversationcontentdatepickeravailabletime `json:"availableTimes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentdatepicker) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentdatepicker) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateMinimum","DateMaximum", }
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
	type Alias Conversationcontentdatepicker
	
	DateMinimum := new(string)
	if o.DateMinimum != nil {
		
		*DateMinimum = timeutil.Strftime(o.DateMinimum, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMinimum = nil
	}
	
	DateMaximum := new(string)
	if o.DateMaximum != nil {
		
		*DateMaximum = timeutil.Strftime(o.DateMaximum, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMaximum = nil
	}
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Subtitle *string `json:"subtitle,omitempty"`
		
		ImageUrl *string `json:"imageUrl,omitempty"`
		
		DateMinimum *string `json:"dateMinimum,omitempty"`
		
		DateMaximum *string `json:"dateMaximum,omitempty"`
		
		AvailableTimes *[]Conversationcontentdatepickeravailabletime `json:"availableTimes,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Subtitle: o.Subtitle,
		
		ImageUrl: o.ImageUrl,
		
		DateMinimum: DateMinimum,
		
		DateMaximum: DateMaximum,
		
		AvailableTimes: o.AvailableTimes,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentdatepicker) UnmarshalJSON(b []byte) error {
	var ConversationcontentdatepickerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentdatepickerMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ConversationcontentdatepickerMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Subtitle, ok := ConversationcontentdatepickerMap["subtitle"].(string); ok {
		o.Subtitle = &Subtitle
	}
    
	if ImageUrl, ok := ConversationcontentdatepickerMap["imageUrl"].(string); ok {
		o.ImageUrl = &ImageUrl
	}
    
	if dateMinimumString, ok := ConversationcontentdatepickerMap["dateMinimum"].(string); ok {
		DateMinimum, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMinimumString)
		o.DateMinimum = &DateMinimum
	}
	
	if dateMaximumString, ok := ConversationcontentdatepickerMap["dateMaximum"].(string); ok {
		DateMaximum, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMaximumString)
		o.DateMaximum = &DateMaximum
	}
	
	if AvailableTimes, ok := ConversationcontentdatepickerMap["availableTimes"].([]interface{}); ok {
		AvailableTimesString, _ := json.Marshal(AvailableTimes)
		json.Unmarshal(AvailableTimesString, &o.AvailableTimes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentdatepicker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
