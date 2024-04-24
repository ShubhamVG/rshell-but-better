import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

part 'connections_event.dart';
part 'connections_state.dart';

class ConnectionsBloc extends Bloc<ConnectionsEvent, ConnectionsState> {
  ConnectionsBloc() : super(const ConnectionsState(connections: <String>[])) {
    on<ConnectionAdded>(_connectionAdded);
    on<ConnectionRemoved>(_connectionRemoved);
    on<ConnectionDropdownUpdate>(_updateDropdownSelected);
  }

  FutureOr<void> _connectionAdded(
    ConnectionAdded event,
    Emitter<ConnectionsState> emit,
  ) {
    emit(state.copyWith(added: event.connection));
  }

  FutureOr<void> _connectionRemoved(
    ConnectionRemoved event,
    Emitter<ConnectionsState> emit,
  ) {
    emit(state.copyWith(removed: event.connection));
  }

  FutureOr<void> _updateDropdownSelected(
    ConnectionDropdownUpdate event,
    Emitter<ConnectionsState> emit,
  ) {
    emit(state.copyWith(selectedp: event.connection));
  }
}
