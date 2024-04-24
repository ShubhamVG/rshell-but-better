import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/connections_bloc/connections_bloc.dart';
import 'response_card.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final TextEditingController entryController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _body(),
    );
  }

  @override
  void dispose() {
    entryController.dispose();
    super.dispose();
  }

  Widget _body() {
    return Column(
      children: [_topBar(), _outputContainer(), _entryArea()],
    );
  }

  Widget _connectionsDropdown() {
    return BlocBuilder<ConnectionsBloc, ConnectionsState>(
      builder: (context, state) {
        return DropdownButton(
          value: (state.selected == "") ? null : state.selected,
          items: state.connections.map((item) {
            return DropdownMenuItem(value: item, child: Text(item));
          }).toList(),
          onChanged: (selectedVal) {
            if (selectedVal == null) return;
            context
                .read<ConnectionsBloc>()
                .add(ConnectionDropdownUpdate(selectedVal));
          },
        );
      },
    );
  }

  Widget _entryArea() {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Row(
        children: [
          Expanded(
            child: TextField(
              controller: entryController,
              keyboardType: TextInputType.multiline,
              minLines: 1,
              maxLines: 10,
              decoration: const InputDecoration(
                labelText: "Command to execute",
              ),
            ),
          ),
          ElevatedButton(onPressed: () {}, child: const Text("Send")), // TODO
        ],
      ),
    );
  }

  // TODO
  Widget _outputContainer() {
    final res1 = Response(200, "This is response 1 or something idk",
        DateTime.timestamp().toString());
    final res2 = Response(
        200,
        "This is resjfdsfsj fdsfsdf sd f sd fsfsdfsdg sg s gfdgfd hdhg gfhjfgh gfhfgh gponse 1 or something idk",
        DateTime.timestamp().toString());
    final res3 = Response(200, "Thdk", DateTime.timestamp().toString());
    final List<ResponseCard> responseCards = [
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
      ResponseCard(res: res1),
      ResponseCard(res: res2),
      ResponseCard(res: res3),
    ];

    return Expanded(
      child: SingleChildScrollView(
        child: Column(
          children: responseCards,
        ),
      ),
    );
  }

  Widget _topBar() {
    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: Row(
        children: [
          const Text("Connections:"),
          const SizedBox(
            width: 5,
          ),
          _connectionsDropdown(),
        ],
      ),
    );
  }
}
