import 'dart:html';

import 'page.dart';
import 'data.dart';

void main() {
  document.title = "FileDepsTreatment Block Chain Treatment";

  final page = Page();

  final data = Data(page);
  data.addNewWallet('Jim');
  data.addNewWallet('Jill');
  data.addNewWallet('Jim');
  data.addNewWallet('Sal');

  page.setupPage(data);
}
