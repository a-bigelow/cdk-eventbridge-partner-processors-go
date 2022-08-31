//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

// Building without runtime type checking enabled, so all the below just return nil

func validateStripeEventProcessor_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_StripeEventProcessor) validateSetInvocationAlarmParameters(val InvocationAlarm) error {
	return nil
}

func (j *jsiiProxy_StripeEventProcessor) validateSetPartnerEventsFunctionParameters(val awslambda.Function) error {
	return nil
}

func validateNewStripeEventProcessorParameters(scope constructs.Construct, id *string, props *StripeProps) error {
	return nil
}

