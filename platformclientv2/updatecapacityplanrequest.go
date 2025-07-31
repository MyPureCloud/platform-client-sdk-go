package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatecapacityplanrequest
type Updatecapacityplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the capacity plan
	Name *string `json:"name,omitempty"`

	// Description - Description of the capacity plan
	Description *string `json:"description,omitempty"`

	// StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartBusinessUnitDate *time.Time `json:"startBusinessUnitDate,omitempty"`

	// EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndBusinessUnitDate *time.Time `json:"endBusinessUnitDate,omitempty"`

	// Forecast - The selected forecast for this capacity plan
	Forecast *Valuewrapperbushorttermforecastreference `json:"forecast,omitempty"`

	// FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
	FullTimeEquivalentWeeklyHours *float64 `json:"fullTimeEquivalentWeeklyHours,omitempty"`

	// UseLatestPlanningGroupStaffingGroupAssociation - Whether to use latest staffing group to planning group association
	UseLatestPlanningGroupStaffingGroupAssociation *bool `json:"useLatestPlanningGroupStaffingGroupAssociation,omitempty"`

	// Metadata - The metadata of this capacity plan
	Metadata *Capacityplanmetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updatecapacityplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Updatecapacityplanrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "StartBusinessUnitDate","EndBusinessUnitDate", }

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
	type Alias Updatecapacityplanrequest
	
	StartBusinessUnitDate := new(string)
	if o.StartBusinessUnitDate != nil {
		*StartBusinessUnitDate = timeutil.Strftime(o.StartBusinessUnitDate, "%Y-%m-%d")
	} else {
		StartBusinessUnitDate = nil
	}
	
	EndBusinessUnitDate := new(string)
	if o.EndBusinessUnitDate != nil {
		*EndBusinessUnitDate = timeutil.Strftime(o.EndBusinessUnitDate, "%Y-%m-%d")
	} else {
		EndBusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		StartBusinessUnitDate *string `json:"startBusinessUnitDate,omitempty"`
		
		EndBusinessUnitDate *string `json:"endBusinessUnitDate,omitempty"`
		
		Forecast *Valuewrapperbushorttermforecastreference `json:"forecast,omitempty"`
		
		FullTimeEquivalentWeeklyHours *float64 `json:"fullTimeEquivalentWeeklyHours,omitempty"`
		
		UseLatestPlanningGroupStaffingGroupAssociation *bool `json:"useLatestPlanningGroupStaffingGroupAssociation,omitempty"`
		
		Metadata *Capacityplanmetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		StartBusinessUnitDate: StartBusinessUnitDate,
		
		EndBusinessUnitDate: EndBusinessUnitDate,
		
		Forecast: o.Forecast,
		
		FullTimeEquivalentWeeklyHours: o.FullTimeEquivalentWeeklyHours,
		
		UseLatestPlanningGroupStaffingGroupAssociation: o.UseLatestPlanningGroupStaffingGroupAssociation,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Updatecapacityplanrequest) UnmarshalJSON(b []byte) error {
	var UpdatecapacityplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatecapacityplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdatecapacityplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := UpdatecapacityplanrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if startBusinessUnitDateString, ok := UpdatecapacityplanrequestMap["startBusinessUnitDate"].(string); ok {
		StartBusinessUnitDate, _ := time.Parse("2006-01-02", startBusinessUnitDateString)
		o.StartBusinessUnitDate = &StartBusinessUnitDate
	}
	
	if endBusinessUnitDateString, ok := UpdatecapacityplanrequestMap["endBusinessUnitDate"].(string); ok {
		EndBusinessUnitDate, _ := time.Parse("2006-01-02", endBusinessUnitDateString)
		o.EndBusinessUnitDate = &EndBusinessUnitDate
	}
	
	if Forecast, ok := UpdatecapacityplanrequestMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if FullTimeEquivalentWeeklyHours, ok := UpdatecapacityplanrequestMap["fullTimeEquivalentWeeklyHours"].(float64); ok {
		o.FullTimeEquivalentWeeklyHours = &FullTimeEquivalentWeeklyHours
	}
    
	if UseLatestPlanningGroupStaffingGroupAssociation, ok := UpdatecapacityplanrequestMap["useLatestPlanningGroupStaffingGroupAssociation"].(bool); ok {
		o.UseLatestPlanningGroupStaffingGroupAssociation = &UseLatestPlanningGroupStaffingGroupAssociation
	}
    
	if Metadata, ok := UpdatecapacityplanrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatecapacityplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
