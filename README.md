# Improv
`Improv` is a `multi-language` [actor framework](https://en.wikipedia.org/wiki/Actor_model) for handling asynchronous distributed systems in a simple and efficient way. It is designed to be used in a `microservices` architecture, but it can be used in any kind of distributed system.

`Improv` is based on `HTTP/2`. It uses `HTTP/2` for communication between `actors`, and it uses `HTTP/2` for communication between `modules` and `services`. This allows you to build distributed systems that are efficient, scalable, and fault-tolerant.


## Features

- **Simple**: `Improv` is designed to be simple and easy to use. It is based on the `actor model`, which is a simple and powerful model for handling concurrency.
- **Efficient**: `Improv` is designed to be efficient. It is based on the `actor model`, which is a lightweight model for handling concurrency.
- **Scalable**: `Improv` is designed to be scalable. It is based on the `actor model`, which is a scalable model for handling concurrency.
- **Fault-tolerant**: `Improv` is designed to be fault-tolerant. It is based on the `actor model`, which is a fault-tolerant model for handling concurrency.
- **Distributed**: `Improv` is designed to be distributed. It is based on the `actor model`, which is a distributed model for handling concurrency.
- **Multi-language**: `Improv` is designed to be multi-language. It is based on the `actor model`, which is a multi-language model for handling concurrency.
- **Secure**: `Improv` is designed with security **in mind** (it's not perfect). All `messages` are encrypted with `AES` by default, however it can be disabled. Only the `actor` that receives the `message` can decrypt it.
- **Open-source**: `Improv` is open-source. You can use it for free and contribute to its development.


## Inspiration
Before you decide if you want to use `Improv`, you should know what went into it's development.

First and foremost, you must know, that `Improv` is inspired by the `actor model`. The `actor model` is a model of concurrent computation that treats `actors` as the universal primitives of concurrent computation. In response to a `message` that it receives, an `actor` can make local decisions, create more `actors`, send more `messages`, and determine how to respond to the next `message` received. `Actors` may modify their own private state, but can only affect each other indirectly through `messages` (avoiding the need for any locks).

The `actor model` has been around for a long time, and it has been used in many different systems. It is a simple and powerful model for handling concurrency, and it is well-suited for building distributed systems.

`Improv` is inspired by the `actor model`, but it is not a direct implementation of it. Instead, `Improv` is a `multi-language` framework that is designed to make it easy to build distributed systems using the `actor model`.

This allows you to build distributed systems in a simple and efficient way, without having to worry about the low-level details of concurrency and distributed systems.

You are able to divide each `actor` into `modules` and `services`. This allows you to build complex systems in a modular way, and to reuse code across different `actors`. 

`Improv` is designed to allow you to truly use the right tool for the job. Divide your system into `actors` and `modules` and use the right language for each part of your system. This allows you to build systems that are efficient, scalable, and fault-tolerant, without having to compromise on the language that you use.



### `Improv` is inspired by the following projects:
- [Akka](https://akka.io/): Akka is a toolkit and runtime for building highly concurrent, distributed, and resilient message-driven applications on the JVM.
- [Actix](https://actix.rs/): Actix is a powerful, pragmatic, and extremely fast web framework for Rust.


## Getting started

### Pick a language
`Improv` is a `multi-language` framework. You can use it in any language that you like. Here are some of the languages that `Improv` supports:

- `Rust`
- `Python`
- `Go`
- `TypeScript`
- `Zig` (maybe if I decide to learn it.)

### Install `Improv`
To install `Improv`, you need to install the `Improv` package for your language. Here are the installation instructions for each language:
