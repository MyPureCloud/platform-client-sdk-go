package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediaendpointstatistics
type Mediaendpointstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Trunk - Trunk information utilized when creating the media endpoint
	Trunk *Mediastatisticstrunkinfo `json:"trunk,omitempty"`

	// Station - Station information associated with media endpoint
	Station *Namedentity `json:"station,omitempty"`

	// User - User information associated media endpoint
	User *Namedentity `json:"user,omitempty"`

	// Ice - The ICE protocol statistics and details. Reference: https://www.rfc-editor.org/rfc/rfc5245
	Ice *Mediaicestatistics `json:"ice,omitempty"`

	// Rtp - Statistics of sent and received RTP. Reference: https://www.rfc-editor.org/rfc/rfc3550
	Rtp *Mediartpstatistics `json:"rtp,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mediaendpointstatistics) SetField(field string, fieldValue interface{}) {
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

func (o Mediaendpointstatistics) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Mediaendpointstatistics
	
	return json.Marshal(&struct { 
		Trunk *Mediastatisticstrunkinfo `json:"trunk,omitempty"`
		
		Station *Namedentity `json:"station,omitempty"`
		
		User *Namedentity `json:"user,omitempty"`
		
		Ice *Mediaicestatistics `json:"ice,omitempty"`
		
		Rtp *Mediartpstatistics `json:"rtp,omitempty"`
		Alias
	}{ 
		Trunk: o.Trunk,
		
		Station: o.Station,
		
		User: o.User,
		
		Ice: o.Ice,
		
		Rtp: o.Rtp,
		Alias:    (Alias)(o),
	})
}

func (o *Mediaendpointstatistics) UnmarshalJSON(b []byte) error {
	var MediaendpointstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &MediaendpointstatisticsMap)
	if err != nil {
		return err
	}
	
	if Trunk, ok := MediaendpointstatisticsMap["trunk"].(map[string]interface{}); ok {
		TrunkString, _ := json.Marshal(Trunk)
		json.Unmarshal(TrunkString, &o.Trunk)
	}
	
	if Station, ok := MediaendpointstatisticsMap["station"].(map[string]interface{}); ok {
		StationString, _ := json.Marshal(Station)
		json.Unmarshal(StationString, &o.Station)
	}
	
	if User, ok := MediaendpointstatisticsMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Ice, ok := MediaendpointstatisticsMap["ice"].(map[string]interface{}); ok {
		IceString, _ := json.Marshal(Ice)
		json.Unmarshal(IceString, &o.Ice)
	}
	
	if Rtp, ok := MediaendpointstatisticsMap["rtp"].(map[string]interface{}); ok {
		RtpString, _ := json.Marshal(Rtp)
		json.Unmarshal(RtpString, &o.Rtp)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediaendpointstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
