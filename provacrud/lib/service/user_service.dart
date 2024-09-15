import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:provacrud/model/user_model.dart';

class UserService {
  Future<http.Response> postData(
      {required String name, required int id}) async {
    try {
      final response = await http.post(
        Uri.parse('http://localhost:8888/api/user'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, dynamic>{'name': name, 'id': id}),
      );

      return response;
    } catch (e) {
      throw Exception('Error: $e');
    }
  }

  Future<List<UserModel>> fetchData() async {
    try {
      final request = await http.get(
        Uri.parse('http://localhost:8888/api/user'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
      );

      if (request.statusCode == 200) {
        final Map<String, dynamic> jsonBody = jsonDecode(request.body);
        var list = jsonBody['data'] as List;
        List<UserModel> users = list.map((u) => UserModel.fromJson(u)).toList();

        return users;
      } else {
        throw Exception('Failed to load user data');
      }
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }

  Future<void> updateData({required int id, required String name}) async {
    try {
      // ignore: unused_local_variable
      final response = await http.patch(
        Uri.parse('http://localhost:8888/api/user/$id'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{'name': name}),
      );
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }

  Future<void> deleteData({required int id}) async {
    try {
      // ignore: unused_local_variable
      final response = await http.delete(
        Uri.parse('http://localhost:8888/api/user/$id'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
      );
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }
}
