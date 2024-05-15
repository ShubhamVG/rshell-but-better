import 'dart:io';
import 'dart:typed_data';

import '../bloc/connections_bloc/connections_bloc.dart';

late final Socket socket;

Future<void> tcpStuff({required ConnectionsBloc connectionsBloc}) async {
  const String serverIp = '127.0.0.1';
  const int serverPort = 8080;

  // TODO: Add error handler
  socket = await Socket.connect(serverIp, serverPort);

  socket.listen((data) {
    final String str = String.fromCharCodes(data);
    connectionsBloc.add(ConnectionAdded(str));
  }, onDone: () {
    socket.flush();
  });
}

Future<void> sendData({
  required Uint8List payload,
}) async {
  socket.write(payload);
}
