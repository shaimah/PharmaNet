import 'dart:convert';
import 'package:http/http.dart' as http;

class AgentApi {
  final String baseUrl;

  AgentApi({required this.baseUrl});

  Future<List<Map<String, dynamic>>> searchInventory(String query) async {
    final url = Uri.parse('$baseUrl/v1/inventory/search?q=$query&limit=100');
    final response = await http.get(url, headers: {
      'X-Agent-Token': 'CHANGE_ME', // token from your agent
    });

    if (response.statusCode == 200) {
      final List data = jsonDecode(response.body);
      return data.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to fetch inventory');
    }
  }
}
