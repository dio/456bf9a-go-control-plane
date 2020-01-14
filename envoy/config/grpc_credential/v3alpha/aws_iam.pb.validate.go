// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/grpc_credential/v3alpha/aws_iam.proto

package envoy_config_grpc_credential_v3alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _aws_iam_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AwsIamConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AwsIamConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetServiceName()) < 1 {
		return AwsIamConfigValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Region

	return nil
}

// AwsIamConfigValidationError is the validation error returned by
// AwsIamConfig.Validate if the designated constraints aren't met.
type AwsIamConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AwsIamConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AwsIamConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AwsIamConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AwsIamConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AwsIamConfigValidationError) ErrorName() string { return "AwsIamConfigValidationError" }

// Error satisfies the builtin error interface
func (e AwsIamConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAwsIamConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AwsIamConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AwsIamConfigValidationError{}
