import 'dart:html';
import 'dart:collection';

import 'package:validators/sanitizers.dart' as sani;
import 'package:FileDepsTreatment/blockchain.dart';

import 'callback.dart';

class Page {
  CallBack _callBack;
  InputElement _newWalletName;
  SelectElement _transactFromName;
  SelectElement _transactToName;
  InputElement _transactAmount;
  TableElement _balanceTable;

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

    final group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    document.body.append(group);
  }

  void _setupMiningGroup() {
    final text = new DivElement()
      ..innerText = 'Start mining the next block with all the wallets:'
      ..style.marginBottom = '4px';

    final startButton = new InputElement()
      ..type = 'submit'
      ..value = 'Start'
      ..style.marginRight = '4px'
      ..onClick.listen(_onStartMining);

    final cancelButton = new InputElement()
      ..type = 'submit'
      ..value = 'Cancel'
      ..onClick.listen(_onCancelMining);

    final nameDiv = new DivElement()
      ..style.display = 'flex'
      ..append(startButton)
      ..append(cancelButton);

    final group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    document.body.append(group);
  }

  void _setupBalanceGroup() {
    final text = new DivElement()
      ..innerText = 'The balances for all wallets:'
      ..style.marginBottom = '4px';

    _balanceTable = new TableElement();

    final group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(_balanceTable);

    document.body.append(group);
  }

  void _setupPendingGroup() {
    // TODO: FINISH
  }

  void _setupChainGroup() {
    // TODO: FINISH
  }

  void _onNewWalletAdd(_) {
    final name = sani.escape(_newWalletName.value);
    if (name.isEmpty) return;
    _callBack.addNewWallet(name);
    _newWalletName.value = '';
  }

  void _onMakeTransaction(_) {
    final fromName = _transactFromName.value;
    final toName = _transactToName.value;
    final amount = double.tryParse(_transactAmount.value);
    if (amount == null || amount <= 0.0) return;
    _callBack.newTransaction(fromName, toName, amount);
    _transactAmount.value = '';
  }

  void _onStartMining(_) => _callBack.startMining();

  void _onCancelMining(_) => _callBack.cancelMining();

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

  void updateBalances(Map<Wallet, double> balances) {
    final rows = List<TableRowElement>();
    final header = TableRowElement()
      ..append(TableCellElement()..text = 'name')
      ..append(TableCellElement()..text = 'amount')
      ..style.fontWeight = 'bold';
    rows.add(header);
    balances.forEach((wallet, amount) {
      final row = TableRowElement()
        ..append(TableCellElement()..text = wallet.name)
        ..append(TableCellElement()..text = '$amount');
      rows.add(row);
    });
    _balanceTable.children.clear();
    _balanceTable.children.addAll(rows);
  }

  void updatePending(UnmodifiableListView<Transaction> pending) {
    // TODO: FINISH
  }

  void updateChain(UnmodifiableListView<Block> chain) {
    // TODO: FINISH
  }
}
