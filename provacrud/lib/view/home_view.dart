import 'package:flutter/material.dart';
import 'package:provacrud/service/user_service.dart';
import 'package:provacrud/view/users_view.dart';

class HomeView extends StatefulWidget {
  const HomeView({super.key});

  @override
  State<HomeView> createState() => _HomeViewState();
}

class _HomeViewState extends State<HomeView> {
  final UserService _userService = UserService();
  var nameController = TextEditingController();
  var idController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Adicionar novo usuário'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: idController,
                decoration: const InputDecoration(labelText: 'Id'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: nameController,
                decoration: const InputDecoration(labelText: 'Nome'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () {
                  _userService.postData(
                      name: nameController.text,
                      id: int.tryParse(idController.text) ?? 0);
                  nameController.clear();
                  idController.clear();
                },
                child: const Text('Enviar'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(builder: (context) => const UsersView()),
                  );
                },
                child: const Text('Ir para listagem'),
              ),
            )
          ],
        ),
      ),
    );
  }
}
