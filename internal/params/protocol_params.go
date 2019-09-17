package params

import "math/big"

const (
	// GasLimitBoundDivisor ...
	GasLimitBoundDivisor uint64 = 1024 // The bound divisor of the gas limit, used in update calculations.
	// MinGasLimit ...
	MinGasLimit uint64 = 5000 // Minimum the gas limit may ever be.
	// GenesisGasLimit ...
	GenesisGasLimit uint64 = 4712388 // Gas limit of the Genesis block.

	// MaximumExtraDataSize ...
	MaximumExtraDataSize uint64 = 32 // Maximum size extra data may be after Genesis.
	// ExpByteGas ...
	ExpByteGas uint64 = 10 // Times ceil(log256(exponent)) for the EXP instruction.
	// SloadGas ...
	SloadGas uint64 = 50 // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	// CallValueTransferGas ...
	CallValueTransferGas uint64 = 9000 // Paid for CALL when the value transfer is non-zero.
	// CallNewAccountGas ...
	CallNewAccountGas uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	// TxGas ...
	TxGas uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	// TxGasContractCreation ...
	TxGasContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	// TxDataZeroGas ...
	TxDataZeroGas uint64 = 4 // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	// QuadCoeffDiv ...
	QuadCoeffDiv uint64 = 512 // Divisor for the quadratic particle of the memory cost equation.
	// LogDataGas ...
	LogDataGas uint64 = 8 // Per byte in a LOG* operation's data.
	// CallStipend ...
	CallStipend uint64 = 2300 // Free gas given at beginning of call.

	// Sha3Gas ...
	Sha3Gas uint64 = 30 // Once per SHA3 operation.
	// Sha3WordGas ...
	Sha3WordGas uint64 = 6 // Once per word of the SHA3 operation's data.

	// SstoreSetGas ...
	SstoreSetGas uint64 = 20000 // Once per SLOAD operation.
	// SstoreResetGas ...
	SstoreResetGas uint64 = 5000 // Once per SSTORE operation if the zeroness changes from zero.
	// SstoreClearGas ...
	SstoreClearGas uint64 = 5000 // Once per SSTORE operation if the zeroness doesn't change.
	// SstoreRefundGas ...
	SstoreRefundGas uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.

	// NetSstoreNoopGas ...
	NetSstoreNoopGas uint64 = 200 // Once per SSTORE operation if the value doesn't change.
	// NetSstoreInitGas ...
	NetSstoreInitGas uint64 = 20000 // Once per SSTORE operation from clean zero.
	// NetSstoreCleanGas ...
	NetSstoreCleanGas uint64 = 5000 // Once per SSTORE operation from clean non-zero.
	// NetSstoreDirtyGas ...
	NetSstoreDirtyGas uint64 = 200 // Once per SSTORE operation from dirty.

	// NetSstoreClearRefund ...
	NetSstoreClearRefund uint64 = 15000 // Once per SSTORE operation for clearing an originally existing storage slot
	// NetSstoreResetRefund ...
	NetSstoreResetRefund uint64 = 4800 // Once per SSTORE operation for resetting to the original non-zero value
	// NetSstoreResetClearRefund ...
	NetSstoreResetClearRefund uint64 = 19800 // Once per SSTORE operation for resetting to the original zero value

	// JumpdestGas ...
	JumpdestGas uint64 = 1 // Refunded gas, once per SSTORE operation if the zeroness changes to zero.
	// EpochDuration ...
	EpochDuration uint64 = 30000 // Duration between proof-of-work epochs.
	// CallGas ...
	CallGas uint64 = 40 // Once per CALL operation & message call transaction.
	// CreateDataGas ...
	CreateDataGas uint64 = 200 //
	// CallCreateDepth ...
	CallCreateDepth uint64 = 1024 // Maximum depth of call/create stack.
	// ExpGas ...
	ExpGas uint64 = 10 // Once per EXP instruction
	// LogGas ...
	LogGas uint64 = 375 // Per LOG* operation.
	// CopyGas ...
	CopyGas uint64 = 3 //
	// StackLimit ...
	StackLimit uint64 = 1024 // Maximum size of VM stack allowed.
	// TierStepGas ...
	TierStepGas uint64 = 0 // Once per operation, for a selection of them.
	// LogTopicGas ...
	LogTopicGas uint64 = 375 // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	// CreateGas ...
	CreateGas uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	// Create2Gas ...
	Create2Gas uint64 = 32000 // Once per CREATE2 operation
	// SuicideRefundGas ...
	SuicideRefundGas uint64 = 24000 // Refunded following a suicide operation.
	// MemoryGas ...
	MemoryGas uint64 = 3 // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	// TxDataNonZeroGas ...
	TxDataNonZeroGas uint64 = 68 // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.

	// MaxCodeSize ...
	MaxCodeSize = 24576 // Maximum bytecode to permit for a contract

	// Precompiled contract gas prices

	// EcrecoverGas ...
	EcrecoverGas uint64 = 3000 // Elliptic curve sender recovery gas price
	// Sha256BaseGas ...
	Sha256BaseGas uint64 = 60 // Base price for a SHA256 operation
	// Sha256PerWordGas ...
	Sha256PerWordGas uint64 = 12 // Per-word price for a SHA256 operation
	// Ripemd160BaseGas ...
	Ripemd160BaseGas uint64 = 600 // Base price for a RIPEMD160 operation
	// Ripemd160PerWordGas ...
	Ripemd160PerWordGas uint64 = 120 // Per-word price for a RIPEMD160 operation
	// IdentityBaseGas ...
	IdentityBaseGas uint64 = 15 // Base price for a data copy operation
	// IdentityPerWordGas ...
	IdentityPerWordGas uint64 = 3 // Per-work price for a data copy operation
	// ModExpQuadCoeffDiv ...
	ModExpQuadCoeffDiv uint64 = 20 // Divisor for the quadratic particle of the big int modular exponentiation
	// Bn256AddGas ...
	Bn256AddGas uint64 = 500 // Gas needed for an elliptic curve addition
	// Bn256ScalarMulGas ...
	Bn256ScalarMulGas uint64 = 40000 // Gas needed for an elliptic curve scalar multiplication
	// Bn256PairingBaseGas ...
	Bn256PairingBaseGas uint64 = 100000 // Base price for an elliptic curve pairing check
	// Bn256PairingPerPointGas ...
	Bn256PairingPerPointGas uint64 = 80000 // Per-point price for an elliptic curve pairing check
)

var (
	// DifficultyBoundDivisor ...
	DifficultyBoundDivisor = big.NewInt(2048) // The bound divisor of the difficulty, used in the update calculations.
	// GenesisDifficulty ...
	GenesisDifficulty = big.NewInt(131072) // Difficulty of the Genesis block.
	// MinimumDifficulty ...
	MinimumDifficulty = big.NewInt(131072) // The minimum that the difficulty may ever be.
	// DurationLimit ...
	DurationLimit = big.NewInt(13) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
)
