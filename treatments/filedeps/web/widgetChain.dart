import 'dart:collection';
import 'dart:html';

import 'package:FileDepsTreatment/blockchain.dart';
import 'package:validators/sanitizers.dart' as sani;

import 'callback.dart';
import 'widget.dart';

class WidgetChain implements Widget {
  final CallBack _callBack;
  DivElement _group;
  DivElement _chainBlocks;

  WidgetChain(this._callBack) {
    final text = new DivElement()
      ..innerText = 'The blocks which make up the current chain:'
      ..style.marginBottom = '4px';

    _chainBlocks = new DivElement()..style.marginLeft = '10px';

    _group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(_chainBlocks);
  }
  DivElement get widget => _group;

  void updateChain(UnmodifiableListView<Block> chain) {
    final transactions = List<DivElement>();

    for (Block block in chain) {
      final text = new DivElement()
        ..innerText = sani.escape(block.toString())
        ..style.border = '1px solid black'
        ..style.borderLeft = '6px solid darkred'
        ..style.padding = '4px'
        ..style.marginBottom = '6px'
        ..style.whiteSpace = 'pre'
        ..style.marginBottom = '4px';
      transactions.add(text);
    }

    _chainBlocks.children.clear();
    _chainBlocks.children.addAll(transactions);
  }
}
