import 'package:crypto/crypto.dart' as crypto;
import 'package:cryptography/cryptography.dart' as cryptography;

import 'bytedata.dart';

/// The maximum integer which can be used for the block nonce.
const int maxNonce = 2 ^ 53 - 1;

/// The amount to transfer to the miner's account as a reward for finding the block.
const double minersReward = 1.0;

/// The amount of time to delay between mining steps.
const miningStepPause = Duration(microseconds: 1);

/// The number of attempts to take per mining step.
/// This isn't really fair in a real block chain because the mining should be
/// done truely multi-threaded or on different machines, as is it will give
/// the first run miners a slightly higher probability of successfully mining.
const miningStepSize = 100;

/// The difficulty that the miner has to match to successfully mine the block.
/// The resulting hash must start with this value to be valid.
final ByteData difficulty = ByteData.zeros(1);

/// The algorithm to hash values with.
/// see https://pub.dev/documentation/crypto/latest/crypto/sha256-constant.html
final crypto.Hash hashAlgorithm = crypto.sha256;

/// The algorithm to sign and validate with.
/// see https://pub.dev/documentation/cryptography/1.4.1/#digital-signature-with-ed25519
final cryptography.SignatureAlgorithm signatureAlgorithm = cryptography.ed25519;
