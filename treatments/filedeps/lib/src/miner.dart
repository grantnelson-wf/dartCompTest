import 'dart:math';

import 'block.dart';
import 'constants.dart';

/// The miner is used for an account to mine new blocks.
class Miner {
  bool _mining;

  /// The address of the wallet this miner belongs to.
  final String address;

  /// The difficulty string that the miner has to match to successfully mine the block.
  final String difficulty;

  /// Creates a new block miner for the given [address] with the required [difficulty].
  Miner(String this.address, String this.difficulty) {
    _mining = false;
  }

  /// This will generate the next nonce to use while mining.
  /// This can be overridden to allow for multiple machines to take different
  /// subspaces of the integer space search through.
  int generateNewNonce() => Random().nextInt(maxInt);

  /// This will modify the nonce and rehash the given block until the
  /// difficulty challange has been reached. The given block will be modified.
  /// The future will return true if a solution is found, false if cancelled.
  Future<bool> mine(Block block) async {
    _mining = true;
    block.minerAddress = address;
    while (_mining) {
      block.nonce = generateNewNonce();
      block.hash = block.calculateHash();
      if (block.hash.startsWith(difficulty)) return true;
    }
    return false;
  }

  /// Indicates if a mine is currently in progress.
  bool get mining => _mining;

  /// This will stop the current mine if one is running.
  void cancel() => _mining = false;
}
