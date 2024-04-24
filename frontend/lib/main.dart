import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'bloc/connections_bloc/connections_bloc.dart';
import 'tcp/tcp.dart';
import 'ui/home_screen.dart';

void main() async {
  final ConnectionsBloc connectionsBloc = ConnectionsBloc();
  await tcpStuff(connectionsBloc: connectionsBloc);
  runApp(MainApp(connectionsBloc: connectionsBloc));
}

class MainApp extends StatelessWidget {
  final ConnectionsBloc connectionsBloc;

  const MainApp({super.key, required this.connectionsBloc});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [BlocProvider(create: (_) => connectionsBloc)],
      child: MaterialApp(
        home: const HomeScreen(),
        theme: ThemeData.dark(),
      ),
    );
  }
}
