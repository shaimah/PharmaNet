import 'package:flutter/material.dart';

class InventoryProvider extends ChangeNotifier {
  List<Map<String, dynamic>> _inventory = [];

  List<Map<String, dynamic>> get inventory => _inventory;

  void updateInventory(List<Map<String, dynamic>> newInventory) {
    _inventory = newInventory;
    notifyListeners();
  }
}
