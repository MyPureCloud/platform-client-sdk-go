package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediartpstatistics
type Mediartpstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PacketsReceived - Number of packets received, including all invalid, duplicate, and discarded packets
	PacketsReceived *int `json:"packetsReceived,omitempty"`

	// PacketsSent - Number of packets sent
	PacketsSent *int `json:"packetsSent,omitempty"`

	// RtpEventsReceived - Number of RFC#2833 packets received
	RtpEventsReceived *int `json:"rtpEventsReceived,omitempty"`

	// RtpEventsSent - Number of RFC#2833 packets sent
	RtpEventsSent *int `json:"rtpEventsSent,omitempty"`

	// EstimatedAverageMos - The estimated average MOS score
	EstimatedAverageMos *float64 `json:"estimatedAverageMos,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mediartpstatistics) SetField(field string, fieldValue interface{}) {
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

func (o Mediartpstatistics) MarshalJSON() ([]byte, error) {
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
	type Alias Mediartpstatistics
	
	return json.Marshal(&struct { 
		PacketsReceived *int `json:"packetsReceived,omitempty"`
		
		PacketsSent *int `json:"packetsSent,omitempty"`
		
		RtpEventsReceived *int `json:"rtpEventsReceived,omitempty"`
		
		RtpEventsSent *int `json:"rtpEventsSent,omitempty"`
		
		EstimatedAverageMos *float64 `json:"estimatedAverageMos,omitempty"`
		Alias
	}{ 
		PacketsReceived: o.PacketsReceived,
		
		PacketsSent: o.PacketsSent,
		
		RtpEventsReceived: o.RtpEventsReceived,
		
		RtpEventsSent: o.RtpEventsSent,
		
		EstimatedAverageMos: o.EstimatedAverageMos,
		Alias:    (Alias)(o),
	})
}

func (o *Mediartpstatistics) UnmarshalJSON(b []byte) error {
	var MediartpstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &MediartpstatisticsMap)
	if err != nil {
		return err
	}
	
	if PacketsReceived, ok := MediartpstatisticsMap["packetsReceived"].(float64); ok {
		PacketsReceivedInt := int(PacketsReceived)
		o.PacketsReceived = &PacketsReceivedInt
	}
	
	if PacketsSent, ok := MediartpstatisticsMap["packetsSent"].(float64); ok {
		PacketsSentInt := int(PacketsSent)
		o.PacketsSent = &PacketsSentInt
	}
	
	if RtpEventsReceived, ok := MediartpstatisticsMap["rtpEventsReceived"].(float64); ok {
		RtpEventsReceivedInt := int(RtpEventsReceived)
		o.RtpEventsReceived = &RtpEventsReceivedInt
	}
	
	if RtpEventsSent, ok := MediartpstatisticsMap["rtpEventsSent"].(float64); ok {
		RtpEventsSentInt := int(RtpEventsSent)
		o.RtpEventsSent = &RtpEventsSentInt
	}
	
	if EstimatedAverageMos, ok := MediartpstatisticsMap["estimatedAverageMos"].(float64); ok {
		o.EstimatedAverageMos = &EstimatedAverageMos
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mediartpstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
