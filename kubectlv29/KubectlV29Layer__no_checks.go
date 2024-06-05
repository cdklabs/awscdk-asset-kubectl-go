//go:build no_runtime_type_checking

package kubectlv29

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlV29Layer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlV29Layer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlV29Layer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlV29Layer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKubectlV29Layer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlV29Layer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlV29Layer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlV29Layer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKubectlV29Layer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKubectlV29LayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

