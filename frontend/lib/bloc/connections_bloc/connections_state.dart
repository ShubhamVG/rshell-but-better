part of 'connections_bloc.dart';

class ConnectionsState extends Equatable {
  final List<String> connections;
  final String selected;

  const ConnectionsState({required this.connections, this.selected = ""});

  ConnectionsState copyWith({
    String? added,
    String? removed,
    String? selectedp,
  }) {
    List<String> connectionsCopyUpdated = List.from(connections);
    if (added != null) {
      connectionsCopyUpdated.add(added);
    }

    if (removed != null) {
      connectionsCopyUpdated.remove(removed);
    }

    return ConnectionsState(
      connections: connectionsCopyUpdated,
      selected: selectedp ?? selected,
    );
  }

  @override
  List<Object> get props => [connections, selected];
}
