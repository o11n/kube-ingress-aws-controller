package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

type elbv2MockOutputs struct {
	registerTargets      *apiResponse
	deregisterTargets    *apiResponse
	describeTags         *apiResponse
	describeTargetGroups *apiResponse
	describeTargetHealth *apiResponse
}

type mockElbv2Client struct {
	elbv2iface.ELBV2API
	outputs  elbv2MockOutputs
	rtinputs []*elbv2.RegisterTargetsInput
	dtinputs []*elbv2.DeregisterTargetsInput
}

func (m *mockElbv2Client) RegisterTargets(in *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	m.rtinputs = append(m.rtinputs, in)
	if out, ok := m.outputs.registerTargets.response.(*elbv2.RegisterTargetsOutput); ok {
		return out, m.outputs.registerTargets.err
	}
	return nil, m.outputs.registerTargets.err
}

func mockRTOutput() *elbv2.RegisterTargetsOutput {
	return &elbv2.RegisterTargetsOutput{}
}

func (m *mockElbv2Client) DeregisterTargets(in *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	m.dtinputs = append(m.dtinputs, in)
	if out, ok := m.outputs.deregisterTargets.response.(*elbv2.DeregisterTargetsOutput); ok {
		return out, m.outputs.deregisterTargets.err
	}
	return nil, m.outputs.deregisterTargets.err
}

func (m *mockElbv2Client) DescribeTags(tags *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	if out, ok := m.outputs.describeTags.response.(*elbv2.DescribeTagsOutput); ok {
		return out, m.outputs.describeTags.err
	}
	return nil, m.outputs.describeTags.err
}

func (m *mockElbv2Client) DescribeTargetGroupsPagesWithContext(ctx aws.Context, in *elbv2.DescribeTargetGroupsInput, f func(resp *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool, opt ...request.Option) error {
	if out, ok := m.outputs.describeTargetGroups.response.(*elbv2.DescribeTargetGroupsOutput); ok {
		f(out, true)
	}
	return m.outputs.describeTargetGroups.err
}

func (m *mockElbv2Client) DescribeTargetHealth(*elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	if out, ok := m.outputs.describeTargetHealth.response.(*elbv2.DescribeTargetHealthOutput); ok {
		return out, m.outputs.describeTargetHealth.err
	}
	return nil, m.outputs.describeTargetHealth.err
}

func mockDTOutput() *elbv2.DeregisterTargetsOutput {
	return &elbv2.DeregisterTargetsOutput{}
}
