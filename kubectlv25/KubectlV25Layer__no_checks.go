//go:build no_runtime_type_checking

// A Lambda Layer that contains kubectl v1.25
package kubectlv25

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlV25Layer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlV25Layer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlV25Layer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlV25Layer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKubectlV25Layer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlV25Layer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlV25Layer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlV25Layer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKubectlV25LayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

