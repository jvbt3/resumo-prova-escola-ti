import 'package:flutter/material.dart';
import 'package:provacrud/service/user_service.dart';

class UpdateView extends StatefulWidget {
  final int userId;
  final String name;

  UpdateView({
    super.key,
    required this.userId,
    required this.name,
  });

  @override
  State<UpdateView> createState() => _UpdateViewState();
}

class _UpdateViewState extends State<UpdateView> {
  final UserService _userService = UserService();
  late TextEditingController nameController;
  late TextEditingController idController;

  @override
  void initState() {
    super.initState();
    nameController = TextEditingController(text: widget.name);
    idController = TextEditingController(text: widget.userId.toString());
  }

  @override
  void dispose() {
    nameController.dispose();
    idController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Atualizar usuário'),
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
                readOnly: true,
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
                onPressed: () async {
                  await _userService.updateData(
                    id: widget.userId,
                    name: nameController.text,
                  );
                  Navigator.pop(context);
                },
                child: const Text('Enviar'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
