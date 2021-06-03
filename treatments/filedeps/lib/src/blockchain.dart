import 'package:cryptography/cryptography.dart';
import 'package:event/event.dart';

import 'block.dart';
import 'constants.dart';
import 'miner.dart';
import 'transaction.dart';
import 'cancelable.dart';

/// A basic block chain.
class BlockChain {
  List<Block> _chain;
  List<Transaction> _pending;
  Event _onGrowth;

  /// Creates a new block chain with the given settings.
  BlockChain() {
    _chain = List<Block>();
    _pending = List<Transaction>();
    _onGrowth = Event();
  }

  /// This event is fired when a block is added to the chain.
  /// see https://pub.dev/packages/event
  Event get onGrowth => _onGrowth;

  /// Determines the current balance for the wallet with the given address.
  /// This does not include any pending transactions.
  double balance(String address) {
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
  List<String> allAddresses() {
    final addresses = Set<String>();
    for (Block block in _chain) {
      for (Transaction transaction in block.transactions) {
        addresses.add(transaction.toAddress);
        addresses.add(transaction.fromAddress);
      }
      addresses.add(block.minerAddress);
    }
    return List<String>.from(addresses);
  }

  /// Creates a new transaction and adds it to the pending transactions.
  /// Returns true if the transaction was added, false if not.
  Future<bool> createTransaction(SimpleKeyPair fromKeys, String toAddress, double amount) async {
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
    return true;
  }

  /// Gets the previous hash value or an empty string if chain is empty.
  String get _previousHash => _chain.last?.hash ?? '';

  /// Constructs the next block to start mining.
  ///
  /// Usually the miners get to choose which set of pending transactions
  /// should be included since they can't all be included in the block.
  /// With actual block chains overdraw protection should be inforced too.
  /// Transactions must be valid to be put into pending, so they should be still valid.
  Block get _nextBlock => _pending.isEmpty ? null : new Block(_previousHash, _pending);

  /// Constructs and starts mining the new block.
  /// If null is returned then there are no transactions to mine,
  /// otherwise the running miner is returned which can be cancelled.
  Future<Cancelable> mineNextBlock(String minerAddress) async {
    final block = _nextBlock;
    final miner = Miner(minerAddress);
    miner.mine(block).then(_appendBlock);
    return miner;
  }

  /// Appends the given block into the list if it is
  /// valid and follows the previous block.
  void _appendBlock(Block block) async {
    if (block == null) return;
    if (block.previousHash != _previousHash) return;
    if (!await block.isValid) return;

    // Block is accepted onto the chain.
    _chain.add(block);

    // Remove any pending transactions which were in the block.
    for (Transaction transaction in block.transactions) {
      if (_pending.contains(transaction)) _pending.remove(transaction);
    }

    // Notify the change.
    _onGrowth.broadcast();
  }

  /// Determines if the block chain is valid or not.
  Future<bool> get isValid async {
    var previousHash = '';
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
