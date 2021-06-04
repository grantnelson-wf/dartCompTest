import 'dart:html';

import 'callback.dart';
import 'widget.dart';

class WidgetMining implements Widget {
  final CallBack _callBack;
  DivElement _group;

  WidgetMining(this._callBack) {
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

    _group = new DivElement()
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);
  }

  DivElement get widget => _group;

  void _onStartMining(_) => _callBack.startMining();

  void _onCancelMining(_) => _callBack.cancelMining();
}
