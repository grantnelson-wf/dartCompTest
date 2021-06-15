import 'dart:html';

import 'package:LibDepsTreatment/site.dart';

void main() {
  document.title = "LibDepsTreatment Block Chain Treatment";

  final site = Site();
  site.addNewWallet('Jim');
  site.addNewWallet('Jill');
  site.addNewWallet('Jack');
  site.addNewWallet('Sal');
}
