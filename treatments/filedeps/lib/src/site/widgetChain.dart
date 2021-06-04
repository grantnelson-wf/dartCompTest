import 'dart:collection';
import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';
import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';
import 'widget.dart';

/// A widget for showing the state of the blocks in the chain.
class WidgetChain implements Widget {
  final CallBack _callBack;
  DivElement _group;
  DivElement _chainBlocks;

  /// Creates a new block chain widget.
  WidgetChain(this._callBack) {
    final text = DivElement()
      ..innerText = 'The blocks which make up the current chain:'
      ..style.marginBottom = '4px';

    _chainBlocks = DivElement()..style.marginLeft = '10px';

    _group = DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(_chainBlocks);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Creates a new cell element for the transaction table.
  TableCellElement _newCell(String text) => TableCellElement()
    ..text = text
    ..style.paddingLeft = '5px'
    ..style.paddingRight = '5px'
    ..style.borderLeft = '1px solid black'
    ..style.borderRight = '1px solid black';

  /// Creates a new key/value row for the header data for a block.
  TableRowElement _newKeyValueRow(String key, String value) => TableRowElement()
    ..append(TableCellElement()
      ..text = key
      ..style.fontWeight = 'bold'
      ..style.paddingLeft = '5px'
      ..style.paddingRight = '5px')
    ..append(TableCellElement()
      ..text = value
      ..style.paddingLeft = '5px'
      ..style.paddingRight = '5px');

  /// Updates the display fo the block chain.
  void updateChain(UnmodifiableListView<Block> chain) {
    final blocks = List<DivElement>();

    for (Block block in chain) {
      final blockData = TableElement()
        ..style.border = 'none'
        ..append(_newKeyValueRow('timestamp:', block.timestamp.toString()))
        ..append(_newKeyValueRow('previous hash:', sani.escape(block.previousHash.toString())))
        ..append(_newKeyValueRow('hash:', sani.escape(block.hash.toString())))
        ..append(_newKeyValueRow('nonce:', '${block.nonce}'))
        ..append(_newKeyValueRow('miner address:', _callBack.nameForAddress(block.minerAddress)));

      final transactions = new TableElement()
        ..style.marginLeft = '10px'
        ..style.border = 'none'
        ..style.borderCollapse = 'collapse';

      final header = TableRowElement()
        ..append(_newCell('timestamp'))
        ..append(_newCell('from'))
        ..append(_newCell('to'))
        ..append(_newCell('amount'))
        ..style.fontWeight = 'bold'
        ..style.borderBottom = '1px solid black';
      transactions.append(header);

      for (Transaction transaction in block.transactions) {
        final row = TableRowElement()
          ..append(_newCell(transaction.timestamp.toString()))
          ..append(_newCell(_callBack.nameForAddress(transaction.fromAddress)))
          ..append(_newCell(_callBack.nameForAddress(transaction.toAddress)))
          ..append(_newCell('${transaction.amount}'));
        transactions.append(row);
      }

      final text = DivElement()
        ..style.border = '1px solid black'
        ..style.borderLeft = '6px solid darkred'
        ..style.padding = '4px'
        ..style.marginBottom = '6px'
        ..style.whiteSpace = 'pre'
        ..style.marginBottom = '4px'
        ..append(blockData)
        ..append(transactions);

      blocks.add(text);
    }

    _chainBlocks.children.clear();
    _chainBlocks.children.addAll(blocks);
  }
}
