import 'dart:html';

import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';

class Page {
  CallBack _callBack;
  InputElement _newWalletName;
  SelectElement _transactFromName;
  SelectElement _transactToName;

  Page() {}

  void setupPage(CallBack callBack) {
    _callBack = callBack;
    _setupNewWalletGroup();
    _setupTransactionGroup();
    _setupMiningGroup();
    _setupBalanceGroup();
    _setupPendingGroup();
    _setupChainGroup();
  }

  void _setupNewWalletGroup() {
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

    final group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    document.body.append(group);
  }

  void _setupTransactionGroup() {
    final text = new DivElement()
      ..innerText = 'Create a new transaction which will pend until the next block:'
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

    final addButton = new InputElement()
      ..type = 'submit'
      ..value = 'Transact'
      ..onClick.listen((_) {
        //callBack.addNewWallet(nameInput.value);
        //nameInput.value = '';
      });

    final nameDiv = new DivElement()
      ..style.display = 'flex'
      ..append(fromText)
      ..append(_transactFromName)
      ..append(toText)
      ..append(_transactToName)
      ..append(addButton);

    final group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    document.body.append(group);
  }

  void _setupMiningGroup() {}

  void _setupBalanceGroup() {}

  void _setupPendingGroup() {}

  void _setupChainGroup() {}

  void _onNewWalletAdd(_) {
    final name = sani.escape(_newWalletName.value);
    _callBack.addNewWallet(name);
    _newWalletName.value = '';
  }

  void newWalletNames(Iterable<String> names) {
    _transactFromName.options.clear();
    _transactToName.options.clear();
    for (String name in names) {
      _transactFromName.options.add(OptionElement()
        ..value = name
        ..label = name);
      _transactToName.options.add(OptionElement()
        ..value = name
        ..label = name);
    }
  }
}
