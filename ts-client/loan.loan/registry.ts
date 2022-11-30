import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { MsgApproveLoan } from "./types/loan/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    
];

export { msgTypes }