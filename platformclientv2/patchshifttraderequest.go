package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchshifttraderequest
type Patchshifttraderequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ReceivingUserId - Update the ID of the receiving user to direct the request at a specific user, or set the wrapped id to null to open up a trade to be matched by any user.
	ReceivingUserId *Valuewrapperstring `json:"receivingUserId,omitempty"`

	// Expiration - Update the expiration time for this shift trade.
	Expiration *Valuewrapperdate `json:"expiration,omitempty"`

	// AcceptableIntervals - Update the acceptable intervals the initiating user is willing to accept in trade. Setting the enclosed list to empty will make this a one sided trade request
	AcceptableIntervals *Listwrapperinterval `json:"acceptableIntervals,omitempty"`

	// Metadata - Version metadata
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Patchshifttraderequest) SetField(field string, fieldValue interface{}) {
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

func (o Patchshifttraderequest) MarshalJSON() ([]byte, error) {
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
	type Alias Patchshifttraderequest
	
	return json.Marshal(&struct { 
		ReceivingUserId *Valuewrapperstring `json:"receivingUserId,omitempty"`
		
		Expiration *Valuewrapperdate `json:"expiration,omitempty"`
		
		AcceptableIntervals *Listwrapperinterval `json:"acceptableIntervals,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		ReceivingUserId: o.ReceivingUserId,
		
		Expiration: o.Expiration,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Patchshifttraderequest) UnmarshalJSON(b []byte) error {
	var PatchshifttraderequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchshifttraderequestMap)
	if err != nil {
		return err
	}
	
	if ReceivingUserId, ok := PatchshifttraderequestMap["receivingUserId"].(map[string]interface{}); ok {
		ReceivingUserIdString, _ := json.Marshal(ReceivingUserId)
		json.Unmarshal(ReceivingUserIdString, &o.ReceivingUserId)
	}
	
	if Expiration, ok := PatchshifttraderequestMap["expiration"].(map[string]interface{}); ok {
		ExpirationString, _ := json.Marshal(Expiration)
		json.Unmarshal(ExpirationString, &o.Expiration)
	}
	
	if AcceptableIntervals, ok := PatchshifttraderequestMap["acceptableIntervals"].(map[string]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if Metadata, ok := PatchshifttraderequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
