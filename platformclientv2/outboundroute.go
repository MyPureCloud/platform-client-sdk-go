package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundroute
type Outboundroute struct { 
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

	// ClassificationTypes - The site associated to the outbound route.
	ClassificationTypes *[]string `json:"classificationTypes,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Distribution
	Distribution *string `json:"distribution,omitempty"`

	// ExternalTrunkBases - Trunk base settings of trunkType \"EXTERNAL\".  This base must also be set on an edge logical interface for correct routing.
	ExternalTrunkBases *[]Domainentityref `json:"externalTrunkBases,omitempty"`

	// Site - The site associated to the outbound route.
	Site *Site `json:"site,omitempty"`

	// Managed - Is this outbound route being managed remotely.
	Managed *bool `json:"managed,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundroute) SetField(field string, fieldValue interface{}) {
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

func (o Outboundroute) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundroute
	
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
		
		ClassificationTypes *[]string `json:"classificationTypes,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Distribution *string `json:"distribution,omitempty"`
		
		ExternalTrunkBases *[]Domainentityref `json:"externalTrunkBases,omitempty"`
		
		Site *Site `json:"site,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
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
		
		ClassificationTypes: o.ClassificationTypes,
		
		Enabled: o.Enabled,
		
		Distribution: o.Distribution,
		
		ExternalTrunkBases: o.ExternalTrunkBases,
		
		Site: o.Site,
		
		Managed: o.Managed,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundroute) UnmarshalJSON(b []byte) error {
	var OutboundrouteMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundrouteMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutboundrouteMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OutboundrouteMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := OutboundrouteMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := OutboundrouteMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := OutboundrouteMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := OutboundrouteMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OutboundrouteMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := OutboundrouteMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := OutboundrouteMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := OutboundrouteMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := OutboundrouteMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := OutboundrouteMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if ClassificationTypes, ok := OutboundrouteMap["classificationTypes"].([]interface{}); ok {
		ClassificationTypesString, _ := json.Marshal(ClassificationTypes)
		json.Unmarshal(ClassificationTypesString, &o.ClassificationTypes)
	}
	
	if Enabled, ok := OutboundrouteMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Distribution, ok := OutboundrouteMap["distribution"].(string); ok {
		o.Distribution = &Distribution
	}
    
	if ExternalTrunkBases, ok := OutboundrouteMap["externalTrunkBases"].([]interface{}); ok {
		ExternalTrunkBasesString, _ := json.Marshal(ExternalTrunkBases)
		json.Unmarshal(ExternalTrunkBasesString, &o.ExternalTrunkBases)
	}
	
	if Site, ok := OutboundrouteMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if Managed, ok := OutboundrouteMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
    
	if SelfUri, ok := OutboundrouteMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
