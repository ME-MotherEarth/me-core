package core

// NodeType represents the node's role in the network
type NodeType string

// NodeTypeObserver signals that a node is running as observer node
const NodeTypeObserver NodeType = "observer"

// NodeTypeValidator signals that a node is running as validator node
const NodeTypeValidator NodeType = "validator"

// pkPrefixSize specifies the max numbers of chars to be displayed from one publc key
const pkPrefixSize = 12

// FileModeUserReadWrite represents the permission for a file which allows the user for reading and writing
const FileModeUserReadWrite = 0600

// FileModeReadWrite represents the permission for a file which allows reading and writing for user and group and read
// for others
const FileModeReadWrite = 0664

// MetachainShardId will be used to identify a shard ID as metachain
const MetachainShardId = uint32(0xFFFFFFFF)

// AllShardId will be used to identify that a message is for all shards
const AllShardId = uint32(0xFFFFFFF0)

// MegabyteSize represents the size in bytes of a megabyte
const MegabyteSize = 1024 * 1024

// BuiltInFunctionClaimDeveloperRewards is the key for the claim developer rewards built-in function
const BuiltInFunctionClaimDeveloperRewards = "ClaimDeveloperRewards"

// BuiltInFunctionChangeOwnerAddress is the key for the change owner built in function built-in function
const BuiltInFunctionChangeOwnerAddress = "ChangeOwnerAddress"

// BuiltInFunctionSetUserName is the key for the set user name built-in function
const BuiltInFunctionSetUserName = "SetUserName"

// BuiltInFunctionSaveKeyValue is the key for the save key value built-in function
const BuiltInFunctionSaveKeyValue = "SaveKeyValue"

// BuiltInFunctionMECTTransfer is the key for the Mother Earth Core Token transfer built-in function
const BuiltInFunctionMECTTransfer = "MECTTransfer"

// BuiltInFunctionMECTBurn is the key for the Mother Earth Core Token burn built-in function
const BuiltInFunctionMECTBurn = "MECTBurn"

// BuiltInFunctionMECTFreeze is the key for the Mother Earth Core Token freeze built-in function
const BuiltInFunctionMECTFreeze = "MECTFreeze"

// BuiltInFunctionMECTUnFreeze is the key for the Mother Earth Core Token unfreeze built-in function
const BuiltInFunctionMECTUnFreeze = "MECTUnFreeze"

// BuiltInFunctionMECTWipe is the key for the Mother Earth Core Token wipe built-in function
const BuiltInFunctionMECTWipe = "MECTWipe"

// BuiltInFunctionMECTPause is the key for the Mother Earth Core Token pause built-in function
const BuiltInFunctionMECTPause = "MECTPause"

// BuiltInFunctionMECTUnPause is the key for the Mother Earth Core Token unpause built-in function
const BuiltInFunctionMECTUnPause = "MECTUnPause"

// BuiltInFunctionSetMECTRole is the key for the Mother Earth Core Token set built-in function
const BuiltInFunctionSetMECTRole = "MECTSetRole"

// BuiltInFunctionUnSetMECTRole is the key for the Mother Earth Core Token unset built-in function
const BuiltInFunctionUnSetMECTRole = "MECTUnSetRole"

// BuiltInFunctionMECTSetLimitedTransfer is the key for the Mother Earth Core Token built-in function which sets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionMECTSetLimitedTransfer = "MECTSetLimitedTransfer"

// BuiltInFunctionMECTUnSetLimitedTransfer is the key for the Mother Earth Core Token built-in function which unsets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionMECTUnSetLimitedTransfer = "MECTUnSetLimitedTransfer"

// BuiltInFunctionMECTLocalMint is the key for the Mother Earth Core Token local mint built-in function
const BuiltInFunctionMECTLocalMint = "MECTLocalMint"

// BuiltInFunctionMECTLocalBurn is the key for the Mother Earth Core Token local burn built-in function
const BuiltInFunctionMECTLocalBurn = "MECTLocalBurn"

// BuiltInFunctionMECTNFTTransfer is the key for the Mother Earth Core Token NFT transfer built-in function
const BuiltInFunctionMECTNFTTransfer = "MECTNFTTransfer"

// BuiltInFunctionMECTNFTCreate is the key for the Mother Earth Core Token NFT create built-in function
const BuiltInFunctionMECTNFTCreate = "MECTNFTCreate"

// BuiltInFunctionMECTNFTAddQuantity is the key for the Mother Earth Core Token NFT add quantity built-in function
const BuiltInFunctionMECTNFTAddQuantity = "MECTNFTAddQuantity"

// BuiltInFunctionMECTNFTCreateRoleTransfer is the key for the Mother Earth Core Token create role transfer function
const BuiltInFunctionMECTNFTCreateRoleTransfer = "MECTNFTCreateRoleTransfer"

// BuiltInFunctionMECTNFTBurn is the key for the Mother Earth Core Token NFT burn built-in function
const BuiltInFunctionMECTNFTBurn = "MECTNFTBurn"

// BuiltInFunctionMECTNFTAddURI is the key for the Mother Earth Core Token NFT add URI built-in function
const BuiltInFunctionMECTNFTAddURI = "MECTNFTAddURI"

// BuiltInFunctionMECTNFTUpdateAttributes is the key for the Mother Earth Core Token NFT update attributes built-in function
const BuiltInFunctionMECTNFTUpdateAttributes = "MECTNFTUpdateAttributes"

// BuiltInFunctionMultiMECTNFTTransfer is the key for the Mother Earth Core Token multi transfer built-in function
const BuiltInFunctionMultiMECTNFTTransfer = "MultiMECTNFTTransfer"

