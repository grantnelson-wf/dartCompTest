import 'dart:html';

import 'package:FileDepsTreatment/site.dart';

void main() {
  document.title = "FileDepsTreatment Block Chain Treatment";

  final site = Site();
  site.addNewWallet('Jim');
  site.addNewWallet('Jill');
  site.addNewWallet('Jack');
  site.addNewWallet('Sal');
}
