package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dashboardconfigurationqueryrequest
type Dashboardconfigurationqueryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DashboardConfigurationIds - The user supplied dashboard configuration ids
	DashboardConfigurationIds *[]string `json:"dashboardConfigurationIds,omitempty"`

	// PageNumber - The page number of the queried response
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize - The number of entities to return of the queried response. The max is 25
	PageSize *int `json:"pageSize,omitempty"`

	// SortBy - The order in which response will be sorted
	SortBy *string `json:"sortBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dashboardconfigurationqueryrequest) SetField(field string, fieldValue interface{}) {
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

func (o Dashboardconfigurationqueryrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Dashboardconfigurationqueryrequest
	
	return json.Marshal(&struct { 
		DashboardConfigurationIds *[]string `json:"dashboardConfigurationIds,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		Alias
	}{ 
		DashboardConfigurationIds: o.DashboardConfigurationIds,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		SortBy: o.SortBy,
		Alias:    (Alias)(o),
	})
}

func (o *Dashboardconfigurationqueryrequest) UnmarshalJSON(b []byte) error {
	var DashboardconfigurationqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DashboardconfigurationqueryrequestMap)
	if err != nil {
		return err
	}
	
	if DashboardConfigurationIds, ok := DashboardconfigurationqueryrequestMap["dashboardConfigurationIds"].([]interface{}); ok {
		DashboardConfigurationIdsString, _ := json.Marshal(DashboardConfigurationIds)
		json.Unmarshal(DashboardConfigurationIdsString, &o.DashboardConfigurationIds)
	}
	
	if PageNumber, ok := DashboardconfigurationqueryrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := DashboardconfigurationqueryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if SortBy, ok := DashboardconfigurationqueryrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dashboardconfigurationqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
