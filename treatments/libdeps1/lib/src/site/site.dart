library site;

import 'dart:collection';
import 'dart:html';

import 'package:LibDepsTreatment/blockchain.dart';
import 'package:event/event.dart' as event;
import 'package:validators/sanitizers.dart' as sani;

part 'callback.dart';
part 'page.dart';
part 'widget.dart';
part 'widgetBalance.dart';
part 'widgetChain.dart';
part 'widgetMining.dart';
part 'widgetNewWallet.dart';
part 'widgetPending.dart';
part 'widgetTransaction.dart';

// This is the main data and event handler for running the block chain example site.
class Site implements CallBack {
  Page _page;
  List<Wallet> _wallets;
  BlockChain _chain;
  MinerGroup _minerGroup;

  /// Creates a new driver for the block chain example site.
  Site._(this._chain, this._minerGroup)
      : _page = Page(),
        _wallets = [] {
    _chain.onNewBlock + _onNewBlock;
    _chain.onNewTransaction + _onNewTransaction;
    _minerGroup.onChanged + _onMinersChanged;
    _page.setupPage(this);
  }

  /// Creates a new driver for the block chain example site.
  factory Site() {
    final chain = BlockChain();
    final minerGroup = MinerGroup(chain);
    return Site._(chain, minerGroup);
  }

  /// Gets the wallets being used in this chain.
  UnmodifiableListView<Wallet> get wallets => UnmodifiableListView(_wallets);

  /// Gets the balance for each wallet.
  Map<Wallet, double> get balances {
    final result = Map<Wallet, double>();
    for (Wallet wallet in _wallets) {
      result[wallet] = wallet.balance(_chain);
    }
    return result;
  }

  /// Find a wallet with the given name.
  Wallet? findWallet(String name) {
    for (Wallet wallet in _wallets) {
      if (wallet.name == name) return wallet;
    }
    return null;
  }

  /// Gets the name of a wallet given the address.
  String nameForAddress(ByteData address) {
    for (Wallet wallet in _wallets) {
      if (wallet.address == address) return wallet.name;
    }
    return 'Unknown';
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
  void newTransaction(String fromName, String toName, double amount) {
    var toWallet = findWallet(toName);
    if (toWallet == null) return;
    findWallet(fromName)?.transfer(_chain, toWallet, amount);
  }

  /// Handles when a new block is added to the chain.
  /// This also indicates the pending transactions and balances might have changed.
  void _onNewBlock(event.EventArgs? _) {
    _page.updateChain(_chain.chain);
    _page.updatePending(_chain.pending);
    _page.updateBalances(balances);
  }

  /// Handles when a new block is added to the chain.
  void _onNewTransaction(event.EventArgs? _) => _page.updatePending(_chain.pending);

  /// Handles when the mining group has changed state.
  void _onMinersChanged(event.EventArgs? _) => _page.updateMiningState(_minerGroup.isMining);
}
