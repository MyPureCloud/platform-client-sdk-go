package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanstaffingrequirementresult
type Capacityplanstaffingrequirementresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BusinessUnit - The business unit to which the capacity plan belongs
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`

	// CapacityPlan - The capacity plan for which requirements are generated
	CapacityPlan *Capacityplanreference `json:"capacityPlan,omitempty"`

	// Status - The status of the requirement generation of the capacity plan
	Status *string `json:"status,omitempty"`

	// ReferenceBusinessUnitDate - The reference date for interval-based data for the requirements. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReferenceBusinessUnitDate *time.Time `json:"referenceBusinessUnitDate,omitempty"`

	// Granularity - Granularity of the intervals
	Granularity *string `json:"granularity,omitempty"`

	// ErrorCode - The error code when status is 'Failed'
	ErrorCode *string `json:"errorCode,omitempty"`

	// DownloadUrl - The URL to get the requirements results for the capacity plan. It will be populated if the status is 'Complete'
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// DownloadTemplate - Staffing requirement results always come through downloadUrl, the schema included here is just for documentation
	DownloadTemplate *Staffingrequirementresultresponsetemplate `json:"downloadTemplate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanstaffingrequirementresult) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanstaffingrequirementresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "ReferenceBusinessUnitDate", }

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
	type Alias Capacityplanstaffingrequirementresult
	
	ReferenceBusinessUnitDate := new(string)
	if o.ReferenceBusinessUnitDate != nil {
		*ReferenceBusinessUnitDate = timeutil.Strftime(o.ReferenceBusinessUnitDate, "%Y-%m-%d")
	} else {
		ReferenceBusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		CapacityPlan *Capacityplanreference `json:"capacityPlan,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ReferenceBusinessUnitDate *string `json:"referenceBusinessUnitDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadTemplate *Staffingrequirementresultresponsetemplate `json:"downloadTemplate,omitempty"`
		Alias
	}{ 
		BusinessUnit: o.BusinessUnit,
		
		CapacityPlan: o.CapacityPlan,
		
		Status: o.Status,
		
		ReferenceBusinessUnitDate: ReferenceBusinessUnitDate,
		
		Granularity: o.Granularity,
		
		ErrorCode: o.ErrorCode,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadTemplate: o.DownloadTemplate,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanstaffingrequirementresult) UnmarshalJSON(b []byte) error {
	var CapacityplanstaffingrequirementresultMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanstaffingrequirementresultMap)
	if err != nil {
		return err
	}
	
	if BusinessUnit, ok := CapacityplanstaffingrequirementresultMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if CapacityPlan, ok := CapacityplanstaffingrequirementresultMap["capacityPlan"].(map[string]interface{}); ok {
		CapacityPlanString, _ := json.Marshal(CapacityPlan)
		json.Unmarshal(CapacityPlanString, &o.CapacityPlan)
	}
	
	if Status, ok := CapacityplanstaffingrequirementresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if referenceBusinessUnitDateString, ok := CapacityplanstaffingrequirementresultMap["referenceBusinessUnitDate"].(string); ok {
		ReferenceBusinessUnitDate, _ := time.Parse("2006-01-02", referenceBusinessUnitDateString)
		o.ReferenceBusinessUnitDate = &ReferenceBusinessUnitDate
	}
	
	if Granularity, ok := CapacityplanstaffingrequirementresultMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if ErrorCode, ok := CapacityplanstaffingrequirementresultMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if DownloadUrl, ok := CapacityplanstaffingrequirementresultMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if DownloadTemplate, ok := CapacityplanstaffingrequirementresultMap["downloadTemplate"].(map[string]interface{}); ok {
		DownloadTemplateString, _ := json.Marshal(DownloadTemplate)
		json.Unmarshal(DownloadTemplateString, &o.DownloadTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanstaffingrequirementresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
