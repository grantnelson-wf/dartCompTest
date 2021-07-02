part of site;

/// A widget for adding new pending transactions to the chain.
class WidgetTransaction implements Widget {
  final CallBack _callBack;
  final DivElement _group;
  final SelectElement _transactFromName;
  final SelectElement _transactToName;
  final InputElement _transactAmount;

  /// Creates a new transaction widget.
  WidgetTransaction._(this._callBack, this._group, this._transactFromName, this._transactToName, this._transactAmount,
      InputElement addButton) {
    addButton.onClick.listen(_onMakeTransaction);
  }

  /// Creates a new transaction widget.
  factory WidgetTransaction(CallBack callBack) {
    final text = new DivElement()
      ..innerText = 'Create a new transaction which will pend until the next block. ' +
          'The wallets must be different and the amount must be greater than zero:'
      ..style.marginBottom = '4px';

    final fromText = new DivElement()
      ..innerText = 'From:'
      ..style.marginRight = '2px';

    final transactFromName = new SelectElement()
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final toText = new DivElement()
      ..innerText = 'To:'
      ..style.marginRight = '2px';

    final transactToName = new SelectElement()
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final transactAmount = new InputElement()
      ..type = 'number'
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final addButton = new InputElement()
      ..type = 'submit'
      ..value = 'Transact';

    final nameDiv = new DivElement()
      ..style.display = 'flex'
      ..append(fromText)
      ..append(transactFromName)
      ..append(toText)
      ..append(transactToName)
      ..append(transactAmount)
      ..append(addButton);

    final group = new DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    return WidgetTransaction._(callBack, group, transactFromName, transactToName, transactAmount, addButton);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Handles when the transact button is clicked.
  void _onMakeTransaction(_) {
    final fromName = _transactFromName.value ?? '';
    final toName = _transactToName.value ?? '';
    final amount = double.tryParse(_transactAmount.value ?? '-1');
    if (amount == null || amount <= 0.0) return;
    _callBack.newTransaction(fromName, toName, amount);
    _transactAmount.value = '';
  }

  /// Updates the names for the wallets with the given set of wallets.
  void updateWalletNames(UnmodifiableListView<Wallet> wallets) {
    List<OptionElement> fromOptions = [];
    List<OptionElement> toOptions = [];
    for (Wallet wallet in wallets) {
      final name = wallet.name;
      fromOptions.add(OptionElement()
        ..value = name
        ..label = name);
      toOptions.add(OptionElement()
        ..value = name
        ..label = name);
    }

    final fromValue = _transactFromName.value ?? '';
    _transactFromName.children.clear();
    _transactFromName.children.addAll(fromOptions);
    if (fromValue.isNotEmpty) _transactFromName.value = fromValue;

    final toValue = _transactToName.value ?? '';
    _transactToName.children.clear();
    _transactToName.children.addAll(toOptions);
    if (toValue.isNotEmpty) _transactToName.value = toValue;
  }
}
