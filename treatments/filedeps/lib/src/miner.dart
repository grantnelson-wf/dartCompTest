import 'dart:math';

import 'block.dart';
import 'constants.dart';
import 'cancelable.dart';

/// The miner is used for an account to mine new blocks.
class Miner implements Cancelable {
  bool _mining;

  /// The address of the wallet this miner belongs to.
  final String address;

  /// Creates a new block miner for the given [address].
  Miner(this.address) {
    _mining = false;
  }

  /// This will generate the next nonce to use while mining.
  /// This can be overridden to allow for multiple machines to take different
  /// subspaces of the integer space search through.
  int generateNewNonce() => Random().nextInt(maxNonce);

  /// This will modify the nonce and rehash the given block until the
  /// difficulty challange has been reached. The given block will be modified.
  /// The future will return the block if a solution is found, null if cancelled.
  Future<Block> mine(Block block) async {
    _mining = true;
    block.minerAddress = address;
    while (_mining) {
      block.nonce = generateNewNonce();
      block.hash = block.calculateHash();
      if (block.hash.startsWith(difficulty)) return block;
    }
    return null;
  }

  /// Indicates if a mine is currently in progress.
  bool get mining => _mining;

  /// This will stop the current mine if one is running.
  void cancel() => _mining = false;
}
