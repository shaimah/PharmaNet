import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../providers/inventory_provider.dart';

class SearchScreen extends StatelessWidget {
  final TextEditingController _controller = TextEditingController();

  @override
  Widget build(BuildContext context) {
    final provider = Provider.of<InventoryProvider>(context);

    return Scaffold(
      appBar: AppBar(title: Text('PharmaNet Inventory Search')),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: _controller,
              decoration: InputDecoration(
                labelText: 'Enter drug name',
                suffixIcon: IconButton(
                  icon: Icon(Icons.search),
                  onPressed: () {
                    provider.search(_controller.text);
                  },
                ),
              ),
            ),
            SizedBox(height: 20),
            provider.isLoading
                ? CircularProgressIndicator()
                : Expanded(
                    child: ListView.builder(
                      itemCount: provider.records.length,
                      itemBuilder: (context, index) {
                        final r = provider.records[index];
                        return ListTile(
                          title: Text(r.productName),
                          subtitle: Text('Qty: ${r.quantity} | Updated: ${r.lastUpdated}'),
                        );
                      },
                    ),
                  ),
          ],
        ),
      ),
    );
  }
}
