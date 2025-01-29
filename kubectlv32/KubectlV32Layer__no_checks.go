//go:build no_runtime_type_checking

package kubectlv32

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlV32Layer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlV32Layer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlV32Layer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlV32Layer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKubectlV32Layer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlV32Layer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlV32Layer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlV32Layer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKubectlV32Layer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKubectlV32LayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

