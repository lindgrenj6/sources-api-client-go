/*
 * Sources
 *
 * Sources
 *
 * API version: 3.0.0
 * Contact: support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sourcesapi

import (
	"encoding/json"
)

// AuthenticationsCollection struct for AuthenticationsCollection
type AuthenticationsCollection struct {
	Meta *CollectionMetadata `json:"meta,omitempty"`
	Links *CollectionLinks `json:"links,omitempty"`
	Data *[]Authentication `json:"data,omitempty"`
}

// NewAuthenticationsCollection instantiates a new AuthenticationsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationsCollection() *AuthenticationsCollection {
	this := AuthenticationsCollection{}
	return &this
}

// NewAuthenticationsCollectionWithDefaults instantiates a new AuthenticationsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationsCollectionWithDefaults() *AuthenticationsCollection {
	this := AuthenticationsCollection{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AuthenticationsCollection) GetMeta() CollectionMetadata {
	if o == nil || o.Meta == nil {
		var ret CollectionMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationsCollection) GetMetaOk() (*CollectionMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AuthenticationsCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMetadata and assigns it to the Meta field.
func (o *AuthenticationsCollection) SetMeta(v CollectionMetadata) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthenticationsCollection) GetLinks() CollectionLinks {
	if o == nil || o.Links == nil {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationsCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthenticationsCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *AuthenticationsCollection) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthenticationsCollection) GetData() []Authentication {
	if o == nil || o.Data == nil {
		var ret []Authentication
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationsCollection) GetDataOk() (*[]Authentication, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthenticationsCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Authentication and assigns it to the Data field.
func (o *AuthenticationsCollection) SetData(v []Authentication) {
	o.Data = &v
}

func (o AuthenticationsCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationsCollection struct {
	value *AuthenticationsCollection
	isSet bool
}

func (v NullableAuthenticationsCollection) Get() *AuthenticationsCollection {
	return v.value
}

func (v *NullableAuthenticationsCollection) Set(val *AuthenticationsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationsCollection(val *AuthenticationsCollection) *NullableAuthenticationsCollection {
	return &NullableAuthenticationsCollection{value: val, isSet: true}
}

func (v NullableAuthenticationsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

