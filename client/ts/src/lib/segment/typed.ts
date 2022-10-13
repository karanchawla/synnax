import { ChannelPayload } from '../channel/payload';
import { ContiguityError, ValidationError } from '../errors';
import {
  Density,
  Size,
  TimeRange,
  TimeSpan,
  TimeStamp,
  TypedArray,
} from '../telem';

import { SegmentPayload } from './payload';

export default class TypedSegment {
  payload: SegmentPayload;
  channel: ChannelPayload;
  view: TypedArray;

  constructor(channel: ChannelPayload, payload: SegmentPayload) {
    this.channel = channel;
    this.payload = payload;
    this.view = new this.channel.dataType.arrayConstructor(
      this.payload.data.buffer
    );
  }

  get start(): TimeStamp {
    return this.payload.start;
  }

  get span(): TimeSpan {
    return this.channel.rate.byteSpan(
      Size.Bytes(this.view.byteLength),
      this.channel.density as Density
    );
  }

  get range(): TimeRange {
    return this.start.spanRange(this.span);
  }

  get end(): TimeStamp {
    return this.range.end;
  }

  get size(): Size {
    return Size.Bytes(this.view.byteLength);
  }

  extend(other: TypedSegment) {
    if (other.channel.key !== this.channel.key) {
      throw new ValidationError(`
        Cannot extend segment because channel keys mismatch.
        Segment Channel Key: ${this.channel.key}
        Other Segment Channel Key: ${other.channel.key}
      `);
    } else if (!this.end.equals(other.start)) {
      throw new ContiguityError(`
      Cannot extend segment because segments are not contiguous.
      Segment End: ${this.end}
      Other Segment Start: ${other.start}
      `);
    }
    const newData = new Uint8Array(
      this.view.byteLength + other.view.byteLength
    );
    newData.set(this.payload.data, 0);
    newData.set(other.payload.data, this.view.byteLength);
    this.payload.data = newData;
    this.view = new this.channel.dataType.arrayConstructor(
      this.payload.data.buffer
    );
  }
}