//go:build no_runtime_type_checking

package cdkeventbridgepartnerprocessors

// Building without runtime type checking enabled, so all the below just return nil

func validateTwilioEventProcessor_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TwilioEventProcessor) validateSetInvocationAlarmParameters(val InvocationAlarm) error {
	return nil
}

func (j *jsiiProxy_TwilioEventProcessor) validateSetPartnerEventsFunctionParameters(val awslambda.Function) error {
	return nil
}

func validateNewTwilioEventProcessorParameters(scope constructs.Construct, id *string, props *TwilioProps) error {
	return nil
}

