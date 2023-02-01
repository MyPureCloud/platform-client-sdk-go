package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonecapabilities
type Phonecapabilities struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Provisions
	Provisions *bool `json:"provisions,omitempty"`

	// Registers
	Registers *bool `json:"registers,omitempty"`

	// DualRegisters
	DualRegisters *bool `json:"dualRegisters,omitempty"`

	// HardwareIdType
	HardwareIdType *string `json:"hardwareIdType,omitempty"`

	// AllowReboot
	AllowReboot *bool `json:"allowReboot,omitempty"`

	// NoRebalance
	NoRebalance *bool `json:"noRebalance,omitempty"`

	// NoCloudProvisioning
	NoCloudProvisioning *bool `json:"noCloudProvisioning,omitempty"`

	// MediaCodecs
	MediaCodecs *[]string `json:"mediaCodecs,omitempty"`

	// Cdm
	Cdm *bool `json:"cdm,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Phonecapabilities) SetField(field string, fieldValue interface{}) {
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

func (o Phonecapabilities) MarshalJSON() ([]byte, error) {
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
	type Alias Phonecapabilities
	
	return json.Marshal(&struct { 
		Provisions *bool `json:"provisions,omitempty"`
		
		Registers *bool `json:"registers,omitempty"`
		
		DualRegisters *bool `json:"dualRegisters,omitempty"`
		
		HardwareIdType *string `json:"hardwareIdType,omitempty"`
		
		AllowReboot *bool `json:"allowReboot,omitempty"`
		
		NoRebalance *bool `json:"noRebalance,omitempty"`
		
		NoCloudProvisioning *bool `json:"noCloudProvisioning,omitempty"`
		
		MediaCodecs *[]string `json:"mediaCodecs,omitempty"`
		
		Cdm *bool `json:"cdm,omitempty"`
		Alias
	}{ 
		Provisions: o.Provisions,
		
		Registers: o.Registers,
		
		DualRegisters: o.DualRegisters,
		
		HardwareIdType: o.HardwareIdType,
		
		AllowReboot: o.AllowReboot,
		
		NoRebalance: o.NoRebalance,
		
		NoCloudProvisioning: o.NoCloudProvisioning,
		
		MediaCodecs: o.MediaCodecs,
		
		Cdm: o.Cdm,
		Alias:    (Alias)(o),
	})
}

func (o *Phonecapabilities) UnmarshalJSON(b []byte) error {
	var PhonecapabilitiesMap map[string]interface{}
	err := json.Unmarshal(b, &PhonecapabilitiesMap)
	if err != nil {
		return err
	}
	
	if Provisions, ok := PhonecapabilitiesMap["provisions"].(bool); ok {
		o.Provisions = &Provisions
	}
    
	if Registers, ok := PhonecapabilitiesMap["registers"].(bool); ok {
		o.Registers = &Registers
	}
    
	if DualRegisters, ok := PhonecapabilitiesMap["dualRegisters"].(bool); ok {
		o.DualRegisters = &DualRegisters
	}
    
	if HardwareIdType, ok := PhonecapabilitiesMap["hardwareIdType"].(string); ok {
		o.HardwareIdType = &HardwareIdType
	}
    
	if AllowReboot, ok := PhonecapabilitiesMap["allowReboot"].(bool); ok {
		o.AllowReboot = &AllowReboot
	}
    
	if NoRebalance, ok := PhonecapabilitiesMap["noRebalance"].(bool); ok {
		o.NoRebalance = &NoRebalance
	}
    
	if NoCloudProvisioning, ok := PhonecapabilitiesMap["noCloudProvisioning"].(bool); ok {
		o.NoCloudProvisioning = &NoCloudProvisioning
	}
    
	if MediaCodecs, ok := PhonecapabilitiesMap["mediaCodecs"].([]interface{}); ok {
		MediaCodecsString, _ := json.Marshal(MediaCodecs)
		json.Unmarshal(MediaCodecsString, &o.MediaCodecs)
	}
	
	if Cdm, ok := PhonecapabilitiesMap["cdm"].(bool); ok {
		o.Cdm = &Cdm
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phonecapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
