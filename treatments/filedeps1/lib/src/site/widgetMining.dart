import 'dart:html';

import 'callback.dart';
import 'widget.dart';

/// A widget for starting and cancelling the mining group.
class WidgetMining implements Widget {
  final CallBack _callBack;
  final DivElement _group;
  final InputElement _startButton;
  final InputElement _cancelButton;

  /// Creates the mining group control widget.
  WidgetMining._(this._callBack, this._group, this._startButton, this._cancelButton) {
    _startButton.onClick.listen(_onStartMining);
    _cancelButton.onClick.listen(_onCancelMining);
  }

  /// Creates the mining group control widget.
  factory WidgetMining(CallBack callBack) {
    final text = DivElement()
      ..innerText = 'Start mining the next block with all the wallets:'
      ..style.marginBottom = '4px';

    final startButton = InputElement()
      ..type = 'submit'
      ..value = 'Start'
      ..style.marginRight = '4px';

    final cancelButton = InputElement()
      ..type = 'submit'
      ..value = 'Cancel'
      ..disabled = true;

    final nameDiv = DivElement()
      ..style.display = 'flex'
      ..append(startButton)
      ..append(cancelButton);

    final group = DivElement()
      ..style.backgroundColor = 'white'
      ..style.border = '1px solid black'
      ..style.borderLeft = '6px solid darkred'
      ..style.padding = '4px'
      ..style.marginBottom = '6px'
      ..append(text)
      ..append(nameDiv);

    return WidgetMining._(callBack, group, startButton, cancelButton);
  }

  /// Gets the div element containing the widget.
  DivElement get widget => _group;

  /// Handles when start mining is clicked.
  void _onStartMining(_) => _callBack.startMining();

  /// Handles when cancel mining is clicked.
  void _onCancelMining(_) => _callBack.cancelMining();

  /// Updates the mining state to indicate if mining is running or not.
  void updateMiningState(bool mining) {
    _startButton.disabled = mining;
    _cancelButton.disabled = !mining;
  }
}
