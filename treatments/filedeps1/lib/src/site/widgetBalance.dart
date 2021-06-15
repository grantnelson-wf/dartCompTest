import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';

import 'callback.dart';
import 'widget.dart';

/// A widget for showing the balances for every wallet.
class WidgetBalance implements Widget {
  final CallBack _callBack;
  DivElement _group;
  TableElement _balanceTable;

  /// Creates a new balance widget.
  WidgetBalance(this._callBack) {
    final text = new DivElement()
      ..innerText = 'The balances for all wallets:'
      ..style.marginBottom = '4px';

    _balanceTable = TableElement()
      ..style.marginLeft = '10px'
      ..style.border = 'none'
      ..style.borderCollapse = 'collapse';

    _group = DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(_balanceTable);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Updates the balances for each wallet.
  void updateBalances(Map<Wallet, double> balances) {
    final rows = List<TableRowElement>();

    balances.forEach((wallet, amount) {
      final row = TableRowElement()
        ..append(TableCellElement()
          ..text = '${wallet.name}:' // already escaped
          ..style.fontWeight = 'bold'
          ..style.paddingLeft = '5px'
          ..style.paddingRight = '5px')
        ..append(TableCellElement()
          ..text = '$amount'
          ..style.paddingLeft = '5px'
          ..style.paddingRight = '5px');
      rows.add(row);
    });

    _balanceTable.children.clear();
    _balanceTable.children.addAll(rows);
  }
}