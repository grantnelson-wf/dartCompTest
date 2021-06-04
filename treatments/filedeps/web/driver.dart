import 'dart:collection';

import 'package:FileDepsTreatment/blockchain.dart';
import 'package:event/event.dart' as event;

import 'callback.dart';
import 'page.dart';

class Driver implements CallBack {
  final Page _page;
  List<Wallet> _wallets;
  BlockChain _chain;
  MinerGroup _minerGroup;

  Driver(this._page) {
    _wallets = List<Wallet>();
    _chain = BlockChain();
    _chain.onNewBlock + _onNewBlock;
    _chain.onNewTransaction + _onNewTransaction;
    _minerGroup = MinerGroup(_chain);
  }

  UnmodifiableListView<Wallet> get wallets => UnmodifiableListView(_wallets);

  Map<Wallet, double> get balances {
    final result = Map<Wallet, double>();
    for (Wallet wallet in _wallets) {
      result[wallet] = wallet.balance(_chain);
    }
    return result;
  }

  Wallet findWallet(String name) {
    for (Wallet wallet in _wallets) {
      if (wallet.name == name) return wallet;
    }
    return null;
  }

  /// Adds a new wallet with the given name.
  Future addNewWallet(String name) async {
    _wallets.add(await Wallet.create(name));
    _page.updateWalletNames(wallets);
    _page.updateBalances(balances);
  }

  /// Starts all the wallets mining.
  void startMining() {
    List<Wallet> shuffled = List<Wallet>.from(_wallets)..shuffle();
    for (Wallet wallet in shuffled) {
      _minerGroup.start(wallet.address);
    }
  }

  /// Cancel all the wallets mining.
  void cancelMining() => _minerGroup.cancel();

  /// Creates a new pending transaction between the wallets with the given names.
  void newTransaction(String fromName, String toName, double amount) =>
      findWallet(fromName)?.transfer(_chain, findWallet(toName), amount);

  /// Handles when a new block is added to the chain.
  void _onNewBlock(event.EventArgs _) => _page.updateChain(_chain.chain);

  /// Handles when a new block is added to the chain.
  void _onNewTransaction(event.EventArgs _) => _page.updatePending(_chain.pending);
}
