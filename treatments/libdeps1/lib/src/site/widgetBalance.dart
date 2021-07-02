part of site;

/// A widget for showing the balances for every wallet.
class WidgetBalance implements Widget {
  final DivElement _group;
  final TableElement _balanceTable;

  /// Creates a new balance widget.
  WidgetBalance._(this._group, this._balanceTable);

  /// Creates a new balance widget.
  factory WidgetBalance() {
    final text = new DivElement()
      ..innerText = 'The balances for all wallets:'
      ..style.marginBottom = '4px';

    final balanceTable = TableElement()
      ..style.marginLeft = '10px'
      ..style.border = 'none'
      ..style.borderCollapse = 'collapse';

    final group = DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(balanceTable);

    return WidgetBalance._(group, balanceTable);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Updates the balances for each wallet.
  void updateBalances(Map<Wallet, double> balances) {
    List<TableRowElement> rows = [];

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
