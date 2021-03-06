// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

const (

	// ErrCodeInsufficientDeliveryPolicyException for service response error code
	// "InsufficientDeliveryPolicyException".
	//
	// Your Amazon S3 bucket policy does not permit AWS Config to write to it.
	ErrCodeInsufficientDeliveryPolicyException = "InsufficientDeliveryPolicyException"

	// ErrCodeInsufficientPermissionsException for service response error code
	// "InsufficientPermissionsException".
	//
	// Indicates one of the following errors:
	//
	//    * The rule cannot be created because the IAM role assigned to AWS Config
	//    lacks permissions to perform the config:Put* action.
	//
	//    * The AWS Lambda function cannot be invoked. Check the function ARN, and
	//    check the function's permissions.
	ErrCodeInsufficientPermissionsException = "InsufficientPermissionsException"

	// ErrCodeInvalidConfigurationRecorderNameException for service response error code
	// "InvalidConfigurationRecorderNameException".
	//
	// You have provided a configuration recorder name that is not valid.
	ErrCodeInvalidConfigurationRecorderNameException = "InvalidConfigurationRecorderNameException"

	// ErrCodeInvalidDeliveryChannelNameException for service response error code
	// "InvalidDeliveryChannelNameException".
	//
	// The specified delivery channel name is not valid.
	ErrCodeInvalidDeliveryChannelNameException = "InvalidDeliveryChannelNameException"

	// ErrCodeInvalidExpressionException for service response error code
	// "InvalidExpressionException".
	//
	// The syntax of the query is incorrect.
	ErrCodeInvalidExpressionException = "InvalidExpressionException"

	// ErrCodeInvalidLimitException for service response error code
	// "InvalidLimitException".
	//
	// The specified limit is outside the allowable range.
	ErrCodeInvalidLimitException = "InvalidLimitException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The specified next token is invalid. Specify the nextToken string that was
	// returned in the previous response to get the next page of results.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One or more of the specified parameters are invalid. Verify that your parameters
	// are valid and try again.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRecordingGroupException for service response error code
	// "InvalidRecordingGroupException".
	//
	// AWS Config throws an exception if the recording group does not contain a
	// valid list of resource types. Invalid values might also be incorrectly formatted.
	ErrCodeInvalidRecordingGroupException = "InvalidRecordingGroupException"

	// ErrCodeInvalidResultTokenException for service response error code
	// "InvalidResultTokenException".
	//
	// The specified ResultToken is invalid.
	ErrCodeInvalidResultTokenException = "InvalidResultTokenException"

	// ErrCodeInvalidRoleException for service response error code
	// "InvalidRoleException".
	//
	// You have provided a null or empty role ARN.
	ErrCodeInvalidRoleException = "InvalidRoleException"

	// ErrCodeInvalidS3KeyPrefixException for service response error code
	// "InvalidS3KeyPrefixException".
	//
	// The specified Amazon S3 key prefix is not valid.
	ErrCodeInvalidS3KeyPrefixException = "InvalidS3KeyPrefixException"

	// ErrCodeInvalidSNSTopicARNException for service response error code
	// "InvalidSNSTopicARNException".
	//
	// The specified Amazon SNS topic does not exist.
	ErrCodeInvalidSNSTopicARNException = "InvalidSNSTopicARNException"

	// ErrCodeInvalidTimeRangeException for service response error code
	// "InvalidTimeRangeException".
	//
	// The specified time range is not valid. The earlier time is not chronologically
	// before the later time.
	ErrCodeInvalidTimeRangeException = "InvalidTimeRangeException"

	// ErrCodeLastDeliveryChannelDeleteFailedException for service response error code
	// "LastDeliveryChannelDeleteFailedException".
	//
	// You cannot delete the delivery channel you specified because the configuration
	// recorder is running.
	ErrCodeLastDeliveryChannelDeleteFailedException = "LastDeliveryChannelDeleteFailedException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// For StartConfigRulesEvaluation API, this exception is thrown if an evaluation
	// is in progress or if you call the StartConfigRulesEvaluation API more than
	// once per minute.
	//
	// For PutConfigurationAggregator API, this exception is thrown if the number
	// of accounts and aggregators exceeds the limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMaxNumberOfConfigRulesExceededException for service response error code
	// "MaxNumberOfConfigRulesExceededException".
	//
	// Failed to add the AWS Config rule because the account already contains the
	// maximum number of 150 rules. Consider deleting any deactivated rules before
	// you add new rules.
	ErrCodeMaxNumberOfConfigRulesExceededException = "MaxNumberOfConfigRulesExceededException"

	// ErrCodeMaxNumberOfConfigurationRecordersExceededException for service response error code
	// "MaxNumberOfConfigurationRecordersExceededException".
	//
	// You have reached the limit of the number of recorders you can create.
	ErrCodeMaxNumberOfConfigurationRecordersExceededException = "MaxNumberOfConfigurationRecordersExceededException"

	// ErrCodeMaxNumberOfDeliveryChannelsExceededException for service response error code
	// "MaxNumberOfDeliveryChannelsExceededException".
	//
	// You have reached the limit of the number of delivery channels you can create.
	ErrCodeMaxNumberOfDeliveryChannelsExceededException = "MaxNumberOfDeliveryChannelsExceededException"

	// ErrCodeMaxNumberOfRetentionConfigurationsExceededException for service response error code
	// "MaxNumberOfRetentionConfigurationsExceededException".
	//
	// Failed to add the retention configuration because a retention configuration
	// with that name already exists.
	ErrCodeMaxNumberOfRetentionConfigurationsExceededException = "MaxNumberOfRetentionConfigurationsExceededException"

	// ErrCodeNoAvailableConfigurationRecorderException for service response error code
	// "NoAvailableConfigurationRecorderException".
	//
	// There are no configuration recorders available to provide the role needed
	// to describe your resources. Create a configuration recorder.
	ErrCodeNoAvailableConfigurationRecorderException = "NoAvailableConfigurationRecorderException"

	// ErrCodeNoAvailableDeliveryChannelException for service response error code
	// "NoAvailableDeliveryChannelException".
	//
	// There is no delivery channel available to record configurations.
	ErrCodeNoAvailableDeliveryChannelException = "NoAvailableDeliveryChannelException"

	// ErrCodeNoAvailableOrganizationException for service response error code
	// "NoAvailableOrganizationException".
	//
	// Organization does is no longer available.
	ErrCodeNoAvailableOrganizationException = "NoAvailableOrganizationException"

	// ErrCodeNoRunningConfigurationRecorderException for service response error code
	// "NoRunningConfigurationRecorderException".
	//
	// There is no configuration recorder running.
	ErrCodeNoRunningConfigurationRecorderException = "NoRunningConfigurationRecorderException"

	// ErrCodeNoSuchBucketException for service response error code
	// "NoSuchBucketException".
	//
	// The specified Amazon S3 bucket does not exist.
	ErrCodeNoSuchBucketException = "NoSuchBucketException"

	// ErrCodeNoSuchConfigRuleException for service response error code
	// "NoSuchConfigRuleException".
	//
	// One or more AWS Config rules in the request are invalid. Verify that the
	// rule names are correct and try again.
	ErrCodeNoSuchConfigRuleException = "NoSuchConfigRuleException"

	// ErrCodeNoSuchConfigurationAggregatorException for service response error code
	// "NoSuchConfigurationAggregatorException".
	//
	// You have specified a configuration aggregator that does not exist.
	ErrCodeNoSuchConfigurationAggregatorException = "NoSuchConfigurationAggregatorException"

	// ErrCodeNoSuchConfigurationRecorderException for service response error code
	// "NoSuchConfigurationRecorderException".
	//
	// You have specified a configuration recorder that does not exist.
	ErrCodeNoSuchConfigurationRecorderException = "NoSuchConfigurationRecorderException"

	// ErrCodeNoSuchDeliveryChannelException for service response error code
	// "NoSuchDeliveryChannelException".
	//
	// You have specified a delivery channel that does not exist.
	ErrCodeNoSuchDeliveryChannelException = "NoSuchDeliveryChannelException"

	// ErrCodeNoSuchRemediationConfigurationException for service response error code
	// "NoSuchRemediationConfigurationException".
	//
	// You specified an AWS Config rule without a remediation configuration.
	ErrCodeNoSuchRemediationConfigurationException = "NoSuchRemediationConfigurationException"

	// ErrCodeNoSuchRetentionConfigurationException for service response error code
	// "NoSuchRetentionConfigurationException".
	//
	// You have specified a retention configuration that does not exist.
	ErrCodeNoSuchRetentionConfigurationException = "NoSuchRetentionConfigurationException"

	// ErrCodeOrganizationAccessDeniedException for service response error code
	// "OrganizationAccessDeniedException".
	//
	// No permission to call the EnableAWSServiceAccess API.
	ErrCodeOrganizationAccessDeniedException = "OrganizationAccessDeniedException"

	// ErrCodeOrganizationAllFeaturesNotEnabledException for service response error code
	// "OrganizationAllFeaturesNotEnabledException".
	//
	// The configuration aggregator cannot be created because organization does
	// not have all features enabled.
	ErrCodeOrganizationAllFeaturesNotEnabledException = "OrganizationAllFeaturesNotEnabledException"

	// ErrCodeOversizedConfigurationItemException for service response error code
	// "OversizedConfigurationItemException".
	//
	// The configuration item size is outside the allowable range.
	ErrCodeOversizedConfigurationItemException = "OversizedConfigurationItemException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The rule is currently being deleted or the rule is deleting your evaluation
	// results. Try your request again later.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotDiscoveredException for service response error code
	// "ResourceNotDiscoveredException".
	//
	// You have specified a resource that is either unknown or has not been discovered.
	ErrCodeResourceNotDiscoveredException = "ResourceNotDiscoveredException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// You have specified a resource that does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// You have reached the limit of the number of tags you can use. You have more
	// than 50 tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The requested action is not valid.
	ErrCodeValidationException = "ValidationException"
)
