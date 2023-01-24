package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetricnetworks
type Edgemetricstopicedgemetricnetworks struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Ifname
	Ifname *string `json:"ifname,omitempty"`

	// SentBytesPerSec
	SentBytesPerSec *int `json:"sentBytesPerSec,omitempty"`

	// ReceivedBytesPerSec
	ReceivedBytesPerSec *int `json:"receivedBytesPerSec,omitempty"`

	// BandwidthBitsPerSec
	BandwidthBitsPerSec *int `json:"bandwidthBitsPerSec,omitempty"`

	// UtilizationPct
	UtilizationPct *float32 `json:"utilizationPct,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgemetricstopicedgemetricnetworks) SetField(field string, fieldValue interface{}) {
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

func (o Edgemetricstopicedgemetricnetworks) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Edgemetricstopicedgemetricnetworks
	
	return json.Marshal(&struct { 
		Ifname *string `json:"ifname,omitempty"`
		
		SentBytesPerSec *int `json:"sentBytesPerSec,omitempty"`
		
		ReceivedBytesPerSec *int `json:"receivedBytesPerSec,omitempty"`
		
		BandwidthBitsPerSec *int `json:"bandwidthBitsPerSec,omitempty"`
		
		UtilizationPct *float32 `json:"utilizationPct,omitempty"`
		Alias
	}{ 
		Ifname: o.Ifname,
		
		SentBytesPerSec: o.SentBytesPerSec,
		
		ReceivedBytesPerSec: o.ReceivedBytesPerSec,
		
		BandwidthBitsPerSec: o.BandwidthBitsPerSec,
		
		UtilizationPct: o.UtilizationPct,
		Alias:    (Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetricnetworks) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricnetworksMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricnetworksMap)
	if err != nil {
		return err
	}
	
	if Ifname, ok := EdgemetricstopicedgemetricnetworksMap["ifname"].(string); ok {
		o.Ifname = &Ifname
	}
    
	if SentBytesPerSec, ok := EdgemetricstopicedgemetricnetworksMap["sentBytesPerSec"].(float64); ok {
		SentBytesPerSecInt := int(SentBytesPerSec)
		o.SentBytesPerSec = &SentBytesPerSecInt
	}
	
	if ReceivedBytesPerSec, ok := EdgemetricstopicedgemetricnetworksMap["receivedBytesPerSec"].(float64); ok {
		ReceivedBytesPerSecInt := int(ReceivedBytesPerSec)
		o.ReceivedBytesPerSec = &ReceivedBytesPerSecInt
	}
	
	if BandwidthBitsPerSec, ok := EdgemetricstopicedgemetricnetworksMap["bandwidthBitsPerSec"].(float64); ok {
		BandwidthBitsPerSecInt := int(BandwidthBitsPerSec)
		o.BandwidthBitsPerSec = &BandwidthBitsPerSecInt
	}
	
	if UtilizationPct, ok := EdgemetricstopicedgemetricnetworksMap["utilizationPct"].(float64); ok {
		UtilizationPctFloat32 := float32(UtilizationPct)
		o.UtilizationPct = &UtilizationPctFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricnetworks) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
