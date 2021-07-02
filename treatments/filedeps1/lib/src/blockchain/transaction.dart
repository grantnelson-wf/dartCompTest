import 'dart:convert';

import 'package:cryptography/cryptography.dart' as cryptography;

import 'bytedata.dart';
import 'constants.dart';

// The description of the transfer of some amount from one address to another.
class Transaction {
  /// This is when this transaction was created.
  final DateTime timestamp;

  /// The address to take the amount from.
  /// This address must have signed this transaction to be valid.
  final ByteData fromAddress;

  /// The address to give the amount to.
  final ByteData toAddress;

  /// The amount being transfered between the addresses.
  final double amount;

  /// The signature of the from address to ensure this
  /// transaction was approved of by the person giving money (fromAddress).
  ByteData _signature;

  /// Creates an existing transaction with the given values.
  Transaction(this.timestamp, this.fromAddress, this.toAddress, this.amount, [ByteData? signature = null])
      : _signature = signature ?? ByteData.empty();

  /// Creates a new transaction and signs it for the wallet with the given key pair.
  static Future<Transaction> createAndSign(
      cryptography.SimpleKeyPair fromKeys, ByteData toAddress, double amount) async {
    final fromAddress = ByteData((await fromKeys.extractPublicKey()).bytes);
    final transaction = new Transaction(DateTime.now(), fromAddress, toAddress, amount);
    final signature = await signatureAlgorithm.sign(transaction.serialize(), keyPair: fromKeys);
    transaction._signature = ByteData(signature.bytes);
    return transaction;
  }

  /// The signature for this transaction.
  ByteData get signature => _signature;

  /// Serializes this transaction into bytes.
  /// This will NOT serialize the signature since it is used when signing,
  /// meaning this can NOT be used for transmiting transactions to other chains.
  List<int> serialize() => []
    ..addAll(utf8.encode(timestamp.toIso8601String()))
    ..addAll(fromAddress.bytes)
    ..addAll(toAddress.bytes)
    ..addAll(utf8.encode(amount.toString()));

  /// Indicates that this transaction and the signature is valid.
  Future<bool> get isValid async {
    if (amount <= 0.0 || toAddress == fromAddress) return false;

    final key = new cryptography.SimplePublicKey(fromAddress.bytes, type: signatureAlgorithm.keyPairType);
    final signed = cryptography.Signature(signature.bytes, publicKey: key);
    return signatureAlgorithm.verify(serialize(), signature: signed);
  }

  /// Gets the human readable string for debugging.
  @override
  String toString() => '$timestamp, $fromAddress, $toAddress, $amount';

  /// Used to determine if the given [other] is equivalent to this transaction.
  @override
  bool operator ==(Object other) =>
      other is Transaction &&
      other.timestamp == timestamp &&
      other.fromAddress == fromAddress &&
      other.toAddress == toAddress &&
      other.amount == amount && // should use an epsilon comparer if serializing via string.
      other._signature == _signature;
}
