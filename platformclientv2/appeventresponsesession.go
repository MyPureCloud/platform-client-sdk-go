package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Appeventresponsesession
type Appeventresponsesession struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DurationInSeconds - Indicates how long the customer has been in the app within this session.
	DurationInSeconds *int `json:"durationInSeconds,omitempty"`

	// EventCount - The count of all events recorded during this session.
	EventCount *int `json:"eventCount,omitempty"`

	// ScreenviewCount - The count of all screen views recorded during this session.
	ScreenviewCount *int `json:"screenviewCount,omitempty"`

	// Referrer - The referrer of the first event in the app session.
	Referrer *Referrer `json:"referrer,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// CreatedDate - UTC timestamp of the session's first event, that is when the session starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Appeventresponsesession) SetField(field string, fieldValue interface{}) {
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

func (o Appeventresponsesession) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Appeventresponsesession
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DurationInSeconds *int `json:"durationInSeconds,omitempty"`
		
		EventCount *int `json:"eventCount,omitempty"`
		
		ScreenviewCount *int `json:"screenviewCount,omitempty"`
		
		Referrer *Referrer `json:"referrer,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DurationInSeconds: o.DurationInSeconds,
		
		EventCount: o.EventCount,
		
		ScreenviewCount: o.ScreenviewCount,
		
		Referrer: o.Referrer,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Appeventresponsesession) UnmarshalJSON(b []byte) error {
	var AppeventresponsesessionMap map[string]interface{}
	err := json.Unmarshal(b, &AppeventresponsesessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AppeventresponsesessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DurationInSeconds, ok := AppeventresponsesessionMap["durationInSeconds"].(float64); ok {
		DurationInSecondsInt := int(DurationInSeconds)
		o.DurationInSeconds = &DurationInSecondsInt
	}
	
	if EventCount, ok := AppeventresponsesessionMap["eventCount"].(float64); ok {
		EventCountInt := int(EventCount)
		o.EventCount = &EventCountInt
	}
	
	if ScreenviewCount, ok := AppeventresponsesessionMap["screenviewCount"].(float64); ok {
		ScreenviewCountInt := int(ScreenviewCount)
		o.ScreenviewCount = &ScreenviewCountInt
	}
	
	if Referrer, ok := AppeventresponsesessionMap["referrer"].(map[string]interface{}); ok {
		ReferrerString, _ := json.Marshal(Referrer)
		json.Unmarshal(ReferrerString, &o.Referrer)
	}
	
	if SelfUri, ok := AppeventresponsesessionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := AppeventresponsesessionMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Appeventresponsesession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
