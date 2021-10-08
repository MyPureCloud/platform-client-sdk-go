package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ivr - Defines the phone numbers, operating hours, and the Architect flows to execute for an IVR.
type Ivr struct { 
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


	// Dnis - The phone number(s) to contact the IVR by.  Each phone number must be unique and not in use by another resource.  For example, a user and an iVR cannot have the same phone number.
	Dnis *[]string `json:"dnis,omitempty"`


	// OpenHoursFlow - The Architect flow to execute during the hours an organization is open.
	OpenHoursFlow *Domainentityref `json:"openHoursFlow,omitempty"`


	// ClosedHoursFlow - The Architect flow to execute during the hours an organization is closed.
	ClosedHoursFlow *Domainentityref `json:"closedHoursFlow,omitempty"`


	// HolidayHoursFlow - The Architect flow to execute during an organization's holiday hours.
	HolidayHoursFlow *Domainentityref `json:"holidayHoursFlow,omitempty"`


	// ScheduleGroup - The schedule group defining the open and closed hours for an organization.  If this is provided, an open flow and a closed flow must be specified as well.
	ScheduleGroup *Domainentityref `json:"scheduleGroup,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Ivr) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ivr
	
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
		
		Dnis *[]string `json:"dnis,omitempty"`
		
		OpenHoursFlow *Domainentityref `json:"openHoursFlow,omitempty"`
		
		ClosedHoursFlow *Domainentityref `json:"closedHoursFlow,omitempty"`
		
		HolidayHoursFlow *Domainentityref `json:"holidayHoursFlow,omitempty"`
		
		ScheduleGroup *Domainentityref `json:"scheduleGroup,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
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
		
		Dnis: o.Dnis,
		
		OpenHoursFlow: o.OpenHoursFlow,
		
		ClosedHoursFlow: o.ClosedHoursFlow,
		
		HolidayHoursFlow: o.HolidayHoursFlow,
		
		ScheduleGroup: o.ScheduleGroup,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Ivr) UnmarshalJSON(b []byte) error {
	var IvrMap map[string]interface{}
	err := json.Unmarshal(b, &IvrMap)
	if err != nil {
		return err
	}
	
	if Id, ok := IvrMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := IvrMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := IvrMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := IvrMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := IvrMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := IvrMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := IvrMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := IvrMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := IvrMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := IvrMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := IvrMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := IvrMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if Dnis, ok := IvrMap["dnis"].([]interface{}); ok {
		DnisString, _ := json.Marshal(Dnis)
		json.Unmarshal(DnisString, &o.Dnis)
	}
	
	if OpenHoursFlow, ok := IvrMap["openHoursFlow"].(map[string]interface{}); ok {
		OpenHoursFlowString, _ := json.Marshal(OpenHoursFlow)
		json.Unmarshal(OpenHoursFlowString, &o.OpenHoursFlow)
	}
	
	if ClosedHoursFlow, ok := IvrMap["closedHoursFlow"].(map[string]interface{}); ok {
		ClosedHoursFlowString, _ := json.Marshal(ClosedHoursFlow)
		json.Unmarshal(ClosedHoursFlowString, &o.ClosedHoursFlow)
	}
	
	if HolidayHoursFlow, ok := IvrMap["holidayHoursFlow"].(map[string]interface{}); ok {
		HolidayHoursFlowString, _ := json.Marshal(HolidayHoursFlow)
		json.Unmarshal(HolidayHoursFlowString, &o.HolidayHoursFlow)
	}
	
	if ScheduleGroup, ok := IvrMap["scheduleGroup"].(map[string]interface{}); ok {
		ScheduleGroupString, _ := json.Marshal(ScheduleGroup)
		json.Unmarshal(ScheduleGroupString, &o.ScheduleGroup)
	}
	
	if SelfUri, ok := IvrMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ivr) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
