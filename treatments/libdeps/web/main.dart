import 'dart:html';

import 'package:LibDepsTreatment/site.dart';

void main() {
  document.title = "LibDepsTreatment Block Chain Treatment";

  final driver = site.Driver();
  driver.addNewWallet('Jim');
  driver.addNewWallet('Jill');
  driver.addNewWallet('Jack');
  driver.addNewWallet('Sal');
}
