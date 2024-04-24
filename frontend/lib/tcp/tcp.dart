import 'dart:io';
// import 'dart:typed_data';

import '../bloc/connections_bloc/connections_bloc.dart';

Future<void> tcpStuff({required ConnectionsBloc connectionsBloc}) async {
  const String serverIp = '127.0.0.1';
  const int serverPort = 8080;

  Socket.connect(serverIp, serverPort).then((socket) {
    socket.listen((data) {
      final String str = String.fromCharCodes(data);
      connectionsBloc.add(ConnectionAdded(str));
    });
  });
}
