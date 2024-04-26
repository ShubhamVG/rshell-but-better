import 'dart:io';
import 'dart:typed_data';

import '../bloc/connections_bloc/connections_bloc.dart';

late final Socket socket;

Future<void> tcpStuff({required ConnectionsBloc connectionsBloc}) async {
  const String serverIp = '127.0.0.1';
  const int serverPort = 8080;

  socket = await Socket.connect(serverIp, serverPort);

  socket.listen((data) {
    final String str = String.fromCharCodes(data);
    connectionsBloc.add(ConnectionAdded(str));
  });
}

Future<void> sendData({
  required int taskCode,
  required Uint8List payload,
}) async {
  payload.insert(0, taskCode);
  socket.write(payload);
}
