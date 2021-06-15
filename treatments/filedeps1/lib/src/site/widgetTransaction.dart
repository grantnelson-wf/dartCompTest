import 'dart:collection';
import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';

import 'callback.dart';
import 'widget.dart';

/// A widget for adding new pending transactions to the chain.
class WidgetTransaction implements Widget {
  final CallBack _callBack;
  DivElement _group;
  SelectElement _transactFromName;
  SelectElement _transactToName;
  InputElement _transactAmount;

  /// Creates a new transaction widget.
  WidgetTransaction(this._callBack) {
    final text = new DivElement()
      ..innerText = 'Create a new transaction which will pend until the next block. ' +
          'The wallets must be different and the amount must be greater than zero:'
      ..style.marginBottom = '4px';

    final fromText = new DivElement()
      ..innerText = 'From:'
      ..style.marginRight = '2px';

    _transactFromName = new SelectElement()
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final toText = new DivElement()
      ..innerText = 'To:'
      ..style.marginRight = '2px';

    _transactToName = new SelectElement()
      ..style.width = '200px'
      ..style.marginRight = '4px';

    _transactAmount = new InputElement()
      ..type = 'number'
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final addButton = new InputElement()
      ..type = 'submit'
      ..value = 'Transact'
      ..onClick.listen(_onMakeTransaction);

    final nameDiv = new DivElement()
      ..style.display = 'flex'
      ..append(fromText)
      ..append(_transactFromName)
      ..append(toText)
      ..append(_transactToName)
      ..append(_transactAmount)
      ..append(addButton);

    _group = new DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Handles when the transact button is clicked.
  void _onMakeTransaction(_) {
    final fromName = _transactFromName.value;
    final toName = _transactToName.value;
    final amount = double.tryParse(_transactAmount.value);
    if (amount == null || amount <= 0.0) return;
    _callBack.newTransaction(fromName, toName, amount);
    _transactAmount.value = '';
  }

  /// Updates the names for the wallets with the given set of wallets.
  void updateWalletNames(UnmodifiableListView<Wallet> wallets) {
    final fromOptions = List<OptionElement>();
    final toOptions = List<OptionElement>();
    for (Wallet wallet in wallets) {
      final name = wallet.name;
      fromOptions.add(OptionElement()
        ..value = name
        ..label = name);
      toOptions.add(OptionElement()
        ..value = name
        ..label = name);
    }

    final fromValue = _transactFromName.value;
    _transactFromName.children.clear();
    _transactFromName.children.addAll(fromOptions);
    if (fromValue.isNotEmpty) _transactFromName.value = fromValue;

    final toValue = _transactToName.value;
    _transactToName.children.clear();
    _transactToName.children.addAll(toOptions);
    if (toValue.isNotEmpty) _transactToName.value = toValue;
  }
}
