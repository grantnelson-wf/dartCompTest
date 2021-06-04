part of blockchain;

/// Wraps an integer list to make it work for addresses and hashes.
class ByteData {
  final List<int> bytes;

  /// Creates a new data object.
  ByteData(this.bytes);

  /// Creates new data which is just the [count] number of zeros.
  factory ByteData.zeros(int count) => ByteData(List<int>.filled(count, 0));

  /// Determines if this data starts with the [other] data.
  bool startsWith(ByteData other) {
    final count = other.bytes.length;
    if (bytes.length < count) return false;
    for (int i = 0; i < count; i++) {
      if (bytes[i] != other.bytes[i]) return false;
    }
    return true;
  }

  /// Gets a human readable string for debugging.
  @override
  String toString() => '[${base64.encode(bytes)}]';

  /// Used to determine if the given [other] is equivalent to this data.
  @override
  bool operator ==(Object other) {
    if (other is! ByteData) return false;
    final obytes = (other as ByteData).bytes;
    final count = obytes.length;
    if (bytes.length != count) return false;
    for (int i = 0; i < count; i++) {
      if (bytes[i] != obytes[i]) return false;
    }
    return true;
  }
}
