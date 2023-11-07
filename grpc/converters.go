package grpc

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetDoubleValue(value *float64) *wrapperspb.DoubleValue {
	if value == nil {
		return nil
	}

	return wrapperspb.Double(*value)
}

func GetDouble(value *wrapperspb.DoubleValue) *float64 {
	if value == nil {
		return nil
	}

	newValue := value.Value
	return &newValue
}

func GetInt32Value(value *int32) *wrapperspb.Int32Value {
	if value == nil {
		return nil
	}

	return wrapperspb.Int32(*value)
}

func GetInt32(value *wrapperspb.Int32Value) *int32 {
	if value == nil {
		return nil
	}

	newValue := value.Value
	return &newValue
}

func GetStringValue(value *string) *wrapperspb.StringValue {
	if value == nil {
		return nil
	}

	return wrapperspb.String(*value)
}

func GetString(value *wrapperspb.StringValue) *string {
	if value == nil {
		return nil
	}

	newValue := value.Value
	return &newValue
}

func GetTimestampValue(value *time.Time) *timestamppb.Timestamp {
	if value == nil {
		return nil
	}

	return timestamppb.New(*value)
}

func GetTime(value *timestamppb.Timestamp) *time.Time {
	if value == nil {
		return nil
	}

	newValue := value.AsTime()
	return &newValue
}
