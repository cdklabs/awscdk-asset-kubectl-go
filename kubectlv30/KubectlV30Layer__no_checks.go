//go:build no_runtime_type_checking

package kubectlv30

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlV30Layer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlV30Layer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlV30Layer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlV30Layer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKubectlV30Layer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlV30Layer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlV30Layer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlV30Layer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKubectlV30Layer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKubectlV30LayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

