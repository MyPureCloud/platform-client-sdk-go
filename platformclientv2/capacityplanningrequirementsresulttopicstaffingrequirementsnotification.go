package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanningrequirementsresulttopicstaffingrequirementsnotification
type Capacityplanningrequirementsresulttopicstaffingrequirementsnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CapacityPlan
	CapacityPlan *Capacityplanningrequirementsresulttopiccapacityplanreference `json:"capacityPlan,omitempty"`

	// BusinessUnit
	BusinessUnit *Capacityplanningrequirementsresulttopicbusinessunit `json:"businessUnit,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// ReferenceBusinessUnitDate
	ReferenceBusinessUnitDate *time.Time `json:"referenceBusinessUnitDate,omitempty"`

	// Granularity
	Granularity *string `json:"granularity,omitempty"`

	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanningrequirementsresulttopicstaffingrequirementsnotification) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanningrequirementsresulttopicstaffingrequirementsnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReferenceBusinessUnitDate", }
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
	type Alias Capacityplanningrequirementsresulttopicstaffingrequirementsnotification
	
	ReferenceBusinessUnitDate := new(string)
	if o.ReferenceBusinessUnitDate != nil {
		
		*ReferenceBusinessUnitDate = timeutil.Strftime(o.ReferenceBusinessUnitDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceBusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		CapacityPlan *Capacityplanningrequirementsresulttopiccapacityplanreference `json:"capacityPlan,omitempty"`
		
		BusinessUnit *Capacityplanningrequirementsresulttopicbusinessunit `json:"businessUnit,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ReferenceBusinessUnitDate *string `json:"referenceBusinessUnitDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		Alias
	}{ 
		CapacityPlan: o.CapacityPlan,
		
		BusinessUnit: o.BusinessUnit,
		
		Status: o.Status,
		
		ReferenceBusinessUnitDate: ReferenceBusinessUnitDate,
		
		Granularity: o.Granularity,
		
		DownloadUrl: o.DownloadUrl,
		
		ErrorCode: o.ErrorCode,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanningrequirementsresulttopicstaffingrequirementsnotification) UnmarshalJSON(b []byte) error {
	var CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap)
	if err != nil {
		return err
	}
	
	if CapacityPlan, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["capacityPlan"].(map[string]interface{}); ok {
		CapacityPlanString, _ := json.Marshal(CapacityPlan)
		json.Unmarshal(CapacityPlanString, &o.CapacityPlan)
	}
	
	if BusinessUnit, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if Status, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if referenceBusinessUnitDateString, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["referenceBusinessUnitDate"].(string); ok {
		ReferenceBusinessUnitDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceBusinessUnitDateString)
		o.ReferenceBusinessUnitDate = &ReferenceBusinessUnitDate
	}
	
	if Granularity, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if DownloadUrl, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if ErrorCode, ok := CapacityplanningrequirementsresulttopicstaffingrequirementsnotificationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanningrequirementsresulttopicstaffingrequirementsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
