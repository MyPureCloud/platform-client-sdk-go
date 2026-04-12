package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmshifttradingjobcompleteeventtradeentity
type Wfmshifttradingjobcompleteeventtradeentity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// ExpirationDate
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// AcceptableIntervals
	AcceptableIntervals *[]Wfmshifttradingjobcompleteeventwfmdatetimeinterval `json:"acceptableIntervals,omitempty"`

	// OneSided
	OneSided *bool `json:"oneSided,omitempty"`

	// Initiating
	Initiating *Wfmshifttradingjobcompleteeventtradeside `json:"initiating,omitempty"`

	// Receiving
	Receiving *Wfmshifttradingjobcompleteeventtradeside `json:"receiving,omitempty"`

	// TargetUserId
	TargetUserId *string `json:"targetUserId,omitempty"`

	// ReviewedBy
	ReviewedBy *string `json:"reviewedBy,omitempty"`

	// ReviewedDate
	ReviewedDate *string `json:"reviewedDate,omitempty"`

	// Metadata
	Metadata *Wfmshifttradingjobcompleteeventwfmversionmetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmshifttradingjobcompleteeventtradeentity) SetField(field string, fieldValue interface{}) {
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

func (o Wfmshifttradingjobcompleteeventtradeentity) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmshifttradingjobcompleteeventtradeentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		
		AcceptableIntervals *[]Wfmshifttradingjobcompleteeventwfmdatetimeinterval `json:"acceptableIntervals,omitempty"`
		
		OneSided *bool `json:"oneSided,omitempty"`
		
		Initiating *Wfmshifttradingjobcompleteeventtradeside `json:"initiating,omitempty"`
		
		Receiving *Wfmshifttradingjobcompleteeventtradeside `json:"receiving,omitempty"`
		
		TargetUserId *string `json:"targetUserId,omitempty"`
		
		ReviewedBy *string `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		Metadata *Wfmshifttradingjobcompleteeventwfmversionmetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		ExpirationDate: o.ExpirationDate,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		OneSided: o.OneSided,
		
		Initiating: o.Initiating,
		
		Receiving: o.Receiving,
		
		TargetUserId: o.TargetUserId,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: o.ReviewedDate,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmshifttradingjobcompleteeventtradeentity) UnmarshalJSON(b []byte) error {
	var WfmshifttradingjobcompleteeventtradeentityMap map[string]interface{}
	err := json.Unmarshal(b, &WfmshifttradingjobcompleteeventtradeentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmshifttradingjobcompleteeventtradeentityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := WfmshifttradingjobcompleteeventtradeentityMap["state"].(string); ok {
		o.State = &State
	}
    
	if ExpirationDate, ok := WfmshifttradingjobcompleteeventtradeentityMap["expirationDate"].(string); ok {
		o.ExpirationDate = &ExpirationDate
	}
    
	if AcceptableIntervals, ok := WfmshifttradingjobcompleteeventtradeentityMap["acceptableIntervals"].([]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if OneSided, ok := WfmshifttradingjobcompleteeventtradeentityMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
    
	if Initiating, ok := WfmshifttradingjobcompleteeventtradeentityMap["initiating"].(map[string]interface{}); ok {
		InitiatingString, _ := json.Marshal(Initiating)
		json.Unmarshal(InitiatingString, &o.Initiating)
	}
	
	if Receiving, ok := WfmshifttradingjobcompleteeventtradeentityMap["receiving"].(map[string]interface{}); ok {
		ReceivingString, _ := json.Marshal(Receiving)
		json.Unmarshal(ReceivingString, &o.Receiving)
	}
	
	if TargetUserId, ok := WfmshifttradingjobcompleteeventtradeentityMap["targetUserId"].(string); ok {
		o.TargetUserId = &TargetUserId
	}
    
	if ReviewedBy, ok := WfmshifttradingjobcompleteeventtradeentityMap["reviewedBy"].(string); ok {
		o.ReviewedBy = &ReviewedBy
	}
    
	if ReviewedDate, ok := WfmshifttradingjobcompleteeventtradeentityMap["reviewedDate"].(string); ok {
		o.ReviewedDate = &ReviewedDate
	}
    
	if Metadata, ok := WfmshifttradingjobcompleteeventtradeentityMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmshifttradingjobcompleteeventtradeentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
