import 'dart:convert';

import 'package:cryptography/cryptography.dart';

import 'constants.dart';

// The description of the transfer of some amount from one address to another.
class Transaction {
  /// This is when this transaction was created.
  final DateTime timestamp;

  /// The address to take the amount from.
  /// This address must have signed this transaction to be valid.
  final String fromAddress;

  /// The address to give the amount to.
  final String toAddress;

  /// The amount being transfered between the addresses.
  final double amount;

  /// The signature of the from address to ensure this
  /// transaction was approved of by the person giving money (fromAddress).
  String _signature;

  /// Creates an existing transaction with the given values.
  Transaction(this.timestamp, this.fromAddress, this.toAddress, this.amount, [String signature = '']) {
    _signature = signature;
  }

  /// Creates a new transaction and signs it for the wallet with the given key pair.
  static Future<Transaction> createAndSign(SimpleKeyPair fromKeys, String toAddress, double amount) async {
    final fromAddress = String.fromCharCodes(fromKeys.publicKey.bytes);
    final transaction = new Transaction(DateTime.now(), fromAddress, toAddress, amount);
    transaction._signature = await signatureAlgorithm.sign(transaction.serialize(), fromKeys);
    return transaction;
  }

  /// The signature for this transaction.
  String get signature => _signature;

  /// Serializes this transaction into bytes.
  List<int> serialize() => List<int>()
    ..addAll(utf8.encode(timestamp.toIso8601String()))
    ..addAll(utf8.encode(fromAddress))
    ..addAll(utf8.encode(toAddress))
    ..addAll(utf8.encode(amount.toString()));

  /// Indicates that this transaction and the signature is valid.
  Future<bool> get isValid async {
    if (fromAddress.isEmpty || amount <= 0.0 || signature.isEmpty || toAddress == fromAddress) return false;
    return signatureAlgorithm.verify(serialize(), signature: signature);
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
