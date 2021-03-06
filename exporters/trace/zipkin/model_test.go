package zipkin

import (
	"testing"
	"time"

	zkmodel "github.com/openzipkin/zipkin-go/model"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	export "go.opentelemetry.io/otel/sdk/export/trace"
)

func TestModelConversion(t *testing.T) {
	inputBatch := []*export.SpanData{
		// typical span data
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindServer,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data with no parent (same as typical, but has
		// invalid parent)
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{},
			SpanKind:     trace.SpanKindServer,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data of unspecified kind
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindUnspecified,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data of internal kind
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindInternal,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data of client kind
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindClient,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data of producer kind
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindProducer,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data of consumer kind
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindConsumer,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data with no events
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindServer,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.Uint64("attr1", 42),
				key.String("attr2", "bar"),
			},
			MessageEvents: nil,
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
		// span data with an "error" attribute set to "false"
		{
			SpanContext: core.SpanContext{
				TraceID: core.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				SpanID:  core.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			},
			ParentSpanID: core.SpanID{0x3F, 0x3E, 0x3D, 0x3C, 0x3B, 0x3A, 0x39, 0x38},
			SpanKind:     trace.SpanKindServer,
			Name:         "foo",
			StartTime:    time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			EndTime:      time.Date(2020, time.March, 11, 19, 25, 0, 0, time.UTC),
			Attributes: []core.KeyValue{
				key.String("error", "false"),
			},
			MessageEvents: []export.Event{
				{
					Time: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Name: "ev1",
					Attributes: []core.KeyValue{
						key.Uint64("eventattr1", 123),
					},
				},
				{
					Time:       time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Name:       "ev2",
					Attributes: nil,
				},
			},
			StatusCode:    codes.NotFound,
			StatusMessage: "404, file not found",
		},
	}

	expectedOutputBatch := []zkmodel.SpanModel{
		// model for typical span data
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "SERVER",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data with no parent
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: nil,
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "SERVER",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data of unspecified kind
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data of internal kind
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data of client kind
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "CLIENT",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data of producer kind
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "PRODUCER",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data of consumer kind
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "CONSUMER",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data with no events
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "SERVER",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations:    nil,
			Tags: map[string]string{
				"attr1":                 "42",
				"attr2":                 "bar",
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
		// model for span data with an "error" attribute set to "false"
		{
			SpanContext: zkmodel.SpanContext{
				TraceID: zkmodel.TraceID{
					High: 0x001020304050607,
					Low:  0x8090a0b0c0d0e0f,
				},
				ID:       zkmodel.ID(0xfffefdfcfbfaf9f8),
				ParentID: zkmodelIDPtr(0x3f3e3d3c3b3a3938),
				Debug:    false,
				Sampled:  nil,
				Err:      nil,
			},
			Name:           "foo",
			Kind:           "SERVER",
			Timestamp:      time.Date(2020, time.March, 11, 19, 24, 0, 0, time.UTC),
			Duration:       time.Minute,
			Shared:         false,
			LocalEndpoint:  nil,
			RemoteEndpoint: nil,
			Annotations: []zkmodel.Annotation{
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 30, 0, time.UTC),
					Value:     `ev1: {"eventattr1":123}`,
				},
				{
					Timestamp: time.Date(2020, time.March, 11, 19, 24, 45, 0, time.UTC),
					Value:     "ev2",
				},
			},
			Tags: map[string]string{
				"ot.status_code":        "NotFound",
				"ot.status_description": "404, file not found",
			},
		},
	}
	gottenOutputBatch := toZipkinSpanModels(inputBatch)
	require.Equal(t, expectedOutputBatch, gottenOutputBatch)
}

func zkmodelIDPtr(n uint64) *zkmodel.ID {
	id := zkmodel.ID(n)
	return &id
}
