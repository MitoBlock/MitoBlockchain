/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "mitoblockchaindev.mitoblockchaindev";

export interface MsgCreateDiscountToken {
  creator: string;
  timestamp: string;
  activityName: string;
  score: string;
  message: string;
  discountValue: string;
  eligibleCompanies: string;
  itemType: string;
  expiryDate: string;
  id: number;
}

export interface MsgCreateDiscountTokenResponse {
  id: number;
}

export interface MsgCreateDiscountTokenStatus {
  creator: string;
  tokenID: number;
  timestamp: string;
  status: string;
  id: number;
}

export interface MsgCreateDiscountTokenStatusResponse {
  id: number;
}

export interface MsgDeleteDiscountTokenStatus {
  creator: string;
  discountTokenStatusID: number;
  tokenID: number;
  id: number;
}

export interface MsgDeleteDiscountTokenStatusResponse {
  id: number;
}

export interface MsgCreateMembershipToken {
  creator: string;
  timestamp: string;
  activityName: string;
  score: string;
  message: string;
  membershipDuration: string;
  eligibleCompanies: string;
  expiryDate: string;
  id: number;
}

export interface MsgCreateMembershipTokenResponse {
  id: number;
}

export interface MsgCreateMembershipTokenStatus {
  creator: string;
  tokenID: number;
  timestamp: string;
  status: string;
  id: number;
}

export interface MsgCreateMembershipTokenStatusResponse {
  id: number;
}

export interface MsgDeleteMembershipTokenStatus {
  creator: string;
  membershipTokenStatusID: number;
  tokenID: number;
  id: number;
}

export interface MsgDeleteMembershipTokenStatusResponse {
  id: number;
}

const baseMsgCreateDiscountToken: object = {
  creator: "",
  timestamp: "",
  activityName: "",
  score: "",
  message: "",
  discountValue: "",
  eligibleCompanies: "",
  itemType: "",
  expiryDate: "",
  id: 0,
};

