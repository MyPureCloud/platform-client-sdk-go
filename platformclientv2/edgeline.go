package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeline
type Edgeline struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// Description - The resource's description.
	Description *string `json:"description,omitempty"`

	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`

	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`

	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`

	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`

	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`

	// Schema
	Schema *Domainentityref `json:"schema,omitempty"`

	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`

	// Edge
	Edge *Edge `json:"edge,omitempty"`

	// EdgeGroup
	EdgeGroup *Edgegroup `json:"edgeGroup,omitempty"`

	// LineType
	LineType *string `json:"lineType,omitempty"`

	// Endpoint
	Endpoint *Endpoint `json:"endpoint,omitempty"`

	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`

	// LogicalInterfaceId
	LogicalInterfaceId *string `json:"logicalInterfaceId,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgeline) SetField(field string, fieldValue interface{}) {
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

func (o Edgeline) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Edgeline
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		Schema *Domainentityref `json:"schema,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		Edge *Edge `json:"edge,omitempty"`
		
		EdgeGroup *Edgegroup `json:"edgeGroup,omitempty"`
		
		LineType *string `json:"lineType,omitempty"`
		
		Endpoint *Endpoint `json:"endpoint,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		LogicalInterfaceId *string `json:"logicalInterfaceId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		Schema: o.Schema,
		
		Properties: o.Properties,
		
		Edge: o.Edge,
		
		EdgeGroup: o.EdgeGroup,
		
		LineType: o.LineType,
		
		Endpoint: o.Endpoint,
		
		IpAddress: o.IpAddress,
		
		LogicalInterfaceId: o.LogicalInterfaceId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Edgeline) UnmarshalJSON(b []byte) error {
	var EdgelineMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelineMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgelineMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EdgelineMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := EdgelineMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := EdgelineMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := EdgelineMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := EdgelineMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EdgelineMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := EdgelineMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := EdgelineMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := EdgelineMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := EdgelineMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := EdgelineMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if Schema, ok := EdgelineMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if Properties, ok := EdgelineMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Edge, ok := EdgelineMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if EdgeGroup, ok := EdgelineMap["edgeGroup"].(map[string]interface{}); ok {
		EdgeGroupString, _ := json.Marshal(EdgeGroup)
		json.Unmarshal(EdgeGroupString, &o.EdgeGroup)
	}
	
	if LineType, ok := EdgelineMap["lineType"].(string); ok {
		o.LineType = &LineType
	}
    
	if Endpoint, ok := EdgelineMap["endpoint"].(map[string]interface{}); ok {
		EndpointString, _ := json.Marshal(Endpoint)
		json.Unmarshal(EndpointString, &o.Endpoint)
	}
	
	if IpAddress, ok := EdgelineMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if LogicalInterfaceId, ok := EdgelineMap["logicalInterfaceId"].(string); ok {
		o.LogicalInterfaceId = &LogicalInterfaceId
	}
    
	if SelfUri, ok := EdgelineMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeline) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
