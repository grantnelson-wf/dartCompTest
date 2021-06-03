import 'package:FileDepsTreatment/blockchain.dart';

import 'callback.dart';
import 'page.dart';

class Data implements CallBack {
  final Page _page;
  List<Wallet> _wallets;

  Data(this._page) {
    _wallets = List<Wallet>();
  }

  Iterable<String> get _walletNames sync* {
    for (Wallet wallet in _wallets) {
      yield wallet.name;
    }
  }

  Future addNewWallet(String name) async {
    _wallets.add(await Wallet.create(name));
    _page.newWalletNames(_walletNames);
  }

  void startMining() {}

  void newTransaction(String fromName, String toName, double amount) {}
}
