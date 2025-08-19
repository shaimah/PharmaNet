import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class StockRecord {
  final String productId;
  final String productName;
  final int quantity;
  final String lastUpdated;

  StockRecord({
    required this.productId,
    required this.productName,
    required this.quantity,
    required this.lastUpdated,
  });

  factory StockRecord.fromJson(Map<String, dynamic> json) {
    return StockRecord(
      productId: json['ProductID'] ?? '',
      productName: json['ProductName'] ?? '',
      quantity: json['Quantity'] ?? 0,
      lastUpdated: json['LastUpdated'] ?? '',
    );
  }
}

class InventoryProvider extends ChangeNotifier {
  List<StockRecord> records = [];
  bool isLoading = false;

  Future<void> search(String query) async {
    isLoading = true;
    notifyListeners();

    final url = Uri.parse('http://127.0.0.1:8081/v1/inventory/search?query=$query&limit=100');
    final response = await http.get(url, headers: {'X-Agent-Token': 'CHANGE_ME'});

    if (response.statusCode == 200) {
      final List data = json.decode(response.body);
      records = data.map((e) => StockRecord.fromJson(e)).toList();
    } else {
      records = [];
    }

    isLoading = false;
    notifyListeners();
  }
}
