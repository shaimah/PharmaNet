import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'screens/search_screen.dart';
import 'providers/inventory_provider.dart';

void main() {
  runApp(
    ChangeNotifierProvider(
      create: (_) => InventoryProvider(),
      child: PharmaNetApp(),
    ),
  );
}

class PharmaNetApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'PharmaNet UI',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: SearchScreen(),
    );
  }
}
