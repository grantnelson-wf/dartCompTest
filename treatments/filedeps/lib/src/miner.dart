import 'dart:math';

import 'block.dart';
import 'bytedata.dart';
import 'cancelable.dart';
import 'constants.dart';

/// The miner is used for an account to mine new blocks.
class Miner implements Cancelable {
  bool _mining;

  /// The address of the wallet this miner belongs to.
  final ByteData address;

  /// Creates a new block miner for the given [address].
  Miner(this.address) {
    _mining = false;
  }

  /// This will generate the next nonce to use while mining.
  /// This can be overridden to allow for multiple machines to take different
  /// subspaces of the integer space search through.
  int generateNewNonce() => Random().nextInt(maxNonce);

  /// Computes a single step of the mine.
  Future _grind(Block block) async {
    block.nonce = generateNewNonce();
    block.hash = block.calculateHash();
  }

  /// This will modify the nonce and rehash the given block until the
  /// difficulty challange has been reached. The given block will be modified.
  /// The future will return the block if a solution is found, null if cancelled.
  Future<Block> mine(Block block) async {
    _mining = true;
    var attempts = 0; // TODO: REMOVE
    block.minerAddress = address;
    while (_mining) {
      await _grind(block);
      print(">mine: $address => ${block.hash}"); // TODO: REMOVE
      if (block.hash.startsWith(difficulty)) return block;
      attempts++; // TODO: REMOVE
      if (attempts > 100) break; // TODO: REMOVE
    }
    return null;
  }

  /// Indicates if a mine is currently in progress.
  bool get mining => _mining;

  /// This will stop the current mine if one is running.
  void cancel() {
    print(">stop: $address"); // TODO: REMOVE
    _mining = false;
  }
}
