# Solution

Move this code to a separate new method (or function) and replace the old code with a call to the method.

```ts
printOwing(): void {
  printBanner();
  printDetails(getOutstanding());
}

printDetails(outstanding: number): void {
  console.log("name: " + name);
  console.log("amount: " + outstanding);
}
```
