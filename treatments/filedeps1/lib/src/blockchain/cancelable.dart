/// Interface for a set of async tasks which can be cancelled between those tasks.
class Cancelable {
  /// Should set a flag so that when the current task finishes
  /// all remaining tasks are discarded.
  void cancel() {}
}
