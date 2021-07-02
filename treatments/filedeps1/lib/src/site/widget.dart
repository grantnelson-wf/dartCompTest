import 'dart:html';

/// The interface needed for a widget to be added to this page.
abstract class Widget {
  /// Gets the div element containing the widget.
  DivElement get widget;
}