export const MsgCreateDiscountToken = {
  encode(
    message: MsgCreateDiscountToken,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.timestamp !== "") {
      writer.uint32(18).string(message.timestamp);
    }
    if (message.activityName !== "") {
      writer.uint32(26).string(message.activityName);
    }
    if (message.score !== "") {
      writer.uint32(34).string(message.score);
    }
    if (message.message !== "") {
      writer.uint32(42).string(message.message);
    }
    if (message.discountValue !== "") {
      writer.uint32(50).string(message.discountValue);
    }
    if (message.eligibleCompanies !== "") {
      writer.uint32(58).string(message.eligibleCompanies);
    }
    if (message.itemType !== "") {
      writer.uint32(66).string(message.itemType);
    }
    if (message.expiryDate !== "") {
      writer.uint32(74).string(message.expiryDate);
    }
    if (message.id !== 0) {
      writer.uint32(80).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDiscountToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDiscountToken } as MsgCreateDiscountToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.timestamp = reader.string();
          break;
        case 3:
          message.activityName = reader.string();
          break;
        case 4:
          message.score = reader.string();
          break;
        case 5:
          message.message = reader.string();
          break;
        case 6:
          message.discountValue = reader.string();
          break;
        case 7:
          message.eligibleCompanies = reader.string();
          break;
        case 8:
          message.itemType = reader.string();
          break;
        case 9:
          message.expiryDate = reader.string();
          break;
        case 10:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDiscountToken {
    const message = { ...baseMsgCreateDiscountToken } as MsgCreateDiscountToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateDiscountToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
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
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateDiscountToken>
  ): MsgCreateDiscountToken {
    const message = { ...baseMsgCreateDiscountToken } as MsgCreateDiscountToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateDiscountTokenResponse: object = { id: 0 };

export const MsgCreateDiscountTokenResponse = {
  encode(
    message: MsgCreateDiscountTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateDiscountTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateDiscountTokenResponse,
    } as MsgCreateDiscountTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDiscountTokenResponse {
    const message = {
      ...baseMsgCreateDiscountTokenResponse,
    } as MsgCreateDiscountTokenResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateDiscountTokenResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateDiscountTokenResponse>
  ): MsgCreateDiscountTokenResponse {
    const message = {
      ...baseMsgCreateDiscountTokenResponse,
    } as MsgCreateDiscountTokenResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateDiscountTokenStatus: object = {
  creator: "",
  tokenID: 0,
  timestamp: "",
  status: "",
  id: 0,
};

export const MsgCreateDiscountTokenStatus = {
  encode(
    message: MsgCreateDiscountTokenStatus,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenID !== 0) {
      writer.uint32(16).uint64(message.tokenID);
    }
    if (message.timestamp !== "") {
      writer.uint32(26).string(message.timestamp);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    if (message.id !== 0) {
      writer.uint32(40).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateDiscountTokenStatus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateDiscountTokenStatus,
    } as MsgCreateDiscountTokenStatus;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenID = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.timestamp = reader.string();
          break;
        case 4:
          message.status = reader.string();
          break;
        case 5:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDiscountTokenStatus {
    const message = {
      ...baseMsgCreateDiscountTokenStatus,
    } as MsgCreateDiscountTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = Number(object.tokenID);
    } else {
      message.tokenID = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateDiscountTokenStatus): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenID !== undefined && (obj.tokenID = message.tokenID);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.status !== undefined && (obj.status = message.status);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateDiscountTokenStatus>
  ): MsgCreateDiscountTokenStatus {
    const message = {
      ...baseMsgCreateDiscountTokenStatus,
    } as MsgCreateDiscountTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = object.tokenID;
    } else {
      message.tokenID = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateDiscountTokenStatusResponse: object = { id: 0 };

export const MsgCreateDiscountTokenStatusResponse = {
  encode(
    message: MsgCreateDiscountTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateDiscountTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateDiscountTokenStatusResponse,
    } as MsgCreateDiscountTokenStatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDiscountTokenStatusResponse {
    const message = {
      ...baseMsgCreateDiscountTokenStatusResponse,
    } as MsgCreateDiscountTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateDiscountTokenStatusResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateDiscountTokenStatusResponse>
  ): MsgCreateDiscountTokenStatusResponse {
    const message = {
      ...baseMsgCreateDiscountTokenStatusResponse,
    } as MsgCreateDiscountTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteDiscountTokenStatus: object = {
  creator: "",
  discountTokenStatusID: 0,
  tokenID: 0,
  id: 0,
};

export const MsgDeleteDiscountTokenStatus = {
  encode(
    message: MsgDeleteDiscountTokenStatus,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.discountTokenStatusID !== 0) {
      writer.uint32(16).uint64(message.discountTokenStatusID);
    }
    if (message.tokenID !== 0) {
      writer.uint32(24).uint64(message.tokenID);
    }
    if (message.id !== 0) {
      writer.uint32(32).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteDiscountTokenStatus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteDiscountTokenStatus,
    } as MsgDeleteDiscountTokenStatus;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.discountTokenStatusID = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.tokenID = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDiscountTokenStatus {
    const message = {
      ...baseMsgDeleteDiscountTokenStatus,
    } as MsgDeleteDiscountTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.discountTokenStatusID !== undefined &&
      object.discountTokenStatusID !== null
    ) {
      message.discountTokenStatusID = Number(object.discountTokenStatusID);
    } else {
      message.discountTokenStatusID = 0;
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = Number(object.tokenID);
    } else {
      message.tokenID = 0;
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteDiscountTokenStatus): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.discountTokenStatusID !== undefined &&
      (obj.discountTokenStatusID = message.discountTokenStatusID);
    message.tokenID !== undefined && (obj.tokenID = message.tokenID);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteDiscountTokenStatus>
  ): MsgDeleteDiscountTokenStatus {
    const message = {
      ...baseMsgDeleteDiscountTokenStatus,
    } as MsgDeleteDiscountTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.discountTokenStatusID !== undefined &&
      object.discountTokenStatusID !== null
    ) {
      message.discountTokenStatusID = object.discountTokenStatusID;
    } else {
      message.discountTokenStatusID = 0;
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = object.tokenID;
    } else {
      message.tokenID = 0;
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteDiscountTokenStatusResponse: object = { id: 0 };

export const MsgDeleteDiscountTokenStatusResponse = {
  encode(
    message: MsgDeleteDiscountTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteDiscountTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteDiscountTokenStatusResponse,
    } as MsgDeleteDiscountTokenStatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDiscountTokenStatusResponse {
    const message = {
      ...baseMsgDeleteDiscountTokenStatusResponse,
    } as MsgDeleteDiscountTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteDiscountTokenStatusResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteDiscountTokenStatusResponse>
  ): MsgDeleteDiscountTokenStatusResponse {
    const message = {
      ...baseMsgDeleteDiscountTokenStatusResponse,
    } as MsgDeleteDiscountTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateMembershipToken: object = {
  creator: "",
  timestamp: "",
  activityName: "",
  score: "",
  message: "",
  membershipDuration: "",
  eligibleCompanies: "",
  expiryDate: "",
  id: 0,
};

export const MsgCreateMembershipToken = {
  encode(
    message: MsgCreateMembershipToken,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.timestamp !== "") {
      writer.uint32(18).string(message.timestamp);
    }
    if (message.activityName !== "") {
      writer.uint32(26).string(message.activityName);
    }
    if (message.score !== "") {
      writer.uint32(34).string(message.score);
    }
    if (message.message !== "") {
      writer.uint32(42).string(message.message);
    }
    if (message.membershipDuration !== "") {
      writer.uint32(50).string(message.membershipDuration);
    }
    if (message.eligibleCompanies !== "") {
      writer.uint32(58).string(message.eligibleCompanies);
    }
    if (message.expiryDate !== "") {
      writer.uint32(66).string(message.expiryDate);
    }
    if (message.id !== 0) {
      writer.uint32(72).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateMembershipToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMembershipToken,
    } as MsgCreateMembershipToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.timestamp = reader.string();
          break;
        case 3:
          message.activityName = reader.string();
          break;
        case 4:
          message.score = reader.string();
          break;
        case 5:
          message.message = reader.string();
          break;
        case 6:
          message.membershipDuration = reader.string();
          break;
        case 7:
          message.eligibleCompanies = reader.string();
          break;
        case 8:
          message.expiryDate = reader.string();
          break;
        case 9:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMembershipToken {
    const message = {
      ...baseMsgCreateMembershipToken,
    } as MsgCreateMembershipToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (
      object.membershipDuration !== undefined &&
      object.membershipDuration !== null
    ) {
      message.membershipDuration = String(object.membershipDuration);
    } else {
      message.membershipDuration = "";
    }
    if (
      object.eligibleCompanies !== undefined &&
      object.eligibleCompanies !== null
    ) {
      message.eligibleCompanies = String(object.eligibleCompanies);
    } else {
      message.eligibleCompanies = "";
    }
    if (object.expiryDate !== undefined && object.expiryDate !== null) {
      message.expiryDate = String(object.expiryDate);
    } else {
      message.expiryDate = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateMembershipToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.activityName !== undefined &&
      (obj.activityName = message.activityName);
    message.score !== undefined && (obj.score = message.score);
    message.message !== undefined && (obj.message = message.message);
    message.membershipDuration !== undefined &&
      (obj.membershipDuration = message.membershipDuration);
    message.eligibleCompanies !== undefined &&
      (obj.eligibleCompanies = message.eligibleCompanies);
    message.expiryDate !== undefined && (obj.expiryDate = message.expiryDate);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateMembershipToken>
  ): MsgCreateMembershipToken {
    const message = {
      ...baseMsgCreateMembershipToken,
    } as MsgCreateMembershipToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (
      object.membershipDuration !== undefined &&
      object.membershipDuration !== null
    ) {
      message.membershipDuration = object.membershipDuration;
    } else {
      message.membershipDuration = "";
    }
    if (
      object.eligibleCompanies !== undefined &&
      object.eligibleCompanies !== null
    ) {
      message.eligibleCompanies = object.eligibleCompanies;
    } else {
      message.eligibleCompanies = "";
    }
    if (object.expiryDate !== undefined && object.expiryDate !== null) {
      message.expiryDate = object.expiryDate;
    } else {
      message.expiryDate = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateMembershipTokenResponse: object = { id: 0 };

export const MsgCreateMembershipTokenResponse = {
  encode(
    message: MsgCreateMembershipTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateMembershipTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMembershipTokenResponse,
    } as MsgCreateMembershipTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMembershipTokenResponse {
    const message = {
      ...baseMsgCreateMembershipTokenResponse,
    } as MsgCreateMembershipTokenResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateMembershipTokenResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateMembershipTokenResponse>
  ): MsgCreateMembershipTokenResponse {
    const message = {
      ...baseMsgCreateMembershipTokenResponse,
    } as MsgCreateMembershipTokenResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateMembershipTokenStatus: object = {
  creator: "",
  tokenID: 0,
  timestamp: "",
  status: "",
  id: 0,
};

export const MsgCreateMembershipTokenStatus = {
  encode(
    message: MsgCreateMembershipTokenStatus,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenID !== 0) {
      writer.uint32(16).uint64(message.tokenID);
    }
    if (message.timestamp !== "") {
      writer.uint32(26).string(message.timestamp);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    if (message.id !== 0) {
      writer.uint32(40).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateMembershipTokenStatus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMembershipTokenStatus,
    } as MsgCreateMembershipTokenStatus;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenID = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.timestamp = reader.string();
          break;
        case 4:
          message.status = reader.string();
          break;
        case 5:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMembershipTokenStatus {
    const message = {
      ...baseMsgCreateMembershipTokenStatus,
    } as MsgCreateMembershipTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = Number(object.tokenID);
    } else {
      message.tokenID = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateMembershipTokenStatus): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenID !== undefined && (obj.tokenID = message.tokenID);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.status !== undefined && (obj.status = message.status);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateMembershipTokenStatus>
  ): MsgCreateMembershipTokenStatus {
    const message = {
      ...baseMsgCreateMembershipTokenStatus,
    } as MsgCreateMembershipTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = object.tokenID;
    } else {
      message.tokenID = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCreateMembershipTokenStatusResponse: object = { id: 0 };

export const MsgCreateMembershipTokenStatusResponse = {
  encode(
    message: MsgCreateMembershipTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateMembershipTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMembershipTokenStatusResponse,
    } as MsgCreateMembershipTokenStatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMembershipTokenStatusResponse {
    const message = {
      ...baseMsgCreateMembershipTokenStatusResponse,
    } as MsgCreateMembershipTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateMembershipTokenStatusResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateMembershipTokenStatusResponse>
  ): MsgCreateMembershipTokenStatusResponse {
    const message = {
      ...baseMsgCreateMembershipTokenStatusResponse,
    } as MsgCreateMembershipTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteMembershipTokenStatus: object = {
  creator: "",
  membershipTokenStatusID: 0,
  tokenID: 0,
  id: 0,
};

export const MsgDeleteMembershipTokenStatus = {
  encode(
    message: MsgDeleteMembershipTokenStatus,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.membershipTokenStatusID !== 0) {
      writer.uint32(16).uint64(message.membershipTokenStatusID);
    }
    if (message.tokenID !== 0) {
      writer.uint32(24).uint64(message.tokenID);
    }
    if (message.id !== 0) {
      writer.uint32(32).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteMembershipTokenStatus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteMembershipTokenStatus,
    } as MsgDeleteMembershipTokenStatus;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.membershipTokenStatusID = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 3:
          message.tokenID = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteMembershipTokenStatus {
    const message = {
      ...baseMsgDeleteMembershipTokenStatus,
    } as MsgDeleteMembershipTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.membershipTokenStatusID !== undefined &&
      object.membershipTokenStatusID !== null
    ) {
      message.membershipTokenStatusID = Number(object.membershipTokenStatusID);
    } else {
      message.membershipTokenStatusID = 0;
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = Number(object.tokenID);
    } else {
      message.tokenID = 0;
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteMembershipTokenStatus): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.membershipTokenStatusID !== undefined &&
      (obj.membershipTokenStatusID = message.membershipTokenStatusID);
    message.tokenID !== undefined && (obj.tokenID = message.tokenID);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteMembershipTokenStatus>
  ): MsgDeleteMembershipTokenStatus {
    const message = {
      ...baseMsgDeleteMembershipTokenStatus,
    } as MsgDeleteMembershipTokenStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.membershipTokenStatusID !== undefined &&
      object.membershipTokenStatusID !== null
    ) {
      message.membershipTokenStatusID = object.membershipTokenStatusID;
    } else {
      message.membershipTokenStatusID = 0;
    }
    if (object.tokenID !== undefined && object.tokenID !== null) {
      message.tokenID = object.tokenID;
    } else {
      message.tokenID = 0;
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteMembershipTokenStatusResponse: object = { id: 0 };

export const MsgDeleteMembershipTokenStatusResponse = {
  encode(
    message: MsgDeleteMembershipTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteMembershipTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteMembershipTokenStatusResponse,
    } as MsgDeleteMembershipTokenStatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteMembershipTokenStatusResponse {
    const message = {
      ...baseMsgDeleteMembershipTokenStatusResponse,
    } as MsgDeleteMembershipTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteMembershipTokenStatusResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteMembershipTokenStatusResponse>
  ): MsgDeleteMembershipTokenStatusResponse {
    const message = {
      ...baseMsgDeleteMembershipTokenStatusResponse,
    } as MsgDeleteMembershipTokenStatusResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDiscountToken(
    request: MsgCreateDiscountToken
  ): Promise<MsgCreateDiscountTokenResponse>;
  CreateDiscountTokenStatus(
    request: MsgCreateDiscountTokenStatus
  ): Promise<MsgCreateDiscountTokenStatusResponse>;
  DeleteDiscountTokenStatus(
    request: MsgDeleteDiscountTokenStatus
  ): Promise<MsgDeleteDiscountTokenStatusResponse>;
  CreateMembershipToken(
    request: MsgCreateMembershipToken
  ): Promise<MsgCreateMembershipTokenResponse>;
  CreateMembershipTokenStatus(
    request: MsgCreateMembershipTokenStatus
  ): Promise<MsgCreateMembershipTokenStatusResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteMembershipTokenStatus(
    request: MsgDeleteMembershipTokenStatus
  ): Promise<MsgDeleteMembershipTokenStatusResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateDiscountToken(
    request: MsgCreateDiscountToken
  ): Promise<MsgCreateDiscountTokenResponse> {
    const data = MsgCreateDiscountToken.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Msg",
      "CreateDiscountToken",
      data
    );
    return promise.then((data) =>
      MsgCreateDiscountTokenResponse.decode(new Reader(data))
    );
  }

  CreateDiscountTokenStatus(
    request: MsgCreateDiscountTokenStatus
  ): Promise<MsgCreateDiscountTokenStatusResponse> {
    const data = MsgCreateDiscountTokenStatus.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Msg",
      "CreateDiscountTokenStatus",
      data
    );
    return promise.then((data) =>
      MsgCreateDiscountTokenStatusResponse.decode(new Reader(data))
    );
  }

  DeleteDiscountTokenStatus(
    request: MsgDeleteDiscountTokenStatus
  ): Promise<MsgDeleteDiscountTokenStatusResponse> {
    const data = MsgDeleteDiscountTokenStatus.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Msg",
      "DeleteDiscountTokenStatus",
      data
    );
    return promise.then((data) =>
      MsgDeleteDiscountTokenStatusResponse.decode(new Reader(data))
    );
  }

  CreateMembershipToken(
    request: MsgCreateMembershipToken
  ): Promise<MsgCreateMembershipTokenResponse> {
    const data = MsgCreateMembershipToken.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Msg",
      "CreateMembershipToken",
      data
    );
    return promise.then((data) =>
      MsgCreateMembershipTokenResponse.decode(new Reader(data))
    );
  }

  CreateMembershipTokenStatus(
    request: MsgCreateMembershipTokenStatus
  ): Promise<MsgCreateMembershipTokenStatusResponse> {
    const data = MsgCreateMembershipTokenStatus.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Msg",
      "CreateMembershipTokenStatus",
      data
    );
    return promise.then((data) =>
      MsgCreateMembershipTokenStatusResponse.decode(new Reader(data))
    );
  }

  DeleteMembershipTokenStatus(
    request: MsgDeleteMembershipTokenStatus
  ): Promise<MsgDeleteMembershipTokenStatusResponse> {
    const data = MsgDeleteMembershipTokenStatus.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Msg",
      "DeleteMembershipTokenStatus",
      data
    );
    return promise.then((data) =>
      MsgDeleteMembershipTokenStatusResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
