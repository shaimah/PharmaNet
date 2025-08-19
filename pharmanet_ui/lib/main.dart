import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'providers/inventory_provider.dart';
import 'screens/search_screen.dart';
import 'services/agent_api.dart';

void main() {
  runApp(const PharmaNetApp());
}

class PharmaNetApp extends StatelessWidget {
  const PharmaNetApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => InventoryProvider(),
      child: MaterialApp(
        title: 'PharmaNet UI',
        home: SearchScreen(
          api: AgentApi(baseUrl: 'http://127.0.0.1:8081'),
        ),
      ),
    );
  }
}
