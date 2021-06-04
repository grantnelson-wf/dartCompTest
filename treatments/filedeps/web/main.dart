import 'dart:html';

import 'page.dart';
import 'driver.dart';

void main() {
  document.title = "FileDepsTreatment Block Chain Treatment";

  final page = Page();

  final driver = Driver(page);
  driver.addNewWallet('Jim');
  driver.addNewWallet('Jill');
  driver.addNewWallet('Jack');
  driver.addNewWallet('Sal');

  page.setupPage(driver);
  page.updateWalletNames(driver.wallets);
  page.updateBalances(driver.balances);
}
