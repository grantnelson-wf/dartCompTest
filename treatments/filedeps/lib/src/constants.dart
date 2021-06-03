import 'package:crypto/crypto.dart' as crypto;
import 'package:cryptography/cryptography.dart' as cryptography;

/// The maximum integer which can be used for the block nonce.
const int maxNonce = 2 ^ 53 - 1;

/// The difficulty string that the miner has to match to successfully mine the block.
const String difficulty = '000';

/// The amount to transfer to the miner's account as a reward for finding the block.
const double minersReward = 1.0;

/// The algorithm to hash values with.
/// see https://pub.dev/documentation/crypto/latest/crypto/sha256-constant.html
final crypto.Hash hashAlgorithm = crypto.sha256;

/// The algorithm to sign and validate with.
/// see https://pub.dev/documentation/cryptography/1.4.1/#digital-signature-with-ed25519
final cryptography.SignatureAlgorithm signatureAlgorithm = cryptography.ed25519;
