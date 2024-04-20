import 'package:flutter/material.dart';

typedef RequestArgsRecord = (int, String, String);

class Response {
  final int statusCode;
  final String payload;
  final String timestamp;

  static final Map<RequestArgsRecord, Response> _cache =
      <RequestArgsRecord, Response>{};

  factory Response(int statusCode, String payload, String timestamp) {
    final RequestArgsRecord packet = (statusCode, payload, timestamp);

    if (_cache.containsKey(packet)) {
      return _cache[(statusCode, payload, timestamp)]!;
    }

    final Response req = Response._internal(statusCode, payload, timestamp);
    _cache[packet] = req;
    return req;
  }

  const Response._internal(this.statusCode, this.payload, this.timestamp);
}

class ResponseCard extends StatelessWidget {
  final Response res;
  const ResponseCard({super.key, required this.res});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(15.0),
      child: Stack(
        children: [
          Column(
            children: [
              Align(
                alignment: Alignment.topLeft,
                child: Text("Status Code: ${res.statusCode}"),
              ),
              Align(
                alignment: Alignment.centerLeft,
                child: SelectableText(res.payload),
              ),
            ],
          ),
          Align(
            alignment: Alignment.topRight,
            child: Text(res.timestamp),
          )
        ],
      ),
    );
  }
}
