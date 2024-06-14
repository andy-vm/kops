// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a unique customer managed [KMS key] in your Amazon Web Services account and
// Region. You can use a KMS key in cryptographic operations, such as encryption
// and signing. Some Amazon Web Services services let you use KMS keys that you
// create and manage to protect your service resources.
//
// A KMS key is a logical representation of a cryptographic key. In addition to
// the key material used in cryptographic operations, a KMS key includes metadata,
// such as the key ID, key policy, creation date, description, and key state. For
// details, see [Managing keys]in the Key Management Service Developer Guide
//
// Use the parameters of CreateKey to specify the type of KMS key, the source of
// its key material, its key policy, description, tags, and other properties.
//
// KMS has replaced the term customer master key (CMK) with KMS key and KMS key.
// The concept has not changed. To prevent breaking changes, KMS is keeping some
// variations of this term.
//
// To create different types of KMS keys, use the following guidance:
//
// Symmetric encryption KMS key By default, CreateKey creates a symmetric
// encryption KMS key with key material that KMS generates. This is the basic and
// most widely used type of KMS key, and provides the best performance.
//
// To create a symmetric encryption KMS key, you don't need to specify any
// parameters. The default value for KeySpec , SYMMETRIC_DEFAULT , the default
// value for KeyUsage , ENCRYPT_DECRYPT , and the default value for Origin ,
// AWS_KMS , create a symmetric encryption KMS key with KMS key material.
//
// If you need a key for basic encryption and decryption or you are creating a KMS
// key to protect your resources in an Amazon Web Services service, create a
// symmetric encryption KMS key. The key material in a symmetric encryption key
// never leaves KMS unencrypted. You can use a symmetric encryption KMS key to
// encrypt and decrypt data up to 4,096 bytes, but they are typically used to
// generate data keys and data keys pairs. For details, see GenerateDataKeyand GenerateDataKeyPair.
//
// Asymmetric KMS keys To create an asymmetric KMS key, use the KeySpec parameter
// to specify the type of key material in the KMS key. Then, use the KeyUsage
// parameter to determine whether the KMS key will be used to encrypt and decrypt
// or sign and verify. You can't change these properties after the KMS key is
// created.
//
// Asymmetric KMS keys contain an RSA key pair, Elliptic Curve (ECC) key pair, or
// an SM2 key pair (China Regions only). The private key in an asymmetric KMS key
// never leaves KMS unencrypted. However, you can use the GetPublicKeyoperation to download
// the public key so it can be used outside of KMS. Each KMS key can have only one
// key usage. KMS keys with RSA key pairs can be used to encrypt and decrypt data
// or sign and verify messages (but not both). KMS keys with NIST-recommended ECC
// key pairs can be used to sign and verify messages or derive shared secrets (but
// not both). KMS keys with ECC_SECG_P256K1 can be used only to sign and verify
// messages. KMS keys with SM2 key pairs (China Regions only) can be used to either
// encrypt and decrypt data, sign and verify messages, or derive shared secrets
// (you must choose one key usage type). For information about asymmetric KMS keys,
// see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// HMAC KMS key To create an HMAC KMS key, set the KeySpec parameter to a key spec
// value for HMAC KMS keys. Then set the KeyUsage parameter to GENERATE_VERIFY_MAC
// . You must set the key usage even though GENERATE_VERIFY_MAC is the only valid
// key usage value for HMAC KMS keys. You can't change these properties after the
// KMS key is created.
//
// HMAC KMS keys are symmetric keys that never leave KMS unencrypted. You can use
// HMAC keys to generate (GenerateMac ) and verify (VerifyMac ) HMAC codes for messages up to 4096
// bytes.
//
// Multi-Region primary keys Imported key material To create a multi-Region
// primary key in the local Amazon Web Services Region, use the MultiRegion
// parameter with a value of True . To create a multi-Region replica key, that is,
// a KMS key with the same key ID and key material as a primary key, but in a
// different Amazon Web Services Region, use the ReplicateKeyoperation. To change a replica
// key to a primary key, and its primary key to a replica key, use the UpdatePrimaryRegionoperation.
//
// You can create multi-Region KMS keys for all supported KMS key types: symmetric
// encryption KMS keys, HMAC KMS keys, asymmetric encryption KMS keys, and
// asymmetric signing KMS keys. You can also create multi-Region keys with imported
// key material. However, you can't create multi-Region keys in a custom key store.
//
// This operation supports multi-Region keys, an KMS feature that lets you create
// multiple interoperable KMS keys in different Amazon Web Services Regions.
// Because these KMS keys have the same key ID, key material, and other metadata,
// you can use them interchangeably to encrypt data in one Amazon Web Services
// Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// To import your own key material into a KMS key, begin by creating a KMS key
// with no key material. To do this, use the Origin parameter of CreateKey with a
// value of EXTERNAL . Next, use GetParametersForImport operation to get a public key and import token.
// Use the wrapping public key to encrypt your key material. Then, use ImportKeyMaterialwith your
// import token to import the key material. For step-by-step instructions, see [Importing Key Material]in
// the Key Management Service Developer Guide .
//
// You can import key material into KMS keys of all supported KMS key types:
// symmetric encryption KMS keys, HMAC KMS keys, asymmetric encryption KMS keys,
// and asymmetric signing KMS keys. You can also create multi-Region keys with
// imported key material. However, you can't import key material into a KMS key in
// a custom key store.
//
// To create a multi-Region primary key with imported key material, use the Origin
// parameter of CreateKey with a value of EXTERNAL and the MultiRegion parameter
// with a value of True . To create replicas of the multi-Region primary key, use
// the ReplicateKeyoperation. For instructions, see [Importing key material into multi-Region keys]. For more information about multi-Region
// keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// Custom key store A [custom key store] lets you protect your Amazon Web Services resources using
// keys in a backing key store that you own and manage. When you request a
// cryptographic operation with a KMS key in a custom key store, the operation is
// performed in the backing key store using its cryptographic keys.
//
// KMS supports [CloudHSM key stores] backed by an CloudHSM cluster and [external key stores] backed by an external key
// manager outside of Amazon Web Services. When you create a KMS key in an CloudHSM
// key store, KMS generates an encryption key in the CloudHSM cluster and
// associates it with the KMS key. When you create a KMS key in an external key
// store, you specify an existing encryption key in the external key manager.
//
// Some external key managers provide a simpler method for creating a KMS key in
// an external key store. For details, see your external key manager documentation.
//
// Before you create a KMS key in a custom key store, the ConnectionState of the
// key store must be CONNECTED . To connect the custom key store, use the ConnectCustomKeyStore
// operation. To find the ConnectionState , use the DescribeCustomKeyStores operation.
//
// To create a KMS key in a custom key store, use the CustomKeyStoreId . Use the
// default KeySpec value, SYMMETRIC_DEFAULT , and the default KeyUsage value,
// ENCRYPT_DECRYPT to create a symmetric encryption key. No other key type is
// supported in a custom key store.
//
// To create a KMS key in an [CloudHSM key store], use the Origin parameter with a value of
// AWS_CLOUDHSM . The CloudHSM cluster that is associated with the custom key store
// must have at least two active HSMs in different Availability Zones in the Amazon
// Web Services Region.
//
// To create a KMS key in an [external key store], use the Origin parameter with a value of
// EXTERNAL_KEY_STORE and an XksKeyId parameter that identifies an existing
// external key.
//
// Some external key managers provide a simpler method for creating a KMS key in
// an external key store. For details, see your external key manager documentation.
//
// Cross-account use: No. You cannot use this operation to create a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:CreateKey] (IAM policy). To use the Tags parameter, [kms:TagResource] (IAM policy).
// For examples and information about related permissions, see [Allow a user to create KMS keys]in the Key
// Management Service Developer Guide.
//
// Related operations:
//
// # DescribeKey
//
// # ListKeys
//
// # ScheduleKeyDeletion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [Managing keys]: https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html
// [KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms-keys
// [Allow a user to create KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies.html#iam-policy-example-create-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
// [kms:TagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [kms:CreateKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Importing key material into multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-import.html
// [Importing Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html
func (c *Client) CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) {
	if params == nil {
		params = &CreateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKey", params, optFns, c.addOperationCreateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyInput struct {

	// Skips ("bypasses") the key policy lockout safety check. The default value is
	// false.
	//
	// Setting this value to true increases the risk that the KMS key becomes
	// unmanageable. Do not set this value to true indiscriminately.
	//
	// For more information, see [Default key policy] in the Key Management Service Developer Guide.
	//
	// Use this parameter only when you intend to prevent the principal that is making
	// the request from making a subsequent [PutKeyPolicy]request on the KMS key.
	//
	// [Default key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key
	// [PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html
	BypassPolicyLockoutSafetyCheck bool

	// Creates the KMS key in the specified [custom key store]. The ConnectionState of the custom key
	// store must be CONNECTED . To find the CustomKeyStoreID and ConnectionState use
	// the DescribeCustomKeyStoresoperation.
	//
	// This parameter is valid only for symmetric encryption KMS keys in a single
	// Region. You cannot create any other type of KMS key in a custom key store.
	//
	// When you create a KMS key in an CloudHSM key store, KMS generates a
	// non-exportable 256-bit symmetric key in its associated CloudHSM cluster and
	// associates it with the KMS key. When you create a KMS key in an external key
	// store, you must use the XksKeyId parameter to specify an external key that
	// serves as key material for the KMS key.
	//
	// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html
	CustomKeyStoreId *string

	// Instead, use the KeySpec parameter.
	//
	// The KeySpec and CustomerMasterKeySpec parameters work the same way. Only the
	// names differ. We recommend that you use KeySpec parameter in your code.
	// However, to avoid breaking changes, KMS supports both parameters.
	//
	// Deprecated: This parameter has been deprecated. Instead, use the KeySpec
	// parameter.
	CustomerMasterKeySpec types.CustomerMasterKeySpec

	// A description of the KMS key. Use a description that helps you decide whether
	// the KMS key is appropriate for a task. The default value is an empty string (no
	// description).
	//
	// Do not include confidential or sensitive information in this field. This field
	// may be displayed in plaintext in CloudTrail logs and other output.
	//
	// To set or change the description after the key is created, use UpdateKeyDescription.
	Description *string

	// Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT ,
	// creates a KMS key with a 256-bit AES-GCM key that is used for encryption and
	// decryption, except in China Regions, where it creates a 128-bit symmetric key
	// that uses SM4 encryption. For help choosing a key spec for your KMS key, see [Choosing a KMS key type]in
	// the Key Management Service Developer Guide .
	//
	// The KeySpec determines whether the KMS key contains a symmetric key or an
	// asymmetric key pair. It also determines the algorithms that the KMS key
	// supports. You can't change the KeySpec after the KMS key is created. To further
	// restrict the algorithms that can be used with the KMS key, use a condition key
	// in its key policy or IAM policy. For more information, see [kms:EncryptionAlgorithm], [kms:MacAlgorithm] or [kms:Signing Algorithm] in the Key
	// Management Service Developer Guide .
	//
	// [Amazon Web Services services that are integrated with KMS]use symmetric encryption KMS keys to protect your data. These services do not
	// support asymmetric KMS keys or HMAC KMS keys.
	//
	// KMS supports the following key specs for KMS keys:
	//
	//   - Symmetric encryption key (default)
	//
	//   - SYMMETRIC_DEFAULT
	//
	//   - HMAC keys (symmetric)
	//
	//   - HMAC_224
	//
	//   - HMAC_256
	//
	//   - HMAC_384
	//
	//   - HMAC_512
	//
	//   - Asymmetric RSA key pairs (encryption and decryption -or- signing and
	//   verification)
	//
	//   - RSA_2048
	//
	//   - RSA_3072
	//
	//   - RSA_4096
	//
	//   - Asymmetric NIST-recommended elliptic curve key pairs (signing and
	//   verification -or- deriving shared secrets)
	//
	//   - ECC_NIST_P256 (secp256r1)
	//
	//   - ECC_NIST_P384 (secp384r1)
	//
	//   - ECC_NIST_P521 (secp521r1)
	//
	//   - Other asymmetric elliptic curve key pairs (signing and verification)
	//
	//   - ECC_SECG_P256K1 (secp256k1), commonly used for cryptocurrencies.
	//
	//   - SM2 key pairs (encryption and decryption -or- signing and verification -or-
	//   deriving shared secrets)
	//
	//   - SM2 (China Regions only)
	//
	// [kms:EncryptionAlgorithm]: https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm
	// [kms:Signing Algorithm]: https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm
	// [kms:MacAlgorithm]: https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm
	// [Choosing a KMS key type]: https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose
	// [Amazon Web Services services that are integrated with KMS]: http://aws.amazon.com/kms/features/#AWS_Service_Integration
	KeySpec types.KeySpec

	// Determines the [cryptographic operations] for which you can use the KMS key. The default value is
	// ENCRYPT_DECRYPT . This parameter is optional when you are creating a symmetric
	// encryption KMS key; otherwise, it is required. You can't change the KeyUsage
	// value after the KMS key is created.
	//
	// Select only one valid value.
	//
	//   - For symmetric encryption KMS keys, omit the parameter or specify
	//   ENCRYPT_DECRYPT .
	//
	//   - For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC .
	//
	//   - For asymmetric KMS keys with RSA key pairs, specify ENCRYPT_DECRYPT or
	//   SIGN_VERIFY .
	//
	//   - For asymmetric KMS keys with NIST-recommended elliptic curve key pairs,
	//   specify SIGN_VERIFY or KEY_AGREEMENT .
	//
	//   - For asymmetric KMS keys with ECC_SECG_P256K1 key pairs specify SIGN_VERIFY .
	//
	//   - For asymmetric KMS keys with SM2 key pairs (China Regions only), specify
	//   ENCRYPT_DECRYPT , SIGN_VERIFY , or KEY_AGREEMENT .
	//
	// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations
	KeyUsage types.KeyUsageType

	// Creates a multi-Region primary key that you can replicate into other Amazon Web
	// Services Regions. You cannot change this value after you create the KMS key.
	//
	// For a multi-Region key, set this parameter to True . For a single-Region KMS
	// key, omit this parameter or set it to False . The default value is False .
	//
	// This operation supports multi-Region keys, an KMS feature that lets you create
	// multiple interoperable KMS keys in different Amazon Web Services Regions.
	// Because these KMS keys have the same key ID, key material, and other metadata,
	// you can use them interchangeably to encrypt data in one Amazon Web Services
	// Region and decrypt it in a different Amazon Web Services Region without
	// re-encrypting the data or making a cross-Region call. For more information about
	// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
	//
	// This value creates a primary key, not a replica. To create a replica key, use
	// the ReplicateKeyoperation.
	//
	// You can create a symmetric or asymmetric multi-Region key, and you can create a
	// multi-Region key with imported key material. However, you cannot create a
	// multi-Region key in a custom key store.
	//
	// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
	MultiRegion *bool

	// The source of the key material for the KMS key. You cannot change the origin
	// after you create the KMS key. The default is AWS_KMS , which means that KMS
	// creates the key material.
	//
	// To [create a KMS key with no key material] (for imported key material), set this value to EXTERNAL . For more
	// information about importing key material into KMS, see [Importing Key Material]in the Key Management
	// Service Developer Guide. The EXTERNAL origin value is valid only for symmetric
	// KMS keys.
	//
	// To [create a KMS key in an CloudHSM key store] and create its key material in the associated CloudHSM cluster, set this
	// value to AWS_CLOUDHSM . You must also use the CustomKeyStoreId parameter to
	// identify the CloudHSM key store. The KeySpec value must be SYMMETRIC_DEFAULT .
	//
	// To [create a KMS key in an external key store], set this value to EXTERNAL_KEY_STORE . You must also use the
	// CustomKeyStoreId parameter to identify the external key store and the XksKeyId
	// parameter to identify the associated external key. The KeySpec value must be
	// SYMMETRIC_DEFAULT .
	//
	// [create a KMS key in an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keys.html
	// [create a KMS key in an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html
	// [Importing Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
	// [create a KMS key with no key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html
	Origin types.OriginType

	// The key policy to attach to the KMS key.
	//
	// If you provide a key policy, it must meet the following criteria:
	//
	//   - The key policy must allow the calling principal to make a subsequent
	//   PutKeyPolicy request on the KMS key. This reduces the risk that the KMS key
	//   becomes unmanageable. For more information, see [Default key policy]in the Key Management Service
	//   Developer Guide. (To omit this condition, set BypassPolicyLockoutSafetyCheck
	//   to true.)
	//
	//   - Each statement in the key policy must contain one or more principals. The
	//   principals in the key policy must exist and be visible to KMS. When you create a
	//   new Amazon Web Services principal, you might need to enforce a delay before
	//   including the new principal in a key policy because the new principal might not
	//   be immediately visible to KMS. For more information, see [Changes that I make are not always immediately visible]in the Amazon Web
	//   Services Identity and Access Management User Guide.
	//
	// If you do not provide a key policy, KMS attaches a default key policy to the
	// KMS key. For more information, see [Default key policy]in the Key Management Service Developer
	// Guide.
	//
	// The key policy size quota is 32 kilobytes (32768 bytes).
	//
	// For help writing and formatting a JSON policy document, see the [IAM JSON Policy Reference] in the
	// Identity and Access Management User Guide .
	//
	// [IAM JSON Policy Reference]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
	// [Default key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default
	// [Changes that I make are not always immediately visible]: https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency
	Policy *string

	// Assigns one or more tags to the KMS key. Use this parameter to tag the KMS key
	// when it is created. To tag an existing KMS key, use the TagResourceoperation.
	//
	// Do not include confidential or sensitive information in this field. This field
	// may be displayed in plaintext in CloudTrail logs and other output.
	//
	// Tagging or untagging a KMS key can allow or deny permission to the KMS key. For
	// details, see [ABAC for KMS]in the Key Management Service Developer Guide.
	//
	// To use this parameter, you must have [kms:TagResource] permission in an IAM policy.
	//
	// Each tag consists of a tag key and a tag value. Both the tag key and the tag
	// value are required, but the tag value can be an empty (null) string. You cannot
	// have more than one tag on a KMS key with the same tag key. If you specify an
	// existing tag key with a different tag value, KMS replaces the current tag value
	// with the specified one.
	//
	// When you add tags to an Amazon Web Services resource, Amazon Web Services
	// generates a cost allocation report with usage and costs aggregated by tags. Tags
	// can also be used to control access to a KMS key. For details, see [Tagging Keys].
	//
	// [kms:TagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
	// [Tagging Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
	// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
	Tags []types.Tag

	// Identifies the [external key] that serves as key material for the KMS key in an [external key store]. Specify the
	// ID that the [external key store proxy]uses to refer to the external key. For help, see the documentation
	// for your external key store proxy.
	//
	// This parameter is required for a KMS key with an Origin value of
	// EXTERNAL_KEY_STORE . It is not valid for KMS keys with any other Origin value.
	//
	// The external key must be an existing 256-bit AES symmetric encryption key
	// hosted outside of Amazon Web Services in an external key manager associated with
	// the external key store specified by the CustomKeyStoreId parameter. This key
	// must be enabled and configured to perform encryption and decryption. Each KMS
	// key in an external key store must use a different external key. For details, see
	// [Requirements for a KMS key in an external key store]in the Key Management Service Developer Guide.
	//
	// Each KMS key in an external key store is associated two backing keys. One is
	// key material that KMS generates. The other is the external key specified by this
	// parameter. When you use the KMS key in an external key store to encrypt data,
	// the encryption operation is performed first by KMS using the KMS key material,
	// and then by the external key manager using the specified external key, a process
	// known as double encryption. For details, see [Double encryption]in the Key Management Service
	// Developer Guide.
	//
	// [external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
	// [Double encryption]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-double-encryption
	// [external key]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-external-key
	// [Requirements for a KMS key in an external key store]: https://docs.aws.amazon.com/create-xks-keys.html#xks-key-requirements
	// [external key store proxy]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-xks-proxy
	XksKeyId *string

	noSmithyDocumentSerde
}

type CreateKeyOutput struct {

	// Metadata associated with the KMS key.
	KeyMetadata *types.KeyMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKey"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpCreateKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKey",
	}
}
