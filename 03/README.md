# 03 - NOP what I was expecting

![](../img/gamer.svg)

Great job on implementing the `HALT` instruction! We now have a _programmable_ computer system, even though the programs we can write are rather simple. This is the minimal valid G-machine program:

```
HALT
```

In fact, that's also the _maximal_ program right now, since while we can write longer programs by repeating the `HALT` instruction, the extra instructions have no effect.

We ran your prototype by the Marketing group, and the feedback was generally positive, but they asked if you couldn't add at least one more instruction, so that we can write and sell useful software for the machine.

## The NOP instruction

The next instruction to implement will be `NOP`, short for No OPeration, which does nothing. This might sound a bit similar to the `HALT` instruction, which does nothing and halts, but there _is_ a difference: the `NOP` instruction doesn't halt! Let's assign it opcode 1.

So let's do another thought experiment. What happens if we write the opcode for the `NOP` instruction into memory address zero, and start the machine? (Think about it before you read on.)

Well, we know P starts at zero, so the first thing the machine will do is read the instruction at address zero, which is `NOP`. Since this has no effect, the fetch-execute cycle will continue, and the machine will fetch the instruction at address 1, which is `HALT`. And the machine should stop, with the program counter P containing the value `2`.

To put it another, equivalent, way, we're submitting the following program to the machine:

```
NOP
HALT
```

Let's make it work!

**TASK:** Write a test function `TestNop` which does the following:

1. Creates a new G-machine.
2. Sets the contents of the first memory location to 1.
3. Calls `Run()` on the machine.
4. Tests that the machine's `P` register contains the value `2`. If not, the test should fail with a message like `"want P == 2, got ..."`

The test should fail, we expect, because we haven't yet implemented the `NOP` instruction. If we've strictly obeyed the test-driven development process, we haven't even implemented a loop in the `Run()` method, or read any instructions from memory, because we didn't need to until now. So the test should fail because P contains `1` instead of `2`. (If it fails for any other reason, keep working, until it fails for that reason.)

**TASK:** Write the minimum code necessary to make the test pass. _Now_ it's necessary to write a loop, and read the next opcode from memory, and take different actions depending on its value. If we'd done this before, even though the tests didn't require it, we would have committed the sin of premature engineering.

Once this test passes, we can do a little refactoring.

**TASK:** Define integer constants `OpHALT` and `OpNOP`, with the values 0 and 1 respectively.

Refactor the tests and the `gmachine` package to use these constants (for example, in `TestNop`, we should set the contents of address zero to `OpNOP`, instead of a literal `1`.)

Use the tests to make sure that your refactoring didn't break anything.

## Done?

When the tests pass, you're done! Go on to the next exercise:

* [Getting in Shape](../02/README.md)

## What's this?

This is one of a set of Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [The G-machine](../README.md).

<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>
