//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An Asset construct that contains kubectl, for use in Lambda Layers
package kubectlv22

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlAsset) validateAddResourceMetadataParameters(resource awscdk.CfnResource, resourceProperty *string) error {
	return nil
}

func (k *jsiiProxy_KubectlAsset) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateKubectlAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKubectlAssetParameters(scope constructs.Construct, id *string, options *awss3assets.AssetOptions) error {
	return nil
}

