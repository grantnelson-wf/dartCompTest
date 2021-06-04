import 'dart:collection';
import 'dart:convert';

import 'bytedata.dart';
import 'constants.dart';
import 'transaction.dart';

/// A block stores a single block in the chain.
/// It contains a set of transactions since the prior chain.
class Block {
  DateTime _timestamp;
  List<Transaction> _transactions;
  ByteData _previousHash;
  ByteData _hash;
  int _nonce;
  ByteData _minerAddress;

  /// Creates a new block with the given previous block's hash
  /// and the transactions for this block.
  Block(ByteData prevHash, [List<Transaction> transactions]) {
    _timestamp = DateTime.now();
    _transactions = transactions ?? List<Transaction>();
    _previousHash = prevHash;
    _hash = null;
    _nonce = 0;
    _minerAddress = null;
  }

  /// Timestamp is the time this block was created at.
  DateTime get timestamp => _timestamp;

  /// Gets the list of the transactions.
  UnmodifiableListView<Transaction> get transactions => UnmodifiableListView(_transactions);

  /// Gets the previous block's hash value.
  ByteData get previousHash => _previousHash;

  /// Gets or sets this block's current hash value.
  ByteData get hash => _hash;
  set hash(ByteData hash) => _hash = hash;

  /// Gets or sets the address of the account/wallet which mined this block.
  ByteData get minerAddress => _minerAddress;
  set minerAddress(ByteData address) => _minerAddress = address;

  /// Gets or sets this block's nonce value.
  int get nonce => _nonce;
  set nonce(int nonce) => _nonce = nonce;

  /// Serializes this block into bytes.
  /// This will NOT serialize the hash (nor transaction signatures) since
  /// it is used when calculating the hash, meaning this can NOT be used
  /// for transmiting blocks to other chains.
  List<int> serialize() {
    var data = List<int>()
      ..addAll(_previousHash.bytes)
      ..addAll(utf8.encode(_timestamp.toIso8601String()))
      ..addAll(utf8.encode(_nonce.toString()))
      ..addAll(_minerAddress.bytes);
    for (Transaction transaction in _transactions) {
      data.addAll(transaction.serialize());
    }
    return data;
  }

  /// Calculates the hash for this whole block,
  /// excluding the hash value itself (and transaction signatures).
  ByteData calculateHash() => ByteData(hashAlgorithm.convert(serialize()).bytes);

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
