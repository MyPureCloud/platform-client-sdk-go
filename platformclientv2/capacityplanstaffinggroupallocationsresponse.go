package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanstaffinggroupallocationsresponse
type Capacityplanstaffinggroupallocationsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CapacityPlan - The capacity plan to which the staffing groups belong
	CapacityPlan *Capacityplanreference `json:"capacityPlan,omitempty"`

	// FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
	FullTimeEquivalentWeeklyHours *float64 `json:"fullTimeEquivalentWeeklyHours,omitempty"`

	// DownloadUrl - The URL to download the staffing group allocations
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// DownloadTemplate - Staffing groups allocation results always come through downloadUrl, the schema included here is just for documentation
	DownloadTemplate *Staffinggroupallocationsresponsetemplate `json:"downloadTemplate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanstaffinggroupallocationsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanstaffinggroupallocationsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Capacityplanstaffinggroupallocationsresponse
	
	return json.Marshal(&struct { 
		CapacityPlan *Capacityplanreference `json:"capacityPlan,omitempty"`
		
		FullTimeEquivalentWeeklyHours *float64 `json:"fullTimeEquivalentWeeklyHours,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadTemplate *Staffinggroupallocationsresponsetemplate `json:"downloadTemplate,omitempty"`
		Alias
	}{ 
		CapacityPlan: o.CapacityPlan,
		
		FullTimeEquivalentWeeklyHours: o.FullTimeEquivalentWeeklyHours,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadTemplate: o.DownloadTemplate,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanstaffinggroupallocationsresponse) UnmarshalJSON(b []byte) error {
	var CapacityplanstaffinggroupallocationsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanstaffinggroupallocationsresponseMap)
	if err != nil {
		return err
	}
	
	if CapacityPlan, ok := CapacityplanstaffinggroupallocationsresponseMap["capacityPlan"].(map[string]interface{}); ok {
		CapacityPlanString, _ := json.Marshal(CapacityPlan)
		json.Unmarshal(CapacityPlanString, &o.CapacityPlan)
	}
	
	if FullTimeEquivalentWeeklyHours, ok := CapacityplanstaffinggroupallocationsresponseMap["fullTimeEquivalentWeeklyHours"].(float64); ok {
		o.FullTimeEquivalentWeeklyHours = &FullTimeEquivalentWeeklyHours
	}
    
	if DownloadUrl, ok := CapacityplanstaffinggroupallocationsresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if DownloadTemplate, ok := CapacityplanstaffinggroupallocationsresponseMap["downloadTemplate"].(map[string]interface{}); ok {
		DownloadTemplateString, _ := json.Marshal(DownloadTemplate)
		json.Unmarshal(DownloadTemplateString, &o.DownloadTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanstaffinggroupallocationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
