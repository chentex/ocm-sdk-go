/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// STSBuilder contains the data and logic needed to build 'STS' objects.
//
// STS contains the necessary attributes to support role-based authentication on AWS.
type STSBuilder struct {
	bitmap_          uint32
	oidcEndpointURL  string
	externalID       string
	operatorIAMRoles []*OperatorIAMRoleBuilder
	roleARN          string
}

// NewSTS creates a new builder of 'STS' objects.
func NewSTS() *STSBuilder {
	return &STSBuilder{}
}

// OIDCEndpointURL sets the value of the 'OIDC_endpoint_URL' attribute to the given value.
//
//
func (b *STSBuilder) OIDCEndpointURL(value string) *STSBuilder {
	b.oidcEndpointURL = value
	b.bitmap_ |= 1
	return b
}

// ExternalID sets the value of the 'external_ID' attribute to the given value.
//
//
func (b *STSBuilder) ExternalID(value string) *STSBuilder {
	b.externalID = value
	b.bitmap_ |= 2
	return b
}

// OperatorIAMRoles sets the value of the 'operator_IAM_roles' attribute to the given values.
//
//
func (b *STSBuilder) OperatorIAMRoles(values ...*OperatorIAMRoleBuilder) *STSBuilder {
	b.operatorIAMRoles = make([]*OperatorIAMRoleBuilder, len(values))
	copy(b.operatorIAMRoles, values)
	b.bitmap_ |= 4
	return b
}

// RoleARN sets the value of the 'role_ARN' attribute to the given value.
//
//
func (b *STSBuilder) RoleARN(value string) *STSBuilder {
	b.roleARN = value
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *STSBuilder) Copy(object *STS) *STSBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.oidcEndpointURL = object.oidcEndpointURL
	b.externalID = object.externalID
	if object.operatorIAMRoles != nil {
		b.operatorIAMRoles = make([]*OperatorIAMRoleBuilder, len(object.operatorIAMRoles))
		for i, v := range object.operatorIAMRoles {
			b.operatorIAMRoles[i] = NewOperatorIAMRole().Copy(v)
		}
	} else {
		b.operatorIAMRoles = nil
	}
	b.roleARN = object.roleARN
	return b
}

// Build creates a 'STS' object using the configuration stored in the builder.
func (b *STSBuilder) Build() (object *STS, err error) {
	object = new(STS)
	object.bitmap_ = b.bitmap_
	object.oidcEndpointURL = b.oidcEndpointURL
	object.externalID = b.externalID
	if b.operatorIAMRoles != nil {
		object.operatorIAMRoles = make([]*OperatorIAMRole, len(b.operatorIAMRoles))
		for i, v := range b.operatorIAMRoles {
			object.operatorIAMRoles[i], err = v.Build()
			if err != nil {
				return
			}
		}
	}
	object.roleARN = b.roleARN
	return
}
