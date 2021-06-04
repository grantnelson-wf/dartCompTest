import 'package:cryptography/cryptography.dart' as cryptography;

import 'bytedata.dart';
import 'blockchain.dart';
import 'cancelable.dart';
import 'minergroup.dart';
import 'constants.dart';

/// A wallet represents an account for someone making transactions in this block chain.
class Wallet {
  String _name;
  final cryptography.KeyPair _keys;

  Wallet._(this._name, this._keys);

  /// Creates a new wallet with the given name and a brand new public/private key.
  /// The new public key is the address for this wallet.
  static Future<Wallet> create(String name) async {
    final keys = await signatureAlgorithm.newKeyPair();
    return Wallet._(name, keys);
  }

  /// Gets the unique address for this wallet.
  ByteData get address => ByteData(_keys.publicKey.bytes);

  /// Gets or sets the name to show for this wallet.
  String get name => _name;
  set name(String name) => _name = name;

  /// Gets the balance of this wallet as currently known by the block chain.
  /// This does not take into account any pending transactions.
  double balance(BlockChain chain) => chain.balance(address);

  /// Transfers an amount from this wallet to the given receiver wallet.
  /// Returns true if the transaction was added, false if not.
  Future<bool> transfer(BlockChain chain, Wallet receiver, double amount) async =>
      chain.createTransaction(_keys, receiver?.address ?? '', amount ?? 0.0);

  /// Starts mining the next block by joining the given miner group.
  /// Returns the cancelable to cancel mining or null if there is nothing to mine.
  Future<Cancelable> mine(MinerGroup group) => group.start(address);

  /// Gets a human readable string for debugging.
  @override
  String toString() => '$_name ($address)';
}
