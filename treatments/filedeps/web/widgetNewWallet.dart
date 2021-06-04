import 'dart:html';

import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';
import 'widget.dart';

class WidgetNewWallet implements Widget {
  final CallBack _callBack;
  DivElement _group;
  InputElement _newWalletName;

  WidgetNewWallet(this._callBack) {
    final text = new DivElement()
      ..innerText = 'Enter a name and click add to create a new wallet:'
      ..style.marginBottom = '4px';

    _newWalletName = new InputElement()
      ..type = 'text'
      ..style.width = '200px'
      ..style.marginRight = '4px';

    final addButton = new InputElement()
      ..type = 'submit'
      ..value = 'Add'
      ..onClick.listen(_onNewWalletAdd);

    final nameDiv = new DivElement()
      ..style.display = 'flex'
      ..append(_newWalletName)
      ..append(addButton);

    _group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);
  }

  DivElement get widget => _group;

  void _onNewWalletAdd(_) {
    final name = sani.escape(_newWalletName.value);
    if (name.isEmpty) return;
    _callBack.addNewWallet(name);
    _newWalletName.value = '';
  }
}
