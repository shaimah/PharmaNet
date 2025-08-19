import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../providers/inventory_provider.dart';
import '../services/agent_api.dart';

class SearchScreen extends StatefulWidget {
  final AgentApi api;
  const SearchScreen({super.key, required this.api});

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  final TextEditingController _controller = TextEditingController();
  bool loading = false;

  void search() async {
    setState(() => loading = true);
    try {
      final results = await widget.api.searchInventory(_controller.text);
      Provider.of<InventoryProvider>(context, listen: false)
          .updateInventory(results);
    } catch (e) {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(e.toString())));
    } finally {
      setState(() => loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    final inventory = Provider.of<InventoryProvider>(context).inventory;

    return Scaffold(
      appBar: AppBar(title: const Text('PharmaNet Inventory')),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            Row(
              children: [
                Expanded(
                  child: TextField(
                    controller: _controller,
                    decoration: const InputDecoration(
                        labelText: 'Enter drug name'),
                  ),
                ),
                const SizedBox(width: 8),
                ElevatedButton(onPressed: search, child: const Text('Search'))
              ],
            ),
            const SizedBox(height: 16),
            loading
                ? const CircularProgressIndicator()
                : Expanded(
                    child: ListView.builder(
                      itemCount: inventory.length,
                      itemBuilder: (context, index) {
                        final item = inventory[index];
                        return ListTile(
                          title: Text(item['ProductName'] ?? ''),
                          subtitle: Text(
                              'Qty: ${item['Quantity'] ?? 0} | Expiry: ${item['ExpiryDate'] ?? '-'}'),
                        );
                      },
                    ),
                  )
          ],
        ),
      ),
    );
  }
}
