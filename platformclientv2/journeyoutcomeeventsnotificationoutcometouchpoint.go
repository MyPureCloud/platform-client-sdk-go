package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcometouchpoint
type Journeyoutcomeeventsnotificationoutcometouchpoint struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Channels
	Channels *[]Journeyoutcomeeventsnotificationoutcometouchpointchannel `json:"channels,omitempty"`

	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ActionMap
	ActionMap *Journeyoutcomeeventsnotificationactionmap `json:"actionMap,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyoutcomeeventsnotificationoutcometouchpoint) SetField(field string, fieldValue interface{}) {
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

func (o Journeyoutcomeeventsnotificationoutcometouchpoint) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyoutcomeeventsnotificationoutcometouchpoint
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channels *[]Journeyoutcomeeventsnotificationoutcometouchpointchannel `json:"channels,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ActionMap *Journeyoutcomeeventsnotificationactionmap `json:"actionMap,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Channels: o.Channels,
		
		CreatedDate: CreatedDate,
		
		ActionMap: o.ActionMap,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcometouchpoint) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcometouchpointMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcometouchpointMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Channels, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["channels"].([]interface{}); ok {
		ChannelsString, _ := json.Marshal(Channels)
		json.Unmarshal(ChannelsString, &o.Channels)
	}
	
	if createdDateString, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if ActionMap, ok := JourneyoutcomeeventsnotificationoutcometouchpointMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcometouchpoint) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
