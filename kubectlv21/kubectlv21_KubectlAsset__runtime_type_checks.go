//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// An Asset construct that contains kubectl, for use in Lambda Layers
package kubectlv21

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (k *jsiiProxy_KubectlAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	if resourceProperty == nil {
		return fmt.Errorf("parameter resourceProperty is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KubectlAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateKubectlAsset_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewKubectlAssetParameters(scope constructs.Construct, id *string, options *awss3assets.AssetOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

