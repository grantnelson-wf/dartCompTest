import 'dart:html';

import 'package:FileDepsTreatment/site.dart';

void main() {
  document.title = "FileDepsTreatment Block Chain Treatment";

  final driver = Driver();
  driver.addNewWallet('Jim');
  driver.addNewWallet('Jill');
  driver.addNewWallet('Jack');
  driver.addNewWallet('Sal');
}
