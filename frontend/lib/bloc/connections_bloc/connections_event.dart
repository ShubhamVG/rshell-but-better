part of 'connections_bloc.dart';

sealed class ConnectionsEvent extends Equatable {
  final String connection;

  const ConnectionsEvent(this.connection);

  @override
  List<Object> get props => [connection];
}

class ConnectionAdded extends ConnectionsEvent {
  const ConnectionAdded(super.connection);
}

class ConnectionRemoved extends ConnectionsEvent {
  const ConnectionRemoved(super.connection);
}

class ConnectionDropdownUpdate extends ConnectionsEvent {
  const ConnectionDropdownUpdate(super.connection);
}
