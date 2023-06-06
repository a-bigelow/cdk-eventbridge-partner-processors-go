package cdkeventbridgepartnerprocessors

import (
	_init_ "github.com/a-bigelow/cdk-eventbridge-partner-processors-go/cdkeventbridgepartnerprocessors/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// CDK wrapper for the GitHub Eventbridge processor.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github
//
type GitHubEventProcessor interface {
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

// The jsii proxy struct for GitHubEventProcessor
type jsiiProxy_GitHubEventProcessor struct {
	jsiiProxy_PartnerProcessor
}

func (j *jsiiProxy_GitHubEventProcessor) InvocationAlarm() InvocationAlarm {
	var returns InvocationAlarm
	_jsii_.Get(
		j,
		"invocationAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubEventProcessor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubEventProcessor) PartnerEventsFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"partnerEventsFunction",
		&returns,
	)
	return returns
}


func NewGitHubEventProcessor(scope constructs.Construct, id *string, props *GitHubProps) GitHubEventProcessor {
	_init_.Initialize()

	if err := validateNewGitHubEventProcessorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubEventProcessor{}

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.GitHubEventProcessor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGitHubEventProcessor_Override(g GitHubEventProcessor, scope constructs.Construct, id *string, props *GitHubProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.GitHubEventProcessor",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GitHubEventProcessor)SetInvocationAlarm(val InvocationAlarm) {
	if err := j.validateSetInvocationAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationAlarm",
		val,
	)
}

func (j *jsiiProxy_GitHubEventProcessor)SetPartnerEventsFunction(val awslambda.Function) {
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
func GitHubEventProcessor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitHubEventProcessor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-eventbridge-partner-processors.GitHubEventProcessor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubEventProcessor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

