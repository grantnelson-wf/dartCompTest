import 'package:cryptography/cryptography.dart' as cryptography;

import 'bytedata.dart';
import 'blockchain.dart';
import 'cancelable.dart';
import 'minergroup.dart';
import 'constants.dart';

/// A wallet represents an account for someone making transactions in this block chain.
class Wallet {
  String _name;
  final ByteData _address;
  final cryptography.SimpleKeyPair _keys;

  Wallet._(this._name, this._address, this._keys);

  /// Creates a new wallet with the given name and a brand new public/private key.
  /// The new public key is the address for this wallet.
  static Future<Wallet> create(String name) async {
    final keys = await signatureAlgorithm.newKeyPair();
    final address = ByteData((await keys.extractPublicKey()).bytes);
    return Wallet._(name, address, keys);
  }

  /// Gets the unique address for this wallet.
  ByteData get address => _address;

  /// Gets or sets the name to show for this wallet.
  String get name => _name;
  set name(String name) => _name = name;

  /// Gets the balance of this wallet as currently known by the block chain.
  /// This does not take into account any pending transactions.
  double balance(BlockChain chain) => chain.balance(address);

  /// Transfers an amount from this wallet to the given receiver wallet.
  /// Returns true if the transaction was added, false if not.
  Future<bool> transfer(BlockChain chain, Wallet receiver, double amount) async =>
      chain.createTransaction(_keys, receiver.address, amount);

  /// Starts mining the next block by joining the given miner group.
  /// Returns the cancelable to cancel mining or null if there is nothing to mine.
  Future<Cancelable?> mine(MinerGroup group) => group.start(address);

  /// Gets a human readable string for debugging.
  @override
  String toString() => '$_name ($address)';
}
