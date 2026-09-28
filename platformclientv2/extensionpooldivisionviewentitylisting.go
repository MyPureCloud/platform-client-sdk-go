package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Extensionpooldivisionviewentitylisting
type Extensionpooldivisionviewentitylisting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Entities
	Entities *[]Extensionpooldivisionview `json:"entities,omitempty"`

	// PageSize
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`

	// Total
	Total *int `json:"total,omitempty"`

	// TotalNumberOfEntities - The total organization-wide number of entities.
	TotalNumberOfEntities *int `json:"totalNumberOfEntities,omitempty"`

	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// LastUri
	LastUri *string `json:"lastUri,omitempty"`

	// NextUri
	NextUri *string `json:"nextUri,omitempty"`

	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

	// PageCount
	PageCount *int `json:"pageCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Extensionpooldivisionviewentitylisting) SetField(field string, fieldValue interface{}) {
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

func (o Extensionpooldivisionviewentitylisting) MarshalJSON() ([]byte, error) {
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
	type Alias Extensionpooldivisionviewentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Extensionpooldivisionview `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		TotalNumberOfEntities *int `json:"totalNumberOfEntities,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		TotalNumberOfEntities: o.TotalNumberOfEntities,
		
		FirstUri: o.FirstUri,
		
		SelfUri: o.SelfUri,
		
		LastUri: o.LastUri,
		
		NextUri: o.NextUri,
		
		PreviousUri: o.PreviousUri,
		
		PageCount: o.PageCount,
		Alias:    (Alias)(o),
	})
}

func (o *Extensionpooldivisionviewentitylisting) UnmarshalJSON(b []byte) error {
	var ExtensionpooldivisionviewentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &ExtensionpooldivisionviewentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ExtensionpooldivisionviewentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := ExtensionpooldivisionviewentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := ExtensionpooldivisionviewentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := ExtensionpooldivisionviewentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if TotalNumberOfEntities, ok := ExtensionpooldivisionviewentitylistingMap["totalNumberOfEntities"].(float64); ok {
		TotalNumberOfEntitiesInt := int(TotalNumberOfEntities)
		o.TotalNumberOfEntities = &TotalNumberOfEntitiesInt
	}
	
	if FirstUri, ok := ExtensionpooldivisionviewentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if SelfUri, ok := ExtensionpooldivisionviewentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if LastUri, ok := ExtensionpooldivisionviewentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if NextUri, ok := ExtensionpooldivisionviewentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if PreviousUri, ok := ExtensionpooldivisionviewentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if PageCount, ok := ExtensionpooldivisionviewentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Extensionpooldivisionviewentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
