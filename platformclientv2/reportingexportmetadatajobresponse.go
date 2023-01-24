package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingexportmetadatajobresponse
type Reportingexportmetadatajobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ViewType - The view type of the export metadata
	ViewType *string `json:"viewType,omitempty"`

	// DateLimitations - The date limitations of the export metadata
	DateLimitations *string `json:"dateLimitations,omitempty"`

	// RequiredFilters - The list of required filters for the export metadata
	RequiredFilters *[]string `json:"requiredFilters,omitempty"`

	// SupportedFilters - The list of supported filters for the export metadata
	SupportedFilters *[]string `json:"supportedFilters,omitempty"`

	// RequiredColumnIds - The list of required column ids for the export metadata
	RequiredColumnIds *[]string `json:"requiredColumnIds,omitempty"`

	// DependentColumnIds - The list of dependent column ids for the export metadata
	DependentColumnIds *map[string][]string `json:"dependentColumnIds,omitempty"`

	// AvailableColumnIds - The list of available column ids for the export metadata
	AvailableColumnIds *[]string `json:"availableColumnIds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingexportmetadatajobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Reportingexportmetadatajobresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Reportingexportmetadatajobresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ViewType *string `json:"viewType,omitempty"`
		
		DateLimitations *string `json:"dateLimitations,omitempty"`
		
		RequiredFilters *[]string `json:"requiredFilters,omitempty"`
		
		SupportedFilters *[]string `json:"supportedFilters,omitempty"`
		
		RequiredColumnIds *[]string `json:"requiredColumnIds,omitempty"`
		
		DependentColumnIds *map[string][]string `json:"dependentColumnIds,omitempty"`
		
		AvailableColumnIds *[]string `json:"availableColumnIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ViewType: o.ViewType,
		
		DateLimitations: o.DateLimitations,
		
		RequiredFilters: o.RequiredFilters,
		
		SupportedFilters: o.SupportedFilters,
		
		RequiredColumnIds: o.RequiredColumnIds,
		
		DependentColumnIds: o.DependentColumnIds,
		
		AvailableColumnIds: o.AvailableColumnIds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Reportingexportmetadatajobresponse) UnmarshalJSON(b []byte) error {
	var ReportingexportmetadatajobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingexportmetadatajobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportingexportmetadatajobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReportingexportmetadatajobresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ViewType, ok := ReportingexportmetadatajobresponseMap["viewType"].(string); ok {
		o.ViewType = &ViewType
	}
    
	if DateLimitations, ok := ReportingexportmetadatajobresponseMap["dateLimitations"].(string); ok {
		o.DateLimitations = &DateLimitations
	}
    
	if RequiredFilters, ok := ReportingexportmetadatajobresponseMap["requiredFilters"].([]interface{}); ok {
		RequiredFiltersString, _ := json.Marshal(RequiredFilters)
		json.Unmarshal(RequiredFiltersString, &o.RequiredFilters)
	}
	
	if SupportedFilters, ok := ReportingexportmetadatajobresponseMap["supportedFilters"].([]interface{}); ok {
		SupportedFiltersString, _ := json.Marshal(SupportedFilters)
		json.Unmarshal(SupportedFiltersString, &o.SupportedFilters)
	}
	
	if RequiredColumnIds, ok := ReportingexportmetadatajobresponseMap["requiredColumnIds"].([]interface{}); ok {
		RequiredColumnIdsString, _ := json.Marshal(RequiredColumnIds)
		json.Unmarshal(RequiredColumnIdsString, &o.RequiredColumnIds)
	}
	
	if DependentColumnIds, ok := ReportingexportmetadatajobresponseMap["dependentColumnIds"].(map[string]interface{}); ok {
		DependentColumnIdsString, _ := json.Marshal(DependentColumnIds)
		json.Unmarshal(DependentColumnIdsString, &o.DependentColumnIds)
	}
	
	if AvailableColumnIds, ok := ReportingexportmetadatajobresponseMap["availableColumnIds"].([]interface{}); ok {
		AvailableColumnIdsString, _ := json.Marshal(AvailableColumnIds)
		json.Unmarshal(AvailableColumnIdsString, &o.AvailableColumnIds)
	}
	
	if SelfUri, ok := ReportingexportmetadatajobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingexportmetadatajobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
