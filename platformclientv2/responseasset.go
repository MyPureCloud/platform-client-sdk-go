package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseasset
type Responseasset struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// ContentLength - response asset size in bytes
	ContentLength *int `json:"contentLength,omitempty"`


	// ContentLocation - response asset location.
	ContentLocation *string `json:"contentLocation,omitempty"`


	// ContentType - MIME type of response asset
	ContentType *string `json:"contentType,omitempty"`


	// DateCreated - Created date of the response asset. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - User who created the response asset
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// DateModified - Last modified date of the response asset. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - User who last modified the response asset
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Responses - Canned responses actively using this asset
	Responses *[]Domainentityref `json:"responses,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Responseasset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responseasset
	
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
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		ContentLocation *string `json:"contentLocation,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Responses *[]Domainentityref `json:"responses,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		ContentLength: o.ContentLength,
		
		ContentLocation: o.ContentLocation,
		
		ContentType: o.ContentType,
		
		DateCreated: DateCreated,
		
		CreatedBy: o.CreatedBy,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		Responses: o.Responses,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Responseasset) UnmarshalJSON(b []byte) error {
	var ResponseassetMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseassetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ResponseassetMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ResponseassetMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := ResponseassetMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if ContentLength, ok := ResponseassetMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if ContentLocation, ok := ResponseassetMap["contentLocation"].(string); ok {
		o.ContentLocation = &ContentLocation
	}
	
	if ContentType, ok := ResponseassetMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if dateCreatedString, ok := ResponseassetMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if CreatedBy, ok := ResponseassetMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateModifiedString, ok := ResponseassetMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := ResponseassetMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Responses, ok := ResponseassetMap["responses"].([]interface{}); ok {
		ResponsesString, _ := json.Marshal(Responses)
		json.Unmarshal(ResponsesString, &o.Responses)
	}
	
	if SelfUri, ok := ResponseassetMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responseasset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
