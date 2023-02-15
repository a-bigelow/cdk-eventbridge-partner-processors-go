//go:build no_runtime_type_checking

// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

// Building without runtime type checking enabled, so all the below just return nil

func validatePartnerProcessor_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PartnerProcessor) validateSetInvocationAlarmParameters(val InvocationAlarm) error {
	return nil
}

func (j *jsiiProxy_PartnerProcessor) validateSetPartnerEventsFunctionParameters(val awslambda.Function) error {
	return nil
}

func validateNewPartnerProcessorParameters(scope constructs.Construct, id *string, props *PartnerFunctionProps) error {
	return nil
}

