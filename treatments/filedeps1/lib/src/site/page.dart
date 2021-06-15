import 'dart:collection';
import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';

import 'callback.dart';
import 'widgetBalance.dart';
import 'widgetChain.dart';
import 'widgetMining.dart';
import 'widgetNewWallet.dart';
import 'widgetPending.dart';
import 'widgetTransaction.dart';

/// The page handles a collection of widgets for the block chain example site.
class Page {
  WidgetNewWallet _newWallet;
  WidgetTransaction _transaction;
  WidgetMining _mining;
  WidgetBalance _balance;
  WidgetPending _pending;
  WidgetChain _chain;

  /// Creates a new page instance but does not setup anything yet.
  Page() {}

  /// Sets up the page with the callback to the driver.
  void setupPage(CallBack callBack) {
    _newWallet = WidgetNewWallet(callBack);
    _transaction = WidgetTransaction(callBack);
    _mining = WidgetMining(callBack);
    _balance = WidgetBalance(callBack);
    _pending = WidgetPending(callBack);
    _chain = WidgetChain(callBack);

    document.body
      ..style.backgroundColor = 'lightgrey'
      ..append(_newWallet.widget)
      ..append(_transaction.widget)
      ..append(_mining.widget)
      ..append(_balance.widget)
      ..append(_pending.widget)
      ..append(_chain.widget);
  }

  /// Updates the names for the wallets with the given set of wallets.
  void updateWalletNames(UnmodifiableListView<Wallet> wallets) => _transaction?.updateWalletNames(wallets);

  /// Updates the balances for each wallet.
  void updateBalances(Map<Wallet, double> balances) => _balance?.updateBalances(balances);

  /// Updates the mining state to indicate if mining is running or not.
  void updateMiningState(bool mining) => _mining?.updateMiningState(mining);

  /// Updates the list of pending transactions.
  void updatePending(UnmodifiableListView<Transaction> pending) => _pending?.updatePending(pending);

  /// Updates the display fo the block chain.
  void updateChain(UnmodifiableListView<Block> chain) => _chain?.updateChain(chain);
}
