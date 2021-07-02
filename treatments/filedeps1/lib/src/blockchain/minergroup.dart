import 'package:event/event.dart' as event;

import 'blockchain.dart';
import 'bytedata.dart';
import 'cancelable.dart';
import 'miner.dart';

/// A collection of miners which are running together to try to
/// mine the next chain. Obviously in a real block chain this
/// would want to be broken across many machines.
class MinerGroup implements Cancelable {
  List<Miner> _miners;
  event.Event _onChanged;

  /// The chain that this miner group is working on.
  final BlockChain chain;

  /// Creates a new miner group for the given chain.
  MinerGroup(this.chain)
      : _miners = [],
        _onChanged = event.Event() {
    this.chain.onNewBlock + _onNewBlock;
  }

  /// Starts mining the new block for the chain.
  /// If null is returned then there are no transactions to mine,
  /// otherwise the running miner is returned which can be cancelled.
  Future<Cancelable?> start(ByteData minerAddress) async {
    final block = chain.nextBlock;
    if (block == null) return null;

    final miner = Miner(minerAddress);
    _miners.add(miner);
    _onChanged.broadcast();

    // Kick off mining but do not await the result.
    miner.mine(block).then(chain.appendBlock);
    return miner;
  }

  /// Indicates that the number of miners has changed.
  event.Event get onChanged => _onChanged;

  /// Indicates if there are any miners running.
  bool get isMining => _miners.isNotEmpty;

  /// Gets the number of miners running.
  int get minerCount => _miners.length;

  /// Handles when a new block is added to the chain.
  void _onNewBlock(event.EventArgs? _) => cancel();

  /// Cancels all running miners.
  /// Miners will finish current attempt before quitting.
  void cancel() {
    for (Miner miner in _miners) {
      miner.cancel();
    }
    _miners.clear();
    _onChanged.broadcast();
  }
}
