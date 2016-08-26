// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package service wraps SSM service
package service

import (
	"testing"

	"github.com/aws/amazon-ssm-agent/agent/log"
	ssmSvc "github.com/aws/amazon-ssm-agent/agent/ssm"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	instanceID = "i-test"
)

var ssmMock = ssmSvc.NewMockDefault()
var logMock = log.NewMockLog()

func TestListAssociations(t *testing.T) {
	associationName := "test"
	associationList := []*ssm.Association{
		&ssm.Association{
			Name: &associationName,
		}}

	listAssociationsOutput := ssm.ListAssociationsOutput{
		Associations: associationList,
	}
	getDocumentOutput := ssm.GetDocumentOutput{
		Name: &associationName,
	}

	ssmMock.On("ListAssociations", mock.AnythingOfType("*log.Mock"), mock.AnythingOfType("string")).Return(&listAssociationsOutput, nil)
	ssmMock.On("GetDocument", mock.AnythingOfType("*log.Mock"), mock.AnythingOfType("string")).Return(&getDocumentOutput, nil)

	response, err := ListAssociations(logMock, ssmMock, instanceID)

	assert.NoError(t, err)
	assert.Equal(t, *response.Association.Name, "test")
}
