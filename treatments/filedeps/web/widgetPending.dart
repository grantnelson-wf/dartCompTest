import 'dart:collection';
import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';
import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';
import 'widget.dart';

class WidgetPending implements Widget {
  final CallBack _callBack;
  DivElement _group;
  DivElement _pendingTransactions;

  WidgetPending(this._callBack) {
    final text = new DivElement()
      ..innerText = 'The pending transactions which have not been added to the chain yet via mining a block:'
      ..style.marginBottom = '4px';

    _pendingTransactions = new DivElement()..style.marginLeft = '10px';

    _group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(_pendingTransactions);
  }
  DivElement get widget => _group;

  void updatePending(UnmodifiableListView<Transaction> pending) {
    final transactions = List<DivElement>();

    for (Transaction transaction in pending) {
      final text = new DivElement()
        ..innerText = sani.escape(transaction.toString())
        ..style.marginBottom = '4px';
      transactions.add(text);
    }

    _pendingTransactions.children.clear();
    _pendingTransactions.children.addAll(transactions);
  }
}
