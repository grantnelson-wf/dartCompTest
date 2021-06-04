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

class Page {
  WidgetNewWallet _newWallet;
  WidgetTransaction _transaction;
  WidgetMining _mining;
  WidgetBalance _balance;
  WidgetPending _pending;
  WidgetChain _chain;

  Page() {}

  void setupPage(CallBack callBack) {
    _newWallet = WidgetNewWallet(callBack);
    _transaction = WidgetTransaction(callBack);
    _mining = WidgetMining(callBack);
    _balance = WidgetBalance(callBack);
    _pending = WidgetPending(callBack);
    _chain = WidgetChain(callBack);

    document.body
      ..append(_newWallet.widget)
      ..append(_transaction.widget)
      ..append(_mining.widget)
      ..append(_balance.widget)
      ..append(_pending.widget)
      ..append(_chain.widget);
  }

  void updateWalletNames(UnmodifiableListView<Wallet> wallets) => _transaction?.updateWalletNames(wallets);

  void updateBalances(Map<Wallet, double> balances) => _balance?.updateBalances(balances);

  void updatePending(UnmodifiableListView<Transaction> pending) => _pending?.updatePending(pending);

  void updateChain(UnmodifiableListView<Block> chain) => _chain?.updateChain(chain);
}
