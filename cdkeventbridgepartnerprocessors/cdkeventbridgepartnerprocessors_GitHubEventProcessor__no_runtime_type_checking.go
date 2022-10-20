//go:build no_runtime_type_checking

// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

// Building without runtime type checking enabled, so all the below just return nil

func validateGitHubEventProcessor_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GitHubEventProcessor) validateSetInvocationAlarmParameters(val InvocationAlarm) error {
	return nil
}

func (j *jsiiProxy_GitHubEventProcessor) validateSetPartnerEventsFunctionParameters(val awslambda.Function) error {
	return nil
}

func validateNewGitHubEventProcessorParameters(scope constructs.Construct, id *string, props *GitHubProps) error {
	return nil
}

