// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/metrics/v3alpha/stats.proto

package envoy_config_metrics_v3alpha

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
var _stats_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on StatsSink with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StatsSink) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	switch m.ConfigType.(type) {

	case *StatsSink_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatsSinkValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StatsSinkValidationError is the validation error returned by
// StatsSink.Validate if the designated constraints aren't met.
type StatsSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsSinkValidationError) ErrorName() string { return "StatsSinkValidationError" }

// Error satisfies the builtin error interface
func (e StatsSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatsSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsSinkValidationError{}

// Validate checks the field values on StatsConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StatsConfig) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStatsTags() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatsConfigValidationError{
					field:  fmt.Sprintf("StatsTags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUseAllDefaultTags()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatsConfigValidationError{
				field:  "UseAllDefaultTags",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStatsMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatsConfigValidationError{
				field:  "StatsMatcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StatsConfigValidationError is the validation error returned by
// StatsConfig.Validate if the designated constraints aren't met.
type StatsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsConfigValidationError) ErrorName() string { return "StatsConfigValidationError" }

// Error satisfies the builtin error interface
func (e StatsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsConfigValidationError{}

// Validate checks the field values on StatsMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StatsMatcher) Validate() error {
	if m == nil {
		return nil
	}

	switch m.StatsMatcher.(type) {

	case *StatsMatcher_RejectAll:
		// no validation rules for RejectAll

	case *StatsMatcher_ExclusionList:

		if v, ok := interface{}(m.GetExclusionList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatsMatcherValidationError{
					field:  "ExclusionList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StatsMatcher_InclusionList:

		if v, ok := interface{}(m.GetInclusionList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatsMatcherValidationError{
					field:  "InclusionList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return StatsMatcherValidationError{
			field:  "StatsMatcher",
			reason: "value is required",
		}

	}

	return nil
}

// StatsMatcherValidationError is the validation error returned by
// StatsMatcher.Validate if the designated constraints aren't met.
type StatsMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsMatcherValidationError) ErrorName() string { return "StatsMatcherValidationError" }

// Error satisfies the builtin error interface
func (e StatsMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatsMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsMatcherValidationError{}

// Validate checks the field values on TagSpecifier with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TagSpecifier) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TagName

	switch m.TagValue.(type) {

	case *TagSpecifier_Regex:

		if len(m.GetRegex()) > 1024 {
			return TagSpecifierValidationError{
				field:  "Regex",
				reason: "value length must be at most 1024 bytes",
			}
		}

	case *TagSpecifier_FixedValue:
		// no validation rules for FixedValue

	}

	return nil
}

// TagSpecifierValidationError is the validation error returned by
// TagSpecifier.Validate if the designated constraints aren't met.
type TagSpecifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagSpecifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagSpecifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagSpecifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagSpecifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagSpecifierValidationError) ErrorName() string { return "TagSpecifierValidationError" }

// Error satisfies the builtin error interface
func (e TagSpecifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagSpecifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagSpecifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagSpecifierValidationError{}

// Validate checks the field values on StatsdSink with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StatsdSink) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Prefix

	switch m.StatsdSpecifier.(type) {

	case *StatsdSink_Address:

		if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatsdSinkValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StatsdSink_TcpClusterName:
		// no validation rules for TcpClusterName

	default:
		return StatsdSinkValidationError{
			field:  "StatsdSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// StatsdSinkValidationError is the validation error returned by
// StatsdSink.Validate if the designated constraints aren't met.
type StatsdSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsdSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsdSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsdSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsdSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsdSinkValidationError) ErrorName() string { return "StatsdSinkValidationError" }

// Error satisfies the builtin error interface
func (e StatsdSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatsdSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsdSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsdSinkValidationError{}

// Validate checks the field values on DogStatsdSink with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DogStatsdSink) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Prefix

	switch m.DogStatsdSpecifier.(type) {

	case *DogStatsdSink_Address:

		if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DogStatsdSinkValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DogStatsdSinkValidationError{
			field:  "DogStatsdSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// DogStatsdSinkValidationError is the validation error returned by
// DogStatsdSink.Validate if the designated constraints aren't met.
type DogStatsdSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DogStatsdSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DogStatsdSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DogStatsdSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DogStatsdSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DogStatsdSinkValidationError) ErrorName() string { return "DogStatsdSinkValidationError" }

// Error satisfies the builtin error interface
func (e DogStatsdSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDogStatsdSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DogStatsdSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DogStatsdSinkValidationError{}

// Validate checks the field values on HystrixSink with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HystrixSink) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NumBuckets

	return nil
}

// HystrixSinkValidationError is the validation error returned by
// HystrixSink.Validate if the designated constraints aren't met.
type HystrixSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HystrixSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HystrixSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HystrixSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HystrixSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HystrixSinkValidationError) ErrorName() string { return "HystrixSinkValidationError" }

// Error satisfies the builtin error interface
func (e HystrixSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHystrixSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HystrixSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HystrixSinkValidationError{}
