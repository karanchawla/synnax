#  Copyright 2023 Synnax Labs, Inc.
#
#  Use of this software is governed by the Business Source License included in the file
#  licenses/BSL.txt.
#
#  As of the Change Date specified in that file, in accordance with the Business Source
#  License, use of this software will be governed by the Apache License, Version 2.0,
#  included in the file licenses/APL.txt.

from alamos.instrumentation import Traceable
from opentelemetry.propagate import get_global_textmap
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import (
    BatchSpanProcessor,
    ConsoleSpanExporter,
)

from alamos import Tracer, Instrumentation, trace

provider = TracerProvider()
processor = BatchSpanProcessor(ConsoleSpanExporter())
provider.add_span_processor(processor)


class TestTrace:
    def test_initialize(self):
        """
        Should initialize the tracer.
        """
        tracer = Tracer(otel_provider=provider, otel_propagator=get_global_textmap())
        assert tracer is not None

    def test_trace(self, instrumentation: Instrumentation):
        """Should not raise an exception."""
        with instrumentation.T.prod("test") as span:
            assert span.key == "test"
            pass

    def test_trace_decorator(self, instrumentation: Instrumentation):
        """Should not raise an exception"""

        class Foo:

            def _(self) -> Traceable:
                return self

            def __init__(self, ins: Instrumentation):
                self.instrumentation = ins

            @trace("prod")
            def decorated(self) -> str:
                return "hello"

        Foo(instrumentation).decorated()


class TestPropagate:
    def test_propagate_depropagate(self, instrumentation: Instrumentation):
        """Should correctly inject the span context into the carrier."""

        class Carrier(dict):
            def set(key, value):
                self[key] = value

        carrier = Carrier()

        with instrumentation.T.prod("test"):
            instrumentation.T.propagate(carrier)

        assert "traceparent" in carrier
