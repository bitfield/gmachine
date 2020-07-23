# 01 - Soul of a New Machine

![](../img/soldering.svg)

Welcome to your first day as Vice-President of Virtual Processors! You will find a key to the executive washroom on your desk, and free candy and snacks are available in the cafeteria. Please note there is no smoking anywhere in the building.

Your first job is to begin the design of a new virtual CPU, called the _G-machine_. Don't worry, we'll be tackling this project in easy stages. Let's first set out what exactly is required.

You will be developing a Go library which implements the G-machine. Users should be able to import your library and use it to write programs which run on the G-machine. We will develop a minimum viable product first, and gradually add more features as we go.

We will be using a simplified model of a computer system in which there are three main components:

* A _CPU_ (Central Processing Unit) which executes instructions in sequence and has _registers_ which store data while it's being processed
* A _memory_ space where the CPU can move data to or from its registers
* A _BIOS_ (Basic Input/Output System) which provides communications facilities like reading or writing to a terminal

At any given moment, the G-machine has a certain _state_: the contents of its registers, plus the contents of its memory. It has a _clock_ which generates regular 'ticks' of time, and at each tick, or clock cycle, the machine updates its state according to what it happens to be doing.

The first thing users need to be able to do is to create a new G-machine they can use. So you'll be implementing the function `gmachine.New()`, which returns a G-machine in its default initial state, which is:

* Two 64-bit registers, A and P, containing zero
* 1024 64-bit words of memory, containing zeroes

The test is already written for you, so let's get started!

**TASK:** Write the minimum code to make the tests pass.

## Done?

When the tests pass, you're done! Hopefully that was pretty easy. Go on to the next exercise:

* [Halt and Catch Fire](../02/README.md)

## What's this?

This is one of a set of Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [The G-machine](../README.md).

<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>
