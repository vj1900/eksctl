package cloudformation

// AWSIoTTopicRule_DynamoDBv2Action AWS CloudFormation Resource (AWS::IoT::TopicRule.DynamoDBv2Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html
type AWSIoTTopicRule_DynamoDBv2Action struct {

	// PutItem AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html#cfn-iot-topicrule-dynamodbv2action-putitem
	PutItem *AWSIoTTopicRule_PutItemInput `json:"PutItem,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html#cfn-iot-topicrule-dynamodbv2action-rolearn
	RoleArn *StringIntrinsic `json:"RoleArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_DynamoDBv2Action) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.DynamoDBv2Action"
}
