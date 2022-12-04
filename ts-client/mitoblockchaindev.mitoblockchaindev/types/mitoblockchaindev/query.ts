/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../mitoblockchaindev/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { DiscountToken } from "../mitoblockchaindev/discount_token";
import { DiscountTokenStatus } from "../mitoblockchaindev/discount_token_status";
import { MembershipToken } from "../mitoblockchaindev/membership_token";
import { MembershipTokenStatus } from "../mitoblockchaindev/membership_token_status";

export const protobufPackage = "mitoblockchaindev.mitoblockchaindev";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryDiscountTokensRequest {
  /** Adding pagination to request */
  pagination: PageRequest | undefined;
}

export interface QueryDiscountTokensResponse {
  /** Returning a list of discount tokens */
  DiscountToken: DiscountToken[];
  /** Adding pagination to response */
  pagination: PageResponse | undefined;
}

export interface QueryGetDiscountTokenStatusRequest {
  id: number;
}

export interface QueryGetDiscountTokenStatusResponse {
  DiscountTokenStatus: DiscountTokenStatus | undefined;
}

export interface QueryAllDiscountTokenStatusRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDiscountTokenStatusResponse {
  DiscountTokenStatus: DiscountTokenStatus[];
  pagination: PageResponse | undefined;
}

export interface QueryDiscountTokenStatusQRequest {
  id: number;
  /** Adding pagination to request */
  pagination: PageRequest | undefined;
}

export interface QueryDiscountTokenStatusQResponse {
  DiscountToken: DiscountToken | undefined;
  /** Returning a list of discount token status */
  DiscountTokenStatus: DiscountTokenStatus[];
  /** Adding pagination to response */
  pagination: PageResponse | undefined;
}

export interface QueryMembershipTokensRequest {
  /** Adding pagination to request */
  pagination: PageRequest | undefined;
}

export interface QueryMembershipTokensResponse {
  /** Returning a list of posts */
  MembershipToken: MembershipToken[];
  /** Adding pagination to response */
  pagination: PageResponse | undefined;
}

export interface QueryGetMembershipTokenStatusRequest {
  id: number;
}

export interface QueryGetMembershipTokenStatusResponse {
  MembershipTokenStatus: MembershipTokenStatus | undefined;
}

export interface QueryAllMembershipTokenStatusRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMembershipTokenStatusResponse {
  MembershipTokenStatus: MembershipTokenStatus[];
  pagination: PageResponse | undefined;
}

export interface QueryMembershipTokenStatusQRequest {
  id: number;
  /** Adding pagination to request */
  pagination: PageRequest | undefined;
}

