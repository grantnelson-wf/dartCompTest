part of site;

/// A widget for showing pending transactions as a table.
class WidgetPending implements Widget {
  final CallBack _callBack;
  final DivElement _group;
  final TableElement _pendingTransactions;

  /// Creates the pending transaction widget.
  WidgetPending._(this._callBack, this._group, this._pendingTransactions);

  /// Creates the pending transaction widget.
  factory WidgetPending(CallBack callBack) {
    final text = new DivElement()
      ..innerText = 'The pending transactions which have not been added to the chain yet via mining a block:'
      ..style.marginBottom = '4px';

    final pendingTransactions = new TableElement()
      ..style.marginLeft = '10px'
      ..style.border = 'none'
      ..style.borderCollapse = 'collapse';

    final group = new DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(pendingTransactions);

    return WidgetPending._(callBack, group, pendingTransactions);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Creates a new cell element for the pending transaction table.
  TableCellElement _newCell(String text) => TableCellElement()
    ..text = text
    ..style.paddingLeft = '5px'
    ..style.paddingRight = '5px'
    ..style.borderLeft = '1px solid black'
    ..style.borderRight = '1px solid black';

  /// Updates the list of pending transactions.
  void updatePending(UnmodifiableListView<Transaction> pending) {
    List<TableRowElement> rows = [];

    if (pending.isNotEmpty) {
      final header = TableRowElement()
        ..append(_newCell('timestamp'))
        ..append(_newCell('from'))
        ..append(_newCell('to'))
        ..append(_newCell('amount'))
        ..style.fontWeight = 'bold'
        ..style.borderBottom = '1px solid black';
      rows.add(header);

      for (Transaction transaction in pending) {
        final row = TableRowElement()
          ..append(_newCell(transaction.timestamp.toString()))
          ..append(_newCell(_callBack.nameForAddress(transaction.fromAddress)))
          ..append(_newCell(_callBack.nameForAddress(transaction.toAddress)))
          ..append(_newCell('${transaction.amount}'));
        rows.add(row);
      }
    }

    _pendingTransactions.children.clear();
    _pendingTransactions.children.addAll(rows);
  }
}
