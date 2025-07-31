package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanresponse
type Capacityplanresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description - Description of the capacity plan
	Description *string `json:"description,omitempty"`

	// Forecast - The selected forecast for this capacity plan. Null when main forecast is used in the future
	Forecast *Bushorttermforecastreference `json:"forecast,omitempty"`

	// StartBusinessUnitDate - The start date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartBusinessUnitDate *time.Time `json:"startBusinessUnitDate,omitempty"`

	// EndBusinessUnitDate - The end date for the capacity plan relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndBusinessUnitDate *time.Time `json:"endBusinessUnitDate,omitempty"`

	// FullTimeEquivalentWeeklyHours - The weekly hours used to calculate full time equivalent agents
	FullTimeEquivalentWeeklyHours *float64 `json:"fullTimeEquivalentWeeklyHours,omitempty"`

	// Metadata - The metadata of this capacity plan
	Metadata *Capacityplanmetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanresponse) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Capacityplanresponse
	
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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Forecast *Bushorttermforecastreference `json:"forecast,omitempty"`
		
		StartBusinessUnitDate *string `json:"startBusinessUnitDate,omitempty"`
		
		EndBusinessUnitDate *string `json:"endBusinessUnitDate,omitempty"`
		
		FullTimeEquivalentWeeklyHours *float64 `json:"fullTimeEquivalentWeeklyHours,omitempty"`
		
		Metadata *Capacityplanmetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Forecast: o.Forecast,
		
		StartBusinessUnitDate: StartBusinessUnitDate,
		
		EndBusinessUnitDate: EndBusinessUnitDate,
		
		FullTimeEquivalentWeeklyHours: o.FullTimeEquivalentWeeklyHours,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanresponse) UnmarshalJSON(b []byte) error {
	var CapacityplanresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CapacityplanresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CapacityplanresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CapacityplanresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Forecast, ok := CapacityplanresponseMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if startBusinessUnitDateString, ok := CapacityplanresponseMap["startBusinessUnitDate"].(string); ok {
		StartBusinessUnitDate, _ := time.Parse("2006-01-02", startBusinessUnitDateString)
		o.StartBusinessUnitDate = &StartBusinessUnitDate
	}
	
	if endBusinessUnitDateString, ok := CapacityplanresponseMap["endBusinessUnitDate"].(string); ok {
		EndBusinessUnitDate, _ := time.Parse("2006-01-02", endBusinessUnitDateString)
		o.EndBusinessUnitDate = &EndBusinessUnitDate
	}
	
	if FullTimeEquivalentWeeklyHours, ok := CapacityplanresponseMap["fullTimeEquivalentWeeklyHours"].(float64); ok {
		o.FullTimeEquivalentWeeklyHours = &FullTimeEquivalentWeeklyHours
	}
    
	if Metadata, ok := CapacityplanresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := CapacityplanresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
