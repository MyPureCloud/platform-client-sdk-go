package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunit
type Managementunit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// BusinessUnit - The business unit to which this management unit belongs
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`


	// StartDayOfWeek - Start day of week for scheduling and forecasting purposes. Moving to Business Unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// TimeZone - The time zone for the management unit in standard Olson format.  Moving to Business Unit
	TimeZone *string `json:"timeZone,omitempty"`


	// Settings - The configuration settings for this management unit
	Settings *Managementunitsettingsresponse `json:"settings,omitempty"`


	// Metadata - Version info metadata for this management unit. Deprecated, use settings.metadata
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Divisionreference `json:"division,omitempty"`


	// Version - The version of the underlying entity.  Deprecated, use field from settings.metadata instead
	Version *int `json:"version,omitempty"`


	// DateModified - The date and time at which this entity was last modified.  Deprecated, use field from settings.metadata instead. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The user who last modified this entity.  Deprecated, use field from settings.metadata instead
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Managementunit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Managementunit
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Settings *Managementunitsettingsresponse `json:"settings,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		BusinessUnit: o.BusinessUnit,
		
		StartDayOfWeek: o.StartDayOfWeek,
		
		TimeZone: o.TimeZone,
		
		Settings: o.Settings,
		
		Metadata: o.Metadata,
		
		Division: o.Division,
		
		Version: o.Version,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Managementunit) UnmarshalJSON(b []byte) error {
	var ManagementunitMap map[string]interface{}
	err := json.Unmarshal(b, &ManagementunitMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ManagementunitMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ManagementunitMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if BusinessUnit, ok := ManagementunitMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if StartDayOfWeek, ok := ManagementunitMap["startDayOfWeek"].(string); ok {
		o.StartDayOfWeek = &StartDayOfWeek
	}
    
	if TimeZone, ok := ManagementunitMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if Settings, ok := ManagementunitMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	
	if Metadata, ok := ManagementunitMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if Division, ok := ManagementunitMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Version, ok := ManagementunitMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateModifiedString, ok := ManagementunitMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := ManagementunitMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if SelfUri, ok := ManagementunitMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Managementunit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
