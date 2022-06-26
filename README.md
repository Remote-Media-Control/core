# Remote Media Control core

A project for controlling media player over network.

## Core concepts

**Sever** - daemon/app, which start on computer (phone maybe in future), that play music.

**Client** - app, which start anywhere in network availability with server - another computer, phone, etc.

**Client** exposes interface to control music, that play on **computer-server**.

## Current status

This repo (`core`) contain GRPC proto and reference CLI client in go.
There is also a server implementation, but so far its functionality is limited.

[RMC/windows](https://github.com/Remote-Media-Control/windows) contains server impementation for windows, in C# for APS.NET Core.

[RMC/client](https://github.com/Remote-Media-Control/client) contains client implementation, in dart/flutter - so it should work on every platform.

## Protocol

RMC uses `grpc` + `protobuf`, currently over insecure http.

If you wish to implement custom client/server, you can find (and you must) protocol at [proto.proto](proto/proto.proto)
