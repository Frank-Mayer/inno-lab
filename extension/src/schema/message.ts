/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export enum Command {
  GET_INPUT_POSITION = 0,
  RETURN_INPUT_POSITION = 1,
  REGISTER_PROMPT = 2,
  PROMPT_FINISHED = 3,
  UNRECOGNIZED = -1,
}

export function commandFromJSON(object: any): Command {
  switch (object) {
    case 0:
    case "GET_INPUT_POSITION":
      return Command.GET_INPUT_POSITION;
    case 1:
    case "RETURN_INPUT_POSITION":
      return Command.RETURN_INPUT_POSITION;
    case 2:
    case "REGISTER_PROMPT":
      return Command.REGISTER_PROMPT;
    case 3:
    case "PROMPT_FINISHED":
      return Command.PROMPT_FINISHED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Command.UNRECOGNIZED;
  }
}

export function commandToJSON(object: Command): string {
  switch (object) {
    case Command.GET_INPUT_POSITION:
      return "GET_INPUT_POSITION";
    case Command.RETURN_INPUT_POSITION:
      return "RETURN_INPUT_POSITION";
    case Command.REGISTER_PROMPT:
      return "REGISTER_PROMPT";
    case Command.PROMPT_FINISHED:
      return "PROMPT_FINISHED";
    case Command.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Message {
  command: Command;
  x: number;
  y: number;
  prompt: string;
  url: string;
}

function createBaseMessage(): Message {
  return { command: 0, x: 0, y: 0, prompt: "", url: "" };
}

export const Message = {
  encode(message: Message, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.command !== 0) {
      writer.uint32(8).int32(message.command);
    }
    if (message.x !== 0) {
      writer.uint32(16).int32(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(24).int32(message.y);
    }
    if (message.prompt !== "") {
      writer.uint32(34).string(message.prompt);
    }
    if (message.url !== "") {
      writer.uint32(42).string(message.url);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Message {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.command = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.x = reader.int32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.y = reader.int32();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.prompt = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.url = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Message {
    return {
      command: isSet(object.command) ? commandFromJSON(object.command) : 0,
      x: isSet(object.x) ? globalThis.Number(object.x) : 0,
      y: isSet(object.y) ? globalThis.Number(object.y) : 0,
      prompt: isSet(object.prompt) ? globalThis.String(object.prompt) : "",
      url: isSet(object.url) ? globalThis.String(object.url) : "",
    };
  },

  toJSON(message: Message): unknown {
    const obj: any = {};
    if (message.command !== 0) {
      obj.command = commandToJSON(message.command);
    }
    if (message.x !== 0) {
      obj.x = Math.round(message.x);
    }
    if (message.y !== 0) {
      obj.y = Math.round(message.y);
    }
    if (message.prompt !== "") {
      obj.prompt = message.prompt;
    }
    if (message.url !== "") {
      obj.url = message.url;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Message>, I>>(base?: I): Message {
    return Message.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Message>, I>>(object: I): Message {
    const message = createBaseMessage();
    message.command = object.command ?? 0;
    message.x = object.x ?? 0;
    message.y = object.y ?? 0;
    message.prompt = object.prompt ?? "";
    message.url = object.url ?? "";
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
