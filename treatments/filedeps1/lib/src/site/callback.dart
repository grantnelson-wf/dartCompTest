import 'package:FileDepsTreatment/blockchain.dart';

/// This defines the interface that the page can use to call into the driver.
class CallBack {
  /// Adds a new wallet with the given name.
  void addNewWallet(String name) {}

  /// Starts all the wallets mining.
  void startMining() {}

  /// Cancel all the wallets mining.
  void cancelMining() {}

  /// Creates a new pending transaction between the wallets with the given names.
  void newTransaction(String fromName, String toName, double amount) {}

  /// Gets the name of a wallet given the address.
  String nameForAddress(ByteData address) => 'Unknown';
}
