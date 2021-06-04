import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';
import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';
import 'widget.dart';

class WidgetBalance implements Widget {
  final CallBack _callBack;
  DivElement _group;
  TableElement _balanceTable;

  WidgetBalance(this._callBack) {
    final text = new DivElement()
      ..innerText = 'The balances for all wallets:'
      ..style.marginBottom = '4px';

    _balanceTable = new TableElement()..style.marginLeft = '10px';

    _group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(_balanceTable);
  }

  DivElement get widget => _group;

  void updateBalances(Map<Wallet, double> balances) {
    final rows = List<TableRowElement>();

    final header = TableRowElement()
      ..append(TableCellElement()..text = 'name')
      ..append(TableCellElement()..text = 'amount')
      ..append(TableCellElement()..text = 'address')
      ..style.fontWeight = 'bold';
    rows.add(header);

    balances.forEach((wallet, amount) {
      final row = TableRowElement()
        ..append(TableCellElement()..text = wallet.name) // already escaped
        ..append(TableCellElement()..text = '$amount')
        ..append(TableCellElement()..text = sani.escape(wallet.address.toString()));
      rows.add(row);
    });

    _balanceTable.children.clear();
    _balanceTable.children.addAll(rows);
  }
}
