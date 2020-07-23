# 02 - Halt and Catch Fire

![](../img/lightbulb.png)

Just FYI, I ran your draft G-machine design past the executive steering committee, and they're all really stoked! Of course it's early days, but I'm sure this is going to be our next killer product. Let's start filling in some of the details.

## The fetch-execute cycle
The next feature we'll need in our virtual CPU is what's called the _fetch-execute cycle_. Essentially all computers work this way:

1. Fetch the next instruction from memory
2. Execute it.
3. Go to step 1.

## The program counter

Saying 'the _next_ instruction' implies that we have some way of remembering where we currently 'are' in memory. That is to say, we need a _register_ on the G-machine which holds the memory address of the next instruction to execute. This is what the P register is for ('P' stands for 'Program Counter', which is the traditional name for this register).

## Instructions

We also need some concept of what an 'instruction' is. You probably know that _machine language_ is the name we give to the set of instructions which a given CPU can understand. For example, the x86_64 processor understands x86_64 machine language. This is the CPU's 'native' language, if you like. If you write a program in machine language, you can run it directly on the processor. Programs in other languages need to be translated (_compiled_) into the right machine language for the CPU you want to run them on.

## Opcodes

Each instruction is represented by a numeric code, called an _opcode_, where each number 0, 1, 2... represents a distinct instruction. A program for the G-machine consists of a sequence of opcodes, perhaps with some accompanying data.

We can imagine a variety of useful instructions which the G-machine might implement: for example, if we want to do arithmetic, we might need something like an ADD instruction.

## The HALT instruction

For now, let's keep it simple, and implement a single instruction named `HALT`, which does nothing except stop the machine. It's entirely up to us which numeric values to assign to opcodes, and it makes no difference to the machine, but for simplicity let's assign `HALT` the opcode 0.

## The `Run()` method
We'll need a way for users to start the machine running, which is to say performing the fetch-execute cycle, until it's either told to stop, or runs into some kind of error. So let's provide a method on the `Machine` object named `Run()` to do this.

What would happen if we were to call the `Run()` method to start a new machine running, given that its memory and registers contain all zeroes? Well, let's follow the fetch-execute cycle:

1. Fetch the next instruction from memory. That is to say, look at the P register to see what memory address it contains, and read the instruction at that address.
2. Since the P register contains zero, we read the instruction at address zero, which is zero.
3. We increment the P register so that it points to the next memory address to read from (in this case, 1).
4. Execute the current instruction, whose opcode is zero. This is the opcode for the `HALT` instruction, so instead of jumping back to step 1, the `Run()` method should return instead.

So the upshot of all this is that if you call `Run()` on a new machine, it should return almost immediately (because it read and executed the `HALT` instruction), and the state of the machine should be unchanged except that the P register now contains the value `1`.

Let's find out!

**TASK:** Write a test function `TestHalt` which does the following:

1. Creates a new G-machine.
2. Calls `Run()` on the machine.
3. Tests that the machine's `P` register contains the value `1`. If not, the test should fail with a message like `"want P == 1, got ..."`
4. Calls `Run()` again.
5. Tests that P contains `2`.

This test will not compile yet, of course, because we haven't written the `Run()` method. If it fails to compile for any other reason, keep working on it until it fails to compile because of the missing `Run()` method.

**TASK:** Write the _minimum_ code necessary to make the test pass. (I'm serious about this. For example, even though we talked about a fetch-execute _cycle_, you won't need to implement a loop inside the `Run()` method, because the test doesn't require it to loop. All it needs to do is increment the P register and return.)

## Done?

When the tests pass, you're done! Go on to the next exercise:

* [NOP what I was expecting](../03/README.md)

## What's this?

This is one of a set of Go exercises by [John Arundel](https://bitfieldconsulting.com/golang/learn) called [The G-machine](../README.md).

<small>Gopher image by [egonelbre](https://github.com/egonelbre/gophers)</small>