// BuiltInFunctionSetGuardian is the key for setting a guardian built-in function
const BuiltInFunctionSetGuardian = "SetGuardian"

// BuiltInFunctionFreezeAccount is the built-in function key for freezing an account
const BuiltInFunctionFreezeAccount = "FreezeAccount"

// BuiltInFunctionUnfreezeAccount is the built-in function key for unfreezing an account
const BuiltInFunctionUnfreezeAccount = "UnfreezeAccount"

// MECTRoleLocalMint is the constant string for the local role of mint for MECT tokens
const MECTRoleLocalMint = "MECTRoleLocalMint"

// MECTRoleLocalBurn is the constant string for the local role of burn for MECT tokens
const MECTRoleLocalBurn = "MECTRoleLocalBurn"

// MECTRoleNFTCreate is the constant string for the local role of create for MECT NFT tokens
const MECTRoleNFTCreate = "MECTRoleNFTCreate"

// MECTRoleNFTCreateMultiShard is the constant string for the local role of create for MECT NFT tokens multishard
const MECTRoleNFTCreateMultiShard = "MECTRoleNFTCreateMultiShard"

// MECTRoleNFTAddQuantity is the constant string for the local role of adding quantity for existing MECT NFT tokens
const MECTRoleNFTAddQuantity = "MECTRoleNFTAddQuantity"

// MECTRoleNFTBurn is the constant string for the local role of burn for MECT NFT tokens
const MECTRoleNFTBurn = "MECTRoleNFTBurn"

// MECTRoleNFTAddURI is the constant string for the local role of adding a URI for MECT NFT tokens
const MECTRoleNFTAddURI = "MECTRoleNFTAddURI"

// MECTRoleNFTUpdateAttributes is the constant string for the local role of updating attributes for MECT NFT tokens
const MECTRoleNFTUpdateAttributes = "MECTRoleNFTUpdateAttributes"

// MECTRoleTransfer is the constant string for the local role to transfer MECT, only for special tokens
const MECTRoleTransfer = "MECTTransferRole"

// MECTType defines the possible types in case of MECT tokens
type MECTType uint32

const (
	// Fungible defines the token type for MECT fungible tokens
	Fungible MECTType = iota
	// NonFungible defines the token type for MECT non fungible tokens
	NonFungible
)

// FungibleMECT defines the string for the token type of fungible MECT
const FungibleMECT = "FungibleMECT"

// NonFungibleMECT defines the string for the token type of non fungible MECT
const NonFungibleMECT = "NonFungibleMECT"

// SemiFungibleMECT defines the string for the token type of semi fungible MECT
const SemiFungibleMECT = "SemiFungibleMECT"

// MaxRoyalty defines 100% as uint32
const MaxRoyalty = uint32(10000)

// RelayedTransaction is the key for the motherearth meta/gassless/relayed transaction standard
const RelayedTransaction = "relayedTx"

// RelayedTransactionV2 is the key for the optimized motherearth meta/gassless/relayed transaction standard
const RelayedTransactionV2 = "relayedTxV2"

// SCDeployInitFunctionName is the key for the function which is called at smart contract deploy time
const SCDeployInitFunctionName = "_init"

// MotherEarthProtectedKeyPrefix is the key prefix which is protected from writing in the trie - only for special builtin functions
const MotherEarthProtectedKeyPrefix = "MOTHEREARTH"

// DelegationSystemSCKey is the key under which there is data in case of system delegation smart contracts
const DelegationSystemSCKey = "delegation"

// MECTKeyIdentifier is the key prefix for mect tokens
const MECTKeyIdentifier = "mect"

// MECTRoleIdentifier is the key prefix for mect role identifier
const MECTRoleIdentifier = "role"

// MECTNFTLatestNonceIdentifier is the key prefix for mect latest nonce identifier
const MECTNFTLatestNonceIdentifier = "nonce"

// GuardiansKeyIdentifier is the key prefix for guardians
const GuardiansKeyIdentifier = "guardians"

// MaxNumShards represents the maximum number of shards possible in the system
const MaxNumShards = 256

// MinMetaTxExtraGasCost is the constant defined for minimum gas value to be sent in meta transaction
const MinMetaTxExtraGasCost = uint64(1_000_000)

// MaxLeafSize represents maximum amount of data which can be saved under one leaf
const MaxLeafSize = uint64(1 << 26) //64MB

// MaxBufferSizeToSendTrieNodes represents max buffer size to send in bytes used when resolving trie nodes
// Every trie node that has a greater size than this constant is considered a large trie node and should be split in
// smaller chunks
const MaxBufferSizeToSendTrieNodes = 1 << 18 //256KB

// MaxUserNameLength represents the maximum number of bytes a UserName can have
const MaxUserNameLength = 32

// MinLenArgumentsMECTTransfer defines the min length of arguments for the MECT transfer
const MinLenArgumentsMECTTransfer = 2

// MinLenArgumentsMECTNFTTransfer defines the minimum length for mect nft transfer
const MinLenArgumentsMECTNFTTransfer = 4

// MaxLenForMECTIssueMint defines the maximum length in bytes for the issued/minted balance
const MaxLenForMECTIssueMint = 100

// BaseOperationCostString represents the field name for base operation costs
const BaseOperationCostString = "BaseOperationCost"

// BuiltInCostString represents the field name for built in operation costs
const BuiltInCostString = "BuiltInCost"

// MECTSCAddress is the hard-coded address for mect issuing smart contract
var MECTSCAddress = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 255, 255}

// SCDeployIdentifier is the identifier for a smart contract deploy
const SCDeployIdentifier = "SCDeploy"

// SCUpgradeIdentifier is the identifier for a smart contract upgrade
const SCUpgradeIdentifier = "SCUpgrade"
