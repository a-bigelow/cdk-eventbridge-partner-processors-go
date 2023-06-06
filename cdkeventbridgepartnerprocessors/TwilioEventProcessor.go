package cdkeventbridgepartnerprocessors

import (
	_init_ "github.com/a-bigelow/cdk-eventbridge-partner-processors-go/cdkeventbridgepartnerprocessors/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// CDK wrapper for the Twilio Eventbridge processor.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github
//
type TwilioEventProcessor interface {
	PartnerProcessor
	InvocationAlarm() InvocationAlarm
	SetInvocationAlarm(val InvocationAlarm)
	// The tree node.
	Node() constructs.Node
	PartnerEventsFunction() awslambda.Function
	SetPartnerEventsFunction(val awslambda.Function)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TwilioEventProcessor
type jsiiProxy_TwilioEventProcessor struct {
	jsiiProxy_PartnerProcessor
}

func (j *jsiiProxy_TwilioEventProcessor) InvocationAlarm() InvocationAlarm {
	var returns InvocationAlarm
	_jsii_.Get(
		j,
		"invocationAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TwilioEventProcessor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TwilioEventProcessor) PartnerEventsFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"partnerEventsFunction",
		&returns,
	)
	return returns
}


func NewTwilioEventProcessor(scope constructs.Construct, id *string, props *TwilioProps) TwilioEventProcessor {
	_init_.Initialize()

	if err := validateNewTwilioEventProcessorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TwilioEventProcessor{}

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.TwilioEventProcessor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTwilioEventProcessor_Override(t TwilioEventProcessor, scope constructs.Construct, id *string, props *TwilioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.TwilioEventProcessor",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TwilioEventProcessor)SetInvocationAlarm(val InvocationAlarm) {
	if err := j.validateSetInvocationAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationAlarm",
		val,
	)
}

func (j *jsiiProxy_TwilioEventProcessor)SetPartnerEventsFunction(val awslambda.Function) {
	if err := j.validateSetPartnerEventsFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerEventsFunction",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TwilioEventProcessor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTwilioEventProcessor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-eventbridge-partner-processors.TwilioEventProcessor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TwilioEventProcessor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

