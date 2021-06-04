part of blockchain;

/// The miner is used for an account to mine new blocks.
class Miner implements Cancelable {
  bool _mining;
  int _prevNonce;

  /// The address of the wallet this miner belongs to.
  final ByteData address;

  /// Creates a new block miner for the given [address].
  Miner(this.address) {
    _mining = false;
    _prevNonce = -1;
  }

  /// This will generate the next nonce to use while mining.
  /// This can be overridden to allow for multiple machines to take different
  /// subspaces of the integer space search through or be randomized.
  /// Since here we are usually only starting only one miner per address
  /// and per block we can simply step the nonce forward without repeating hashes.
  int generateNewNonce() {
    _prevNonce++;
    return _prevNonce;
  }

  /// This will modify the nonce and rehash the given block until the
  /// difficulty challange has been reached. The given block will be modified.
  /// The future will return the block if a solution is found, null if cancelled.
  Future<Block> mine(Block block) async {
    _mining = true;
    block.minerAddress = address;
    while (_mining) {
      var success = await Future<bool>.delayed(miningStepPause, () {
        for (int i = 0; i < miningStepSize; i++) {
          block.nonce = generateNewNonce();
          block.hash = block.calculateHash();

          if (!_mining) return false;
          if (block.hash.startsWith(difficulty)) return true;
        }
        return false;
      });
      if (success) return block;
    }
    return null;
  }

  /// Indicates if a mine is currently in progress.
  bool get mining => _mining;

  /// This will stop the current mine if one is running.
  void cancel() => _mining = false;
}
