import 'dart:collection';
import 'dart:convert';

import 'package:crypto/crypto.dart';

import 'transaction.dart';

/// A block stores a single block in the chain.
/// It contains a set of transactions since the prior chain.
class Block {
  DateTime _timestamp;
  List<Transaction> _transactions;
  String _previousHash;
  String _hash;
  int _nonce;
  String _minerAddress;

  /// Creates a new block with the given previous block's hash
  /// and the transactions for this block.
  Block(String prevHash, [List<Transaction> transactions]) {
    _timestamp = DateTime.now();
    _transactions = transactions ?? List<Transaction>();
    _previousHash = prevHash;
    _hash = null;
    _nonce = 0;
    _minerAddress = '';
  }

  /// Timestamp is the time this block was created at.
  DateTime get timestamp => _timestamp;

  /// Gets the list of the transactions.
  UnmodifiableListView<Transaction> get transactions => _transactions;

  /// Gets the previous block's hash value.
  String get previousHash => _previousHash;

  /// Gets or sets this block's current hash value.
  String get hash => _hash;
  set hash(String hash) => _hash = hash;

  /// Gets or sets the address of the account/wallet which mined this block.
  String get minerAddress => _minerAddress;
  set minerAddress(String address) => _minerAddress = address;

  /// Gets or sets this block's nonce value.
  int get nonce => _nonce;
  set nonce(int nonce) => _nonce = nonce;

  /// Calculates the hash for this whole block,
  /// excluding the hash value itself.
  String calculateHash() {
    final output = AccumulatorSink<Digest>();
    final input = hashAlgorithm.startChunkedConversion(output)
      ..add(utf8.encode(_previousHash))
      ..add(utf8.encode(_timestamp.toIso8601String()))
      ..add(utf8.encode(_nonce.toString()))
      ..add(utf8.encode(_minerAddress));
    for (Transaction transaction in _transactions) {
      input.add(transaction.serialize());
    }
    input.close();
    return output.events.single;
  }

  /// Determines if this block, its hash, and transactions are valid.
  Future<bool> get isValid async {
    for (Transaction transaction in _transactions) {
      if (!await transaction.isValid) return false;
    }
    return calculateHash() == _hash;
  }

  /// Gets a human readable string of this block for debugging.
  @override
  String toString() {
    final buffer = new StringBuffer()
      ..write('timestamp: $timestamp\n')
      ..write('previousHash: $previousHash\n')
      ..write('hash: $hash\n')
      ..write('nonce: $nonce\n')
      ..write('minerAddress: $minerAddress\n')
      ..write('transactions:\n');
    for (Transaction transaction in _transactions) {
      buffer.write('   $transaction\n');
    }
    return buffer.toString();
  }
}
