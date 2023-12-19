/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

export const protobufPackage = "";

export interface Res {
  prompt: string;
  src: string;
  time: number;
}

function createBaseRes(): Res {
  return { prompt: "", src: "", time: 0 };
}

export const Res = {
  encode(message: Res, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.prompt !== "") {
      writer.uint32(10).string(message.prompt);
    }
    if (message.src !== "") {
      writer.uint32(18).string(message.src);
    }
    if (message.time !== 0) {
      writer.uint32(24).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Res {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRes();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.prompt = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.src = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.time = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Res {
    return {
      prompt: isSet(object.prompt) ? globalThis.String(object.prompt) : "",
      src: isSet(object.src) ? globalThis.String(object.src) : "",
      time: isSet(object.time) ? globalThis.Number(object.time) : 0,
    };
  },

  toJSON(message: Res): unknown {
    const obj: any = {};
    if (message.prompt !== "") {
      obj.prompt = message.prompt;
    }
    if (message.src !== "") {
      obj.src = message.src;
    }
    if (message.time !== 0) {
      obj.time = Math.round(message.time);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Res>, I>>(base?: I): Res {
    return Res.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Res>, I>>(object: I): Res {
    const message = createBaseRes();
    message.prompt = object.prompt ?? "";
    message.src = object.src ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(globalThis.Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
