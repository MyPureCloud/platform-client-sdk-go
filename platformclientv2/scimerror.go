package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimerror - Defines a SCIM error.
type Scimerror struct { 
	// Schemas - The list of schemas for the SCIM error.
	Schemas *[]string `json:"schemas,omitempty"`


	// Status - The HTTP status code returned for the SCIM error.
	Status *string `json:"status,omitempty"`


	// ScimType - The type of SCIM error when httpStatus is a \"400\" error.
	ScimType *string `json:"scimType,omitempty"`


	// Detail - The detailed description of the SCIM error.
	Detail *string `json:"detail,omitempty"`

}

func (o *Scimerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimerror
	
	return json.Marshal(&struct { 
		Schemas *[]string `json:"schemas,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ScimType *string `json:"scimType,omitempty"`
		
		Detail *string `json:"detail,omitempty"`
		*Alias
	}{ 
		Schemas: o.Schemas,
		
		Status: o.Status,
		
		ScimType: o.ScimType,
		
		Detail: o.Detail,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimerror) UnmarshalJSON(b []byte) error {
	var ScimerrorMap map[string]interface{}
	err := json.Unmarshal(b, &ScimerrorMap)
	if err != nil {
		return err
	}
	
	if Schemas, ok := ScimerrorMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if Status, ok := ScimerrorMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if ScimType, ok := ScimerrorMap["scimType"].(string); ok {
		o.ScimType = &ScimType
	}
	
	if Detail, ok := ScimerrorMap["detail"].(string); ok {
		o.Detail = &Detail
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
