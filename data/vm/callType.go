package vm

// CallType specifies the type of SC invocation (in terms of asynchronicity)
type CallType int

const (
	// DirectCall means that the call is an explicit SC invocation originating from a user Transaction
	DirectCall CallType = iota

	// AsynchronousCall means that the invocation was performed from within
	// another SmartContract from another Shard, using asyncCall
	AsynchronousCall

	// AsynchronousCallBack means that an AsynchronousCall was performed
	// previously, and now the control returns to the caller SmartContract's callBack method
	AsynchronousCallBack

	// MECTTransferAndExecute means that there is a smart contract execution after the MECT transfer
	// this is needed in order to skip the check whether a contract is payable or not
	MECTTransferAndExecute

	// ExecOnDestByCaller means that the call is an invocation of a built in function / smart contract from
	// another smart contract but the caller is from the previous caller
	ExecOnDestByCaller
)
