package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Siteconnection
type Siteconnection struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// Managed
	Managed *bool `json:"managed,omitempty"`

	// VarType - Connection method from site to site (Direct, Indirect, CloudProxy
	VarType *string `json:"type,omitempty"`

	// Enabled - Indicates if the current site is linked
	Enabled *bool `json:"enabled,omitempty"`

	// MediaModel - Media model for the current site.
	MediaModel *string `json:"mediaModel,omitempty"`

	// EdgeList - All of the edges to which the site connects
	EdgeList *[]Connectededge `json:"edgeList,omitempty"`

	// CoreSite - The core site
	CoreSite *bool `json:"coreSite,omitempty"`

	// PrimaryCoreSites - List of site ids names and selfUris for the primary core sites
	PrimaryCoreSites *[]Domainentityref `json:"primaryCoreSites,omitempty"`

	// SecondaryCoreSites - List of site ids names and selfUris for the secondary core sites
	SecondaryCoreSites *[]Domainentityref `json:"secondaryCoreSites,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Siteconnection) SetField(field string, fieldValue interface{}) {
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

func (o Siteconnection) MarshalJSON() ([]byte, error) {
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
	type Alias Siteconnection
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MediaModel *string `json:"mediaModel,omitempty"`
		
		EdgeList *[]Connectededge `json:"edgeList,omitempty"`
		
		CoreSite *bool `json:"coreSite,omitempty"`
		
		PrimaryCoreSites *[]Domainentityref `json:"primaryCoreSites,omitempty"`
		
		SecondaryCoreSites *[]Domainentityref `json:"secondaryCoreSites,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		Managed: o.Managed,
		
		VarType: o.VarType,
		
		Enabled: o.Enabled,
		
		MediaModel: o.MediaModel,
		
		EdgeList: o.EdgeList,
		
		CoreSite: o.CoreSite,
		
		PrimaryCoreSites: o.PrimaryCoreSites,
		
		SecondaryCoreSites: o.SecondaryCoreSites,
		Alias:    (Alias)(o),
	})
}

func (o *Siteconnection) UnmarshalJSON(b []byte) error {
	var SiteconnectionMap map[string]interface{}
	err := json.Unmarshal(b, &SiteconnectionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SiteconnectionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SiteconnectionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := SiteconnectionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Managed, ok := SiteconnectionMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
    
	if VarType, ok := SiteconnectionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Enabled, ok := SiteconnectionMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MediaModel, ok := SiteconnectionMap["mediaModel"].(string); ok {
		o.MediaModel = &MediaModel
	}
    
	if EdgeList, ok := SiteconnectionMap["edgeList"].([]interface{}); ok {
		EdgeListString, _ := json.Marshal(EdgeList)
		json.Unmarshal(EdgeListString, &o.EdgeList)
	}
	
	if CoreSite, ok := SiteconnectionMap["coreSite"].(bool); ok {
		o.CoreSite = &CoreSite
	}
    
	if PrimaryCoreSites, ok := SiteconnectionMap["primaryCoreSites"].([]interface{}); ok {
		PrimaryCoreSitesString, _ := json.Marshal(PrimaryCoreSites)
		json.Unmarshal(PrimaryCoreSitesString, &o.PrimaryCoreSites)
	}
	
	if SecondaryCoreSites, ok := SiteconnectionMap["secondaryCoreSites"].([]interface{}); ok {
		SecondaryCoreSitesString, _ := json.Marshal(SecondaryCoreSites)
		json.Unmarshal(SecondaryCoreSitesString, &o.SecondaryCoreSites)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Siteconnection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
