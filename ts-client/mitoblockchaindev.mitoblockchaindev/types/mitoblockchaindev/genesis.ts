/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../mitoblockchaindev/params";
import { DiscountTokenStatus } from "../mitoblockchaindev/discount_token_status";
import { MembershipTokenStatus } from "../mitoblockchaindev/membership_token_status";

export const protobufPackage = "mitoblockchaindev.mitoblockchaindev";

/** GenesisState defines the mitoblockchaindev module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  discountTokenStatusList: DiscountTokenStatus[];
  discountTokenStatusCount: number;
  membershipTokenStatusList: MembershipTokenStatus[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  membershipTokenStatusCount: number;
}

const baseGenesisState: object = {
  discountTokenStatusCount: 0,
  membershipTokenStatusCount: 0,
};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.discountTokenStatusList) {
      DiscountTokenStatus.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.discountTokenStatusCount !== 0) {
      writer.uint32(24).uint64(message.discountTokenStatusCount);
    }
    for (const v of message.membershipTokenStatusList) {
      MembershipTokenStatus.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.membershipTokenStatusCount !== 0) {
      writer.uint32(40).uint64(message.membershipTokenStatusCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.discountTokenStatusList = [];
    message.membershipTokenStatusList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.discountTokenStatusList.push(
            DiscountTokenStatus.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.discountTokenStatusCount = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 4:
          message.membershipTokenStatusList.push(
            MembershipTokenStatus.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.membershipTokenStatusCount = longToNumber(
            reader.uint64() as Long
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.discountTokenStatusList = [];
    message.membershipTokenStatusList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.discountTokenStatusList !== undefined &&
      object.discountTokenStatusList !== null
    ) {
      for (const e of object.discountTokenStatusList) {
        message.discountTokenStatusList.push(DiscountTokenStatus.fromJSON(e));
      }
    }
    if (
      object.discountTokenStatusCount !== undefined &&
      object.discountTokenStatusCount !== null
    ) {
      message.discountTokenStatusCount = Number(
        object.discountTokenStatusCount
      );
    } else {
      message.discountTokenStatusCount = 0;
    }
    if (
      object.membershipTokenStatusList !== undefined &&
      object.membershipTokenStatusList !== null
    ) {
      for (const e of object.membershipTokenStatusList) {
        message.membershipTokenStatusList.push(
          MembershipTokenStatus.fromJSON(e)
        );
      }
    }
    if (
      object.membershipTokenStatusCount !== undefined &&
      object.membershipTokenStatusCount !== null
    ) {
      message.membershipTokenStatusCount = Number(
        object.membershipTokenStatusCount
      );
    } else {
      message.membershipTokenStatusCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.discountTokenStatusList) {
      obj.discountTokenStatusList = message.discountTokenStatusList.map((e) =>
        e ? DiscountTokenStatus.toJSON(e) : undefined
      );
    } else {
      obj.discountTokenStatusList = [];
    }
    message.discountTokenStatusCount !== undefined &&
      (obj.discountTokenStatusCount = message.discountTokenStatusCount);
    if (message.membershipTokenStatusList) {
      obj.membershipTokenStatusList = message.membershipTokenStatusList.map(
        (e) => (e ? MembershipTokenStatus.toJSON(e) : undefined)
      );
    } else {
      obj.membershipTokenStatusList = [];
    }
    message.membershipTokenStatusCount !== undefined &&
      (obj.membershipTokenStatusCount = message.membershipTokenStatusCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.discountTokenStatusList = [];
    message.membershipTokenStatusList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.discountTokenStatusList !== undefined &&
      object.discountTokenStatusList !== null
    ) {
      for (const e of object.discountTokenStatusList) {
        message.discountTokenStatusList.push(
          DiscountTokenStatus.fromPartial(e)
        );
      }
    }
    if (
      object.discountTokenStatusCount !== undefined &&
      object.discountTokenStatusCount !== null
    ) {
      message.discountTokenStatusCount = object.discountTokenStatusCount;
    } else {
      message.discountTokenStatusCount = 0;
    }
    if (
      object.membershipTokenStatusList !== undefined &&
      object.membershipTokenStatusList !== null
    ) {
      for (const e of object.membershipTokenStatusList) {
        message.membershipTokenStatusList.push(
          MembershipTokenStatus.fromPartial(e)
        );
      }
    }
    if (
      object.membershipTokenStatusCount !== undefined &&
      object.membershipTokenStatusCount !== null
    ) {
      message.membershipTokenStatusCount = object.membershipTokenStatusCount;
    } else {
      message.membershipTokenStatusCount = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
