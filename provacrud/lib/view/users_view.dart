import 'package:flutter/material.dart';
import 'package:provacrud/model/user_model.dart';
import 'package:provacrud/service/user_service.dart';
import 'package:provacrud/view/update_view.dart';

class UsersView extends StatefulWidget {
  const UsersView({super.key});

  @override
  State<UsersView> createState() => _UsersViewState();
}

class _UsersViewState extends State<UsersView> {
  final UserService _userService = UserService();
  late Future<List<UserModel>> _usersFuture;

  @override
  void initState() {
    super.initState();
    _usersFuture = _userService.fetchData();
  }

  void _refreshUsers() {
    setState(() {
      _usersFuture = _userService.fetchData();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: FutureBuilder<List<UserModel>>(
        future: _usersFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else if (snapshot.hasError) {
            return Center(
              child: Text('Error: ${snapshot.error}'),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(
              child: Text('Nenhum usuário encontrado'),
            );
          } else {
            List<UserModel> users = snapshot.data!;
            return ListView.builder(
              itemCount: users.length,
              itemBuilder: (context, index) {
                return SizedBox(
                  height: 50,
                  child: Card(
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Name: ${users[index].name}",
                        ),
                        Row(
                          children: [
                            IconButton(
                              onPressed: () async {
                                int id = users[index].id;
                                await _userService.deleteData(id: id);
                                _refreshUsers();
                              },
                              icon: const Icon(Icons.delete),
                            ),
                            IconButton(
                              onPressed: () async {
                                int? id = users[index].id;
                                await Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => UpdateView(
                                      userId: id,
                                      name: users[index].name,
                                    ),
                                  ),
                                );
                                _refreshUsers();
                              },
                              icon: const Icon(Icons.edit),
                            ),
                          ],
                        ),
                      ],
                    ),
                  ),
                );
              },
            );
          }
        },
      ),
    );
  }
}
