class UserModel {
  String name;
  int id;
  UserModel({
    required this.name,
    required this.id,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'name': name,
      'id': id,
    };
  }

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      name: json['name'] as String,
      id: json['id'] as int,
    );
  }
}
