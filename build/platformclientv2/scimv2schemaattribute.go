package platformclientv2
import (
	"encoding/json"
)

// Scimv2schemaattribute - A complex type that defines service provider attributes or subattributes and their qualities.
type Scimv2schemaattribute struct { 
	// Name - The name of the attribute.
	Name *string `json:"name,omitempty"`


	// VarType - The data type of the attribute.
	VarType *string `json:"type,omitempty"`


	// SubAttributes - The list of subattributes for an attribute of the type \"complex\". Uses the same schema as \"attributes\".
	SubAttributes *[]Scimv2schemaattribute `json:"subAttributes,omitempty"`


	// MultiValued - Indicates whether an attribute contains multiple values.
	MultiValued *bool `json:"multiValued,omitempty"`


	// Description - The description of the attribute.
	Description *string `json:"description,omitempty"`


	// Required - Indicates whether an attribute is required.
	Required *bool `json:"required,omitempty"`


	// CanonicalValues - The list of standard values that service providers may use. Service providers may ignore unsupported values.
	CanonicalValues *[]string `json:"canonicalValues,omitempty"`


	// CaseExact - Indicates whether a string attribute is case-sensitive. If set to \"true\", the server preserves case sensitivity. If set to \"false\", the server may change the case. The server also uses case sensitivity when evaluating filters. See section 3.4.2.2 \"Filtering\" in RFC 7644 for details.
	CaseExact *bool `json:"caseExact,omitempty"`


	// Mutability - The circumstances under which an attribute can be defined or redefined. The default is \"readWrite\".
	Mutability *string `json:"mutability,omitempty"`


	// Returned - The circumstances under which an attribute and its values are returned in response to a GET, PUT, POST, or PATCH request.
	Returned *string `json:"returned,omitempty"`


	// Uniqueness - The method by which the service provider enforces the uniqueness of an attribute value. A server can reject a value by returning the HTTP response code 400 (Bad Request). A client can enforce uniqueness to a greater degree than the server provider enforces. For example, a client could make a value unique even though the server has \"uniqueness\" set to \"none\".
	Uniqueness *string `json:"uniqueness,omitempty"`


	// ReferenceTypes - The list of SCIM resource types that may be referenced. Only applies when \"type\" is set to \"reference\".
	ReferenceTypes *[]string `json:"referenceTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2schemaattribute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
