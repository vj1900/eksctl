package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSApiGatewayResource AWS CloudFormation Resource (AWS::ApiGateway::Resource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html
type AWSApiGatewayResource struct {

	// ParentId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	ParentId *StringIntrinsic `json:"ParentId,omitempty"`

	// PathPart AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	PathPart *StringIntrinsic `json:"PathPart,omitempty"`

	// RestApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	RestApiId *StringIntrinsic `json:"RestApiId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayResource) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Resource"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSApiGatewayResource) MarshalJSON() ([]byte, error) {
	type Properties AWSApiGatewayResource
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSApiGatewayResource) UnmarshalJSON(b []byte) error {
	type Properties AWSApiGatewayResource
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSApiGatewayResource(*res.Properties)
	}

	return nil
}

// GetAllAWSApiGatewayResourceResources retrieves all AWSApiGatewayResource items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayResourceResources() map[string]AWSApiGatewayResource {
	results := map[string]AWSApiGatewayResource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSApiGatewayResource:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApiGateway::Resource" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSApiGatewayResource
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSApiGatewayResourceWithName retrieves all AWSApiGatewayResource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayResourceWithName(name string) (AWSApiGatewayResource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSApiGatewayResource:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApiGateway::Resource" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSApiGatewayResource
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSApiGatewayResource{}, errors.New("resource not found")
}