export interface QueryMembershipTokenStatusQResponse {
  MembershipToken: MembershipToken | undefined;
  /** Returning a list of membership token status */
  MembershipTokenStatus: MembershipTokenStatus[];
  /** Adding pagination to response */
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryDiscountTokensRequest: object = {};

export const QueryDiscountTokensRequest = {
  encode(
    message: QueryDiscountTokensRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDiscountTokensRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDiscountTokensRequest,
    } as QueryDiscountTokensRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDiscountTokensRequest {
    const message = {
      ...baseQueryDiscountTokensRequest,
    } as QueryDiscountTokensRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDiscountTokensRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDiscountTokensRequest>
  ): QueryDiscountTokensRequest {
    const message = {
      ...baseQueryDiscountTokensRequest,
    } as QueryDiscountTokensRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDiscountTokensResponse: object = {};

export const QueryDiscountTokensResponse = {
  encode(
    message: QueryDiscountTokensResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.DiscountToken) {
      DiscountToken.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDiscountTokensResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDiscountTokensResponse,
    } as QueryDiscountTokensResponse;
    message.DiscountToken = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.DiscountToken.push(
            DiscountToken.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDiscountTokensResponse {
    const message = {
      ...baseQueryDiscountTokensResponse,
    } as QueryDiscountTokensResponse;
    message.DiscountToken = [];
    if (object.DiscountToken !== undefined && object.DiscountToken !== null) {
      for (const e of object.DiscountToken) {
        message.DiscountToken.push(DiscountToken.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDiscountTokensResponse): unknown {
    const obj: any = {};
    if (message.DiscountToken) {
      obj.DiscountToken = message.DiscountToken.map((e) =>
        e ? DiscountToken.toJSON(e) : undefined
      );
    } else {
      obj.DiscountToken = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDiscountTokensResponse>
  ): QueryDiscountTokensResponse {
    const message = {
      ...baseQueryDiscountTokensResponse,
    } as QueryDiscountTokensResponse;
    message.DiscountToken = [];
    if (object.DiscountToken !== undefined && object.DiscountToken !== null) {
      for (const e of object.DiscountToken) {
        message.DiscountToken.push(DiscountToken.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetDiscountTokenStatusRequest: object = { id: 0 };

export const QueryGetDiscountTokenStatusRequest = {
  encode(
    message: QueryGetDiscountTokenStatusRequest,
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
  ): QueryGetDiscountTokenStatusRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetDiscountTokenStatusRequest,
    } as QueryGetDiscountTokenStatusRequest;
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

  fromJSON(object: any): QueryGetDiscountTokenStatusRequest {
    const message = {
      ...baseQueryGetDiscountTokenStatusRequest,
    } as QueryGetDiscountTokenStatusRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetDiscountTokenStatusRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetDiscountTokenStatusRequest>
  ): QueryGetDiscountTokenStatusRequest {
    const message = {
      ...baseQueryGetDiscountTokenStatusRequest,
    } as QueryGetDiscountTokenStatusRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetDiscountTokenStatusResponse: object = {};

export const QueryGetDiscountTokenStatusResponse = {
  encode(
    message: QueryGetDiscountTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.DiscountTokenStatus !== undefined) {
      DiscountTokenStatus.encode(
        message.DiscountTokenStatus,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetDiscountTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetDiscountTokenStatusResponse,
    } as QueryGetDiscountTokenStatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.DiscountTokenStatus = DiscountTokenStatus.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDiscountTokenStatusResponse {
    const message = {
      ...baseQueryGetDiscountTokenStatusResponse,
    } as QueryGetDiscountTokenStatusResponse;
    if (
      object.DiscountTokenStatus !== undefined &&
      object.DiscountTokenStatus !== null
    ) {
      message.DiscountTokenStatus = DiscountTokenStatus.fromJSON(
        object.DiscountTokenStatus
      );
    } else {
      message.DiscountTokenStatus = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetDiscountTokenStatusResponse): unknown {
    const obj: any = {};
    message.DiscountTokenStatus !== undefined &&
      (obj.DiscountTokenStatus = message.DiscountTokenStatus
        ? DiscountTokenStatus.toJSON(message.DiscountTokenStatus)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetDiscountTokenStatusResponse>
  ): QueryGetDiscountTokenStatusResponse {
    const message = {
      ...baseQueryGetDiscountTokenStatusResponse,
    } as QueryGetDiscountTokenStatusResponse;
    if (
      object.DiscountTokenStatus !== undefined &&
      object.DiscountTokenStatus !== null
    ) {
      message.DiscountTokenStatus = DiscountTokenStatus.fromPartial(
        object.DiscountTokenStatus
      );
    } else {
      message.DiscountTokenStatus = undefined;
    }
    return message;
  },
};

const baseQueryAllDiscountTokenStatusRequest: object = {};

export const QueryAllDiscountTokenStatusRequest = {
  encode(
    message: QueryAllDiscountTokenStatusRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllDiscountTokenStatusRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllDiscountTokenStatusRequest,
    } as QueryAllDiscountTokenStatusRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDiscountTokenStatusRequest {
    const message = {
      ...baseQueryAllDiscountTokenStatusRequest,
    } as QueryAllDiscountTokenStatusRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllDiscountTokenStatusRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllDiscountTokenStatusRequest>
  ): QueryAllDiscountTokenStatusRequest {
    const message = {
      ...baseQueryAllDiscountTokenStatusRequest,
    } as QueryAllDiscountTokenStatusRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllDiscountTokenStatusResponse: object = {};

export const QueryAllDiscountTokenStatusResponse = {
  encode(
    message: QueryAllDiscountTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.DiscountTokenStatus) {
      DiscountTokenStatus.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllDiscountTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllDiscountTokenStatusResponse,
    } as QueryAllDiscountTokenStatusResponse;
    message.DiscountTokenStatus = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.DiscountTokenStatus.push(
            DiscountTokenStatus.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDiscountTokenStatusResponse {
    const message = {
      ...baseQueryAllDiscountTokenStatusResponse,
    } as QueryAllDiscountTokenStatusResponse;
    message.DiscountTokenStatus = [];
    if (
      object.DiscountTokenStatus !== undefined &&
      object.DiscountTokenStatus !== null
    ) {
      for (const e of object.DiscountTokenStatus) {
        message.DiscountTokenStatus.push(DiscountTokenStatus.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllDiscountTokenStatusResponse): unknown {
    const obj: any = {};
    if (message.DiscountTokenStatus) {
      obj.DiscountTokenStatus = message.DiscountTokenStatus.map((e) =>
        e ? DiscountTokenStatus.toJSON(e) : undefined
      );
    } else {
      obj.DiscountTokenStatus = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllDiscountTokenStatusResponse>
  ): QueryAllDiscountTokenStatusResponse {
    const message = {
      ...baseQueryAllDiscountTokenStatusResponse,
    } as QueryAllDiscountTokenStatusResponse;
    message.DiscountTokenStatus = [];
    if (
      object.DiscountTokenStatus !== undefined &&
      object.DiscountTokenStatus !== null
    ) {
      for (const e of object.DiscountTokenStatus) {
        message.DiscountTokenStatus.push(DiscountTokenStatus.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDiscountTokenStatusQRequest: object = { id: 0 };

export const QueryDiscountTokenStatusQRequest = {
  encode(
    message: QueryDiscountTokenStatusQRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDiscountTokenStatusQRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDiscountTokenStatusQRequest,
    } as QueryDiscountTokenStatusQRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDiscountTokenStatusQRequest {
    const message = {
      ...baseQueryDiscountTokenStatusQRequest,
    } as QueryDiscountTokenStatusQRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDiscountTokenStatusQRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDiscountTokenStatusQRequest>
  ): QueryDiscountTokenStatusQRequest {
    const message = {
      ...baseQueryDiscountTokenStatusQRequest,
    } as QueryDiscountTokenStatusQRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDiscountTokenStatusQResponse: object = {};

export const QueryDiscountTokenStatusQResponse = {
  encode(
    message: QueryDiscountTokenStatusQResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.DiscountToken !== undefined) {
      DiscountToken.encode(
        message.DiscountToken,
        writer.uint32(10).fork()
      ).ldelim();
    }
    for (const v of message.DiscountTokenStatus) {
      DiscountTokenStatus.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDiscountTokenStatusQResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDiscountTokenStatusQResponse,
    } as QueryDiscountTokenStatusQResponse;
    message.DiscountTokenStatus = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.DiscountToken = DiscountToken.decode(reader, reader.uint32());
          break;
        case 2:
          message.DiscountTokenStatus.push(
            DiscountTokenStatus.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDiscountTokenStatusQResponse {
    const message = {
      ...baseQueryDiscountTokenStatusQResponse,
    } as QueryDiscountTokenStatusQResponse;
    message.DiscountTokenStatus = [];
    if (object.DiscountToken !== undefined && object.DiscountToken !== null) {
      message.DiscountToken = DiscountToken.fromJSON(object.DiscountToken);
    } else {
      message.DiscountToken = undefined;
    }
    if (
      object.DiscountTokenStatus !== undefined &&
      object.DiscountTokenStatus !== null
    ) {
      for (const e of object.DiscountTokenStatus) {
        message.DiscountTokenStatus.push(DiscountTokenStatus.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDiscountTokenStatusQResponse): unknown {
    const obj: any = {};
    message.DiscountToken !== undefined &&
      (obj.DiscountToken = message.DiscountToken
        ? DiscountToken.toJSON(message.DiscountToken)
        : undefined);
    if (message.DiscountTokenStatus) {
      obj.DiscountTokenStatus = message.DiscountTokenStatus.map((e) =>
        e ? DiscountTokenStatus.toJSON(e) : undefined
      );
    } else {
      obj.DiscountTokenStatus = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDiscountTokenStatusQResponse>
  ): QueryDiscountTokenStatusQResponse {
    const message = {
      ...baseQueryDiscountTokenStatusQResponse,
    } as QueryDiscountTokenStatusQResponse;
    message.DiscountTokenStatus = [];
    if (object.DiscountToken !== undefined && object.DiscountToken !== null) {
      message.DiscountToken = DiscountToken.fromPartial(object.DiscountToken);
    } else {
      message.DiscountToken = undefined;
    }
    if (
      object.DiscountTokenStatus !== undefined &&
      object.DiscountTokenStatus !== null
    ) {
      for (const e of object.DiscountTokenStatus) {
        message.DiscountTokenStatus.push(DiscountTokenStatus.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMembershipTokensRequest: object = {};

export const QueryMembershipTokensRequest = {
  encode(
    message: QueryMembershipTokensRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryMembershipTokensRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryMembershipTokensRequest,
    } as QueryMembershipTokensRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMembershipTokensRequest {
    const message = {
      ...baseQueryMembershipTokensRequest,
    } as QueryMembershipTokensRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMembershipTokensRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMembershipTokensRequest>
  ): QueryMembershipTokensRequest {
    const message = {
      ...baseQueryMembershipTokensRequest,
    } as QueryMembershipTokensRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMembershipTokensResponse: object = {};

export const QueryMembershipTokensResponse = {
  encode(
    message: QueryMembershipTokensResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.MembershipToken) {
      MembershipToken.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryMembershipTokensResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryMembershipTokensResponse,
    } as QueryMembershipTokensResponse;
    message.MembershipToken = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MembershipToken.push(
            MembershipToken.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMembershipTokensResponse {
    const message = {
      ...baseQueryMembershipTokensResponse,
    } as QueryMembershipTokensResponse;
    message.MembershipToken = [];
    if (
      object.MembershipToken !== undefined &&
      object.MembershipToken !== null
    ) {
      for (const e of object.MembershipToken) {
        message.MembershipToken.push(MembershipToken.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMembershipTokensResponse): unknown {
    const obj: any = {};
    if (message.MembershipToken) {
      obj.MembershipToken = message.MembershipToken.map((e) =>
        e ? MembershipToken.toJSON(e) : undefined
      );
    } else {
      obj.MembershipToken = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMembershipTokensResponse>
  ): QueryMembershipTokensResponse {
    const message = {
      ...baseQueryMembershipTokensResponse,
    } as QueryMembershipTokensResponse;
    message.MembershipToken = [];
    if (
      object.MembershipToken !== undefined &&
      object.MembershipToken !== null
    ) {
      for (const e of object.MembershipToken) {
        message.MembershipToken.push(MembershipToken.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetMembershipTokenStatusRequest: object = { id: 0 };

export const QueryGetMembershipTokenStatusRequest = {
  encode(
    message: QueryGetMembershipTokenStatusRequest,
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
  ): QueryGetMembershipTokenStatusRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMembershipTokenStatusRequest,
    } as QueryGetMembershipTokenStatusRequest;
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

  fromJSON(object: any): QueryGetMembershipTokenStatusRequest {
    const message = {
      ...baseQueryGetMembershipTokenStatusRequest,
    } as QueryGetMembershipTokenStatusRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetMembershipTokenStatusRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMembershipTokenStatusRequest>
  ): QueryGetMembershipTokenStatusRequest {
    const message = {
      ...baseQueryGetMembershipTokenStatusRequest,
    } as QueryGetMembershipTokenStatusRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetMembershipTokenStatusResponse: object = {};

export const QueryGetMembershipTokenStatusResponse = {
  encode(
    message: QueryGetMembershipTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.MembershipTokenStatus !== undefined) {
      MembershipTokenStatus.encode(
        message.MembershipTokenStatus,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMembershipTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMembershipTokenStatusResponse,
    } as QueryGetMembershipTokenStatusResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MembershipTokenStatus = MembershipTokenStatus.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMembershipTokenStatusResponse {
    const message = {
      ...baseQueryGetMembershipTokenStatusResponse,
    } as QueryGetMembershipTokenStatusResponse;
    if (
      object.MembershipTokenStatus !== undefined &&
      object.MembershipTokenStatus !== null
    ) {
      message.MembershipTokenStatus = MembershipTokenStatus.fromJSON(
        object.MembershipTokenStatus
      );
    } else {
      message.MembershipTokenStatus = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMembershipTokenStatusResponse): unknown {
    const obj: any = {};
    message.MembershipTokenStatus !== undefined &&
      (obj.MembershipTokenStatus = message.MembershipTokenStatus
        ? MembershipTokenStatus.toJSON(message.MembershipTokenStatus)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMembershipTokenStatusResponse>
  ): QueryGetMembershipTokenStatusResponse {
    const message = {
      ...baseQueryGetMembershipTokenStatusResponse,
    } as QueryGetMembershipTokenStatusResponse;
    if (
      object.MembershipTokenStatus !== undefined &&
      object.MembershipTokenStatus !== null
    ) {
      message.MembershipTokenStatus = MembershipTokenStatus.fromPartial(
        object.MembershipTokenStatus
      );
    } else {
      message.MembershipTokenStatus = undefined;
    }
    return message;
  },
};

const baseQueryAllMembershipTokenStatusRequest: object = {};

export const QueryAllMembershipTokenStatusRequest = {
  encode(
    message: QueryAllMembershipTokenStatusRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllMembershipTokenStatusRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMembershipTokenStatusRequest,
    } as QueryAllMembershipTokenStatusRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMembershipTokenStatusRequest {
    const message = {
      ...baseQueryAllMembershipTokenStatusRequest,
    } as QueryAllMembershipTokenStatusRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMembershipTokenStatusRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMembershipTokenStatusRequest>
  ): QueryAllMembershipTokenStatusRequest {
    const message = {
      ...baseQueryAllMembershipTokenStatusRequest,
    } as QueryAllMembershipTokenStatusRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMembershipTokenStatusResponse: object = {};

export const QueryAllMembershipTokenStatusResponse = {
  encode(
    message: QueryAllMembershipTokenStatusResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.MembershipTokenStatus) {
      MembershipTokenStatus.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllMembershipTokenStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMembershipTokenStatusResponse,
    } as QueryAllMembershipTokenStatusResponse;
    message.MembershipTokenStatus = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MembershipTokenStatus.push(
            MembershipTokenStatus.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMembershipTokenStatusResponse {
    const message = {
      ...baseQueryAllMembershipTokenStatusResponse,
    } as QueryAllMembershipTokenStatusResponse;
    message.MembershipTokenStatus = [];
    if (
      object.MembershipTokenStatus !== undefined &&
      object.MembershipTokenStatus !== null
    ) {
      for (const e of object.MembershipTokenStatus) {
        message.MembershipTokenStatus.push(MembershipTokenStatus.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMembershipTokenStatusResponse): unknown {
    const obj: any = {};
    if (message.MembershipTokenStatus) {
      obj.MembershipTokenStatus = message.MembershipTokenStatus.map((e) =>
        e ? MembershipTokenStatus.toJSON(e) : undefined
      );
    } else {
      obj.MembershipTokenStatus = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMembershipTokenStatusResponse>
  ): QueryAllMembershipTokenStatusResponse {
    const message = {
      ...baseQueryAllMembershipTokenStatusResponse,
    } as QueryAllMembershipTokenStatusResponse;
    message.MembershipTokenStatus = [];
    if (
      object.MembershipTokenStatus !== undefined &&
      object.MembershipTokenStatus !== null
    ) {
      for (const e of object.MembershipTokenStatus) {
        message.MembershipTokenStatus.push(
          MembershipTokenStatus.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMembershipTokenStatusQRequest: object = { id: 0 };

export const QueryMembershipTokenStatusQRequest = {
  encode(
    message: QueryMembershipTokenStatusQRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryMembershipTokenStatusQRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryMembershipTokenStatusQRequest,
    } as QueryMembershipTokenStatusQRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMembershipTokenStatusQRequest {
    const message = {
      ...baseQueryMembershipTokenStatusQRequest,
    } as QueryMembershipTokenStatusQRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMembershipTokenStatusQRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMembershipTokenStatusQRequest>
  ): QueryMembershipTokenStatusQRequest {
    const message = {
      ...baseQueryMembershipTokenStatusQRequest,
    } as QueryMembershipTokenStatusQRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMembershipTokenStatusQResponse: object = {};

export const QueryMembershipTokenStatusQResponse = {
  encode(
    message: QueryMembershipTokenStatusQResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.MembershipToken !== undefined) {
      MembershipToken.encode(
        message.MembershipToken,
        writer.uint32(10).fork()
      ).ldelim();
    }
    for (const v of message.MembershipTokenStatus) {
      MembershipTokenStatus.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryMembershipTokenStatusQResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryMembershipTokenStatusQResponse,
    } as QueryMembershipTokenStatusQResponse;
    message.MembershipTokenStatus = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MembershipToken = MembershipToken.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.MembershipTokenStatus.push(
            MembershipTokenStatus.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMembershipTokenStatusQResponse {
    const message = {
      ...baseQueryMembershipTokenStatusQResponse,
    } as QueryMembershipTokenStatusQResponse;
    message.MembershipTokenStatus = [];
    if (
      object.MembershipToken !== undefined &&
      object.MembershipToken !== null
    ) {
      message.MembershipToken = MembershipToken.fromJSON(
        object.MembershipToken
      );
    } else {
      message.MembershipToken = undefined;
    }
    if (
      object.MembershipTokenStatus !== undefined &&
      object.MembershipTokenStatus !== null
    ) {
      for (const e of object.MembershipTokenStatus) {
        message.MembershipTokenStatus.push(MembershipTokenStatus.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMembershipTokenStatusQResponse): unknown {
    const obj: any = {};
    message.MembershipToken !== undefined &&
      (obj.MembershipToken = message.MembershipToken
        ? MembershipToken.toJSON(message.MembershipToken)
        : undefined);
    if (message.MembershipTokenStatus) {
      obj.MembershipTokenStatus = message.MembershipTokenStatus.map((e) =>
        e ? MembershipTokenStatus.toJSON(e) : undefined
      );
    } else {
      obj.MembershipTokenStatus = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMembershipTokenStatusQResponse>
  ): QueryMembershipTokenStatusQResponse {
    const message = {
      ...baseQueryMembershipTokenStatusQResponse,
    } as QueryMembershipTokenStatusQResponse;
    message.MembershipTokenStatus = [];
    if (
      object.MembershipToken !== undefined &&
      object.MembershipToken !== null
    ) {
      message.MembershipToken = MembershipToken.fromPartial(
        object.MembershipToken
      );
    } else {
      message.MembershipToken = undefined;
    }
    if (
      object.MembershipTokenStatus !== undefined &&
      object.MembershipTokenStatus !== null
    ) {
      for (const e of object.MembershipTokenStatus) {
        message.MembershipTokenStatus.push(
          MembershipTokenStatus.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of DiscountTokens items. */
  DiscountTokens(
    request: QueryDiscountTokensRequest
  ): Promise<QueryDiscountTokensResponse>;
  /** Queries a DiscountTokenStatus by id. */
  DiscountTokenStatus(
    request: QueryGetDiscountTokenStatusRequest
  ): Promise<QueryGetDiscountTokenStatusResponse>;
  /** Queries a list of DiscountTokenStatus items. */
  DiscountTokenStatusAll(
    request: QueryAllDiscountTokenStatusRequest
  ): Promise<QueryAllDiscountTokenStatusResponse>;
  /** Queries a list of DiscountTokenStatusQ items. */
  DiscountTokenStatusQ(
    request: QueryDiscountTokenStatusQRequest
  ): Promise<QueryDiscountTokenStatusQResponse>;
  /** Queries a list of MembershipTokens items. */
  MembershipTokens(
    request: QueryMembershipTokensRequest
  ): Promise<QueryMembershipTokensResponse>;
  /** Queries a MembershipTokenStatus by id. */
  MembershipTokenStatus(
    request: QueryGetMembershipTokenStatusRequest
  ): Promise<QueryGetMembershipTokenStatusResponse>;
  /** Queries a list of MembershipTokenStatus items. */
  MembershipTokenStatusAll(
    request: QueryAllMembershipTokenStatusRequest
  ): Promise<QueryAllMembershipTokenStatusResponse>;
  /** Queries a list of MembershipTokenStatusQ items. */
  MembershipTokenStatusQ(
    request: QueryMembershipTokenStatusQRequest
  ): Promise<QueryMembershipTokenStatusQResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  DiscountTokens(
    request: QueryDiscountTokensRequest
  ): Promise<QueryDiscountTokensResponse> {
    const data = QueryDiscountTokensRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "DiscountTokens",
      data
    );
    return promise.then((data) =>
      QueryDiscountTokensResponse.decode(new Reader(data))
    );
  }

  DiscountTokenStatus(
    request: QueryGetDiscountTokenStatusRequest
  ): Promise<QueryGetDiscountTokenStatusResponse> {
    const data = QueryGetDiscountTokenStatusRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "DiscountTokenStatus",
      data
    );
    return promise.then((data) =>
      QueryGetDiscountTokenStatusResponse.decode(new Reader(data))
    );
  }

  DiscountTokenStatusAll(
    request: QueryAllDiscountTokenStatusRequest
  ): Promise<QueryAllDiscountTokenStatusResponse> {
    const data = QueryAllDiscountTokenStatusRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "DiscountTokenStatusAll",
      data
    );
    return promise.then((data) =>
      QueryAllDiscountTokenStatusResponse.decode(new Reader(data))
    );
  }

  DiscountTokenStatusQ(
    request: QueryDiscountTokenStatusQRequest
  ): Promise<QueryDiscountTokenStatusQResponse> {
    const data = QueryDiscountTokenStatusQRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "DiscountTokenStatusQ",
      data
    );
    return promise.then((data) =>
      QueryDiscountTokenStatusQResponse.decode(new Reader(data))
    );
  }

  MembershipTokens(
    request: QueryMembershipTokensRequest
  ): Promise<QueryMembershipTokensResponse> {
    const data = QueryMembershipTokensRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "MembershipTokens",
      data
    );
    return promise.then((data) =>
      QueryMembershipTokensResponse.decode(new Reader(data))
    );
  }

  MembershipTokenStatus(
    request: QueryGetMembershipTokenStatusRequest
  ): Promise<QueryGetMembershipTokenStatusResponse> {
    const data = QueryGetMembershipTokenStatusRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "MembershipTokenStatus",
      data
    );
    return promise.then((data) =>
      QueryGetMembershipTokenStatusResponse.decode(new Reader(data))
    );
  }

  MembershipTokenStatusAll(
    request: QueryAllMembershipTokenStatusRequest
  ): Promise<QueryAllMembershipTokenStatusResponse> {
    const data = QueryAllMembershipTokenStatusRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "MembershipTokenStatusAll",
      data
    );
    return promise.then((data) =>
      QueryAllMembershipTokenStatusResponse.decode(new Reader(data))
    );
  }

  MembershipTokenStatusQ(
    request: QueryMembershipTokenStatusQRequest
  ): Promise<QueryMembershipTokenStatusQResponse> {
    const data = QueryMembershipTokenStatusQRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mitoblockchaindev.mitoblockchaindev.Query",
      "MembershipTokenStatusQ",
      data
    );
    return promise.then((data) =>
      QueryMembershipTokenStatusQResponse.decode(new Reader(data))
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
