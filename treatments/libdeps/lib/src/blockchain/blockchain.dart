library blockchain;

import 'dart:collection';
import 'dart:convert';
import 'dart:async';

import 'package:cryptography/cryptography.dart' as cryptography;
import 'package:crypto/crypto.dart' as crypto;
import 'package:event/event.dart' as event;

part 'block.dart';
part 'bytedata.dart';
part 'cancelable.dart';
part 'constants.dart';
part 'miner.dart';
part 'minergroup.dart';
part 'transaction.dart';
part 'wallet.dart';

/// A basic block chain.
///
/// This block chain doesn't do any voting among divergent chains or
/// even transmit new blocks and transactions to other chains.
/// This is just a stand-alone example block chain.
class BlockChain {
  List<Block> _chain;
  List<Transaction> _pending;
  event.Event _onNewTransaction;
  event.Event _onNewBlock;

  /// Creates a new block chain with the given settings.
  BlockChain() {
    _chain = List<Block>();
    _pending = List<Transaction>();
    _onNewTransaction = event.Event();
    _onNewBlock = event.Event();
  }

  /// This event is fired when a transaction is added to the pending transactions.
  event.Event get onNewTransaction => _onNewTransaction;

  /// This event is fired when a block is added to the chain.
  /// This also means the list of transactions has been reduced.
  event.Event get onNewBlock => _onNewBlock;

  /// Gets the set of pending transactions.
  UnmodifiableListView<Transaction> get pending => UnmodifiableListView(_pending);

  /// Gets the set of block chain set of chains.
  UnmodifiableListView<Block> get chain => UnmodifiableListView(_chain);

  /// Determines the current balance for the wallet with the given address.
  /// This does not include any pending transactions.
  double balance(ByteData address) {
    var amount = 0.0;
    for (Block block in _chain) {
      for (Transaction transaction in block.transactions) {
        if (transaction.toAddress == address) amount += transaction.amount;
        if (transaction.fromAddress == address) amount -= transaction.amount;
      }
      if (block.minerAddress == address) amount += minersReward;
    }
    return amount;
  }

  /// Gets the list of all addresses that have transactions in the chain.
  /// This does not include any pending transactions.
  List<ByteData> allAddresses() {
    final addresses = Set<ByteData>();
    for (Block block in _chain) {
      for (Transaction transaction in block.transactions) {
        addresses.add(transaction.toAddress);
        addresses.add(transaction.fromAddress);
      }
      addresses.add(block.minerAddress);
    }
    return List<ByteData>.from(addresses);
  }

  /// Creates a new transaction and adds it to the pending transactions.
  /// Returns true if the transaction was added, false if not.
  Future<bool> createTransaction(cryptography.KeyPair fromKeys, ByteData toAddress, double amount) async {
    final transaction = await Transaction.createAndSign(fromKeys, toAddress, amount);
    return addTransaction(transaction);
  }

  /// Adds a new transaction to the pending transactions.
  /// Returns true if the transaction was added, false if not.
  Future<bool> addTransaction(Transaction transaction) async {
    if (transaction == null) return false;
    if (!await transaction.isValid) return false;
    if (_pending.contains(transaction)) return false;

    // Transaction is accepted into pending
    _pending.add(transaction);

    // Notify the change.
    _onNewTransaction.broadcast();
    return true;
  }

  /// Gets the previous hash value or an empty string if chain is empty.
  ByteData get _previousHash => _chain.isNotEmpty ? _chain.last?.hash : ByteData([]);

  /// Constructs the next block to start mining.
  ///
  /// Usually the miners get to choose which set of pending transactions
  /// should be included since they can't all be included in the block.
  /// With actual block chains overdraw protection should be inforced too.
  /// Transactions must be valid to be put into pending, so they should be still valid.
  Block get nextBlock => _pending.isEmpty ? null : new Block(_previousHash, List<Transaction>.from(_pending));

  /// Appends the given block into the list if it is
  /// valid and follows the previous block.
  void appendBlock(Block block) async {
    if (block == null) return;
    if (!await block.isValid) return;
    if (block.previousHash != _previousHash) return;

    // Block is accepted onto the chain.
    _chain.add(block);

    // Remove any pending transactions which were in the block.
    for (Transaction transaction in block.transactions) {
      if (_pending.contains(transaction)) _pending.remove(transaction);
    }

    // Notify the change.
    _onNewBlock.broadcast();
  }

  /// Determines if the block chain is valid or not.
  Future<bool> get isValid async {
    ByteData previousHash;
    for (Block block in _chain) {
      if (block.previousHash == previousHash) return false;
      if (!await block.isValid) return false;
      previousHash = block.previousHash;
    }
    return true;
  }

  /// Gets a human readable string of this block chain for debugging.
  @override
  String toString() {
    final buffer = new StringBuffer();
    buffer.write('chain:\n');
    for (Block block in _chain) {
      final indented = block.toString().replaceAll('\n', '\n   ');
      buffer.write('   $indented\n');
    }
    buffer.write('pending:\n');
    for (Transaction transaction in _pending) {
      buffer.write('   $transaction\n');
    }
    return buffer.toString();
  }
}
