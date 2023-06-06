//go:build no_runtime_type_checking

package cdkeventbridgepartnerprocessors

// Building without runtime type checking enabled, so all the below just return nil

func validateInvocationAlarm_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewInvocationAlarmParameters(scope constructs.Construct, id *string, props *InvocationAlarmProps) error {
	return nil
}

