import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateMembershipTokenStatus } from "./types/mitoblockchaindev/tx";
import { MsgCreateDiscountToken } from "./types/mitoblockchaindev/tx";
import { MsgCreateMembershipToken } from "./types/mitoblockchaindev/tx";
import { MsgDeleteMembershipTokenStatus } from "./types/mitoblockchaindev/tx";
import { MsgDeleteDiscountTokenStatus } from "./types/mitoblockchaindev/tx";
import { MsgCreateDiscountTokenStatus } from "./types/mitoblockchaindev/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/mitoblockchaindev.mitoblockchaindev.MsgCreateMembershipTokenStatus", MsgCreateMembershipTokenStatus],
    ["/mitoblockchaindev.mitoblockchaindev.MsgCreateDiscountToken", MsgCreateDiscountToken],
    ["/mitoblockchaindev.mitoblockchaindev.MsgCreateMembershipToken", MsgCreateMembershipToken],
    ["/mitoblockchaindev.mitoblockchaindev.MsgDeleteMembershipTokenStatus", MsgDeleteMembershipTokenStatus],
    ["/mitoblockchaindev.mitoblockchaindev.MsgDeleteDiscountTokenStatus", MsgDeleteDiscountTokenStatus],
    ["/mitoblockchaindev.mitoblockchaindev.MsgCreateDiscountTokenStatus", MsgCreateDiscountTokenStatus],
    
];

export { msgTypes }