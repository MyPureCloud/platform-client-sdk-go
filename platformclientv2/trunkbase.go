package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkbase
type Trunkbase struct { 
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

	// TrunkMetabase - The meta-base this trunk is based on.
	TrunkMetabase *Domainentityref `json:"trunkMetabase,omitempty"`

	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`

	// TrunkType - The type of this trunk base.
	TrunkType *string `json:"trunkType,omitempty"`

	// Managed - Is this trunk being managed remotely. This property is synchronized with the managed property of the Edge Group to which it is assigned.
	Managed *bool `json:"managed,omitempty"`

	// Site - Used to determine the media regions for inbound and outbound calls through a trunk. Also determines the dial plan to use for calls that came in on a trunk and have to be sent out on it as well.
	Site *Domainentityref `json:"site,omitempty"`

	// InboundSite - Allows a customer to set the site to which inbound calls will be routed
	InboundSite *Domainentityref `json:"inboundSite,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trunkbase) SetField(field string, fieldValue interface{}) {
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

func (o Trunkbase) MarshalJSON() ([]byte, error) {
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
	type Alias Trunkbase
	
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
		
		TrunkMetabase *Domainentityref `json:"trunkMetabase,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		TrunkType *string `json:"trunkType,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
		Site *Domainentityref `json:"site,omitempty"`
		
		InboundSite *Domainentityref `json:"inboundSite,omitempty"`
		
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
		
		TrunkMetabase: o.TrunkMetabase,
		
		Properties: o.Properties,
		
		TrunkType: o.TrunkType,
		
		Managed: o.Managed,
		
		Site: o.Site,
		
		InboundSite: o.InboundSite,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Trunkbase) UnmarshalJSON(b []byte) error {
	var TrunkbaseMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkbaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrunkbaseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TrunkbaseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := TrunkbaseMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := TrunkbaseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := TrunkbaseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := TrunkbaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := TrunkbaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := TrunkbaseMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := TrunkbaseMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := TrunkbaseMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := TrunkbaseMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := TrunkbaseMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if TrunkMetabase, ok := TrunkbaseMap["trunkMetabase"].(map[string]interface{}); ok {
		TrunkMetabaseString, _ := json.Marshal(TrunkMetabase)
		json.Unmarshal(TrunkMetabaseString, &o.TrunkMetabase)
	}
	
	if Properties, ok := TrunkbaseMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if TrunkType, ok := TrunkbaseMap["trunkType"].(string); ok {
		o.TrunkType = &TrunkType
	}
    
	if Managed, ok := TrunkbaseMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
    
	if Site, ok := TrunkbaseMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if InboundSite, ok := TrunkbaseMap["inboundSite"].(map[string]interface{}); ok {
		InboundSiteString, _ := json.Marshal(InboundSite)
		json.Unmarshal(InboundSiteString, &o.InboundSite)
	}
	
	if SelfUri, ok := TrunkbaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkbase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
