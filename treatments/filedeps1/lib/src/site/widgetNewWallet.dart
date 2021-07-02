import 'dart:html';

import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';
import 'widget.dart';

/// A widget for adding new wallets to the driver.
class WidgetNewWallet implements Widget {
  final CallBack _callBack;
  final DivElement _group;
  final InputElement _newWalletName;

  /// Creates a new wallet widget.
  WidgetNewWallet._(this._callBack, this._group, this._newWalletName, InputElement addButton) {
    addButton.onClick.listen(_onNewWalletAdd);
  }

  /// Creates a new wallet widget.
  factory WidgetNewWallet(CallBack callBack) {
    final text = new DivElement()
      ..innerText = 'Enter a name and click add to create a new wallet:'
      ..style.marginBottom = '4px';

    final newWalletName = new InputElement()
      ..type = 'text'
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final addButton = new InputElement()
      ..type = 'submit'
      ..value = 'Add';

    final nameDiv = new DivElement()
      ..style.display = 'flex'
      ..append(newWalletName)
      ..append(addButton);

    final group = new DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    return WidgetNewWallet._(callBack, group, newWalletName, addButton);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Handles when the add new wallet button is clicked.
  void _onNewWalletAdd(_) {
    final name = sani.escape(_newWalletName.value ?? '');
    if (name.isEmpty) return;
    _callBack.addNewWallet(name);
    _newWalletName.value = '';
  }
}
