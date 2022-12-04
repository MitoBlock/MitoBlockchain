/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mitoblockchaindev.mitoblockchaindev";

export interface DiscountToken {
  creator: string;
  id: number;
  timestamp: string;
  activityName: string;
  score: string;
  message: string;
  discountValue: string;
  eligibleCompanies: string;
  itemType: string;
  expiryDate: string;
  createdAt: number;
}

const baseDiscountToken: object = {
  creator: "",
  id: 0,
  timestamp: "",
  activityName: "",
  score: "",
  message: "",
  discountValue: "",
  eligibleCompanies: "",
  itemType: "",
  expiryDate: "",
  createdAt: 0,
};

export const DiscountToken = {
  encode(message: DiscountToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.timestamp !== "") {
      writer.uint32(26).string(message.timestamp);
    }
    if (message.activityName !== "") {
      writer.uint32(34).string(message.activityName);
    }
    if (message.score !== "") {
      writer.uint32(42).string(message.score);
    }
    if (message.message !== "") {
      writer.uint32(50).string(message.message);
    }
    if (message.discountValue !== "") {
      writer.uint32(58).string(message.discountValue);
    }
    if (message.eligibleCompanies !== "") {
      writer.uint32(66).string(message.eligibleCompanies);
    }
    if (message.itemType !== "") {
      writer.uint32(74).string(message.itemType);
    }
    if (message.expiryDate !== "") {
      writer.uint32(82).string(message.expiryDate);
    }
    if (message.createdAt !== 0) {
      writer.uint32(88).int64(message.createdAt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DiscountToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDiscountToken } as DiscountToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.timestamp = reader.string();
          break;
        case 4:
          message.activityName = reader.string();
          break;
        case 5:
          message.score = reader.string();
          break;
        case 6:
          message.message = reader.string();
          break;
        case 7:
          message.discountValue = reader.string();
          break;
        case 8:
          message.eligibleCompanies = reader.string();
          break;
        case 9:
          message.itemType = reader.string();
          break;
        case 10:
          message.expiryDate = reader.string();
          break;
        case 11:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DiscountToken {
    const message = { ...baseDiscountToken } as DiscountToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    if (object.activityName !== undefined && object.activityName !== null) {
      message.activityName = String(object.activityName);
    } else {
      message.activityName = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = String(object.score);
    } else {
      message.score = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.discountValue !== undefined && object.discountValue !== null) {
      message.discountValue = String(object.discountValue);
    } else {
      message.discountValue = "";
    }
    if (
      object.eligibleCompanies !== undefined &&
      object.eligibleCompanies !== null
    ) {
      message.eligibleCompanies = String(object.eligibleCompanies);
    } else {
      message.eligibleCompanies = "";
    }
    if (object.itemType !== undefined && object.itemType !== null) {
      message.itemType = String(object.itemType);
    } else {
      message.itemType = "";
    }
    if (object.expiryDate !== undefined && object.expiryDate !== null) {
      message.expiryDate = String(object.expiryDate);
    } else {
      message.expiryDate = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = Number(object.createdAt);
    } else {
      message.createdAt = 0;
    }
    return message;
  },

  toJSON(message: DiscountToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.activityName !== undefined &&
      (obj.activityName = message.activityName);
    message.score !== undefined && (obj.score = message.score);
    message.message !== undefined && (obj.message = message.message);
    message.discountValue !== undefined &&
      (obj.discountValue = message.discountValue);
    message.eligibleCompanies !== undefined &&
      (obj.eligibleCompanies = message.eligibleCompanies);
    message.itemType !== undefined && (obj.itemType = message.itemType);
    message.expiryDate !== undefined && (obj.expiryDate = message.expiryDate);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    return obj;
  },

  fromPartial(object: DeepPartial<DiscountToken>): DiscountToken {
    const message = { ...baseDiscountToken } as DiscountToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    if (object.activityName !== undefined && object.activityName !== null) {
      message.activityName = object.activityName;
    } else {
      message.activityName = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.discountValue !== undefined && object.discountValue !== null) {
      message.discountValue = object.discountValue;
    } else {
      message.discountValue = "";
    }
    if (
      object.eligibleCompanies !== undefined &&
      object.eligibleCompanies !== null
    ) {
      message.eligibleCompanies = object.eligibleCompanies;
    } else {
      message.eligibleCompanies = "";
    }
    if (object.itemType !== undefined && object.itemType !== null) {
      message.itemType = object.itemType;
    } else {
      message.itemType = "";
    }
    if (object.expiryDate !== undefined && object.expiryDate !== null) {
      message.expiryDate = object.expiryDate;
    } else {
      message.expiryDate = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = object.createdAt;
    } else {
      message.createdAt = 0;
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
